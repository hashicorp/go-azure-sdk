package agentdeployment

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgentDeploymentProvisioningState string

const (
	AgentDeploymentProvisioningStateCanceled  AgentDeploymentProvisioningState = "Canceled"
	AgentDeploymentProvisioningStateCreating  AgentDeploymentProvisioningState = "Creating"
	AgentDeploymentProvisioningStateDeleting  AgentDeploymentProvisioningState = "Deleting"
	AgentDeploymentProvisioningStateFailed    AgentDeploymentProvisioningState = "Failed"
	AgentDeploymentProvisioningStateSucceeded AgentDeploymentProvisioningState = "Succeeded"
	AgentDeploymentProvisioningStateUpdating  AgentDeploymentProvisioningState = "Updating"
)

func PossibleValuesForAgentDeploymentProvisioningState() []string {
	return []string{
		string(AgentDeploymentProvisioningStateCanceled),
		string(AgentDeploymentProvisioningStateCreating),
		string(AgentDeploymentProvisioningStateDeleting),
		string(AgentDeploymentProvisioningStateFailed),
		string(AgentDeploymentProvisioningStateSucceeded),
		string(AgentDeploymentProvisioningStateUpdating),
	}
}

func (s *AgentDeploymentProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAgentDeploymentProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAgentDeploymentProvisioningState(input string) (*AgentDeploymentProvisioningState, error) {
	vals := map[string]AgentDeploymentProvisioningState{
		"canceled":  AgentDeploymentProvisioningStateCanceled,
		"creating":  AgentDeploymentProvisioningStateCreating,
		"deleting":  AgentDeploymentProvisioningStateDeleting,
		"failed":    AgentDeploymentProvisioningStateFailed,
		"succeeded": AgentDeploymentProvisioningStateSucceeded,
		"updating":  AgentDeploymentProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AgentDeploymentProvisioningState(input)
	return &out, nil
}

type AgentDeploymentState string

const (
	AgentDeploymentStateDeleted  AgentDeploymentState = "Deleted"
	AgentDeploymentStateDeleting AgentDeploymentState = "Deleting"
	AgentDeploymentStateFailed   AgentDeploymentState = "Failed"
	AgentDeploymentStateRunning  AgentDeploymentState = "Running"
	AgentDeploymentStateStarting AgentDeploymentState = "Starting"
	AgentDeploymentStateStopped  AgentDeploymentState = "Stopped"
	AgentDeploymentStateStopping AgentDeploymentState = "Stopping"
	AgentDeploymentStateUpdating AgentDeploymentState = "Updating"
)

func PossibleValuesForAgentDeploymentState() []string {
	return []string{
		string(AgentDeploymentStateDeleted),
		string(AgentDeploymentStateDeleting),
		string(AgentDeploymentStateFailed),
		string(AgentDeploymentStateRunning),
		string(AgentDeploymentStateStarting),
		string(AgentDeploymentStateStopped),
		string(AgentDeploymentStateStopping),
		string(AgentDeploymentStateUpdating),
	}
}

func (s *AgentDeploymentState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAgentDeploymentState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAgentDeploymentState(input string) (*AgentDeploymentState, error) {
	vals := map[string]AgentDeploymentState{
		"deleted":  AgentDeploymentStateDeleted,
		"deleting": AgentDeploymentStateDeleting,
		"failed":   AgentDeploymentStateFailed,
		"running":  AgentDeploymentStateRunning,
		"starting": AgentDeploymentStateStarting,
		"stopped":  AgentDeploymentStateStopped,
		"stopping": AgentDeploymentStateStopping,
		"updating": AgentDeploymentStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AgentDeploymentState(input)
	return &out, nil
}

type AgentDeploymentType string

const (
	AgentDeploymentTypeCustom  AgentDeploymentType = "Custom"
	AgentDeploymentTypeHosted  AgentDeploymentType = "Hosted"
	AgentDeploymentTypeManaged AgentDeploymentType = "Managed"
)

func PossibleValuesForAgentDeploymentType() []string {
	return []string{
		string(AgentDeploymentTypeCustom),
		string(AgentDeploymentTypeHosted),
		string(AgentDeploymentTypeManaged),
	}
}

func (s *AgentDeploymentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAgentDeploymentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAgentDeploymentType(input string) (*AgentDeploymentType, error) {
	vals := map[string]AgentDeploymentType{
		"custom":  AgentDeploymentTypeCustom,
		"hosted":  AgentDeploymentTypeHosted,
		"managed": AgentDeploymentTypeManaged,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AgentDeploymentType(input)
	return &out, nil
}

type AgentProtocol string

const (
	AgentProtocolATwoA     AgentProtocol = "A2A"
	AgentProtocolAgent     AgentProtocol = "Agent"
	AgentProtocolResponses AgentProtocol = "Responses"
)

func PossibleValuesForAgentProtocol() []string {
	return []string{
		string(AgentProtocolATwoA),
		string(AgentProtocolAgent),
		string(AgentProtocolResponses),
	}
}

func (s *AgentProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAgentProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAgentProtocol(input string) (*AgentProtocol, error) {
	vals := map[string]AgentProtocol{
		"a2a":       AgentProtocolATwoA,
		"agent":     AgentProtocolAgent,
		"responses": AgentProtocolResponses,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AgentProtocol(input)
	return &out, nil
}
