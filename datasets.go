package powerbi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"path"
	"strconv"

	"github.com/stpabhi/powerbi-go/types"
)

const datasetsBasePath = "datasets"

type DatasetsService service

var _ DatasetsInterface = &DatasetsService{}

type DatasetsInterface interface {
	Group() DatasetGroup
}

// DatasetGroup calls the dataset APIs of one workspace, given by its group ID. An
// empty group ID is My workspace.
type DatasetGroup interface {
	BindToGateway(ctx context.Context, groupID, datasetID string, req types.BindToGatewayRequest) error
	CancelRefresh(ctx context.Context, groupID, datasetID, refreshID string) error
	DeleteDataset(ctx context.Context, groupID, datasetID string) error
	DiscoverGateways(ctx context.Context, groupID, datasetID string) (*types.GatewayList, error)
	ExecuteQueries(ctx context.Context, groupID, datasetID string, req types.DatasetExecuteQueriesRequest) (*types.DatasetExecuteQueriesResponse, error)
	Dataset(ctx context.Context, groupID, datasetID string) (*types.Dataset, error)
	DatasetToDataflowLinks(ctx context.Context, groupID string) (*types.DatasetToDataflowLinksResponse, error)
	DatasetUsers(ctx context.Context, groupID, datasetID string) (*types.DatasetUsersAccess, error)
	Datasets(ctx context.Context, groupID string) (*types.DatasetList, error)
	Datasources(ctx context.Context, groupID, datasetID string) (*types.DatasourceList, error)
	DirectQueryRefreshSchedule(ctx context.Context, groupID, datasetID string) (*types.DirectQueryRefreshSchedule, error)
	GatewayDatasources(ctx context.Context, groupID, datasetID string) (*types.GatewayDatasourceList, error)
	Parameters(ctx context.Context, groupID, datasetID string) (*types.MashupParameterList, error)
	RefreshDataset(ctx context.Context, groupID, datasetID string, req types.DatasetRefreshRequest) (*types.RefreshAccepted, error)
	RefreshExecutionDetails(ctx context.Context, groupID, datasetID, refreshID string) (*types.RefreshExecutionDetails, error)
	RefreshHistory(ctx context.Context, groupID, datasetID string, top int) (*types.RefreshList, error)
}

type datasetGroupService service

// datasetsPath is the datasets collection of a workspace: groups/{groupID}/datasets,
// or datasets in My workspace when groupID is empty. Power BI refuses the
// groups paths for a personal workspace (GroupNotAccessible), so My workspace
// is addressed without a group.
func datasetsPath(groupID string) string {
	if groupID == "" {
		return datasetsBasePath
	}
	return groupsBasePath + "/" + url.PathEscape(groupID) + "/" + datasetsBasePath
}

func (s *DatasetsService) Group() DatasetGroup {
	return &datasetGroupService{s.client}
}

// BindToGateway Binds the specified dataset from the specified workspace to the specified gateway, optionally with a given set of data source IDs.
// If you don't supply a specific data source ID, the dataset will be bound to the first matching data source in the gateway.
func (s *datasetGroupService) BindToGateway(ctx context.Context, groupID, datasetID string, req types.BindToGatewayRequest) error {
	u := fmt.Sprintf("%s/%s/Default.BindToGateway", datasetsPath(groupID), url.PathEscape(datasetID))
	_, resp, err := s.client.postJSON(ctx, u, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// CancelRefresh cancels the specified refresh operation for the specified dataset from the specified workspace.
func (s *datasetGroupService) CancelRefresh(ctx context.Context, groupID, datasetID, refreshID string) error {
	u := fmt.Sprintf("%s/%s/refreshes/%s", datasetsPath(groupID), url.PathEscape(datasetID), url.PathEscape(refreshID))
	_, resp, err := s.client.doRequest(ctx, "DELETE", u, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// DeleteDataset deletes the specified dataset from the specified workspace.
func (s *datasetGroupService) DeleteDataset(ctx context.Context, groupID, datasetID string) error {
	u := fmt.Sprintf("%s/%s", datasetsPath(groupID), url.PathEscape(datasetID))
	_, resp, err := s.client.doRequest(ctx, "DELETE", u, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// DiscoverGateways returns a list of gateways that the specified dataset from the specified workspace can be bound to.
func (s *datasetGroupService) DiscoverGateways(ctx context.Context, groupID, datasetID string) (*types.GatewayList, error) {
	u := fmt.Sprintf("%s/%s/Default.DiscoverGateways", datasetsPath(groupID), url.PathEscape(datasetID))
	_, resp, err := s.client.doRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.GatewayList
	_, err = toObject(resp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// ExecuteQueries executes Data Analysis Expressions (DAX) queries against the provided dataset.
func (s *datasetGroupService) ExecuteQueries(ctx context.Context, groupID, datasetID string, req types.DatasetExecuteQueriesRequest) (*types.DatasetExecuteQueriesResponse, error) {
	u := fmt.Sprintf("%s/%s/executeQueries", datasetsPath(groupID), url.PathEscape(datasetID))
	_, resp, err := s.client.postJSON(ctx, u, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var out types.DatasetExecuteQueriesResponse
	dec := json.NewDecoder(resp.Body)
	if s.client.UseNumber {
		dec.UseNumber()
	}
	if err := dec.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Dataset returns the specified dataset from the specified workspace.
func (s *datasetGroupService) Dataset(ctx context.Context, groupID, datasetID string) (*types.Dataset, error) {
	u := fmt.Sprintf("%s/%s", datasetsPath(groupID), url.PathEscape(datasetID))
	_, resp, err := s.client.doRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.Dataset{})
}

// DatasetToDataflowLinks returns a list of upstream dataflows for datasets from the specified workspace.
func (s *datasetGroupService) DatasetToDataflowLinks(ctx context.Context, groupID string) (*types.DatasetToDataflowLinksResponse, error) {
	u := datasetsPath(groupID) + "/upstreamDataflows"
	_, resp, err := s.client.doRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.DatasetToDataflowLinksResponse{})
}

// DatasetUsers returns a list of principals that have access to the specified dataset.
func (s *datasetGroupService) DatasetUsers(ctx context.Context, groupID, datasetID string) (*types.DatasetUsersAccess, error) {
	u := fmt.Sprintf("%s/%s/users", datasetsPath(groupID), url.PathEscape(datasetID))
	_, resp, err := s.client.doRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.DatasetUsersAccess{})
}

// Datasets returns a list of datasets from the specified workspace.
func (s *datasetGroupService) Datasets(ctx context.Context, groupID string) (*types.DatasetList, error) {
	u := datasetsPath(groupID)
	_, resp, err := s.client.doRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.DatasetList{})
}

// Datasources returns a list of data sources for the specified dataset from the specified workspace.
func (s *datasetGroupService) Datasources(ctx context.Context, groupID, datasetID string) (*types.DatasourceList, error) {
	u := fmt.Sprintf("%s/%s/datasources", datasetsPath(groupID), url.PathEscape(datasetID))
	_, resp, err := s.client.doRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.DatasourceList{})
}

// DirectQueryRefreshSchedule returns the refresh schedule for a specified DirectQuery or LiveConnection dataset from the specified workspace.
func (s *datasetGroupService) DirectQueryRefreshSchedule(ctx context.Context, groupID, datasetID string) (*types.DirectQueryRefreshSchedule, error) {
	u := fmt.Sprintf("%s/%s/directQueryRefreshSchedule", datasetsPath(groupID), url.PathEscape(datasetID))
	_, resp, err := s.client.doRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.DirectQueryRefreshSchedule{})
}

// GatewayDatasources returns a list of gateway data sources for the specified dataset from the specified workspace.
func (s *datasetGroupService) GatewayDatasources(ctx context.Context, groupID, datasetID string) (*types.GatewayDatasourceList, error) {
	u := fmt.Sprintf("%s/%s/Default.GetBoundGatewayDatasources", datasetsPath(groupID), url.PathEscape(datasetID))
	_, resp, err := s.client.doRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.GatewayDatasourceList{})
}

// Parameters returns a list of parameters for the specified dataset from the specified workspace.
func (s *datasetGroupService) Parameters(ctx context.Context, groupID, datasetID string) (*types.MashupParameterList, error) {
	u := fmt.Sprintf("%s/%s/parameters", datasetsPath(groupID), url.PathEscape(datasetID))
	_, resp, err := s.client.doRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.MashupParameterList{})
}

// RefreshDataset triggers a refresh of the specified dataset from the specified workspace.
// Power BI accepts it asynchronously; follow it with RefreshHistory, or with
// RefreshExecutionDetails for an enhanced refresh. A request whose response is lost
// may still have been accepted, so check RefreshHistory before sending it again.
func (s *datasetGroupService) RefreshDataset(ctx context.Context, groupID, datasetID string, req types.DatasetRefreshRequest) (*types.RefreshAccepted, error) {
	u := fmt.Sprintf("%s/%s/refreshes", datasetsPath(groupID), url.PathEscape(datasetID))
	_, resp, err := s.client.postJSON(ctx, u, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	out := &types.RefreshAccepted{RequestID: resp.Header.Get("x-ms-request-id"), Location: resp.Header.Get("Location")}
	// The Location URL ends in the refresh's ID.
	if out.Location != "" {
		if l, err := url.Parse(out.Location); err == nil {
			out.RequestID = path.Base(l.Path)
		}
	}
	return out, nil
}

// RefreshExecutionDetails returns the state of the specified enhanced refresh of the specified dataset.
func (s *datasetGroupService) RefreshExecutionDetails(ctx context.Context, groupID, datasetID, refreshID string) (*types.RefreshExecutionDetails, error) {
	u := fmt.Sprintf("%s/%s/refreshes/%s", datasetsPath(groupID), url.PathEscape(datasetID), url.PathEscape(refreshID))
	_, resp, err := s.client.doRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.RefreshExecutionDetails{})
}

// RefreshHistory returns the refresh history of the specified dataset from the specified
// workspace, most recent first: the last top entries, or the last 60 when top is 0.
func (s *datasetGroupService) RefreshHistory(ctx context.Context, groupID, datasetID string, top int) (*types.RefreshList, error) {
	u := fmt.Sprintf("%s/%s/refreshes", datasetsPath(groupID), url.PathEscape(datasetID))
	if top > 0 {
		u += "?$top=" + strconv.Itoa(top)
	}
	_, resp, err := s.client.doRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.RefreshList{})
}
