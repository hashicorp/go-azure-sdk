package workflowresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationDirection string

const (
	OperationDirectionCancel OperationDirection = "cancel"
	OperationDirectionDo     OperationDirection = "do"
	OperationDirectionUndo   OperationDirection = "undo"
)

func PossibleValuesForOperationDirection() []string {
	return []string{
		string(OperationDirectionCancel),
		string(OperationDirectionDo),
		string(OperationDirectionUndo),
	}
}

type WorkflowStatus string

const (
	WorkflowStatusAborted   WorkflowStatus = "aborted"
	WorkflowStatusActive    WorkflowStatus = "active"
	WorkflowStatusExpired   WorkflowStatus = "expired"
	WorkflowStatusFailed    WorkflowStatus = "failed"
	WorkflowStatusSucceeded WorkflowStatus = "succeeded"
)

func PossibleValuesForWorkflowStatus() []string {
	return []string{
		string(WorkflowStatusAborted),
		string(WorkflowStatusActive),
		string(WorkflowStatusExpired),
		string(WorkflowStatusFailed),
		string(WorkflowStatusSucceeded),
	}
}
