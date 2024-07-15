package containerappssessionpools

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerType string

const (
	ContainerTypeCustomContainer ContainerType = "CustomContainer"
	ContainerTypePythonLTS       ContainerType = "PythonLTS"
)

func PossibleValuesForContainerType() []string {
	return []string{
		string(ContainerTypeCustomContainer),
		string(ContainerTypePythonLTS),
	}
}

func (s *ContainerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseContainerType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseContainerType(input string) (*ContainerType, error) {
	vals := map[string]ContainerType{
		"customcontainer": ContainerTypeCustomContainer,
		"pythonlts":       ContainerTypePythonLTS,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContainerType(input)
	return &out, nil
}

type ExecutionType string

const (
	ExecutionTypeTimed ExecutionType = "Timed"
)

func PossibleValuesForExecutionType() []string {
	return []string{
		string(ExecutionTypeTimed),
	}
}

func (s *ExecutionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExecutionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExecutionType(input string) (*ExecutionType, error) {
	vals := map[string]ExecutionType{
		"timed": ExecutionTypeTimed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExecutionType(input)
	return &out, nil
}

type PoolManagementType string

const (
	PoolManagementTypeDynamic PoolManagementType = "Dynamic"
	PoolManagementTypeManual  PoolManagementType = "Manual"
)

func PossibleValuesForPoolManagementType() []string {
	return []string{
		string(PoolManagementTypeDynamic),
		string(PoolManagementTypeManual),
	}
}

func (s *PoolManagementType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePoolManagementType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePoolManagementType(input string) (*PoolManagementType, error) {
	vals := map[string]PoolManagementType{
		"dynamic": PoolManagementTypeDynamic,
		"manual":  PoolManagementTypeManual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PoolManagementType(input)
	return &out, nil
}

type SessionNetworkStatus string

const (
	SessionNetworkStatusEgressDisabled SessionNetworkStatus = "EgressDisabled"
	SessionNetworkStatusEgressEnabled  SessionNetworkStatus = "EgressEnabled"
)

func PossibleValuesForSessionNetworkStatus() []string {
	return []string{
		string(SessionNetworkStatusEgressDisabled),
		string(SessionNetworkStatusEgressEnabled),
	}
}

func (s *SessionNetworkStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSessionNetworkStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSessionNetworkStatus(input string) (*SessionNetworkStatus, error) {
	vals := map[string]SessionNetworkStatus{
		"egressdisabled": SessionNetworkStatusEgressDisabled,
		"egressenabled":  SessionNetworkStatusEgressEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SessionNetworkStatus(input)
	return &out, nil
}

type SessionPoolProvisioningState string

const (
	SessionPoolProvisioningStateCanceled   SessionPoolProvisioningState = "Canceled"
	SessionPoolProvisioningStateDeleting   SessionPoolProvisioningState = "Deleting"
	SessionPoolProvisioningStateFailed     SessionPoolProvisioningState = "Failed"
	SessionPoolProvisioningStateInProgress SessionPoolProvisioningState = "InProgress"
	SessionPoolProvisioningStateSucceeded  SessionPoolProvisioningState = "Succeeded"
)

func PossibleValuesForSessionPoolProvisioningState() []string {
	return []string{
		string(SessionPoolProvisioningStateCanceled),
		string(SessionPoolProvisioningStateDeleting),
		string(SessionPoolProvisioningStateFailed),
		string(SessionPoolProvisioningStateInProgress),
		string(SessionPoolProvisioningStateSucceeded),
	}
}

func (s *SessionPoolProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSessionPoolProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSessionPoolProvisioningState(input string) (*SessionPoolProvisioningState, error) {
	vals := map[string]SessionPoolProvisioningState{
		"canceled":   SessionPoolProvisioningStateCanceled,
		"deleting":   SessionPoolProvisioningStateDeleting,
		"failed":     SessionPoolProvisioningStateFailed,
		"inprogress": SessionPoolProvisioningStateInProgress,
		"succeeded":  SessionPoolProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SessionPoolProvisioningState(input)
	return &out, nil
}
