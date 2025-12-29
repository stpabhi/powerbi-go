package powerbi

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stpabhi/powerbi-go/types"
)

const groupsBasePath = "groups"

// GroupsService handles communication with the groups related methods of the Power BI API.
// https://learn.microsoft.com/en-us/rest/api/power-bi/groups
type GroupsService service

// Create creates a new workspace.
func (s *GroupsService) Create(ctx context.Context, req types.CreateGroupRequest, opts types.CreateGroupOptions) (*types.Group, error) {
	u := groupsBasePath
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, err
	}

	_, resp, err := s.client.postJSON(ctx, u, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.Group{})
}

// Delete deletes a workspace by ID.
func (s *GroupsService) Delete(ctx context.Context, groupID string) error {
	u := fmt.Sprintf("%s/%s", groupsBasePath, url.PathEscape(groupID))
	_, resp, err := s.client.doRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// Get returns a group (workspace) by ID.
func (s *GroupsService) Get(ctx context.Context, groupID string) (*types.Group, error) {
	u := fmt.Sprintf("%s/%s", groupsBasePath, url.PathEscape(groupID))
	_, resp, err := s.client.doRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.Group{})
}

// List lists workspaces with optional $top, $skip and $filter parameters.
func (s *GroupsService) List(ctx context.Context, opts types.ListGroupsOptions) ([]types.Group, error) {
	u := groupsBasePath
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, err
	}

	_, resp, err := s.client.doRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.GroupList
	_, err = toObject(resp, &result)

	return result.Value, err
}

// Update updates mutable fields of a workspace.
func (s *GroupsService) Update(ctx context.Context, groupID string, req types.UpdateGroupRequest) (*types.Group, error) {
	u := fmt.Sprintf("%s/%s", groupsBasePath, url.PathEscape(groupID))
	_, resp, err := s.client.patchJSON(ctx, u, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.Group{})
}

// AddGroupUser grants access to a workspace.
func (s *GroupsService) AddGroupUser(ctx context.Context, groupID string, user types.GroupUser) error {
	u := fmt.Sprintf("%s/%s/users", groupsBasePath, url.PathEscape(groupID))
	_, resp, err := s.client.postJSON(ctx, u, user)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// DeleteGroupUser removes a user's access from a workspace.
// The user parameter may be an email address or object ID.
func (s *GroupsService) DeleteGroupUser(ctx context.Context, groupID, user string, opts types.DeleteGroupUserOptions) error {
	u := fmt.Sprintf("%s/%s/users/%s", groupsBasePath, url.PathEscape(groupID), url.PathEscape(user))
	u, err := addOptions(u, opts)
	if err != nil {
		return err
	}

	_, resp, err := s.client.doRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// ListGroupUsers lists the users of a workspace.
func (s *GroupsService) ListGroupUsers(ctx context.Context, groupID string, opts types.ListGroupUserOptions) ([]types.GroupUser, error) {
	u := fmt.Sprintf("%s/%s/users", groupsBasePath, url.PathEscape(groupID))
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, err
	}

	_, resp, err := s.client.doRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.GroupUserList
	_, err = toObject(resp, &result)

	return result.Value, err
}

// UpdateGroupUser updates a user's access in a workspace.
func (s *GroupsService) UpdateGroupUser(ctx context.Context, groupID string, user types.GroupUser) error {
	u := fmt.Sprintf("%s/%s/users", groupsBasePath, url.PathEscape(groupID))
	_, resp, err := s.client.putJSON(ctx, u, user)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// AddDashboard creates a new empty dashboard in the specified workspace.
func (s *GroupsService) AddDashboard(ctx context.Context, groupID string, req types.AddDashboardRequest) (*types.Dashboard, error) {
	u := fmt.Sprintf("%s/%s/dashboards", groupsBasePath, url.PathEscape(groupID))
	_, resp, err := s.client.postJSON(ctx, u, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.Dashboard{})
}

// DeleteDashboard deletes the specified dashboard from the specified workspace.
func (s *GroupsService) DeleteDashboard(ctx context.Context, groupID, dashboardID string) error {
	u := fmt.Sprintf("%s/%s/dashboards/%s", groupsBasePath, url.PathEscape(groupID), url.PathEscape(dashboardID))
	_, resp, err := s.client.doRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// GetDashboard returns the specified dashboard from the specified workspace.
func (s *GroupsService) GetDashboard(ctx context.Context, groupID, dashboardID string) (*types.Dashboard, error) {
	u := fmt.Sprintf("%s/%s/dashboards/%s", groupsBasePath, url.PathEscape(groupID), url.PathEscape(dashboardID))
	_, resp, err := s.client.doRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.Dashboard{})
}

// ListDashboards returns a list of dashboards from the specified workspace.
func (s *GroupsService) ListDashboards(ctx context.Context, groupID string) ([]types.Dashboard, error) {
	u := fmt.Sprintf("%s/%s/dashboards", groupsBasePath, url.PathEscape(groupID))
	_, resp, err := s.client.doRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.DashboardList
	_, err = toObject(resp, &result)
	if err != nil {
		return nil, err
	}

	return result.Value, nil
}

// CloneTile the specified tile from the specified workspace.
func (s *GroupsService) CloneTile(ctx context.Context, groupID, dashboardID, tileID string, req types.CloneTileRequest) (*types.Tile, error) {
	u := fmt.Sprintf("%s/%s/dashboards/%s/tiles/%s/Clone", groupsBasePath, url.PathEscape(groupID), url.PathEscape(dashboardID), url.PathEscape(tileID))
	_, resp, err := s.client.postJSON(ctx, u, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.Tile{})
}

// GetTile returns the specified tile within the specified dashboard from the specified workspace.
func (s *GroupsService) GetTile(ctx context.Context, groupID, dashboardID, tileID string) (*types.Tile, error) {
	u := fmt.Sprintf("%s/%s/dashboards/%s/tiles/%s", groupsBasePath, url.PathEscape(groupID), url.PathEscape(dashboardID), url.PathEscape(tileID))
	_, resp, err := s.client.doRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.Tile{})
}

// ListTiles returns a list of tiles within the specified dashboard from the specified workspace.
func (s *GroupsService) ListTiles(ctx context.Context, groupID, dashboardID string) ([]types.Tile, error) {
	u := fmt.Sprintf("%s/%s/dashboards/%s/tiles", groupsBasePath, url.PathEscape(groupID), url.PathEscape(dashboardID))
	_, resp, err := s.client.doRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.TileList
	_, err = toObject(resp, &result)
	if err != nil {
		return nil, err
	}

	return result.Value, nil
}
