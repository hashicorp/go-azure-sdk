package volumes

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

type VolumeDetailedStatus string

const (
	VolumeDetailedStatusActive       VolumeDetailedStatus = "Active"
	VolumeDetailedStatusError        VolumeDetailedStatus = "Error"
	VolumeDetailedStatusProvisioning VolumeDetailedStatus = "Provisioning"
)

func PossibleValuesForVolumeDetailedStatus() []string {
	return []string{
		string(VolumeDetailedStatusActive),
		string(VolumeDetailedStatusError),
		string(VolumeDetailedStatusProvisioning),
	}
}

func (s *VolumeDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVolumeDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVolumeDetailedStatus(input string) (*VolumeDetailedStatus, error) {
	vals := map[string]VolumeDetailedStatus{
		"active":       VolumeDetailedStatusActive,
		"error":        VolumeDetailedStatusError,
		"provisioning": VolumeDetailedStatusProvisioning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VolumeDetailedStatus(input)
	return &out, nil
}

type VolumeProvisioningState string

const (
	VolumeProvisioningStateAccepted     VolumeProvisioningState = "Accepted"
	VolumeProvisioningStateCanceled     VolumeProvisioningState = "Canceled"
	VolumeProvisioningStateFailed       VolumeProvisioningState = "Failed"
	VolumeProvisioningStateProvisioning VolumeProvisioningState = "Provisioning"
	VolumeProvisioningStateSucceeded    VolumeProvisioningState = "Succeeded"
)

func PossibleValuesForVolumeProvisioningState() []string {
	return []string{
		string(VolumeProvisioningStateAccepted),
		string(VolumeProvisioningStateCanceled),
		string(VolumeProvisioningStateFailed),
		string(VolumeProvisioningStateProvisioning),
		string(VolumeProvisioningStateSucceeded),
	}
}

func (s *VolumeProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVolumeProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVolumeProvisioningState(input string) (*VolumeProvisioningState, error) {
	vals := map[string]VolumeProvisioningState{
		"accepted":     VolumeProvisioningStateAccepted,
		"canceled":     VolumeProvisioningStateCanceled,
		"failed":       VolumeProvisioningStateFailed,
		"provisioning": VolumeProvisioningStateProvisioning,
		"succeeded":    VolumeProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VolumeProvisioningState(input)
	return &out, nil
}
