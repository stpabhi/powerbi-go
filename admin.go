package powerbi

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stpabhi/powerbi-go/types"
)

const adminBasePath = "admin"

type AdminService service

var _ Admin = &AdminService{}

type Admin interface {
	Groups() Groups
}

type Groups interface {
	AddUserAsAdmin(ctx context.Context, groupID string, req types.GroupUser) error
	DeleteUserAsAdmin(ctx context.Context, groupID string, userID string, opts types.DeleteUserOptions) error
	GetGroupAsAdmin(ctx context.Context, groupID string, opts types.GroupOptions) (*types.AdminGroup, error)
	GetGroupUsersAsAdmin(ctx context.Context, groupID string) ([]types.GroupUser, error)
	GetGroupsAsAdmin(ctx context.Context, opts types.GroupsOptions) ([]types.AdminGroup, error)
	GetUnusedArtifactsAsAdmin(ctx context.Context, groupID string, opts types.UnusedArtifactsOptions) (*types.UnusedArtifactsResponse, error)
	RestoreDeletedGroupAsAdmin(ctx context.Context, groupID string, req types.GroupRestoreRequest) error
	UpdateGroupAsAdmin(ctx context.Context, groupID string, req types.AdminGroup) error
}
type groupService service

func (s *AdminService) Groups() Groups {
	return &groupService{s.client}
}

// AddUserAsAdmin grants user permissions to the specified workspace.
// This API call only supports adding a user, security group, M365 group and service principal.
func (s *groupService) AddUserAsAdmin(ctx context.Context, groupID string, req types.GroupUser) error {
	u := fmt.Sprintf("%s/%s/%s/users", adminBasePath, groupsBasePath, url.PathEscape(groupID))
	_, resp, err := s.client.postJSON(ctx, u, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// DeleteUserAsAdmin removes user permissions from the specified workspace.
// This API call supports removing a user, security group, M365 group and service principal.
// Please use email address or UPN for user, group object Id for group and app object Id for service principal to delete.
func (s *groupService) DeleteUserAsAdmin(ctx context.Context, groupID string, userID string, opts types.DeleteUserOptions) error {
	u := fmt.Sprintf("%s/%s/%s/%s/users/%s", adminBasePath, groupsBasePath, url.PathEscape(groupID), "users", url.PathEscape(userID))
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

// GetGroupAsAdmin returns a workspace for the organization.
func (s *groupService) GetGroupAsAdmin(ctx context.Context, groupID string, opts types.GroupOptions) (*types.AdminGroup, error) {
	u := fmt.Sprintf("%s/%s/%s", adminBasePath, groupsBasePath, url.PathEscape(groupID))
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, err
	}

	_, resp, err := s.client.doRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.AdminGroup{})
}

// GetGroupUsersAsAdmin returns a list of users that have access to the specified workspace.
func (s *groupService) GetGroupUsersAsAdmin(ctx context.Context, groupID string) ([]types.GroupUser, error) {
	u := fmt.Sprintf("%s/%s/%s/%s", adminBasePath, groupsBasePath, url.PathEscape(groupID), "users")
	_, resp, err := s.client.doRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.GroupUserList
	_, err = toObject(resp, &result)
	if err != nil {
		return nil, err
	}
	return result.Value, nil
}

// GetGroupsAsAdmin returns a list of workspaces for the organization.
func (s *groupService) GetGroupsAsAdmin(ctx context.Context, opts types.GroupsOptions) ([]types.AdminGroup, error) {
	u := fmt.Sprintf("%s/%s", adminBasePath, groupsBasePath)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, err
	}

	_, resp, err := s.client.doRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.AdminGroupList
	_, err = toObject(resp, &result)
	if err != nil {
		return nil, err
	}
	
	return result.Value, nil
}

// GetUnusedArtifactsAsAdmin returns a list of datasets, reports, and dashboards
// that have not been used within 30 days for the specified workspace. This is a preview API call.
func (s *groupService) GetUnusedArtifactsAsAdmin(ctx context.Context, groupID string, opts types.UnusedArtifactsOptions) (*types.UnusedArtifactsResponse, error) {
	u := fmt.Sprintf("%s/%s/%s/%s", adminBasePath, groupsBasePath, url.PathEscape(groupID), "unused")
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, err
	}

	_, resp, err := s.client.doRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.UnusedArtifactsResponse{})
}

// RestoreDeletedGroupAsAdmin restores a deleted workspace.
func (s *groupService) RestoreDeletedGroupAsAdmin(ctx context.Context, groupID string, req types.GroupRestoreRequest) error {
	u := fmt.Sprintf("%s/%s/%s/%s", adminBasePath, groupsBasePath, url.PathEscape(groupID), "restore")
	_, resp, err := s.client.postJSON(ctx, u, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// UpdateGroupAsAdmin updates the properties of the specified workspace.
func (s *groupService) UpdateGroupAsAdmin(ctx context.Context, groupID string, req types.AdminGroup) error {
	u := fmt.Sprintf("%s/%s/%s", adminBasePath, groupsBasePath, url.PathEscape(groupID))
	_, resp, err := s.client.patchJSON(ctx, u, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
