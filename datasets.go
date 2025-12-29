package powerbi

import (
	"context"
	"fmt"
	"net/url"

	"github.com/stpabhi/powerbi-go/types"
)

const datasetsBasePath = "datasets"

type DatasetsService service

var _ DatasetsInterface = &DatasetsService{}

type DatasetsInterface interface {
	Group() DatasetGroup
}

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
}

type datasetGroupService service

func (s *DatasetsService) Group() DatasetGroup {
	return &datasetGroupService{s.client}
}

// BindToGateway Binds the specified dataset from the specified workspace to the specified gateway, optionally with a given set of data source IDs.
// If you don't supply a specific data source ID, the dataset will be bound to the first matching data source in the gateway.
func (s *datasetGroupService) BindToGateway(ctx context.Context, groupID, datasetID string, req types.BindToGatewayRequest) error {
	u := fmt.Sprintf("%s/%s/%s/%s/Default.BindToGateway", groupsBasePath, url.PathEscape(groupID), datasetsBasePath, url.PathEscape(datasetID))
	_, resp, err := s.client.postJSON(ctx, u, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// CancelRefresh cancels the specified refresh operation for the specified dataset from the specified workspace.
func (s *datasetGroupService) CancelRefresh(ctx context.Context, groupID, datasetID, refreshID string) error {
	u := fmt.Sprintf("%s/%s/%s/%s/refreshes/%s", groupsBasePath, url.PathEscape(groupID), datasetsBasePath, url.PathEscape(datasetID), url.PathEscape(refreshID))
	_, resp, err := s.client.doRequest(ctx, "DELETE", u, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// DeleteDataset deletes the specified dataset from the specified workspace.
func (s *datasetGroupService) DeleteDataset(ctx context.Context, groupID, datasetID string) error {
	u := fmt.Sprintf("%s/%s/%s/%s", groupsBasePath, url.PathEscape(groupID), datasetsBasePath, url.PathEscape(datasetID))
	_, resp, err := s.client.doRequest(ctx, "DELETE", u, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// DiscoverGateways returns a list of gateways that the specified dataset from the specified workspace can be bound to.
func (s *datasetGroupService) DiscoverGateways(ctx context.Context, groupID, datasetID string) (*types.GatewayList, error) {
	u := fmt.Sprintf("%s/%s/%s/%s/Default.DiscoverGateways", groupsBasePath, url.PathEscape(groupID), datasetsBasePath, url.PathEscape(datasetID))
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
	u := fmt.Sprintf("%s/%s/%s/%s/executeQueries", groupsBasePath, url.PathEscape(groupID), datasetsBasePath, url.PathEscape(datasetID))
	_, resp, err := s.client.postJSON(ctx, u, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.DatasetExecuteQueriesResponse{})
}

// Dataset returns the specified dataset from the specified workspace.
func (s *datasetGroupService) Dataset(ctx context.Context, groupID, datasetID string) (*types.Dataset, error) {
	u := fmt.Sprintf("%s/%s/%s/%s", groupsBasePath, url.PathEscape(groupID), datasetsBasePath, url.PathEscape(datasetID))
	_, resp, err := s.client.doRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.Dataset{})
}

// DatasetToDataflowLinks returns a list of upstream dataflows for datasets from the specified workspace.
func (s *datasetGroupService) DatasetToDataflowLinks(ctx context.Context, groupID string) (*types.DatasetToDataflowLinksResponse, error) {
	u := fmt.Sprintf("%s/%s/%s/upstreamDataflows", groupsBasePath, url.PathEscape(groupID), datasetsBasePath)
	_, resp, err := s.client.doRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.DatasetToDataflowLinksResponse{})
}

// DatasetUsers returns a list of principals that have access to the specified dataset.
func (s *datasetGroupService) DatasetUsers(ctx context.Context, groupID, datasetID string) (*types.DatasetUsersAccess, error) {
	u := fmt.Sprintf("%s/%s/%s/%s/users", groupsBasePath, url.PathEscape(groupID), datasetsBasePath, url.PathEscape(datasetID))
	_, resp, err := s.client.doRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.DatasetUsersAccess{})
}

// Datasets returns a list of datasets from the specified workspace.
func (s *datasetGroupService) Datasets(ctx context.Context, groupID string) (*types.DatasetList, error) {
	u := fmt.Sprintf("%s/%s/%s", groupsBasePath, url.PathEscape(groupID), datasetsBasePath)
	_, resp, err := s.client.doRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.DatasetList{})
}

// Datasources returns a list of data sources for the specified dataset from the specified workspace.
func (s *datasetGroupService) Datasources(ctx context.Context, groupID, datasetID string) (*types.DatasourceList, error) {
	u := fmt.Sprintf("%s/%s/%s/%s/datasources", groupsBasePath, url.PathEscape(groupID), datasetsBasePath, url.PathEscape(datasetID))
	_, resp, err := s.client.doRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.DatasourceList{})
}

// DirectQueryRefreshSchedule returns the refresh schedule for a specified DirectQuery or LiveConnection dataset from the specified workspace.
func (s *datasetGroupService) DirectQueryRefreshSchedule(ctx context.Context, groupID, datasetID string) (*types.DirectQueryRefreshSchedule, error) {
	u := fmt.Sprintf("%s/%s/%s/%s/directQueryRefreshSchedule", groupsBasePath, url.PathEscape(groupID), datasetsBasePath, url.PathEscape(datasetID))
	_, resp, err := s.client.doRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.DirectQueryRefreshSchedule{})
}

// GatewayDatasources returns a list of gateway data sources for the specified dataset from the specified workspace.
func (s *datasetGroupService) GatewayDatasources(ctx context.Context, groupID, datasetID string) (*types.GatewayDatasourceList, error) {
	u := fmt.Sprintf("%s/%s/%s/%s/Default.GetBoundGatewayDatasources", groupsBasePath, url.PathEscape(groupID), datasetsBasePath, url.PathEscape(datasetID))
	_, resp, err := s.client.doRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.GatewayDatasourceList{})
}

// Parameters returns a list of parameters for the specified dataset from the specified workspace.
func (s *datasetGroupService) Parameters(ctx context.Context, groupID, datasetID string) (*types.MashupParameterList, error) {
	u := fmt.Sprintf("%s/%s/%s/%s/parameters", groupsBasePath, url.PathEscape(groupID), datasetsBasePath, url.PathEscape(datasetID))
	_, resp, err := s.client.doRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.MashupParameterList{})
}
