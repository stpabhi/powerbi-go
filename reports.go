package powerbi

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stpabhi/powerbi-go/types"
)

const reportsBasePath = "reports"

// ReportsService handles communication with the reports related methods of the Power BI API.
// https://learn.microsoft.com/en-us/rest/api/power-bi/reports
type ReportsService service

// BindToGateway binds the report's dataset to a gateway in My workspace.
// POST /reports/{reportId}/Default.BindToGateway
func (s *ReportsService) BindToGateway(ctx context.Context, reportID string, req types.RdlBindToGatewayRequest) error {
	u := fmt.Sprintf("%s/%s/Default.BindToGateway", reportsBasePath, url.PathEscape(reportID))
	_, resp, err := s.client.postJSON(ctx, u, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// BindToGatewayInGroup binds the report's dataset to a gateway in a workspace.
// POST /groups/{groupId}/reports/{reportId}/Default.BindToGateway
func (s *ReportsService) BindToGatewayInGroup(ctx context.Context, groupID, reportID string, req types.RdlBindToGatewayRequest) error {
	u := fmt.Sprintf("%s/%s/%s/%s/Default.BindToGateway", groupsBasePath, url.PathEscape(groupID), reportsBasePath, url.PathEscape(reportID))
	_, resp, err := s.client.postJSON(ctx, u, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// Clone clones a report from My workspace.
// POST /reports/{reportId}/Clone
func (s *ReportsService) Clone(ctx context.Context, reportID string, req types.CloneReportRequest) (*types.Report, error) {
	u := fmt.Sprintf("%s/%s/Clone", reportsBasePath, url.PathEscape(reportID))
	_, resp, err := s.client.postJSON(ctx, u, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.Report{})
}

// CloneInGroup clones the specified report from the specified workspace.
func (s *ReportsService) CloneInGroup(ctx context.Context, groupID, reportID string, req types.CloneReportRequest) (*types.Report, error) {
	u := fmt.Sprintf("%s/%s/%s/%s/Clone", groupsBasePath, url.PathEscape(groupID), reportsBasePath, url.PathEscape(reportID))
	_, resp, err := s.client.postJSON(ctx, u, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.Report{})
}

// Delete deletes the specified report from My workspace.
func (s *ReportsService) Delete(ctx context.Context, reportID string) error {
	u := fmt.Sprintf("%s/%s", reportsBasePath, url.PathEscape(reportID))
	_, resp, err := s.client.doRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// DeleteInGroup deletes the specified report from the specified workspace.
func (s *ReportsService) DeleteInGroup(ctx context.Context, groupID, reportID string) error {
	u := fmt.Sprintf("%s/%s/%s/%s", groupsBasePath, url.PathEscape(groupID), reportsBasePath, url.PathEscape(reportID))
	_, resp, err := s.client.doRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// GetPage the specified page within the specified report from My workspace.
// GET /reports/{reportId}/pages/{pageName}
func (s *ReportsService) GetPage(ctx context.Context, reportID, pageName string) (*types.Page, error) {
	u := fmt.Sprintf("%s/%s/pages/%s", reportsBasePath, url.PathEscape(reportID), url.PathEscape(pageName))
	_, resp, err := s.client.doRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.Page{})
}

// GetPageInGroup the specified page within the specified report from the specified workspace.
func (s *ReportsService) GetPageInGroup(ctx context.Context, groupID, reportID, pageName string) (*types.Page, error) {
	u := fmt.Sprintf("%s/%s/%s/%s/pages/%s", groupsBasePath, url.PathEscape(groupID), reportsBasePath, url.PathEscape(reportID), url.PathEscape(pageName))
	_, resp, err := s.client.doRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.Page{})
}

// ListPages returns a list of pages within the specified report from My workspace.
// GET /reports/{reportId}/pages
func (s *ReportsService) ListPages(ctx context.Context, reportID string) ([]types.Page, error) {
	u := fmt.Sprintf("%s/%s/pages", reportsBasePath, url.PathEscape(reportID))
	_, resp, err := s.client.doRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.PageList
	_, err = toObject(resp, &result)
	return result.Value, err
}

// ListPagesInGroup returns the pages for a report in a workspace.
func (s *ReportsService) ListPagesInGroup(ctx context.Context, groupID, reportID string) ([]types.Page, error) {
	u := fmt.Sprintf("%s/%s/%s/%s/pages", groupsBasePath, url.PathEscape(groupID), reportsBasePath, url.PathEscape(reportID))
	_, resp, err := s.client.doRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.PageList
	_, err = toObject(resp, &result)

	return result.Value, err
}

// Get returns the specified report from My workspace.
func (s *ReportsService) Get(ctx context.Context, reportID string) (*types.Report, error) {
	u := fmt.Sprintf("%s/%s", reportsBasePath, url.PathEscape(reportID))
	_, resp, err := s.client.doRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.Report{})
}

// GetInGroup returns the specified report from the specified workspace.
func (s *ReportsService) GetInGroup(ctx context.Context, groupID, reportID string) (*types.Report, error) {
	u := fmt.Sprintf("%s/%s/%s/%s", groupsBasePath, url.PathEscape(groupID), reportsBasePath, url.PathEscape(reportID))
	_, resp, err := s.client.doRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toObject(resp, &types.Report{})
}

// List a list of reports from My workspace.
func (s *ReportsService) List(ctx context.Context) ([]types.Report, error) {
	u := reportsBasePath
	_, resp, err := s.client.doRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ReportList
	_, err = toObject(resp, &result)

	return result.Value, err
}

// ListInGroup a list of reports from the specified workspace.
func (s *ReportsService) ListInGroup(ctx context.Context, groupID string) ([]types.Report, error) {
	u := fmt.Sprintf("%s/%s/%s", groupsBasePath, url.PathEscape(groupID), reportsBasePath)
	_, resp, err := s.client.doRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ReportList
	_, err = toObject(resp, &result)

	return result.Value, err
}

// Rebind rebinds the specified report from My workspace to the specified dataset.
// POST /reports/{reportId}/Rebind
func (s *ReportsService) Rebind(ctx context.Context, reportID string, req types.RebindReportRequest) error {
	u := fmt.Sprintf("%s/%s/Rebind", reportsBasePath, url.PathEscape(reportID))
	_, resp, err := s.client.postJSON(ctx, u, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// RebindInGroup rebinds the specified report from the specified workspace to the specified dataset.
func (s *ReportsService) RebindInGroup(ctx context.Context, groupID, reportID string, req types.RebindReportRequest) error {
	u := fmt.Sprintf("%s/%s/%s/%s/Rebind", groupsBasePath, url.PathEscape(groupID), reportsBasePath, url.PathEscape(reportID))
	_, resp, err := s.client.postJSON(ctx, u, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
