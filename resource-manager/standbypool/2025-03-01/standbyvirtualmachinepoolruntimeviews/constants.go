package standbyvirtualmachinepoolruntimeviews

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HealthStateCode string

const (
	HealthStateCodeHealthStateDegraded HealthStateCode = "HealthState/degraded"
	HealthStateCodeHealthStateHealthy  HealthStateCode = "HealthState/healthy"
)

func PossibleValuesForHealthStateCode() []string {
	return []string{
		string(HealthStateCodeHealthStateDegraded),
		string(HealthStateCodeHealthStateHealthy),
	}
}

func (s *HealthStateCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHealthStateCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHealthStateCode(input string) (*HealthStateCode, error) {
	vals := map[string]HealthStateCode{
		"healthstate/degraded": HealthStateCodeHealthStateDegraded,
		"healthstate/healthy":  HealthStateCodeHealthStateHealthy,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HealthStateCode(input)
	return &out, nil
}

type PoolVirtualMachineState string

const (
	PoolVirtualMachineStateCreating     PoolVirtualMachineState = "Creating"
	PoolVirtualMachineStateDeallocated  PoolVirtualMachineState = "Deallocated"
	PoolVirtualMachineStateDeallocating PoolVirtualMachineState = "Deallocating"
	PoolVirtualMachineStateDeleting     PoolVirtualMachineState = "Deleting"
	PoolVirtualMachineStateHibernated   PoolVirtualMachineState = "Hibernated"
	PoolVirtualMachineStateHibernating  PoolVirtualMachineState = "Hibernating"
	PoolVirtualMachineStateRunning      PoolVirtualMachineState = "Running"
	PoolVirtualMachineStateStarting     PoolVirtualMachineState = "Starting"
)

func PossibleValuesForPoolVirtualMachineState() []string {
	return []string{
		string(PoolVirtualMachineStateCreating),
		string(PoolVirtualMachineStateDeallocated),
		string(PoolVirtualMachineStateDeallocating),
		string(PoolVirtualMachineStateDeleting),
		string(PoolVirtualMachineStateHibernated),
		string(PoolVirtualMachineStateHibernating),
		string(PoolVirtualMachineStateRunning),
		string(PoolVirtualMachineStateStarting),
	}
}

func (s *PoolVirtualMachineState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePoolVirtualMachineState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePoolVirtualMachineState(input string) (*PoolVirtualMachineState, error) {
	vals := map[string]PoolVirtualMachineState{
		"creating":     PoolVirtualMachineStateCreating,
		"deallocated":  PoolVirtualMachineStateDeallocated,
		"deallocating": PoolVirtualMachineStateDeallocating,
		"deleting":     PoolVirtualMachineStateDeleting,
		"hibernated":   PoolVirtualMachineStateHibernated,
		"hibernating":  PoolVirtualMachineStateHibernating,
		"running":      PoolVirtualMachineStateRunning,
		"starting":     PoolVirtualMachineStateStarting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PoolVirtualMachineState(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
	}
}

func (s *ProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"canceled":  ProvisioningStateCanceled,
		"deleting":  ProvisioningStateDeleting,
		"failed":    ProvisioningStateFailed,
		"succeeded": ProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}
