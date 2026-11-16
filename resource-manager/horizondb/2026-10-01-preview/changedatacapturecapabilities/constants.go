package changedatacapturecapabilities

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangeDataCaptureDesiredState string

const (
	ChangeDataCaptureDesiredStateDisabled ChangeDataCaptureDesiredState = "Disabled"
	ChangeDataCaptureDesiredStateEnabled  ChangeDataCaptureDesiredState = "Enabled"
)

func PossibleValuesForChangeDataCaptureDesiredState() []string {
	return []string{
		string(ChangeDataCaptureDesiredStateDisabled),
		string(ChangeDataCaptureDesiredStateEnabled),
	}
}

func (s *ChangeDataCaptureDesiredState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChangeDataCaptureDesiredState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChangeDataCaptureDesiredState(input string) (*ChangeDataCaptureDesiredState, error) {
	vals := map[string]ChangeDataCaptureDesiredState{
		"disabled": ChangeDataCaptureDesiredStateDisabled,
		"enabled":  ChangeDataCaptureDesiredStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChangeDataCaptureDesiredState(input)
	return &out, nil
}

type ChangeDataCaptureEffectiveState string

const (
	ChangeDataCaptureEffectiveStateDegraded  ChangeDataCaptureEffectiveState = "Degraded"
	ChangeDataCaptureEffectiveStateDisabled  ChangeDataCaptureEffectiveState = "Disabled"
	ChangeDataCaptureEffectiveStateDisabling ChangeDataCaptureEffectiveState = "Disabling"
	ChangeDataCaptureEffectiveStateEnabled   ChangeDataCaptureEffectiveState = "Enabled"
	ChangeDataCaptureEffectiveStateEnabling  ChangeDataCaptureEffectiveState = "Enabling"
	ChangeDataCaptureEffectiveStateFailed    ChangeDataCaptureEffectiveState = "Failed"
)

func PossibleValuesForChangeDataCaptureEffectiveState() []string {
	return []string{
		string(ChangeDataCaptureEffectiveStateDegraded),
		string(ChangeDataCaptureEffectiveStateDisabled),
		string(ChangeDataCaptureEffectiveStateDisabling),
		string(ChangeDataCaptureEffectiveStateEnabled),
		string(ChangeDataCaptureEffectiveStateEnabling),
		string(ChangeDataCaptureEffectiveStateFailed),
	}
}

func (s *ChangeDataCaptureEffectiveState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChangeDataCaptureEffectiveState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChangeDataCaptureEffectiveState(input string) (*ChangeDataCaptureEffectiveState, error) {
	vals := map[string]ChangeDataCaptureEffectiveState{
		"degraded":  ChangeDataCaptureEffectiveStateDegraded,
		"disabled":  ChangeDataCaptureEffectiveStateDisabled,
		"disabling": ChangeDataCaptureEffectiveStateDisabling,
		"enabled":   ChangeDataCaptureEffectiveStateEnabled,
		"enabling":  ChangeDataCaptureEffectiveStateEnabling,
		"failed":    ChangeDataCaptureEffectiveStateFailed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChangeDataCaptureEffectiveState(input)
	return &out, nil
}

type ChangeDataCaptureEngineConfigurationState string

const (
	ChangeDataCaptureEngineConfigurationStateFailed      ChangeDataCaptureEngineConfigurationState = "Failed"
	ChangeDataCaptureEngineConfigurationStateNotReady    ChangeDataCaptureEngineConfigurationState = "NotReady"
	ChangeDataCaptureEngineConfigurationStateReady       ChangeDataCaptureEngineConfigurationState = "Ready"
	ChangeDataCaptureEngineConfigurationStateReconciling ChangeDataCaptureEngineConfigurationState = "Reconciling"
)

func PossibleValuesForChangeDataCaptureEngineConfigurationState() []string {
	return []string{
		string(ChangeDataCaptureEngineConfigurationStateFailed),
		string(ChangeDataCaptureEngineConfigurationStateNotReady),
		string(ChangeDataCaptureEngineConfigurationStateReady),
		string(ChangeDataCaptureEngineConfigurationStateReconciling),
	}
}

func (s *ChangeDataCaptureEngineConfigurationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChangeDataCaptureEngineConfigurationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChangeDataCaptureEngineConfigurationState(input string) (*ChangeDataCaptureEngineConfigurationState, error) {
	vals := map[string]ChangeDataCaptureEngineConfigurationState{
		"failed":      ChangeDataCaptureEngineConfigurationStateFailed,
		"notready":    ChangeDataCaptureEngineConfigurationStateNotReady,
		"ready":       ChangeDataCaptureEngineConfigurationStateReady,
		"reconciling": ChangeDataCaptureEngineConfigurationStateReconciling,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChangeDataCaptureEngineConfigurationState(input)
	return &out, nil
}

type ChangeDataCaptureExtensionState string

const (
	ChangeDataCaptureExtensionStateFailed       ChangeDataCaptureExtensionState = "Failed"
	ChangeDataCaptureExtensionStateInstalling   ChangeDataCaptureExtensionState = "Installing"
	ChangeDataCaptureExtensionStateNotInstalled ChangeDataCaptureExtensionState = "NotInstalled"
	ChangeDataCaptureExtensionStateReady        ChangeDataCaptureExtensionState = "Ready"
	ChangeDataCaptureExtensionStateRemoving     ChangeDataCaptureExtensionState = "Removing"
	ChangeDataCaptureExtensionStateUpgrading    ChangeDataCaptureExtensionState = "Upgrading"
)

func PossibleValuesForChangeDataCaptureExtensionState() []string {
	return []string{
		string(ChangeDataCaptureExtensionStateFailed),
		string(ChangeDataCaptureExtensionStateInstalling),
		string(ChangeDataCaptureExtensionStateNotInstalled),
		string(ChangeDataCaptureExtensionStateReady),
		string(ChangeDataCaptureExtensionStateRemoving),
		string(ChangeDataCaptureExtensionStateUpgrading),
	}
}

func (s *ChangeDataCaptureExtensionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChangeDataCaptureExtensionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChangeDataCaptureExtensionState(input string) (*ChangeDataCaptureExtensionState, error) {
	vals := map[string]ChangeDataCaptureExtensionState{
		"failed":       ChangeDataCaptureExtensionStateFailed,
		"installing":   ChangeDataCaptureExtensionStateInstalling,
		"notinstalled": ChangeDataCaptureExtensionStateNotInstalled,
		"ready":        ChangeDataCaptureExtensionStateReady,
		"removing":     ChangeDataCaptureExtensionStateRemoving,
		"upgrading":    ChangeDataCaptureExtensionStateUpgrading,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChangeDataCaptureExtensionState(input)
	return &out, nil
}

type ChangeDataCaptureReadWriteEndpointState string

const (
	ChangeDataCaptureReadWriteEndpointStateReady       ChangeDataCaptureReadWriteEndpointState = "Ready"
	ChangeDataCaptureReadWriteEndpointStateUnavailable ChangeDataCaptureReadWriteEndpointState = "Unavailable"
)

func PossibleValuesForChangeDataCaptureReadWriteEndpointState() []string {
	return []string{
		string(ChangeDataCaptureReadWriteEndpointStateReady),
		string(ChangeDataCaptureReadWriteEndpointStateUnavailable),
	}
}

func (s *ChangeDataCaptureReadWriteEndpointState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChangeDataCaptureReadWriteEndpointState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChangeDataCaptureReadWriteEndpointState(input string) (*ChangeDataCaptureReadWriteEndpointState, error) {
	vals := map[string]ChangeDataCaptureReadWriteEndpointState{
		"ready":       ChangeDataCaptureReadWriteEndpointStateReady,
		"unavailable": ChangeDataCaptureReadWriteEndpointStateUnavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChangeDataCaptureReadWriteEndpointState(input)
	return &out, nil
}

type ChangeDataCaptureReplicationArchitecture string

const (
	ChangeDataCaptureReplicationArchitectureComputeBased ChangeDataCaptureReplicationArchitecture = "ComputeBased"
	ChangeDataCaptureReplicationArchitectureHybrid       ChangeDataCaptureReplicationArchitecture = "Hybrid"
	ChangeDataCaptureReplicationArchitectureStorageBased ChangeDataCaptureReplicationArchitecture = "StorageBased"
)

func PossibleValuesForChangeDataCaptureReplicationArchitecture() []string {
	return []string{
		string(ChangeDataCaptureReplicationArchitectureComputeBased),
		string(ChangeDataCaptureReplicationArchitectureHybrid),
		string(ChangeDataCaptureReplicationArchitectureStorageBased),
	}
}

func (s *ChangeDataCaptureReplicationArchitecture) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChangeDataCaptureReplicationArchitecture(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChangeDataCaptureReplicationArchitecture(input string) (*ChangeDataCaptureReplicationArchitecture, error) {
	vals := map[string]ChangeDataCaptureReplicationArchitecture{
		"computebased": ChangeDataCaptureReplicationArchitectureComputeBased,
		"hybrid":       ChangeDataCaptureReplicationArchitectureHybrid,
		"storagebased": ChangeDataCaptureReplicationArchitectureStorageBased,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChangeDataCaptureReplicationArchitecture(input)
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
