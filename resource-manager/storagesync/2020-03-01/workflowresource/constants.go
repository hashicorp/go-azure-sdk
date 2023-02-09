package workflowresource

import "strings"

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

func parseOperationDirection(input string) (*OperationDirection, error) {
	vals := map[string]OperationDirection{
		"cancel": OperationDirectionCancel,
		"do":     OperationDirectionDo,
		"undo":   OperationDirectionUndo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperationDirection(input)
	return &out, nil
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

func parseWorkflowStatus(input string) (*WorkflowStatus, error) {
	vals := map[string]WorkflowStatus{
		"aborted":   WorkflowStatusAborted,
		"active":    WorkflowStatusActive,
		"expired":   WorkflowStatusExpired,
		"failed":    WorkflowStatusFailed,
		"succeeded": WorkflowStatusSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkflowStatus(input)
	return &out, nil
}
