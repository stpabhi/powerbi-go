package types

type DataflowUserAccessRight string

const (
	DataflowUserAccessRightNone        DataflowUserAccessRight = "None"
	DataflowUserAccessRightRead        DataflowUserAccessRight = "Read"
	DataflowUserAccessRightReadWrite   DataflowUserAccessRight = "ReadWrite"
	DataflowUserAccessRightReadReshare DataflowUserAccessRight = "ReadReshare"
	DataflowUserAccessRightOwner       DataflowUserAccessRight = "Owner"
)

type DataflowUser struct {
	User                    `json:",inline"`
	DataflowUserAccessRight DataflowUserAccessRight `json:"dataflowUserAccessRight"`
}
