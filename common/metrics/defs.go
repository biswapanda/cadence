package metrics

// types used/defined by the package
type (
	// MetricName is the name of the metric
	MetricName string

	// MetricType is the type of the metric
	MetricType int

	// metricDefinition contains the definition for a metric
	metricDefinition struct {
		metricType MetricType // metric type
		metricName MetricName // metric name
	}

	// scopeDefinition holds the tag definitions for a scope
	scopeDefinition struct {
		operation string            // 'operation' tag for scope
		tags      map[string]string // additional tags for scope
	}

	// ServiceIdx is an index that uniquely identifies the service
	ServiceIdx int
)

// MetricTypes which are supported
const (
	Counter MetricType = iota
	Timer
	Gauge
)

// Service names for all services that emit metrics.
const (
	Common = iota
	Frontend
	History
	Matching
	NumServices
)

// Common tags for all services
const (
	HostnameTagName  = "hostname"
	OperationTagName = "operation"
)

// This package should hold all the metrics and tags for cadence
const (
	UnknownDirectoryTagValue = "Unknown"
)

// Common service base metrics
const (
	RestartCount         = "restarts"
	NumGoRoutinesGauge   = "num-goroutines"
	GoMaxProcsGauge      = "gomaxprocs"
	MemoryAllocatedGauge = "memory.allocated"
	MemoryHeapGauge      = "memory.heap"
	MemoryHeapIdleGauge  = "memory.heapidle"
	MemoryHeapInuseGauge = "memory.heapinuse"
	MemoryStackGauge     = "memory.stack"
	NumGCCounter         = "memory.num-gc"
	GcPauseMsTimer       = "memory.gc-pause-ms"
)

// ServiceMetrics are types for common service base metrics
var ServiceMetrics = map[MetricName]MetricType{
	RestartCount: Counter,
}

// GoRuntimeMetrics represent the runtime stats from go runtime
var GoRuntimeMetrics = map[MetricName]MetricType{
	NumGoRoutinesGauge:   Gauge,
	GoMaxProcsGauge:      Gauge,
	MemoryAllocatedGauge: Gauge,
	MemoryHeapGauge:      Gauge,
	MemoryHeapIdleGauge:  Gauge,
	MemoryHeapInuseGauge: Gauge,
	MemoryStackGauge:     Gauge,
	NumGCCounter:         Counter,
	GcPauseMsTimer:       Timer,
}

// Scopes enum
const (
	// -- Common Operation scopes --

	// CreateShardScope tracks CreateShard calls made by service to persistence layer
	CreateShardScope = iota
	// GetShardScope tracks GetShard calls made by service to persistence layer
	GetShardScope
	// UpdateShardScope tracks UpdateShard calls made by service to persistence layer
	UpdateShardScope
	// CreateWorkflowExecutionScope tracks CreateWorkflowExecution calls made by service to persistence layer
	CreateWorkflowExecutionScope
	// GetWorkflowExecutionScope tracks GetWorkflowExecution calls made by service to persistence layer
	GetWorkflowExecutionScope
	// UpdateWorkflowExecutionScope tracks UpdateWorkflowExecution calls made by service to persistence layer
	UpdateWorkflowExecutionScope
	// DeleteWorkflowExecutionScope tracks DeleteWorkflowExecution calls made by service to persistence layer
	DeleteWorkflowExecutionScope
	// GetTransferTasksScope tracks GetTransferTasks calls made by service to persistence layer
	GetTransferTasksScope
	// CompleteTransferTasksScope tracks CompleteTransferTasks calls made by service to persistence layer
	CompleteTransferTaskScope
	// GetTimerIndexTasksScope tracks GetTimerIndexTasks calls made by service to persistence layer
	GetTimerIndexTasksScope
	// GetWorkflowMutableStateScope tracks GetWorkflowMutableState calls made by service to persistence layer
	GetWorkflowMutableStateScope
	// CreateTaskScope tracks CreateTask calls made by service to persistence layer
	CreateTaskScope
	// GetTasksScope tracks GetTasks calls made by service to persistence layer
	GetTasksScope
	// CompleteTaskScope tracks CompleteTask calls made by service to persistence layer
	CompleteTaskScope
	// LeaseTaskListScope tracks LeaseTaskList calls made by service to persistence layer
	LeaseTaskListScope
	// UpdateTaskListScope tracks UpdateTaskListScope calls made by service to persistence layer
	UpdateTaskListScope
	// HistoryClientStartWorkflowExecutionScope tracks RPC calls to history service
	HistoryClientStartWorkflowExecutionScope
	// HistoryClientRecordActivityTaskHeartbeatScope tracks RPC calls to history service
	HistoryClientRecordActivityTaskHeartbeatScope
	// HistoryClientRespondDecisionTaskCompletedScope tracks RPC calls to history service
	HistoryClientRespondDecisionTaskCompletedScope
	// HistoryClientRespondActivityTaskCompletedScope tracks RPC calls to history service
	HistoryClientRespondActivityTaskCompletedScope
	// HistoryClientRespondActivityTaskFailedScope tracks RPC calls to history service
	HistoryClientRespondActivityTaskFailedScope
	// HistoryClientGetWorkflowExecutionHistoryScope tracks RPC calls to history service
	HistoryClientGetWorkflowExecutionHistoryScope
	// HistoryClientRecordDecisionTaskStartedScope tracks RPC calls to history service
	HistoryClientRecordDecisionTaskStartedScope
	// HistoryClientRecordActivityTaskStartedScope tracks RPC calls to history service
	HistoryClientRecordActivityTaskStartedScope
	// MatchingClientPollForDecisionTaskScope tracks RPC calls to matching service
	MatchingClientPollForDecisionTaskScope
	// MatchingClientPollForActivityTaskScope tracks RPC calls to matching service
	MatchingClientPollForActivityTaskScope
	// MatchingClientAddActivityTaskScope tracks RPC calls to matching service
	MatchingClientAddActivityTaskScope
	// MatchingClientAddDecisionTaskScope tracks RPC calls to matching service
	MatchingClientAddDecisionTaskScope

	NumCommonScopes
)

// -- Operation scopes for Frontend service --
const (
	// StartWorkflowExecutionScope tracks StartWorkflowExecution API calls received by service
	StartWorkflowExecutionScope = iota + NumCommonScopes
	// PollForDecisionTaskScope tracks PollForDecisionTask API calls received by service
	PollForDecisionTaskScope
	// PollForActivityTaskScope tracks PollForActivityTask API calls received by service
	PollForActivityTaskScope
	// RecordActivityTaskHeartbeatScope tracks RecordActivityTaskHeartbeat API calls received by service
	RecordActivityTaskHeartbeatScope
	// RespondDecisionTaskCompletedScope tracks RespondDecisionTaskCompleted API calls received by service
	RespondDecisionTaskCompletedScope
	// RespondActivityTaskCompletedScope tracks RespondActivityTaskCompleted API calls received by service
	RespondActivityTaskCompletedScope
	// RespondActivityTaskFailedScope tracks RespondActivityTaskFailed API calls received by service
	RespondActivityTaskFailedScope
	// GetWorkflowExecutionHistoryScope tracks GetWorkflowExecutionHistory API calls received by service
	GetWorkflowExecutionHistoryScope

	NumFrontendScopes
)

// -- Operation scopes for History service --
const (
	// HistoryStartWorkflowExecutionScope tracks StartWorkflowExecution API calls received by service
	HistoryStartWorkflowExecutionScope = iota + NumCommonScopes
	// HistoryRecordActivityTaskHeartbeatScope tracks RecordActivityTaskHeartbeat API calls received by service
	HistoryRecordActivityTaskHeartbeatScope
	// HistoryRespondDecisionTaskCompletedScope tracks RespondDecisionTaskCompleted API calls received by service
	HistoryRespondDecisionTaskCompletedScope
	// HistoryRespondActivityTaskCompletedScope tracks RespondActivityTaskCompleted API calls received by service
	HistoryRespondActivityTaskCompletedScope
	// HistoryRespondActivityTaskFailedScope tracks RespondActivityTaskFailed API calls received by service
	HistoryRespondActivityTaskFailedScope
	// HistoryGetWorkflowExecutionHistoryScope tracks GetWorkflowExecutionHistory API calls received by service
	HistoryGetWorkflowExecutionHistoryScope
	// HistoryRecordDecisionTaskStartedScope tracks RecordDecisionTaskStarted API calls received by service
	HistoryRecordDecisionTaskStartedScope
	// HistoryRecordActivityTaskStartedScope tracks RecordActivityTaskStarted API calls received by service
	HistoryRecordActivityTaskStartedScope

	NumHistoryScopes
)

// -- Operation scopes for Matching service --
const (
	// PollForDecisionTaskScope tracks PollForDecisionTask API calls received by service
	MatchingPollForDecisionTaskScope = iota + NumCommonScopes
	// PollForActivityTaskScope tracks PollForActivityTask API calls received by service
	MatchingPollForActivityTaskScope
	// MatchingAddActivityTaskScope tracks AddActivityTask API calls received by service
	MatchingAddActivityTaskScope
	// MatchingAddDecisionTaskScope tracks AddDecisionTask API calls received by service
	MatchingAddDecisionTaskScope

	NumMatchingScopes
)

// ScopeDefs record the scopes for all services
var ScopeDefs = map[ServiceIdx]map[int]scopeDefinition{
	// common scope Names
	Common: {
		CreateShardScope:             {operation: "CreateShard"},
		GetShardScope:                {operation: "GetShard"},
		UpdateShardScope:             {operation: "UpdateShard"},
		CreateWorkflowExecutionScope: {operation: "CreateWorkflowExecution"},
		GetWorkflowExecutionScope:    {operation: "GetWorkflowExecution"},
		UpdateWorkflowExecutionScope: {operation: "UpdateWorkflowExecution"},
		DeleteWorkflowExecutionScope: {operation: "DeleteWorkflowExecution"},
		GetTransferTasksScope:        {operation: "GetTransferTasks"},
		CompleteTransferTaskScope:    {operation: "CompleteTransferTask"},
		GetTimerIndexTasksScope:      {operation: "GetTimerIndexTasks"},
		GetWorkflowMutableStateScope: {operation: "GetWorkflowMutableState"},
		CreateTaskScope:              {operation: "CreateTask"},
		GetTasksScope:                {operation: "GetTasks"},
		CompleteTaskScope:            {operation: "CompleteTask"},
		LeaseTaskListScope:           {operation: "LeaseTaskList"},
		UpdateTaskListScope:          {operation: "UpdateTaskList"},

		HistoryClientStartWorkflowExecutionScope:       {operation: "HistoryClientStartWorkflowExecution"},
		HistoryClientRecordActivityTaskHeartbeatScope:  {operation: "HistoryClientRecordActivityTaskHeartbeat"},
		HistoryClientRespondDecisionTaskCompletedScope: {operation: "HistoryClientRespondDecisionTaskCompleted"},
		HistoryClientRespondActivityTaskCompletedScope: {operation: "HistoryClientRespondActivityTaskCompleted"},
		HistoryClientRespondActivityTaskFailedScope:    {operation: "HistoryClientRespondActivityTaskFailed"},
		HistoryClientGetWorkflowExecutionHistoryScope:  {operation: "HistoryClientGetWorkflowExecutionHistory"},
		HistoryClientRecordDecisionTaskStartedScope:    {operation: "HistoryClientRecordDecisionTaskStarted"},
		HistoryClientRecordActivityTaskStartedScope:    {operation: "HistoryClientRecordActivityTaskStarted"},
		MatchingClientPollForDecisionTaskScope:         {operation: "MatchingClientPollForDecisionTask"},
		MatchingClientPollForActivityTaskScope:         {operation: "MatchingClientPollForActivityTask"},
		MatchingClientAddActivityTaskScope:             {operation: "MatchingClientAddActivityTask"},
		MatchingClientAddDecisionTaskScope:             {operation: "MatchingClientAddDecisionTask"},
	},
	// Frontend Scope Names
	Frontend: {
		StartWorkflowExecutionScope:       {operation: "StartWorkflowExecution"},
		PollForDecisionTaskScope:          {operation: "PollForDecisionTask"},
		PollForActivityTaskScope:          {operation: "PollForActivityTask"},
		RecordActivityTaskHeartbeatScope:  {operation: "RecordActivityTaskHeartbeat"},
		RespondDecisionTaskCompletedScope: {operation: "RespondDecisionTaskCompleted"},
		RespondActivityTaskCompletedScope: {operation: "RespondActivityTaskCompleted"},
		RespondActivityTaskFailedScope:    {operation: "RespondActivityTaskFailed"},
		GetWorkflowExecutionHistoryScope:  {operation: "GetWorkflowExecutionHistory"},
	},
	// History Scope Names
	History: {
		HistoryStartWorkflowExecutionScope:       {operation: "StartWorkflowExecution"},
		HistoryRecordActivityTaskHeartbeatScope:  {operation: "RecordActivityTaskHeartbeat"},
		HistoryRespondDecisionTaskCompletedScope: {operation: "RespondDecisionTaskCompleted"},
		HistoryRespondActivityTaskCompletedScope: {operation: "RespondActivityTaskCompleted"},
		HistoryRespondActivityTaskFailedScope:    {operation: "RespondActivityTaskFailed"},
		HistoryGetWorkflowExecutionHistoryScope:  {operation: "GetWorkflowExecutionHistory"},
		HistoryRecordDecisionTaskStartedScope:    {operation: "RecordDecisionTaskStarted"},
		HistoryRecordActivityTaskStartedScope:    {operation: "RecordActivityTaskStarted"},
	},
	// Matching Scope Names
	Matching: {
		MatchingPollForDecisionTaskScope: {operation: "PollForDecisionTask"},
		MatchingPollForActivityTaskScope: {operation: "PollForActivityTask"},
		MatchingAddActivityTaskScope:     {operation: "AddActivityTask"},
		MatchingAddDecisionTaskScope:     {operation: "AddDecisionTask"},
	},
}

// Metrics enum
const (
	WorkflowRequests = iota
	WorkflowFailures
	WorkflowLatency
)

// MetricDefs record the metrics for all services
var MetricDefs = map[ServiceIdx]map[int]metricDefinition{
	Common: {
		WorkflowRequests: {metricName: "requests", metricType: Counter},
		WorkflowFailures: {metricName: "errors", metricType: Counter},
		WorkflowLatency:  {metricName: "latency", metricType: Timer},
	},
	Frontend: {},
	History:  {},
	Matching: {},
}

// ErrorClass is an enum to help with classifying SLA vs. non-SLA errors (SLA = "service level agreement")
type ErrorClass uint8

const (
	// NoError indicates that there is no error (error should be nil)
	NoError = ErrorClass(iota)
	// UserError indicates that this is NOT an SLA-reportable error
	UserError
	// InternalError indicates that this is an SLA-reportable error
	InternalError
)