package racks

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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

type RackDetailedStatus string

const (
	RackDetailedStatusAvailable    RackDetailedStatus = "Available"
	RackDetailedStatusError        RackDetailedStatus = "Error"
	RackDetailedStatusProvisioning RackDetailedStatus = "Provisioning"
)

func PossibleValuesForRackDetailedStatus() []string {
	return []string{
		string(RackDetailedStatusAvailable),
		string(RackDetailedStatusError),
		string(RackDetailedStatusProvisioning),
	}
}

func (s *RackDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRackDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRackDetailedStatus(input string) (*RackDetailedStatus, error) {
	vals := map[string]RackDetailedStatus{
		"available":    RackDetailedStatusAvailable,
		"error":        RackDetailedStatusError,
		"provisioning": RackDetailedStatusProvisioning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RackDetailedStatus(input)
	return &out, nil
}

type RackProvisioningState string

const (
	RackProvisioningStateAccepted     RackProvisioningState = "Accepted"
	RackProvisioningStateCanceled     RackProvisioningState = "Canceled"
	RackProvisioningStateFailed       RackProvisioningState = "Failed"
	RackProvisioningStateProvisioning RackProvisioningState = "Provisioning"
	RackProvisioningStateSucceeded    RackProvisioningState = "Succeeded"
)

func PossibleValuesForRackProvisioningState() []string {
	return []string{
		string(RackProvisioningStateAccepted),
		string(RackProvisioningStateCanceled),
		string(RackProvisioningStateFailed),
		string(RackProvisioningStateProvisioning),
		string(RackProvisioningStateSucceeded),
	}
}

func (s *RackProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRackProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRackProvisioningState(input string) (*RackProvisioningState, error) {
	vals := map[string]RackProvisioningState{
		"accepted":     RackProvisioningStateAccepted,
		"canceled":     RackProvisioningStateCanceled,
		"failed":       RackProvisioningStateFailed,
		"provisioning": RackProvisioningStateProvisioning,
		"succeeded":    RackProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RackProvisioningState(input)
	return &out, nil
}
