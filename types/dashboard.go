package types

type AddDashboardRequest struct {
	Name string `json:"name"`
}

type Dashboard struct {
	AppID         string          `json:"appId,omitempty"`
	DisplayName   string          `json:"displayName"`
	EmbedURL      string          `json:"embedUrl"`
	ID            string          `json:"id"`
	IsReadOnly    bool            `json:"isReadOnly"`
	Subscriptions []Subscription  `json:"subscriptions,omitempty"`
	Users         []DashboardUser `json:"users,omitempty"`
	WebURL        string          `json:"webUrl,omitempty"`
}

type DashboardList struct {
	Value []Dashboard `json:"value"`
}

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
