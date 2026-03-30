package consoles

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConsoleDetailedStatus string

const (
	ConsoleDetailedStatusError ConsoleDetailedStatus = "Error"
	ConsoleDetailedStatusReady ConsoleDetailedStatus = "Ready"
)

func PossibleValuesForConsoleDetailedStatus() []string {
	return []string{
		string(ConsoleDetailedStatusError),
		string(ConsoleDetailedStatusReady),
	}
}

func (s *ConsoleDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConsoleDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConsoleDetailedStatus(input string) (*ConsoleDetailedStatus, error) {
	vals := map[string]ConsoleDetailedStatus{
		"error": ConsoleDetailedStatusError,
		"ready": ConsoleDetailedStatusReady,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConsoleDetailedStatus(input)
	return &out, nil
}

type ConsoleEnabled string

const (
	ConsoleEnabledFalse ConsoleEnabled = "False"
	ConsoleEnabledTrue  ConsoleEnabled = "True"
)

func PossibleValuesForConsoleEnabled() []string {
	return []string{
		string(ConsoleEnabledFalse),
		string(ConsoleEnabledTrue),
	}
}

func (s *ConsoleEnabled) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConsoleEnabled(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConsoleEnabled(input string) (*ConsoleEnabled, error) {
	vals := map[string]ConsoleEnabled{
		"false": ConsoleEnabledFalse,
		"true":  ConsoleEnabledTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConsoleEnabled(input)
	return &out, nil
}

type ConsoleProvisioningState string

const (
	ConsoleProvisioningStateAccepted     ConsoleProvisioningState = "Accepted"
	ConsoleProvisioningStateCanceled     ConsoleProvisioningState = "Canceled"
	ConsoleProvisioningStateFailed       ConsoleProvisioningState = "Failed"
	ConsoleProvisioningStateProvisioning ConsoleProvisioningState = "Provisioning"
	ConsoleProvisioningStateSucceeded    ConsoleProvisioningState = "Succeeded"
)

func PossibleValuesForConsoleProvisioningState() []string {
	return []string{
		string(ConsoleProvisioningStateAccepted),
		string(ConsoleProvisioningStateCanceled),
		string(ConsoleProvisioningStateFailed),
		string(ConsoleProvisioningStateProvisioning),
		string(ConsoleProvisioningStateSucceeded),
	}
}

func (s *ConsoleProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConsoleProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConsoleProvisioningState(input string) (*ConsoleProvisioningState, error) {
	vals := map[string]ConsoleProvisioningState{
		"accepted":     ConsoleProvisioningStateAccepted,
		"canceled":     ConsoleProvisioningStateCanceled,
		"failed":       ConsoleProvisioningStateFailed,
		"provisioning": ConsoleProvisioningStateProvisioning,
		"succeeded":    ConsoleProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConsoleProvisioningState(input)
	return &out, nil
}

type ExtendedLocationType string

const (
	ExtendedLocationTypeCustomLocation ExtendedLocationType = "CustomLocation"
	ExtendedLocationTypeEdgeZone       ExtendedLocationType = "EdgeZone"
)

func PossibleValuesForExtendedLocationType() []string {
	return []string{
		string(ExtendedLocationTypeCustomLocation),
		string(ExtendedLocationTypeEdgeZone),
	}
}

func (s *ExtendedLocationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExtendedLocationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExtendedLocationType(input string) (*ExtendedLocationType, error) {
	vals := map[string]ExtendedLocationType{
		"customlocation": ExtendedLocationTypeCustomLocation,
		"edgezone":       ExtendedLocationTypeEdgeZone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExtendedLocationType(input)
	return &out, nil
}
