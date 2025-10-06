package powerbi

import (
	"context"
	"fmt"
	"net/url"

	"github.com/stpabhi/powerbi-go/types"
)

// EmbedTokenService handles embed token related API methods.
// https://learn.microsoft.com/en-us/rest/api/power-bi/embed-token
type EmbedTokenService service

// GenerateToken generates an embed token for multiple reports, datasets, and target workspaces.
// POST /GenerateToken
func (s *EmbedTokenService) GenerateToken(ctx context.Context, req types.GenerateTokenRequestV2) (*types.EmbedToken, error) {
	u := "GenerateToken"
	_, resp, err := s.client.postJSON(ctx, u, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.EmbedToken{})
}

// GenerateTokenForDashboardsInGroup generates an embed token for a dashboard in a workspace.
// POST /groups/{groupId}/dashboards/{dashboardId}/GenerateToken
func (s *EmbedTokenService) GenerateTokenForDashboardsInGroup(ctx context.Context, groupID, dashboardID string, body types.GenerateTokenRequest) (*types.EmbedToken, error) {
	u := fmt.Sprintf("%s/%s/%s/%s/GenerateToken", groupsBasePath, url.PathEscape(groupID), dashboardsBasePath, url.PathEscape(dashboardID))
	_, resp, err := s.client.postJSON(ctx, u, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.EmbedToken{})
}

// GenerateTokenForDatasetsInGroup generates an embed token based on the specified dataset from the specified workspace.
// POST /groups/{groupId}/datasets/{datasetId}/GenerateToken
func (s *EmbedTokenService) GenerateTokenForDatasetsInGroup(ctx context.Context, groupID, datasetID string, body types.GenerateTokenRequest) (*types.EmbedToken, error) {
	u := fmt.Sprintf("%s/%s/%s/%s/GenerateToken", groupsBasePath, url.PathEscape(groupID), datasetsBasePath, url.PathEscape(datasetID))
	_, resp, err := s.client.postJSON(ctx, u, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.EmbedToken{})
}

// GenerateTokenForReportsCreateInGroup generates an embed token to allow report creation in the specified workspace based on the specified dataset.
// POST /groups/{groupId}/reports/GenerateToken
func (s *EmbedTokenService) GenerateTokenForReportsCreateInGroup(ctx context.Context, groupID string, body types.GenerateTokenRequest) (*types.EmbedToken, error) {
	u := fmt.Sprintf("%s/%s/%s/GenerateToken", groupsBasePath, url.PathEscape(groupID), reportsBasePath)
	_, resp, err := s.client.postJSON(ctx, u, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.EmbedToken{})
}

// GenerateTokenForReportsInGroup generates an embed token to view or edit the specified report from the specified workspace.
// POST /groups/{groupId}/reports/{reportId}/GenerateToken
func (s *EmbedTokenService) GenerateTokenForReportsInGroup(ctx context.Context, groupID, reportID string, req types.GenerateTokenRequest) (*types.EmbedToken, error) {
	u := fmt.Sprintf("%s/%s/%s/%s/GenerateToken", groupsBasePath, url.PathEscape(groupID), reportsBasePath, url.PathEscape(reportID))
	_, resp, err := s.client.postJSON(ctx, u, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.EmbedToken{})
}

// GenerateTokenForTilesInGroup generates an embed token to view the specified tile from the specified workspace.
// POST /groups/{groupId}/dashboards/{dashboardId}/tiles/{tileId}/GenerateToken
func (s *EmbedTokenService) GenerateTokenForTilesInGroup(ctx context.Context, groupID, dashboardID, tileID string, body types.GenerateTokenRequest) (*types.EmbedToken, error) {
	u := fmt.Sprintf("%s/%s/%s/%s/tiles/%s/GenerateToken", groupsBasePath, url.PathEscape(groupID), dashboardsBasePath, url.PathEscape(dashboardID), url.PathEscape(tileID))
	_, resp, err := s.client.postJSON(ctx, u, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.EmbedToken{})
}
