package changedatacaptureconfigurations

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

type ChangeDataCaptureReplicationState string

const (
	ChangeDataCaptureReplicationStateConnecting    ChangeDataCaptureReplicationState = "Connecting"
	ChangeDataCaptureReplicationStateDegraded      ChangeDataCaptureReplicationState = "Degraded"
	ChangeDataCaptureReplicationStateFailed        ChangeDataCaptureReplicationState = "Failed"
	ChangeDataCaptureReplicationStateNotApplicable ChangeDataCaptureReplicationState = "NotApplicable"
	ChangeDataCaptureReplicationStateNotConnected  ChangeDataCaptureReplicationState = "NotConnected"
	ChangeDataCaptureReplicationStatePaused        ChangeDataCaptureReplicationState = "Paused"
	ChangeDataCaptureReplicationStateReconnecting  ChangeDataCaptureReplicationState = "Reconnecting"
	ChangeDataCaptureReplicationStateReplicating   ChangeDataCaptureReplicationState = "Replicating"
	ChangeDataCaptureReplicationStateStopping      ChangeDataCaptureReplicationState = "Stopping"
)

func PossibleValuesForChangeDataCaptureReplicationState() []string {
	return []string{
		string(ChangeDataCaptureReplicationStateConnecting),
		string(ChangeDataCaptureReplicationStateDegraded),
		string(ChangeDataCaptureReplicationStateFailed),
		string(ChangeDataCaptureReplicationStateNotApplicable),
		string(ChangeDataCaptureReplicationStateNotConnected),
		string(ChangeDataCaptureReplicationStatePaused),
		string(ChangeDataCaptureReplicationStateReconnecting),
		string(ChangeDataCaptureReplicationStateReplicating),
		string(ChangeDataCaptureReplicationStateStopping),
	}
}

func (s *ChangeDataCaptureReplicationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChangeDataCaptureReplicationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChangeDataCaptureReplicationState(input string) (*ChangeDataCaptureReplicationState, error) {
	vals := map[string]ChangeDataCaptureReplicationState{
		"connecting":    ChangeDataCaptureReplicationStateConnecting,
		"degraded":      ChangeDataCaptureReplicationStateDegraded,
		"failed":        ChangeDataCaptureReplicationStateFailed,
		"notapplicable": ChangeDataCaptureReplicationStateNotApplicable,
		"notconnected":  ChangeDataCaptureReplicationStateNotConnected,
		"paused":        ChangeDataCaptureReplicationStatePaused,
		"reconnecting":  ChangeDataCaptureReplicationStateReconnecting,
		"replicating":   ChangeDataCaptureReplicationStateReplicating,
		"stopping":      ChangeDataCaptureReplicationStateStopping,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChangeDataCaptureReplicationState(input)
	return &out, nil
}

type ChangeDataCaptureTopologyBindingState string

const (
	ChangeDataCaptureTopologyBindingStateActive         ChangeDataCaptureTopologyBindingState = "Active"
	ChangeDataCaptureTopologyBindingStateConflict       ChangeDataCaptureTopologyBindingState = "Conflict"
	ChangeDataCaptureTopologyBindingStateRebindRequired ChangeDataCaptureTopologyBindingState = "RebindRequired"
	ChangeDataCaptureTopologyBindingStateReconciling    ChangeDataCaptureTopologyBindingState = "Reconciling"
	ChangeDataCaptureTopologyBindingStateSuperseded     ChangeDataCaptureTopologyBindingState = "Superseded"
)

func PossibleValuesForChangeDataCaptureTopologyBindingState() []string {
	return []string{
		string(ChangeDataCaptureTopologyBindingStateActive),
		string(ChangeDataCaptureTopologyBindingStateConflict),
		string(ChangeDataCaptureTopologyBindingStateRebindRequired),
		string(ChangeDataCaptureTopologyBindingStateReconciling),
		string(ChangeDataCaptureTopologyBindingStateSuperseded),
	}
}

func (s *ChangeDataCaptureTopologyBindingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChangeDataCaptureTopologyBindingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChangeDataCaptureTopologyBindingState(input string) (*ChangeDataCaptureTopologyBindingState, error) {
	vals := map[string]ChangeDataCaptureTopologyBindingState{
		"active":         ChangeDataCaptureTopologyBindingStateActive,
		"conflict":       ChangeDataCaptureTopologyBindingStateConflict,
		"rebindrequired": ChangeDataCaptureTopologyBindingStateRebindRequired,
		"reconciling":    ChangeDataCaptureTopologyBindingStateReconciling,
		"superseded":     ChangeDataCaptureTopologyBindingStateSuperseded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChangeDataCaptureTopologyBindingState(input)
	return &out, nil
}

type ChangeDataCaptureUsageProfile string

const (
	ChangeDataCaptureUsageProfileMirroring ChangeDataCaptureUsageProfile = "Mirroring"
)

func PossibleValuesForChangeDataCaptureUsageProfile() []string {
	return []string{
		string(ChangeDataCaptureUsageProfileMirroring),
	}
}

func (s *ChangeDataCaptureUsageProfile) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChangeDataCaptureUsageProfile(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChangeDataCaptureUsageProfile(input string) (*ChangeDataCaptureUsageProfile, error) {
	vals := map[string]ChangeDataCaptureUsageProfile{
		"mirroring": ChangeDataCaptureUsageProfileMirroring,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChangeDataCaptureUsageProfile(input)
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
