package powerbi

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stpabhi/powerbi-go/types"
)

const (
	testGroup   = "f089354e-8366-4e18-aea3-4cb4a3a50b48"
	testDataset = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
)

// testClient returns a client for a fake Power BI that answers with h.
func testClient(t *testing.T, h http.HandlerFunc) *Client {
	t.Helper()
	srv := httptest.NewServer(h)
	t.Cleanup(srv.Close)
	c := NewClient(srv.Client())
	c.BaseURL, _ = url.Parse(srv.URL + "/v1.0/myorg/")
	return c
}

func TestRefreshDataset(t *testing.T) {
	var body map[string]any
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/v1.0/myorg/groups/"+testGroup+"/datasets/"+testDataset+"/refreshes" {
			t.Errorf("request %s %s", r.Method, r.URL.Path)
		}
		b, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(b, &body)
		w.Header().Set("x-ms-request-id", "87f31ef7-1e3a-4006-9b0b-191693e79e9e")
		w.Header().Set("Location", "https://api.powerbi.com/v1.0/myorg/groups/"+testGroup+"/datasets/"+testDataset+"/refreshes/5b0b6d9c-1e3a-4006-9b0b-191693e79e9e")
		w.WriteHeader(http.StatusAccepted)
	})
	got, err := c.Datasets.Group().RefreshDataset(context.Background(), testGroup, testDataset, types.DatasetRefreshRequest{Type: types.DatasetRefreshTypeFull})
	if err != nil {
		t.Fatal(err)
	}
	// The Location names the refresh.
	if got.RequestID != "5b0b6d9c-1e3a-4006-9b0b-191693e79e9e" {
		t.Errorf("request ID %q", got.RequestID)
	}
	// Only what was set is sent: an enhanced refresh must not carry notifyOption.
	if len(body) != 1 || body["type"] != "Full" {
		t.Errorf("body %v", body)
	}
}

func TestRefreshHistory(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("$top") != "1" {
			t.Errorf("query %s", r.URL.RawQuery)
		}
		io.WriteString(w, `{"value":[{"refreshType":"ViaApi","startTime":"2017-06-13T09:25:43.153Z","endTime":"2017-06-13T09:31:43.153Z",
			"serviceExceptionJson":"{\"errorCode\":\"ModelRefreshFailed_CredentialsNotSpecified\"}","status":"Failed","requestId":"11bf290a-346b-48b7-8973-c5df149337ff",
			"refreshAttempts":[{"attemptId":1,"type":"Data"}]}]}`)
	})
	got, err := c.Datasets.Group().RefreshHistory(context.Background(), testGroup, testDataset, 1)
	if err != nil {
		t.Fatal(err)
	}
	if len(got.Value) != 1 || got.Value[0].Status != "Failed" || got.Value[0].ServiceExceptionJSON == "" || len(got.Value[0].RefreshAttempts) != 1 {
		t.Fatalf("history %+v", got)
	}
}

func TestExecuteQueriesKeepsNumbers(t *testing.T) {
	var body map[string]any
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(b, &body)
		io.WriteString(w, `{"results":[{"tables":[{"rows":[{"[n]":9007199254740993,"[d]":0.1}]}]}]}`)
	})
	c.UseNumber = true
	got, err := c.Datasets.Group().ExecuteQueries(context.Background(), testGroup, testDataset, types.DatasetExecuteQueriesRequest{
		Queries: []types.DatasetExecuteQueriesQuery{{Query: `EVALUATE ROW("n", 1)`}},
	})
	if err != nil {
		t.Fatal(err)
	}
	row := got.Results[0].Tables[0].Rows[0].(map[string]any)
	if n, ok := row["[n]"].(json.Number); !ok || n.String() != "9007199254740993" {
		t.Errorf("n = %#v", row["[n]"])
	}
	// impersonatedUserName is left out when empty.
	if _, ok := body["impersonatedUserName"]; ok {
		t.Errorf("body %v", body)
	}
}

func TestErrorsKeepHeaders(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("RequestId", "a1b2")
		w.Header().Set("Retry-After", "20")
		w.WriteHeader(http.StatusTooManyRequests)
		io.WriteString(w, `{"error":{"code":"TooManyRequests"}}`)
	})
	_, err := c.Datasets.Group().RefreshHistory(context.Background(), testGroup, testDataset, 1)
	var e *types.ErrHTTP
	if !errors.As(err, &e) || e.Code != http.StatusTooManyRequests || e.Header.Get("RequestId") != "a1b2" || e.Header.Get("Retry-After") != "20" {
		t.Fatalf("error %#v", err)
	}
}

func TestMaxResponseBytes(t *testing.T) {
	const answer = `{"value":[{"status":"Completed"}]}`
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, answer) })
	c.MaxResponseBytes = int64(len(answer)) - 1
	if _, err := c.Datasets.Group().RefreshHistory(context.Background(), testGroup, testDataset, 1); !errors.Is(err, ErrResponseTooLarge) {
		t.Fatalf("error %v", err)
	}
	// A body exactly at the limit is read whole.
	c.MaxResponseBytes = int64(len(answer))
	if got, err := c.Datasets.Group().RefreshHistory(context.Background(), testGroup, testDataset, 1); err != nil || got.Value[0].Status != "Completed" {
		t.Fatalf("history %+v %v", got, err)
	}
}
