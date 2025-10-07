package types

// Datasource is a Power BI data source.
type Datasource struct {
	DatasourceID      string                      `json:"datasourceId,omitempty"`
	DatasourceType    string                      `json:"datasourceType,omitempty"`
	ConnectionDetails DatasourceConnectionDetails `json:"connectionDetails,omitempty"`
	ConnectionString  string                      `json:"connectionString,omitempty"`
	GatewayID         string                      `json:"gatewayId,omitempty"`
	Name              string                      `json:"name,omitempty"`
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
