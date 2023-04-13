package workflowrunoperations

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ParameterType string

const (
	ParameterTypeArray        ParameterType = "Array"
	ParameterTypeBool         ParameterType = "Bool"
	ParameterTypeFloat        ParameterType = "Float"
	ParameterTypeInt          ParameterType = "Int"
	ParameterTypeNotSpecified ParameterType = "NotSpecified"
	ParameterTypeObject       ParameterType = "Object"
	ParameterTypeSecureObject ParameterType = "SecureObject"
	ParameterTypeSecureString ParameterType = "SecureString"
	ParameterTypeString       ParameterType = "String"
)

func PossibleValuesForParameterType() []string {
	return []string{
		string(ParameterTypeArray),
		string(ParameterTypeBool),
		string(ParameterTypeFloat),
		string(ParameterTypeInt),
		string(ParameterTypeNotSpecified),
		string(ParameterTypeObject),
		string(ParameterTypeSecureObject),
		string(ParameterTypeSecureString),
		string(ParameterTypeString),
	}
}

func parseParameterType(input string) (*ParameterType, error) {
	vals := map[string]ParameterType{
		"array":        ParameterTypeArray,
		"bool":         ParameterTypeBool,
		"float":        ParameterTypeFloat,
		"int":          ParameterTypeInt,
		"notspecified": ParameterTypeNotSpecified,
		"object":       ParameterTypeObject,
		"secureobject": ParameterTypeSecureObject,
		"securestring": ParameterTypeSecureString,
		"string":       ParameterTypeString,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ParameterType(input)
	return &out, nil
}

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

func parseWorkflowStatus(input string) (*WorkflowStatus, error) {
	vals := map[string]WorkflowStatus{
		"aborted":      WorkflowStatusAborted,
		"cancelled":    WorkflowStatusCancelled,
		"failed":       WorkflowStatusFailed,
		"faulted":      WorkflowStatusFaulted,
		"ignored":      WorkflowStatusIgnored,
		"notspecified": WorkflowStatusNotSpecified,
		"paused":       WorkflowStatusPaused,
		"running":      WorkflowStatusRunning,
		"skipped":      WorkflowStatusSkipped,
		"succeeded":    WorkflowStatusSucceeded,
		"suspended":    WorkflowStatusSuspended,
		"timedout":     WorkflowStatusTimedOut,
		"waiting":      WorkflowStatusWaiting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkflowStatus(input)
	return &out, nil
}
