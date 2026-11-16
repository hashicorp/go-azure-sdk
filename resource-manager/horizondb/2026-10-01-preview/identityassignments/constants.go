package identityassignments

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangeDataCaptureIdentityActivationState string

const (
	ChangeDataCaptureIdentityActivationStateActivating   ChangeDataCaptureIdentityActivationState = "Activating"
	ChangeDataCaptureIdentityActivationStateActive       ChangeDataCaptureIdentityActivationState = "Active"
	ChangeDataCaptureIdentityActivationStateDeactivating ChangeDataCaptureIdentityActivationState = "Deactivating"
	ChangeDataCaptureIdentityActivationStateFailed       ChangeDataCaptureIdentityActivationState = "Failed"
	ChangeDataCaptureIdentityActivationStateInactive     ChangeDataCaptureIdentityActivationState = "Inactive"
)

func PossibleValuesForChangeDataCaptureIdentityActivationState() []string {
	return []string{
		string(ChangeDataCaptureIdentityActivationStateActivating),
		string(ChangeDataCaptureIdentityActivationStateActive),
		string(ChangeDataCaptureIdentityActivationStateDeactivating),
		string(ChangeDataCaptureIdentityActivationStateFailed),
		string(ChangeDataCaptureIdentityActivationStateInactive),
	}
}

func (s *ChangeDataCaptureIdentityActivationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChangeDataCaptureIdentityActivationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChangeDataCaptureIdentityActivationState(input string) (*ChangeDataCaptureIdentityActivationState, error) {
	vals := map[string]ChangeDataCaptureIdentityActivationState{
		"activating":   ChangeDataCaptureIdentityActivationStateActivating,
		"active":       ChangeDataCaptureIdentityActivationStateActive,
		"deactivating": ChangeDataCaptureIdentityActivationStateDeactivating,
		"failed":       ChangeDataCaptureIdentityActivationStateFailed,
		"inactive":     ChangeDataCaptureIdentityActivationStateInactive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChangeDataCaptureIdentityActivationState(input)
	return &out, nil
}

type ChangeDataCaptureIdentityAssignmentState string

const (
	ChangeDataCaptureIdentityAssignmentStateAssigned  ChangeDataCaptureIdentityAssignmentState = "Assigned"
	ChangeDataCaptureIdentityAssignmentStateAssigning ChangeDataCaptureIdentityAssignmentState = "Assigning"
	ChangeDataCaptureIdentityAssignmentStateFailed    ChangeDataCaptureIdentityAssignmentState = "Failed"
	ChangeDataCaptureIdentityAssignmentStateRemoving  ChangeDataCaptureIdentityAssignmentState = "Removing"
)

func PossibleValuesForChangeDataCaptureIdentityAssignmentState() []string {
	return []string{
		string(ChangeDataCaptureIdentityAssignmentStateAssigned),
		string(ChangeDataCaptureIdentityAssignmentStateAssigning),
		string(ChangeDataCaptureIdentityAssignmentStateFailed),
		string(ChangeDataCaptureIdentityAssignmentStateRemoving),
	}
}

func (s *ChangeDataCaptureIdentityAssignmentState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChangeDataCaptureIdentityAssignmentState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChangeDataCaptureIdentityAssignmentState(input string) (*ChangeDataCaptureIdentityAssignmentState, error) {
	vals := map[string]ChangeDataCaptureIdentityAssignmentState{
		"assigned":  ChangeDataCaptureIdentityAssignmentStateAssigned,
		"assigning": ChangeDataCaptureIdentityAssignmentStateAssigning,
		"failed":    ChangeDataCaptureIdentityAssignmentStateFailed,
		"removing":  ChangeDataCaptureIdentityAssignmentStateRemoving,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChangeDataCaptureIdentityAssignmentState(input)
	return &out, nil
}

type ChangeDataCaptureIdentityValidationState string

const (
	ChangeDataCaptureIdentityValidationStateFailed     ChangeDataCaptureIdentityValidationState = "Failed"
	ChangeDataCaptureIdentityValidationStateNotStarted ChangeDataCaptureIdentityValidationState = "NotStarted"
	ChangeDataCaptureIdentityValidationStateRunning    ChangeDataCaptureIdentityValidationState = "Running"
	ChangeDataCaptureIdentityValidationStateSucceeded  ChangeDataCaptureIdentityValidationState = "Succeeded"
)

func PossibleValuesForChangeDataCaptureIdentityValidationState() []string {
	return []string{
		string(ChangeDataCaptureIdentityValidationStateFailed),
		string(ChangeDataCaptureIdentityValidationStateNotStarted),
		string(ChangeDataCaptureIdentityValidationStateRunning),
		string(ChangeDataCaptureIdentityValidationStateSucceeded),
	}
}

func (s *ChangeDataCaptureIdentityValidationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChangeDataCaptureIdentityValidationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChangeDataCaptureIdentityValidationState(input string) (*ChangeDataCaptureIdentityValidationState, error) {
	vals := map[string]ChangeDataCaptureIdentityValidationState{
		"failed":     ChangeDataCaptureIdentityValidationStateFailed,
		"notstarted": ChangeDataCaptureIdentityValidationStateNotStarted,
		"running":    ChangeDataCaptureIdentityValidationStateRunning,
		"succeeded":  ChangeDataCaptureIdentityValidationStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChangeDataCaptureIdentityValidationState(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateInProgress   ProvisioningState = "InProgress"
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateFailed),
		string(ProvisioningStateInProgress),
		string(ProvisioningStateProvisioning),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
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
		"canceled":     ProvisioningStateCanceled,
		"failed":       ProvisioningStateFailed,
		"inprogress":   ProvisioningStateInProgress,
		"provisioning": ProvisioningStateProvisioning,
		"succeeded":    ProvisioningStateSucceeded,
		"updating":     ProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}
