package powerbi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/stpabhi/powerbi-go/types"
)

const (
	version   = "1.0.0"
	baseURL   = "https://api.powerbi.com/v1.0/myorg/"
	userAgent = "powerbi-go/" + version
	mediaType = "application/json"
)

// A Client manages communication with the Microsoft Power BI REST APIs.
type Client struct {
	// HTTP client used to communicate with the Microsoft Power BI REST APIs.
	HTTPClient *http.Client

	// Base URL for API requests. BaseURL should always be specified with a trailing slash.
	BaseURL *url.URL

	// User agent for client.
	UserAgent string

	// Reuse a single struct instead of allocating one for each service on the heap.
	common service

	// Add services here
	Groups *GroupsService
}

type service struct {
	client *Client
}

// addOptions adds the parameters in opts as URL query parameters to s. opts
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opts interface{}) (string, error) {
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opts)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

// NewFromToken returns a new Power BI API client with the given API token.
func NewFromToken(token string) *Client {
	sanitizedToken := strings.TrimSpace(token)
	ts := TokenTransport{
		AccessToken: sanitizedToken,
	}

	return NewClient(ts.Client())
}

// NewClient returns a new API client. If a nil httpClient is
// provided, a new http.Client will be used. To use API methods which
// require authentication, provide a http.Client that will perform the
// authentication for you (such as that provided by the golang.org/x/oauth2 library).
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(baseURL)

	c := &Client{HTTPClient: httpClient, BaseURL: baseURL, UserAgent: userAgent}
	c.common.client = c

	// Add services here
	c.Groups = (*GroupsService)(&c.common)

	return c
}

func (c *Client) patchJSON(ctx context.Context, path string, obj any, headerKV ...string) (*http.Request, *http.Response, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, nil, err
	}
	return c.doRequest(ctx, http.MethodPatch, path, bytes.NewBuffer(data), append(headerKV, "Content-Type", mediaType)...)
}

func (c *Client) postJSON(ctx context.Context, path string, obj any, headerKV ...string) (*http.Request, *http.Response, error) {
	var body io.Reader

	switch v := obj.(type) {
	case string:
		if v != "" {
			body = strings.NewReader(v)
		}
	default:
		data, err := json.Marshal(obj)
		if err != nil {
			return nil, nil, err
		}
		body = bytes.NewBuffer(data)
		headerKV = append(headerKV, "Content-Type", mediaType)
	}
	return c.doRequest(ctx, http.MethodPost, path, body, headerKV...)
}

func (c *Client) putJSON(ctx context.Context, path string, obj any, headerKV ...string) (*http.Request, *http.Response, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, nil, err
	}
	return c.doRequest(ctx, http.MethodPut, path, bytes.NewBuffer(data), append(headerKV, "Content-Type", mediaType)...)
}

func (c *Client) doRequest(ctx context.Context, method, path string, body io.Reader, headerKV ...string) (*http.Request, *http.Response, error) {
	u, err := c.BaseURL.Parse(path)
	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return nil, nil, err
	}

	if len(headerKV)%2 != 0 {
		return nil, nil, fmt.Errorf("length of headerKV must be even")
	}
	for i := 0; i < len(headerKV); i += 2 {
		req.Header.Add(headerKV[i], headerKV[i+1])
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	if resp.StatusCode > 399 {
		data, _ := io.ReadAll(resp.Body)
		msg := string(data)
		if len(msg) == 0 {
			msg = resp.Status
		}
		return nil, nil, &types.ErrHTTP{
			Code:    resp.StatusCode,
			Message: msg,
		}
	}
	return req, resp, err
}

func toObject[T any](resp *http.Response, obj T) (def T, _ error) {
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&obj); err != nil {
		return def, err
	}
	return obj, nil
}

func setCredentialsAsHeaders(req *http.Request, token string) *http.Request {
	// To set extra headers, we must make a copy of the Request so
	// that we don't modify the Request we were given. This is required by the
	// specification of http.RoundTripper.
	//
	// Since we are going to modify only req.Header here, we only need a deep copy
	// of req.Header.
	convertedRequest := new(http.Request)
	*convertedRequest = *req
	convertedRequest.Header = make(http.Header, len(req.Header))

	for k, s := range req.Header {
		convertedRequest.Header[k] = append([]string(nil), s...)
	}
	convertedRequest.Header.Set("Authorization", "Bearer "+token)
	return convertedRequest
}

// TokenTransport is an http.RoundTripper that authenticates all requests
// using the provided token.
type TokenTransport struct {
	AccessToken string

	// Transport is the underlying HTTP transport to use when making requests.
	// It will default to http.DefaultTransport if nil.
	Transport http.RoundTripper
}

// RoundTrip implements the RoundTripper interface.
func (t *TokenTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req2 := setCredentialsAsHeaders(req, t.AccessToken)
	return t.transport().RoundTrip(req2)
}

// Client returns an *http.Client that makes requests that are authenticated
// using Authorization Token.
func (t *TokenTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func (t *TokenTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

// PtrTo returns a pointer to the provided input.
func PtrTo[T any](v T) *T {
	return &v
}
