package logicapps

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkflowHealthState string

const (
	WorkflowHealthStateHealthy      WorkflowHealthState = "Healthy"
	WorkflowHealthStateNotSpecified WorkflowHealthState = "NotSpecified"
	WorkflowHealthStateUnhealthy    WorkflowHealthState = "Unhealthy"
	WorkflowHealthStateUnknown      WorkflowHealthState = "Unknown"
)

func PossibleValuesForWorkflowHealthState() []string {
	return []string{
		string(WorkflowHealthStateHealthy),
		string(WorkflowHealthStateNotSpecified),
		string(WorkflowHealthStateUnhealthy),
		string(WorkflowHealthStateUnknown),
	}
}

func (s *WorkflowHealthState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWorkflowHealthState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWorkflowHealthState(input string) (*WorkflowHealthState, error) {
	vals := map[string]WorkflowHealthState{
		"healthy":      WorkflowHealthStateHealthy,
		"notspecified": WorkflowHealthStateNotSpecified,
		"unhealthy":    WorkflowHealthStateUnhealthy,
		"unknown":      WorkflowHealthStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkflowHealthState(input)
	return &out, nil
}

type WorkflowKind string

const (
	WorkflowKindAgentic   WorkflowKind = "Agentic"
	WorkflowKindStateful  WorkflowKind = "Stateful"
	WorkflowKindStateless WorkflowKind = "Stateless"
)

func PossibleValuesForWorkflowKind() []string {
	return []string{
		string(WorkflowKindAgentic),
		string(WorkflowKindStateful),
		string(WorkflowKindStateless),
	}
}

func (s *WorkflowKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWorkflowKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWorkflowKind(input string) (*WorkflowKind, error) {
	vals := map[string]WorkflowKind{
		"agentic":   WorkflowKindAgentic,
		"stateful":  WorkflowKindStateful,
		"stateless": WorkflowKindStateless,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkflowKind(input)
	return &out, nil
}

type WorkflowState string

const (
	WorkflowStateCompleted    WorkflowState = "Completed"
	WorkflowStateDeleted      WorkflowState = "Deleted"
	WorkflowStateDisabled     WorkflowState = "Disabled"
	WorkflowStateEnabled      WorkflowState = "Enabled"
	WorkflowStateNotSpecified WorkflowState = "NotSpecified"
	WorkflowStateSuspended    WorkflowState = "Suspended"
)

func PossibleValuesForWorkflowState() []string {
	return []string{
		string(WorkflowStateCompleted),
		string(WorkflowStateDeleted),
		string(WorkflowStateDisabled),
		string(WorkflowStateEnabled),
		string(WorkflowStateNotSpecified),
		string(WorkflowStateSuspended),
	}
}

func (s *WorkflowState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWorkflowState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWorkflowState(input string) (*WorkflowState, error) {
	vals := map[string]WorkflowState{
		"completed":    WorkflowStateCompleted,
		"deleted":      WorkflowStateDeleted,
		"disabled":     WorkflowStateDisabled,
		"enabled":      WorkflowStateEnabled,
		"notspecified": WorkflowStateNotSpecified,
		"suspended":    WorkflowStateSuspended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkflowState(input)
	return &out, nil
}
