package types

type GroupUserAccessRight string

const (
	GroupUserAccessRightNone        GroupUserAccessRight = "None"
	GroupUserAccessRightAdmin       GroupUserAccessRight = "Admin"
	GroupUserAccessRightMember      GroupUserAccessRight = "Member"
	GroupUserAccessRightContributor GroupUserAccessRight = "Contributor"
	GroupUserAccessRightViewer      GroupUserAccessRight = "Viewer"
)

type PrincipalType string

const (
	PrincipalTypeNone  PrincipalType = "None"
	PrincipalTypeUser  PrincipalType = "User"
	PrincipalTypeGroup PrincipalType = "Group"
	PrincipalTypeApp   PrincipalType = "App"
)

// ServicePrincipalProfile is a Power BI service principal profile.
// Only relevant for Power BI Embedded multi-tenancy solution.
type ServicePrincipalProfile struct {
	DisplayName string `json:"displayName,omitempty"`
	ID          string `json:"id,omitempty"`
}

// GroupUser represents a user or principal with access to a workspace (group).
// For POST/PUT requests, the typical required fields are Identifier,
// GroupUserAccessRight, and PrincipalType.
type GroupUser struct {
	GroupUserAccessRight GroupUserAccessRight     `json:"groupUserAccessRight"`
	Identifier           string                   `json:"identifier"`
	PrincipalType        PrincipalType            `json:"principalType"`
	DisplayName          *string                  `json:"displayName,omitempty"`
	EmailAddress         *string                  `json:"emailAddress,omitempty"`
	GraphID              *string                  `json:"graphId,omitempty"`
	Profile              *ServicePrincipalProfile `json:"profile,omitempty"`
	UserType             *string                  `json:"userType,omitempty"`
}

type DeleteGroupUserOptions struct {
	ProfileID string `url:"profileId,omitempty"`
}

type ListGroupUserOptions struct {
	Skip int `url:"$skip,omitempty"`
	Top  int `url:"$top,omitempty"`
}

type GroupUserList struct {
	Value []GroupUser `json:"value"`
}
