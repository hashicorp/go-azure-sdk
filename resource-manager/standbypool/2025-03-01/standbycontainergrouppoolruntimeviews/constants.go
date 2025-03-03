package standbycontainergrouppoolruntimeviews

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

type PoolContainerGroupState string

const (
	PoolContainerGroupStateCreating PoolContainerGroupState = "Creating"
	PoolContainerGroupStateDeleting PoolContainerGroupState = "Deleting"
	PoolContainerGroupStateRunning  PoolContainerGroupState = "Running"
)

func PossibleValuesForPoolContainerGroupState() []string {
	return []string{
		string(PoolContainerGroupStateCreating),
		string(PoolContainerGroupStateDeleting),
		string(PoolContainerGroupStateRunning),
	}
}

func (s *PoolContainerGroupState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePoolContainerGroupState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePoolContainerGroupState(input string) (*PoolContainerGroupState, error) {
	vals := map[string]PoolContainerGroupState{
		"creating": PoolContainerGroupStateCreating,
		"deleting": PoolContainerGroupStateDeleting,
		"running":  PoolContainerGroupStateRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PoolContainerGroupState(input)
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
