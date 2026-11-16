package changedatacapturedestinations

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

type ChangeDataCaptureDestinationEffectiveState string

const (
	ChangeDataCaptureDestinationEffectiveStateDegraded           ChangeDataCaptureDestinationEffectiveState = "Degraded"
	ChangeDataCaptureDestinationEffectiveStateDisabled           ChangeDataCaptureDestinationEffectiveState = "Disabled"
	ChangeDataCaptureDestinationEffectiveStateDisabling          ChangeDataCaptureDestinationEffectiveState = "Disabling"
	ChangeDataCaptureDestinationEffectiveStateEnabled            ChangeDataCaptureDestinationEffectiveState = "Enabled"
	ChangeDataCaptureDestinationEffectiveStateEnabling           ChangeDataCaptureDestinationEffectiveState = "Enabling"
	ChangeDataCaptureDestinationEffectiveStateFailed             ChangeDataCaptureDestinationEffectiveState = "Failed"
	ChangeDataCaptureDestinationEffectiveStateWaitingForIdentity ChangeDataCaptureDestinationEffectiveState = "WaitingForIdentity"
)

func PossibleValuesForChangeDataCaptureDestinationEffectiveState() []string {
	return []string{
		string(ChangeDataCaptureDestinationEffectiveStateDegraded),
		string(ChangeDataCaptureDestinationEffectiveStateDisabled),
		string(ChangeDataCaptureDestinationEffectiveStateDisabling),
		string(ChangeDataCaptureDestinationEffectiveStateEnabled),
		string(ChangeDataCaptureDestinationEffectiveStateEnabling),
		string(ChangeDataCaptureDestinationEffectiveStateFailed),
		string(ChangeDataCaptureDestinationEffectiveStateWaitingForIdentity),
	}
}

func (s *ChangeDataCaptureDestinationEffectiveState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChangeDataCaptureDestinationEffectiveState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChangeDataCaptureDestinationEffectiveState(input string) (*ChangeDataCaptureDestinationEffectiveState, error) {
	vals := map[string]ChangeDataCaptureDestinationEffectiveState{
		"degraded":           ChangeDataCaptureDestinationEffectiveStateDegraded,
		"disabled":           ChangeDataCaptureDestinationEffectiveStateDisabled,
		"disabling":          ChangeDataCaptureDestinationEffectiveStateDisabling,
		"enabled":            ChangeDataCaptureDestinationEffectiveStateEnabled,
		"enabling":           ChangeDataCaptureDestinationEffectiveStateEnabling,
		"failed":             ChangeDataCaptureDestinationEffectiveStateFailed,
		"waitingforidentity": ChangeDataCaptureDestinationEffectiveStateWaitingForIdentity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChangeDataCaptureDestinationEffectiveState(input)
	return &out, nil
}

type ChangeDataCaptureDestinationType string

const (
	ChangeDataCaptureDestinationTypeFabric ChangeDataCaptureDestinationType = "Fabric"
)

func PossibleValuesForChangeDataCaptureDestinationType() []string {
	return []string{
		string(ChangeDataCaptureDestinationTypeFabric),
	}
}

func (s *ChangeDataCaptureDestinationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChangeDataCaptureDestinationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChangeDataCaptureDestinationType(input string) (*ChangeDataCaptureDestinationType, error) {
	vals := map[string]ChangeDataCaptureDestinationType{
		"fabric": ChangeDataCaptureDestinationTypeFabric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChangeDataCaptureDestinationType(input)
	return &out, nil
}

type ChangeDataCaptureProtocol string

const (
	ChangeDataCaptureProtocolNative ChangeDataCaptureProtocol = "Native"
)

func PossibleValuesForChangeDataCaptureProtocol() []string {
	return []string{
		string(ChangeDataCaptureProtocolNative),
	}
}

func (s *ChangeDataCaptureProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChangeDataCaptureProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChangeDataCaptureProtocol(input string) (*ChangeDataCaptureProtocol, error) {
	vals := map[string]ChangeDataCaptureProtocol{
		"native": ChangeDataCaptureProtocolNative,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChangeDataCaptureProtocol(input)
	return &out, nil
}

type Protocol string

const (
	ProtocolNative Protocol = "Native"
)

func PossibleValuesForProtocol() []string {
	return []string{
		string(ProtocolNative),
	}
}

func (s *Protocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProtocol(input string) (*Protocol, error) {
	vals := map[string]Protocol{
		"native": ProtocolNative,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Protocol(input)
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

type Status string

const (
	StatusSucceeded Status = "Succeeded"
)

func PossibleValuesForStatus() []string {
	return []string{
		string(StatusSucceeded),
	}
}

func (s *Status) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStatus(input string) (*Status, error) {
	vals := map[string]Status{
		"succeeded": StatusSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Status(input)
	return &out, nil
}
