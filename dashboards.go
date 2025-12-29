package powerbi

import (
	"context"
	"fmt"
	"net/url"

	"github.com/stpabhi/powerbi-go/types"
)

const dashboardsBasePath = "dashboards"

// DashboardsService handles communication with the dashboards related methods of the Power BI API.
// https://learn.microsoft.com/en-us/rest/api/power-bi/dashboards
type DashboardsService service

// Add adds a dashboard.
func (s *DashboardsService) Add(ctx context.Context, req types.AddDashboardRequest) (*types.Dashboard, error) {
	u := dashboardsBasePath
	_, resp, err := s.client.postJSON(ctx, u, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.Dashboard{})
}

// CloneTile clones the specified tile from My workspace.
func (s *DashboardsService) CloneTile(ctx context.Context, dashboardID string, tileID string, req types.CloneTileRequest) (*types.Tile, error) {
	u := fmt.Sprintf("%s/%s/tiles/%s/Clone", dashboardsBasePath, dashboardID, tileID)
	_, resp, err := s.client.postJSON(ctx, u, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.Tile{})
}

// Delete deletes the specified dashboard from My workspace.
func (s *DashboardsService) Delete(ctx context.Context, dashboardID string) error {
	u := fmt.Sprintf("%s/dashboards/%s", dashboardsBasePath, url.PathEscape(dashboardID))
	_, resp, err := s.client.doRequest(ctx, "DELETE", u, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// Get gets the specified dashboard from My workspace.
func (s *DashboardsService) Get(ctx context.Context, dashboardID string) (*types.Dashboard, error) {
	u := fmt.Sprintf("%s/dashboards/%s", dashboardsBasePath, url.PathEscape(dashboardID))
	_, resp, err := s.client.doRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.Dashboard{})
}

// List returns a list of dashboards from My workspace.
func (s *DashboardsService) List(ctx context.Context) ([]types.Dashboard, error) {
	u := dashboardsBasePath
	_, resp, err := s.client.doRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []types.Dashboard
	_, err = toObject(resp, &result)
	return result, err
}

// GetTile returns the specified tile within the specified dashboard from My workspace.
func (s *DashboardsService) GetTile(ctx context.Context, dashboardID string, tileID string) (*types.Tile, error) {
	u := fmt.Sprintf("%s/dashboards/%s/tiles/%s", dashboardsBasePath, url.PathEscape(dashboardID), url.PathEscape(tileID))
	_, resp, err := s.client.doRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.Tile{})
}

// ListTiles returns a list of tiles within the specified dashboard from My workspace.
func (s *DashboardsService) ListTiles(ctx context.Context, dashboardID string) ([]types.Tile, error) {
	u := fmt.Sprintf("%s/dashboards/%s/tiles", dashboardsBasePath, url.PathEscape(dashboardID))
	_, resp, err := s.client.doRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []types.Tile
	_, err = toObject(resp, &result)
	return result, err
}
