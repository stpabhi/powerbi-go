package types

type ReportType string

const (
	ReportTypePowerBI   ReportType = "PowerBIReport"
	ReportTypePaginated ReportType = "PaginatedReport"
)

type ReportUserAccessRight string

const (
	ReportUserAccessRightNone        ReportUserAccessRight = "None"
	ReportUserAccessRightRead        ReportUserAccessRight = "Read"
	ReportUserAccessRightReadWrite   ReportUserAccessRight = "ReadWrite"
	ReportUserAccessRightReadReshare ReportUserAccessRight = "ReadReshare"
	ReportUserAccessRightReadCopy    ReportUserAccessRight = "ReadCopy"
	ReportUserAccessRightOwner       ReportUserAccessRight = "Owner"
)

type UserMetadata struct {
	DisplayName   string                   `json:"displayName"`
	EmailAddress  string                   `json:"emailAddress"`
	GraphID       string                   `json:"graphId"`
	Identifier    string                   `json:"identifier"`
	PrincipalType PrincipalType            `json:"principalType"`
	Profile       *ServicePrincipalProfile `json:"profile,omitempty"`
	UserType      string                   `json:"userType"`
}

type SubscriptionUser struct {
	UserMetadata `json:",inline"`
}

type ReportUser struct {
	UserMetadata          `json:",inline"`
	ReportUserAccessRight ReportUserAccessRight `json:"reportUserAccessRight"`
}

// Subscription represents an email subscription for a Power BI item
// (such as a report or a dashboard).
type Subscription struct {
	ArtifactDisplayName    string             `json:"artifactDisplayName,omitempty"`
	ArtifactID             string             `json:"artifactId,omitempty"`
	ArtifactType           string             `json:"artifactType,omitempty"`
	AttachmentFormat       string             `json:"attachmentFormat,omitempty"`
	EndDate                string             `json:"endDate,omitempty"`
	Frequency              string             `json:"frequency,omitempty"`
	ID                     string             `json:"id"`
	IsEnabled              bool               `json:"isEnabled,omitempty"`
	LinkToContent          bool               `json:"linkToContent,omitempty"`
	PreviewImage           bool               `json:"previewImage,omitempty"`
	StartDate              string             `json:"startDate,omitempty"`
	SubArtifactDisplayName string             `json:"subArtifactDisplayName,omitempty"`
	Title                  string             `json:"title,omitempty"`
	Users                  []SubscriptionUser `json:"users,omitempty"`
}

type RdlBindDetail struct {
	DataSourceName     string `json:"dataSourceName"`
	DataSourceObjectID string `json:"dataSourceObjectId"`
}

// RdlBindToGatewayRequest binds the report's dataset to a gateway.
type RdlBindToGatewayRequest struct {
	BindDetails     []RdlBindDetail `json:"bindDetails"`
	GatewayObjectID string          `json:"gatewayObjectId"`
}

// Report represents a Power BI report.
type Report struct {
	AppID            string         `json:"appId,omitempty"`
	DatasetID        string         `json:"datasetId,omitempty"`
	Description      string         `json:"description,omitempty"`
	EmbedURL         string         `json:"embedUrl,omitempty"`
	ID               string         `json:"id"`
	IsOwnedByMe      bool           `json:"isOwnedByMe,omitempty"`
	Name             string         `json:"name"`
	OriginalReportID string         `json:"originalReportId,omitempty"`
	ReportType       ReportType     `json:"reportType"`
	Subscriptions    []Subscription `json:"subscriptions,omitempty"`
	Users            []ReportUser   `json:"users,omitempty"`
	WebURL           string         `json:"webUrl,omitempty"`
}

type ReportList struct {
	Value []Report `json:"value"`
}

// CloneReportRequest represents the request body for cloning a report.
type CloneReportRequest struct {
	Name              string `json:"name"`
	TargetModelID     string `json:"targetModelId,omitempty"`
	TargetWorkspaceID string `json:"targetWorkspaceId,omitempty"`
}

// RebindReportRequest represents the request body for rebinding a report to a dataset.
type RebindReportRequest struct {
	DatasetID string `json:"datasetId"`
}

// Page represents a report page.
type Page struct {
	DisplayName string `json:"displayName"`
	Name        string `json:"name"`
	Order       int    `json:"order"`
}

type PageList struct {
	Value []Page `json:"value"`
}

func (r Report) String() string {
	return Stringify(r)
}
