package powerbi

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stpabhi/powerbi-go/types"
)

// PushDatasetsService handles communication with push datasets APIs.
// https://learn.microsoft.com/en-us/rest/api/power-bi/push-datasets
type PushDatasetsService service

// DeleteRows deletes all rows from the specified table within the specified dataset from My workspace.
// https://learn.microsoft.com/en-us/rest/api/power-bi/push-datasets/datasets-delete-rows
func (s *PushDatasetsService) DeleteRows(ctx context.Context, datasetID, tableName string) error {
	u := fmt.Sprintf("%s/%s/%s/%s/%s", datasetsBasePath, url.PathEscape(datasetID), "tables", url.PathEscape(tableName), "rows")

	_, resp, err := s.client.doRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// DeleteRowsInGroup deletes all rows from the specified table within the specified dataset from the specified workspace.
// https://learn.microsoft.com/en-us/rest/api/power-bi/push-datasets/datasets-delete-rows-in-group
func (s *PushDatasetsService) DeleteRowsInGroup(ctx context.Context, groupID, datasetID, tableName string) error {
	u := fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s", groupsBasePath, url.PathEscape(groupID), datasetsBasePath, url.PathEscape(datasetID), "tables", url.PathEscape(tableName), "rows")

	_, resp, err := s.client.doRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// GetTables returns a list of tables within the specified dataset from My workspace.
// https://learn.microsoft.com/en-us/rest/api/power-bi/push-datasets/datasets-get-tables
func (s *PushDatasetsService) GetTables(ctx context.Context, datasetID string) ([]types.Table, error) {
	u := fmt.Sprintf("%s/%s/%s", datasetsBasePath, url.PathEscape(datasetID), "tables")

	_, resp, err := s.client.doRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.TableList
	_, err = toObject(resp, &result)
	if err != nil {
		return nil, err
	}

	return result.Value, nil
}

// GetTablesInGroup returns a list of tables within the specified dataset from the specified workspace.
// https://learn.microsoft.com/en-us/rest/api/power-bi/push-datasets/datasets-get-tables-in-group
func (s *PushDatasetsService) GetTablesInGroup(ctx context.Context, groupID, datasetID string) ([]types.Table, error) {
	u := fmt.Sprintf("%s/%s/%s/%s/%s", groupsBasePath, url.PathEscape(groupID), datasetsBasePath, url.PathEscape(datasetID), "tables")

	_, resp, err := s.client.doRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.TableList
	_, err = toObject(resp, &result)
	if err != nil {
		return nil, err
	}

	return result.Value, nil
}

// PostDataset creates a new dataset on My workspace.
// https://learn.microsoft.com/en-us/rest/api/power-bi/push-datasets/datasets-post-dataset
func (s *PushDatasetsService) PostDataset(ctx context.Context, req types.CreateDatasetRequest, opts types.DatasetOptions) (*types.Dataset, error) {
	u := datasetsBasePath
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, err
	}

	_, resp, err := s.client.postJSON(ctx, u, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.Dataset{})
}

// PostDatasetInGroup creates a new dataset in the specified workspace.
// https://learn.microsoft.com/en-us/rest/api/power-bi/push-datasets/datasets-post-dataset-in-group
func (s *PushDatasetsService) PostDatasetInGroup(ctx context.Context, groupID string, req types.CreateDatasetRequest, opts types.DatasetOptions) (*types.Dataset, error) {
	u := fmt.Sprintf("%s/%s/%s", groupsBasePath, url.PathEscape(groupID), datasetsBasePath)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, err
	}

	_, resp, err := s.client.postJSON(ctx, u, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.Dataset{})
}

// PostRows adds new data rows to the specified table within the specified dataset from My workspace.
// https://learn.microsoft.com/en-us/rest/api/power-bi/push-datasets/datasets-post-rows
func (s *PushDatasetsService) PostRows(ctx context.Context, datasetID, tableName string, req types.PostRowsRequest) error {
	u := fmt.Sprintf("%s/%s/%s/%s/%s", datasetsBasePath, url.PathEscape(datasetID), "tables", url.PathEscape(tableName), "rows")

	_, resp, err := s.client.postJSON(ctx, u, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// PostRowsInGroup adds new data rows to the specified table within the specified dataset from the specified workspace.
// https://learn.microsoft.com/en-us/rest/api/power-bi/push-datasets/datasets-post-rows-in-group
func (s *PushDatasetsService) PostRowsInGroup(ctx context.Context, groupID, datasetID, tableName string, req types.PostRowsRequest) error {
	u := fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s", groupsBasePath, url.PathEscape(groupID), datasetsBasePath, url.PathEscape(datasetID), "tables", url.PathEscape(tableName), "rows")

	_, resp, err := s.client.postJSON(ctx, u, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// PutTable updates the metadata and schema for the specified table within the specified dataset from My workspace.
// https://learn.microsoft.com/en-us/rest/api/power-bi/push-datasets/datasets-put-table
func (s *PushDatasetsService) PutTable(ctx context.Context, datasetID, tableName string, req types.Table) (*types.Table, error) {
	u := fmt.Sprintf("%s/%s/%s/%s", datasetsBasePath, url.PathEscape(datasetID), "tables", url.PathEscape(tableName))

	_, resp, err := s.client.putJSON(ctx, u, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.Table{})
}

// PutTableInGroup updates the metadata and schema for the specified table within the specified dataset from the specified workspace.
// https://learn.microsoft.com/en-us/rest/api/power-bi/push-datasets/datasets-put-table-in-group
func (s *PushDatasetsService) PutTableInGroup(ctx context.Context, groupID, datasetID, tableName string, req types.Table) (*types.Table, error) {
	u := fmt.Sprintf("%s/%s/%s/%s/%s/%s", groupsBasePath, url.PathEscape(groupID), datasetsBasePath, url.PathEscape(datasetID), "tables", url.PathEscape(tableName))

	_, resp, err := s.client.putJSON(ctx, u, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.Table{})
}
