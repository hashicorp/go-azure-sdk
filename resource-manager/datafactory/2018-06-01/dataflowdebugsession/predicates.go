package dataflowdebugsession

type DataFlowDebugSessionInfoOperationPredicate struct {
	ComputeType            *string
	CoreCount              *int64
	DataFlowName           *string
	IntegrationRuntimeName *string
	LastActivityTime       *string
	NodeCount              *int64
	SessionId              *string
	StartTime              *string
	TimeToLiveInMinutes    *int64
}

func (p DataFlowDebugSessionInfoOperationPredicate) Matches(input DataFlowDebugSessionInfo) bool {

	if p.ComputeType != nil && (input.ComputeType == nil && *p.ComputeType != *input.ComputeType) {
		return false
	}

	if p.CoreCount != nil && (input.CoreCount == nil && *p.CoreCount != *input.CoreCount) {
		return false
	}

	if p.DataFlowName != nil && (input.DataFlowName == nil && *p.DataFlowName != *input.DataFlowName) {
		return false
	}

	if p.IntegrationRuntimeName != nil && (input.IntegrationRuntimeName == nil && *p.IntegrationRuntimeName != *input.IntegrationRuntimeName) {
		return false
	}

	if p.LastActivityTime != nil && (input.LastActivityTime == nil && *p.LastActivityTime != *input.LastActivityTime) {
		return false
	}

	if p.NodeCount != nil && (input.NodeCount == nil && *p.NodeCount != *input.NodeCount) {
		return false
	}

	if p.SessionId != nil && (input.SessionId == nil && *p.SessionId != *input.SessionId) {
		return false
	}

	if p.StartTime != nil && (input.StartTime == nil && *p.StartTime != *input.StartTime) {
		return false
	}

	if p.TimeToLiveInMinutes != nil && (input.TimeToLiveInMinutes == nil && *p.TimeToLiveInMinutes != *input.TimeToLiveInMinutes) {
		return false
	}

	return true
}
