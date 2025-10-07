package types

type DashboardUserAccessRight string

const (
	DashboardUserAccessRightNone        DashboardUserAccessRight = "None"
	DashboardUserAccessRightRead        DashboardUserAccessRight = "Read"
	DashboardUserAccessRightReadWrite   DashboardUserAccessRight = "ReadWrite"
	DashboardUserAccessRightReadReshare DashboardUserAccessRight = "ReadReshare"
	DashboardUserAccessRightReadCopy    DashboardUserAccessRight = "ReadCopy"
	DashboardUserAccessRightOwner       DashboardUserAccessRight = "Owner"
)

type DashboardUser struct {
	User                     `json:",inline"`
	DashboardUserAccessRight DashboardUserAccessRight `json:"dashboardUserAccessRight"`
}
