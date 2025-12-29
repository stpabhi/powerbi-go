package types

type PositionConflictAction string

const (
	PositionConflictActionTail  PositionConflictAction = "Tail"
	PositionConflictActionAbort PositionConflictAction = "Abort"
)

type CloneTileRequest struct {
	PositionConflictAction PositionConflictAction `json:"positionConflictAction,omitempty"`
	TargetDashboardID      string                 `json:"targetDashboardId"`
	TargetModelID          string                 `json:"targetModelId,omitempty"`
	TargetReportID         string                 `json:"targetReportId,omitempty"`
	TargetWorkspaceID      string                 `json:"targetWorkspaceId,omitempty"`
}

type Tile struct {
	ColSpan   int    `json:"colSpan"`
	DatasetID string `json:"datasetId,omitempty"`
	EmbedData string `json:"embedData"`
	EmbedURL  string `json:"embedUrl"`
	ID        string `json:"id"`
	ReportID  string `json:"reportId,omitempty"`
	RowSpan   int    `json:"rowSpan"`
	Title     string `json:"title"`
}

type TileList struct {
	Value []Tile `json:"value"`
}
