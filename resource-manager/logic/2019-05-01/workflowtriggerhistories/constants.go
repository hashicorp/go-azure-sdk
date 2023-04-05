package workflowtriggerhistories

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkflowStatus string

const (
	WorkflowStatusAborted      WorkflowStatus = "Aborted"
	WorkflowStatusCancelled    WorkflowStatus = "Cancelled"
	WorkflowStatusFailed       WorkflowStatus = "Failed"
	WorkflowStatusFaulted      WorkflowStatus = "Faulted"
	WorkflowStatusIgnored      WorkflowStatus = "Ignored"
	WorkflowStatusNotSpecified WorkflowStatus = "NotSpecified"
	WorkflowStatusPaused       WorkflowStatus = "Paused"
	WorkflowStatusRunning      WorkflowStatus = "Running"
	WorkflowStatusSkipped      WorkflowStatus = "Skipped"
	WorkflowStatusSucceeded    WorkflowStatus = "Succeeded"
	WorkflowStatusSuspended    WorkflowStatus = "Suspended"
	WorkflowStatusTimedOut     WorkflowStatus = "TimedOut"
	WorkflowStatusWaiting      WorkflowStatus = "Waiting"
)

func PossibleValuesForWorkflowStatus() []string {
	return []string{
		string(WorkflowStatusAborted),
		string(WorkflowStatusCancelled),
		string(WorkflowStatusFailed),
		string(WorkflowStatusFaulted),
		string(WorkflowStatusIgnored),
		string(WorkflowStatusNotSpecified),
		string(WorkflowStatusPaused),
		string(WorkflowStatusRunning),
		string(WorkflowStatusSkipped),
		string(WorkflowStatusSucceeded),
		string(WorkflowStatusSuspended),
		string(WorkflowStatusTimedOut),
		string(WorkflowStatusWaiting),
	}
}
