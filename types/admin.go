package types

type DeleteUserOptions struct {
	IsGroup   bool   `url:"isGroup,omitempty"`
	ProfileID string `url:"profileId,omitempty"`
}

type GroupOptions struct {
	Expand string `url:"$expand,omitempty"`
}

type GroupsOptions struct {
	GroupOptions `url:",inline"`
	Skip         int    `url:"$skip,omitempty"`
	Top          int    `url:"$top,omitempty"`
	Filter       string `url:"$filter,omitempty"` // TODO filter doesn't work; support OData filter syntax
}

type GroupType string

const (
	GroupTypeAdminWorkspace GroupType = "AdminWorkspace"
	GroupTypePersonalGroup  GroupType = "PersonalGroup"
	GroupTypePersonal       GroupType = "Personal"
	GroupTypeGroup          GroupType = "Group"
	GroupTypeWorkspace      GroupType = "Workspace"
)

type AdminGroup struct {
	Group                     `json:",inline"`
	Dashboards                []AdminDashboard `json:"dashboards,omitempty"`
	Dataflows                 []AdminDataflow  `json:"dataflows,omitempty"`
	Datasets                  []AdminDataset   `json:"datasets,omitempty"`
	Description               string           `json:"description,omitempty"`
	HasWorkspaceLevelSettings bool             `json:"hasWorkspaceLevelSettings,omitempty"`
	PipelineID                string           `json:"pipelineId,omitempty"`
	Reports                   []AdminReport    `json:"reports,omitempty"`
	State                     string           `json:"state,omitempty"`
	Type                      GroupType        `json:"type,omitempty"`
	Users                     []GroupUser      `json:"users,omitempty"`
	Workbooks                 []Workbook       `json:"workbooks,omitempty"`
}

type AdminDashboard struct {
	AppID         string          `json:"appId,omitempty"`
	DisplayName   string          `json:"displayName,omitempty"`
	EmbedURL      string          `json:"embedUrl,omitempty"`
	ID            string          `json:"id,omitempty"`
	IsReadOnly    bool            `json:"isReadOnly,omitempty"`
	Subscriptions []Subscription  `json:"subscriptions,omitempty"`
	Tiles         []AdminTile     `json:"tiles,omitempty"`
	Users         []DashboardUser `json:"users,omitempty"`
	WebURL        string          `json:"webUrl,omitempty"`
	WorkspaceID   string          `json:"workspaceId,omitempty"`
}

type AdminTile struct {
	ColSpan   int    `json:"colSpan,omitempty"`
	DatasetID string `json:"datasetId,omitempty"`
	EmbedData string `json:"embedData,omitempty"`
	EmbedURL  string `json:"embedUrl,omitempty"`
	ID        string `json:"id,omitempty"`
	ReportID  string `json:"reportId,omitempty"`
	RowSpan   int    `json:"rowSpan,omitempty"`
	Title     string `json:"title,omitempty"`
}

type AdminDataflow struct {
	ConfiguredBy string         `json:"configuredBy,omitempty"`
	Description  string         `json:"description,omitempty"`
	ModelURL     string         `json:"modelUrl,omitempty"`
	Name         string         `json:"name,omitempty"`
	ObjectID     string         `json:"objectId,omitempty"`
	Users        []DataflowUser `json:"users,omitempty"`
	WorkspaceID  string         `json:"workspaceId,omitempty"`
}

type AdminDataset struct {
	ContentProviderType              string              `json:"contentProviderType,omitempty"`
	Encryption                       Encryption          `json:"encryption,omitempty"`
	IsEffectiveIdentityRequired      bool                `json:"isEffectiveIdentityRequired,omitempty"`
	IsEffectiveIdentityRolesRequired bool                `json:"isEffectiveIdentityRolesRequired,omitempty"`
	IsInPlaceSharingEnabled          bool                `json:"isInPlaceSharingEnabled,omitempty"`
	IsOnPremGatewayRequired          bool                `json:"isOnPremGatewayRequired,omitempty"`
	IsRefreshable                    bool                `json:"isRefreshable,omitempty"`
	AddRowsAPIEnabled                bool                `json:"addRowsAPIEnabled,omitempty"`
	ConfiguredBy                     string              `json:"configuredBy,omitempty"`
	CreateReportEmbedURL             string              `json:"createReportEmbedUrl,omitempty"`
	CreatedDate                      string              `json:"createdDate,omitempty"`
	Description                      string              `json:"description,omitempty"`
	ID                               string              `json:"id,omitempty"`
	Name                             string              `json:"name,omitempty"`
	QNAEmbedURL                      string              `json:"qnaEmbedUrl,omitempty"`
	QueryScaleOutSettings            string              `json:"queryScaleOutSettings,omitempty"`
	TargetStorageMode                string              `json:"targetStorageMode,omitempty"`
	UpstreamDataflows                []DependentDataflow `json:"upstreamDataflows,omitempty"`
	Users                            []DatasetUser       `json:"users,omitempty"`
	WebURL                           string              `json:"webUrl,omitempty"`
	WorkspaceID                      string              `json:"workspaceId,omitempty"`
}

type AdminReport struct {
	Report           `json:",inline"`
	CreatedBy        string `json:"createdBy,omitempty"`
	CreatedDateTime  string `json:"createdDateTime,omitempty"`
	ModifiedBy       string `json:"modifiedBy,omitempty"`
	ModifiedDateTime string `json:"modifiedDateTime,omitempty"`
	WorkspaceID      string `json:"workspaceId,omitempty"`
}
type EncryptionStatus string

const (
	EncryptionStatusUnknown                EncryptionStatus = "Unknown"
	EncryptionStatusNotSupported           EncryptionStatus = "EncryptionStatus"
	EncryptionStatusInSyncWithWorkspace    EncryptionStatus = "InSyncWithWorkspace"
	EncryptionStatusNotInSyncWithWorkspace EncryptionStatus = "NotInSyncWithWorkspace"
)

type Encryption struct {
	EncryptionStatus EncryptionStatus `json:"encryptionStatus,omitempty"`
}

type DependentDataflow struct {
	GroupID          string `json:"groupId,omitempty"`
	TargetDataflowID string `json:"targetDataflowId,omitempty"`
}

type Workbook struct {
	DatasetID string `json:"datasetId,omitempty"`
	Name      string `json:"name,omitempty"`
}

type UnusedArtifactsResponse struct {
	ContinuationToken      string                 `json:"continuationToken,omitempty"`
	ContinuationURI        string                 `json:"continuationUri,omitempty"`
	UnusedArtifactEntities []UnusedArtifactEntity `json:"unusedArtifactEntities,omitempty"`
}

type UnusedArtifactEntity struct {
	ArtifactID           string `json:"artifactId,omitempty"`
	ArtifactSizeInMB     int    `json:"artifactSizeInMB,omitempty"`
	ArtifactType         string `json:"artifactType,omitempty"`
	CreatedDateTime      string `json:"createdDateTime,omitempty"`
	DisplayName          string `json:"displayName,omitempty"`
	LastAccessedDateTime string `json:"lastAccessedDateTime,omitempty"`
}

type UnusedArtifactsOptions struct {
	ContinuationToken string `url:"continuationToken,omitempty"`
}

type GroupRestoreRequest struct {
	EmailAddress string `json:"emailAddress"`
	Name         string `json:"name"`
}

type AdminGroupList struct {
	Value []AdminGroup `json:"value"`
}

func (g AdminGroup) String() string {
	return Stringify(g)
}
