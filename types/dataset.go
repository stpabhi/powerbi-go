package types

type Dataset struct {
	ContentProviderType              string                        `json:"contentProviderType,omitempty"`
	Encryption                       *Encryption                   `json:"encryption,omitempty"`
	IsEffectiveIdentityRequired      bool                          `json:"isEffectiveIdentityRequired,omitempty"`
	IsEffectiveIdentityRolesRequired bool                          `json:"isEffectiveIdentityRolesRequired,omitempty"`
	IsInPlaceSharingEnabled          bool                          `json:"isInPlaceSharingEnabled,omitempty"`
	IsOnPremGatewayRequired          bool                          `json:"isOnPremGatewayRequired,omitempty"`
	IsRefreshable                    bool                          `json:"isRefreshable,omitempty"`
	AddRowsAPIEnabled                bool                          `json:"addRowsAPIEnabled,omitempty"`
	ConfiguredBy                     string                        `json:"configuredBy,omitempty"`
	CreateReportEmbedURL             string                        `json:"createReportEmbedUrl,omitempty"`
	CreatedDate                      string                        `json:"createdDate,omitempty"`
	Description                      string                        `json:"description,omitempty"`
	ID                               string                        `json:"id,omitempty"`
	Name                             string                        `json:"name,omitempty"`
	QNAEmbedURL                      string                        `json:"qnaEmbedUrl,omitempty"`
	QueryScaleOutSettings            *DatasetQueryScaleOutSettings `json:"queryScaleOutSettings,omitempty"`
	Tags                             []string                      `json:"tags,omitempty"`
	TargetStorageMode                string                        `json:"targetStorageMode,omitempty"`
	UpstreamDataflows                []DependentDataflow           `json:"upstreamDataflows,omitempty"`
	Users                            []DatasetUser                 `json:"users,omitempty"`
	WebURL                           string                        `json:"webUrl,omitempty"`
}

type DatasetQueryScaleOutSettings struct {
	AutoSyncReadOnlyReplicas bool `json:"autoSyncReadOnlyReplicas,omitempty"`
	MaxReadOnlyReplicas      int  `json:"maxReadOnlyReplicas,omitempty"`
}

// Datasource is a Power BI data source.
type Datasource struct {
	DatasourceID      string                       `json:"datasourceId,omitempty"`
	DatasourceType    string                       `json:"datasourceType,omitempty"`
	ConnectionDetails *DatasourceConnectionDetails `json:"connectionDetails,omitempty"`
	ConnectionString  string                       `json:"connectionString,omitempty"`
	GatewayID         string                       `json:"gatewayId,omitempty"`
	Name              string                       `json:"name,omitempty"`
}

type DatasourceConnectionDetails struct {
	Account      string `json:"account,omitempty"`
	ClassInfo    string `json:"classInfo,omitempty"`
	Database     string `json:"database,omitempty"`
	Domain       string `json:"domain,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
	Kind         string `json:"kind,omitempty"`
	LoginServer  string `json:"loginServer,omitempty"`
	Path         string `json:"path,omitempty"`
	Server       string `json:"server,omitempty"`
	URL          string `json:"url,omitempty"`
}

type DatasetUserAccessRight string

const (
	DatasetUserAccessRightNone                    DatasetUserAccessRight = "None"
	DatasetUserAccessRightRead                    DatasetUserAccessRight = "Read"
	DatasetUserAccessRightReadWrite               DatasetUserAccessRight = "ReadWrite"
	DatasetUserAccessRightReadReshare             DatasetUserAccessRight = "ReadReshare"
	DatasetUserAccessRightReadWriteReshare        DatasetUserAccessRight = "ReadWriteReshare"
	DatasetUserAccessRightReadExplore             DatasetUserAccessRight = "ReadExplore"
	DatasetUserAccessRightReadReshareExplore      DatasetUserAccessRight = "ReadReshareExplore"
	DatasetUserAccessRightReadWriteExplore        DatasetUserAccessRight = "ReadWriteExplore"
	DatasetUserAccessRightReadWriteReshareExplore DatasetUserAccessRight = "ReadWriteReshareExplore"
)

type DatasetUser struct {
	User                   `json:",inline"`
	DatasetUserAccessRight DatasetUserAccessRight `json:"datasetUserAccessRight"`
}

// DeleteRowsOptions controls query parameters for deleting rows from a table in a push dataset.
type DeleteRowsOptions struct {
	Filter string `url:"$filter,omitempty"`
}

type DefaultRetentionPolicy string

const (
	DefaultRetentionPolicyNone      DefaultRetentionPolicy = "None"
	DefaultRetentionPolicyBasicFIFO DefaultRetentionPolicy = "basicFIFO"
)

type DatasetOptions struct {
	DefaultRetentionPolicy DefaultRetentionPolicy `url:"defaultRetentionPolicy,omitempty"`
}

type DatasetMode string

const (
	DatasetModeAsAzure       DatasetMode = "AsAzure"
	DatasetModeAsOnPrem      DatasetMode = "AsOnPrem"
	DatasetModePush          DatasetMode = "Push"
	DatasetModeStreaming     DatasetMode = "Streaming"
	DatasetModePushStreaming DatasetMode = "PushStreaming"
)

type CrossFilteringBehavior string

const (
	CrossFilteringBehaviorOneDirection   CrossFilteringBehavior = "OneDirection"
	CrossFilteringBehaviorBothDirections CrossFilteringBehavior = "BothDirections"
	CrossFilteringBehaviorAutomatic      CrossFilteringBehavior = "Automatic"
)

type Relationship struct {
	CrossFilteringBehavior CrossFilteringBehavior `json:"crossFilteringBehavior,omitempty"`
	FromColumn             string                 `json:"fromColumn,omitempty"`
	FromTable              string                 `json:"fromTable,omitempty"`
	Name                   string                 `json:"name,omitempty"`
	ToColumn               string                 `json:"toColumn,omitempty"`
	ToTable                string                 `json:"toTable,omitempty"`
}

type CreateDatasetRequest struct {
	Datasources   []Datasource   `json:"datasources,omitempty"`
	DefaultMode   DatasetMode    `json:"defaultMode,omitempty"`
	Name          string         `json:"name"`
	Relationships []Relationship `json:"relationships,omitempty"`
	Tables        []Table        `json:"tables"`
}

// Column is a dataset column.
type Column struct {
	DataCategory string `json:"dataCategory,omitempty"`
	DataType     string `json:"dataType,omitempty"`
	FormatString string `json:"formatString,omitempty"`
	IsHidden     bool   `json:"isHidden,omitempty"`
	Name         string `json:"name,omitempty"`
	SortByColumn string `json:"sortByColumn,omitempty"`
	SummarizeBy  string `json:"summarizeBy,omitempty"`
}

type Row struct {
	ID string `json:"id,omitempty"`
}

type PostRowsRequest struct {
	Rows []map[string]any `json:"rows"`
}

type Measure struct {
	Description  string `json:"description,omitempty"`
	Expression   string `json:"expression,omitempty"`
	FormatString string `json:"formatString,omitempty"`
	IsHidden     bool   `json:"isHidden,omitempty"`
	Name         string `json:"name,omitempty"`
}

type ASMashupExpression struct {
	Expression string `json:"expression,omitempty"`
}

// Table is a dataset table.
type Table struct {
	Columns     []Column             `json:"columns,omitempty"`
	Description string               `json:"description,omitempty"`
	IsHidden    bool                 `json:"isHidden,omitempty"`
	Measures    []Measure            `json:"measures,omitempty"`
	Name        string               `json:"name,omitempty"`
	Rows        []Row                `json:"rows,omitempty"`
	Source      []ASMashupExpression `json:"source,omitempty"`
}

// TableList is Power BI table collection
type TableList struct {
	Value []Table `json:"value"`
}
