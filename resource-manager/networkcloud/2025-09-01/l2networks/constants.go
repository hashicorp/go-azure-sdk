package l2networks

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

type HybridAksPluginType string

const (
	HybridAksPluginTypeDPDK     HybridAksPluginType = "DPDK"
	HybridAksPluginTypeOSDevice HybridAksPluginType = "OSDevice"
	HybridAksPluginTypeSRIOV    HybridAksPluginType = "SRIOV"
)

func PossibleValuesForHybridAksPluginType() []string {
	return []string{
		string(HybridAksPluginTypeDPDK),
		string(HybridAksPluginTypeOSDevice),
		string(HybridAksPluginTypeSRIOV),
	}
}

func (s *HybridAksPluginType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHybridAksPluginType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHybridAksPluginType(input string) (*HybridAksPluginType, error) {
	vals := map[string]HybridAksPluginType{
		"dpdk":     HybridAksPluginTypeDPDK,
		"osdevice": HybridAksPluginTypeOSDevice,
		"sriov":    HybridAksPluginTypeSRIOV,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HybridAksPluginType(input)
	return &out, nil
}

type L2NetworkDetailedStatus string

const (
	L2NetworkDetailedStatusAvailable    L2NetworkDetailedStatus = "Available"
	L2NetworkDetailedStatusError        L2NetworkDetailedStatus = "Error"
	L2NetworkDetailedStatusProvisioning L2NetworkDetailedStatus = "Provisioning"
)

func PossibleValuesForL2NetworkDetailedStatus() []string {
	return []string{
		string(L2NetworkDetailedStatusAvailable),
		string(L2NetworkDetailedStatusError),
		string(L2NetworkDetailedStatusProvisioning),
	}
}

func (s *L2NetworkDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseL2NetworkDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseL2NetworkDetailedStatus(input string) (*L2NetworkDetailedStatus, error) {
	vals := map[string]L2NetworkDetailedStatus{
		"available":    L2NetworkDetailedStatusAvailable,
		"error":        L2NetworkDetailedStatusError,
		"provisioning": L2NetworkDetailedStatusProvisioning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := L2NetworkDetailedStatus(input)
	return &out, nil
}

type L2NetworkProvisioningState string

const (
	L2NetworkProvisioningStateAccepted     L2NetworkProvisioningState = "Accepted"
	L2NetworkProvisioningStateCanceled     L2NetworkProvisioningState = "Canceled"
	L2NetworkProvisioningStateFailed       L2NetworkProvisioningState = "Failed"
	L2NetworkProvisioningStateProvisioning L2NetworkProvisioningState = "Provisioning"
	L2NetworkProvisioningStateSucceeded    L2NetworkProvisioningState = "Succeeded"
)

func PossibleValuesForL2NetworkProvisioningState() []string {
	return []string{
		string(L2NetworkProvisioningStateAccepted),
		string(L2NetworkProvisioningStateCanceled),
		string(L2NetworkProvisioningStateFailed),
		string(L2NetworkProvisioningStateProvisioning),
		string(L2NetworkProvisioningStateSucceeded),
	}
}

func (s *L2NetworkProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseL2NetworkProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseL2NetworkProvisioningState(input string) (*L2NetworkProvisioningState, error) {
	vals := map[string]L2NetworkProvisioningState{
		"accepted":     L2NetworkProvisioningStateAccepted,
		"canceled":     L2NetworkProvisioningStateCanceled,
		"failed":       L2NetworkProvisioningStateFailed,
		"provisioning": L2NetworkProvisioningStateProvisioning,
		"succeeded":    L2NetworkProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := L2NetworkProvisioningState(input)
	return &out, nil
}
