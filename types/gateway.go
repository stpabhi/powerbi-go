package types

type GatewayPublicKey struct {
	Exponent string `json:"exponent"`
	Modulus  string `json:"modulus"`
}

type Gateway struct {
	GatewayAnnotation string           `json:"gatewayAnnotation"`
	GatewayStatus     string           `json:"gatewayStatus"`
	ID                string           `json:"id"`
	Name              string           `json:"name"`
	PublicKey         GatewayPublicKey `json:"publicKey"`
	Type              string           `json:"type"`
}

type GatewayList struct {
	Value []Gateway `json:"value"`
}

type GatewayDatasourceCredentialDetails struct {
	UseEndUserOAuth2Credentials bool `json:"useEndUserOAuth2Credentials"`
}

type CredentialType string

const (
	CredentialTypeBasic     CredentialType = "Basic"
	CredentialTypeWindows   CredentialType = "Windows"
	CredentialTypeAnonymous CredentialType = "Anonymous"
	CredentialTypeOAuth2    CredentialType = "OAuth2"
	CredentialTypeKey       CredentialType = "Key"
	CredentialTypeSAS       CredentialType = "SAS"
)

type GatewayDatasource struct {
	ConnectionDetails string                             `json:"connectionDetails"`
	CredentialDetails GatewayDatasourceCredentialDetails `json:"credentialDetails"`
	CredentialType    CredentialType                     `json:"credentialType"`
	DatasourceName    string                             `json:"datasourceName"`
	DatasourceType    string                             `json:"datasourceType"`
	GatewayID         string                             `json:"gatewayId"`
	ID                string                             `json:"id"`
}

type GatewayDatasourceList struct {
	Value []GatewayDatasource `json:"value"`
}
