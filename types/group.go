package types

// CreateGroupRequest is the payload to create a workspace.
type CreateGroupRequest struct {
	Name        string `json:"name"`
	workspaceV2 bool   `url:"workspaceV2,default=true"`
}

// UpdateGroupRequest is the payload to update a workspace's mutable fields.
type UpdateGroupRequest struct {
	Name                        string                      `json:"name"`
	DefaultDatasetStorageFormat DefaultDatasetStorageFormat `json:"defaultDatasetStorageFormat,omitempty"`
}

type DefaultDatasetStorageFormat string

const (
	DefaultDatasetStorageFormatSmall DefaultDatasetStorageFormat = "Small"
	DefaultDatasetStorageFormatLarge DefaultDatasetStorageFormat = "Large"
)

type AzureResource struct {
	ID             string `json:"id"`
	ResourceGroup  string `json:"resourceGroup"`
	ResourceName   string `json:"resourceName"`
	SubscriptionID string `json:"subscriptionId"`
}

// Group represents a Power BI workspace (group).
// Fields are expanded to cover commonly returned properties. Unknown fields
// will be ignored by json unmarshalling if not present in responses.
type Group struct {
	CapacityID                  *string                     `json:"capacityId,omitempty"`
	DataflowStorageID           *string                     `json:"dataflowStorageId,omitempty"`
	DefaultDatasetStorageFormat DefaultDatasetStorageFormat `json:"defaultDatasetStorageFormat,omitempty"`
	ID                          string                      `json:"id"`
	IsOnDedicatedCapacity       bool                        `json:"isOnDedicatedCapacity,omitempty"`
	IsReadOnly                  bool                        `json:"isReadOnly,omitempty"`
	LogAnalyticsWorkspace       *AzureResource              `json:"logAnalyticsWorkspace,omitempty"`
	Name                        string                      `json:"name"`
}

type WorkspaceV2 bool

const (
	WorkspaceV2Enabled WorkspaceV2 = true
)

// CreateGroupOptions controls the query for creating a group.
type CreateGroupOptions struct {
	WorkspaceV2 WorkspaceV2 `url:"workspaceV2,omitempty"`
}

// ListGroupsOptions controls the query for listing groups.
type ListGroupsOptions struct {
	Filter string `url:"$filter,omitempty"` // TODO filter doesn't work; support OData filter syntax
	Skip   int    `url:"$skip,omitempty"`
	Top    int    `url:"$top,omitempty"`
}

// GroupList is a paged list of groups.
type GroupList struct {
	Value []Group `json:"value"`
}

func (g Group) String() string {
	return Stringify(g)
}
