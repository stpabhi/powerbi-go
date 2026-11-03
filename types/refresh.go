package types

// DatasetRefreshRequest is a refresh request. A request that sets anything
// other than NotifyOption is an enhanced refresh, which needs Premium,
// Premium Per User or Fabric capacity and must leave NotifyOption empty.
type DatasetRefreshRequest struct {
	NotifyOption       NotifyOption            `json:"notifyOption,omitempty"`
	Type               DatasetRefreshType      `json:"type,omitempty"`
	CommitMode         DatasetCommitMode       `json:"commitMode,omitempty"`
	MaxParallelism     int                     `json:"maxParallelism,omitempty"`
	RetryCount         int                     `json:"retryCount,omitempty"`
	Objects            []DatasetRefreshObjects `json:"objects,omitempty"`
	ApplyRefreshPolicy *bool                   `json:"applyRefreshPolicy,omitempty"`
	EffectiveDate      string                  `json:"effectiveDate,omitempty"`
	// Timeout is hh:mm:ss.
	Timeout string `json:"timeout,omitempty"`
}

type NotifyOption string

const (
	NotifyOptionNoNotification   NotifyOption = "NoNotification"
	NotifyOptionMailOnFailure    NotifyOption = "MailOnFailure"
	NotifyOptionMailOnCompletion NotifyOption = "MailOnCompletion"
)

type DatasetRefreshType string

const (
	DatasetRefreshTypeFull        DatasetRefreshType = "Full"
	DatasetRefreshTypeClearValues DatasetRefreshType = "ClearValues"
	DatasetRefreshTypeCalculate   DatasetRefreshType = "Calculate"
	DatasetRefreshTypeDataOnly    DatasetRefreshType = "DataOnly"
	DatasetRefreshTypeAutomatic   DatasetRefreshType = "Automatic"
	DatasetRefreshTypeDefragment  DatasetRefreshType = "Defragment"
)

type DatasetCommitMode string

const (
	DatasetCommitModeTransactional DatasetCommitMode = "Transactional"
	DatasetCommitModePartialBatch  DatasetCommitMode = "PartialBatch"
)

// DatasetRefreshObjects is a table, or one partition of it, to refresh.
type DatasetRefreshObjects struct {
	Table     string `json:"table,omitempty"`
	Partition string `json:"partition,omitempty"`
}

// RefreshAccepted is Power BI's answer to an accepted refresh request.
type RefreshAccepted struct {
	// RequestID identifies the refresh, as RefreshExecutionDetails and
	// CancelRefresh take it.
	RequestID string
	// Location is the refresh's URL, for an enhanced refresh.
	Location string
}

// RefreshList is a dataset's refresh history, most recent first.
type RefreshList struct {
	Value []Refresh `json:"value"`
}

// Refresh is one refresh history entry.
type Refresh struct {
	RequestID   string `json:"requestId,omitempty"`
	RefreshType string `json:"refreshType,omitempty"`
	StartTime   string `json:"startTime,omitempty"`
	// EndTime is empty while the refresh runs.
	EndTime string `json:"endTime,omitempty"`
	// Status is Unknown while a refresh runs, then Completed, Failed,
	// Cancelled or Disabled.
	Status string `json:"status,omitempty"`
	// ExtendedStatus is set for an enhanced refresh, such as NotStarted or
	// InProgress.
	ExtendedStatus string `json:"extendedStatus,omitempty"`
	// ServiceExceptionJSON is the failure's error code as JSON text.
	ServiceExceptionJSON string           `json:"serviceExceptionJson,omitempty"`
	RefreshAttempts      []RefreshAttempt `json:"refreshAttempts,omitempty"`
}

type RefreshAttempt struct {
	AttemptID            int    `json:"attemptId"`
	StartTime            string `json:"startTime,omitempty"`
	EndTime              string `json:"endTime,omitempty"`
	ServiceExceptionJSON string `json:"serviceExceptionJson,omitempty"`
	Type                 string `json:"type,omitempty"`
}

// RefreshExecutionDetails is the state of one enhanced refresh.
type RefreshExecutionDetails struct {
	StartTime          string                `json:"startTime,omitempty"`
	EndTime            string                `json:"endTime,omitempty"`
	Type               string                `json:"type,omitempty"`
	CommitMode         string                `json:"commitMode,omitempty"`
	Status             string                `json:"status,omitempty"`
	ExtendedStatus     string                `json:"extendedStatus,omitempty"`
	CurrentRefreshType string                `json:"currentRefreshType,omitempty"`
	NumberOfAttempts   int                   `json:"numberOfAttempts,omitempty"`
	Objects            []RefreshObjectStatus `json:"objects,omitempty"`
	Messages           []RefreshMessage      `json:"messages,omitempty"`
}

type RefreshObjectStatus struct {
	Table     string `json:"table,omitempty"`
	Partition string `json:"partition,omitempty"`
	Status    string `json:"status,omitempty"`
}

type RefreshMessage struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Type    string `json:"type,omitempty"`
}
