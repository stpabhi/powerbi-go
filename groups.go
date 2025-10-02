package powerbi

import (
	"context"
	"net/http"

	"github.com/stpabhi/powerbi-go/types"
)

const groupsBasePath = "groups"

// GroupsService handles communication with the groups related methods of the Power BI API.
// https://learn.microsoft.com/en-us/rest/api/power-bi/groups
type GroupsService service

func (s *GroupsService) List(ctx context.Context) ([]types.Group, error) {
	url := groupsBasePath
	_, resp, err := s.client.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var result types.GroupList
	_, err = toObject(resp, &result)
	return result.Value, err
}
