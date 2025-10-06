package types

// DatasourceIdentity is effective identity for connecting DirectQuery data sources with single sign-on (SSO) enabled.
// https://learn.microsoft.com/en-us/rest/api/power-bi/embed-token/generate-token#datasourceidentity
type DatasourceIdentity struct {
	IdentityBlob string               `json:"identityBlob,omitempty"`
	Datasources  []DatasourceSelector `json:"datasources,omitempty"`
}

// DatasourceSelector is an object that uniquely identifies a single data source by its connection details.
// https://learn.microsoft.com/en-us/rest/api/power-bi/embed-token/generate-token#datasourceselector
type DatasourceSelector struct {
	DatasourceType    string                      `json:"datasourceType,omitempty"`
	ConnectionDetails DatasourceConnectionDetails `json:"connectionDetails,omitempty"`
}

// EffectiveIdentity defines the user identity and roles.
// https://learn.microsoft.com/en-us/rest/api/power-bi/embed-token/generate-token#effectiveidentity
type EffectiveIdentity struct {
	Username         string       `json:"username,omitempty"`
	Roles            []string     `json:"roles,omitempty"`
	Datasets         []string     `json:"datasets,omitempty"`
	CustomData       string       `json:"customData,omitempty"`
	IdentityBlob     IdentityBlob `json:"identityBlob,omitempty"`
	AuditableContext string       `json:"auditableContext,omitempty"`
	Reports          []string     `json:"reports,omitempty"`
}

type EmbedToken struct {
	Expiration string `json:"expiration"`
	Token      string `json:"token"`
	TokenID    string `json:"tokenId"`
}

type TokenAccessLevel string

const (
	TokenAccessLevelView   TokenAccessLevel = "View"
	TokenAccessLevelEdit   TokenAccessLevel = "Edit"
	TokenAccessLevelCreate TokenAccessLevel = "Create"
)

type GenerateTokenRequest struct {
	AccessLevel       TokenAccessLevel    `json:"accessLevel,omitempty"`
	AllowSaveAs       bool                `json:"allowSaveAs,omitempty"`
	DatasetID         string              `json:"datasetId,omitempty"`
	Identities        []EffectiveIdentity `json:"identities,omitempty"`
	LifetimeInMinutes int                 `json:"lifetimeInMinutes,omitempty"`
}

// GenerateTokenRequestV2
// https://learn.microsoft.com/en-us/rest/api/power-bi/embed-token/generate-token#generatetokenrequestv2
type GenerateTokenRequestV2 struct {
	Datasets             []GenerateTokenRequestV2Dataset         `json:"datasets,omitempty"`
	DatasourceIdentities []DatasourceIdentity                    `json:"datasourceIdentities,omitempty"`
	Identities           []EffectiveIdentity                     `json:"identities,omitempty"`
	LifetimeInMinutes    int                                     `json:"lifetimeInMinutes,omitempty"`
	Reports              []GenerateTokenRequestV2Report          `json:"reports,omitempty"`
	TargetWorkspaces     []GenerateTokenRequestV2TargetWorkspace `json:"targetWorkspaces,omitempty"`
}

type GenerateTokenRequestV2Dataset struct {
	ID              string          `json:"id,omitempty"`
	XMLAPermissions XMLAPermissions `json:"xmlaPermissions,omitempty"`
}
type GenerateTokenRequestV2Report struct {
	AllowEdit bool   `json:"allowEdit,omitempty"`
	ID        string `json:"id,omitempty"`
}
type GenerateTokenRequestV2TargetWorkspace struct {
	ID string `json:"id,omitempty"`
}

// IdentityBlob is a blob for specifying an identity. Only supported for datasets with a DirectQuery connection to Azure SQL
// https://learn.microsoft.com/en-us/rest/api/power-bi/embed-token/generate-token#identityblob
type IdentityBlob struct {
	Value string `json:"value,omitempty"`
}

type XMLAPermissions string

const (
	XMLAPermissionsOff      XMLAPermissions = "Off"
	XMLAPermissionsReadOnly XMLAPermissions = "ReadOnly"
)
