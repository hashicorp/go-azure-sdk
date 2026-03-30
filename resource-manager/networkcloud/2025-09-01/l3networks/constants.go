package l3networks

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

type HybridAksIPamEnabled string

const (
	HybridAksIPamEnabledFalse HybridAksIPamEnabled = "False"
	HybridAksIPamEnabledTrue  HybridAksIPamEnabled = "True"
)

func PossibleValuesForHybridAksIPamEnabled() []string {
	return []string{
		string(HybridAksIPamEnabledFalse),
		string(HybridAksIPamEnabledTrue),
	}
}

func (s *HybridAksIPamEnabled) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHybridAksIPamEnabled(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHybridAksIPamEnabled(input string) (*HybridAksIPamEnabled, error) {
	vals := map[string]HybridAksIPamEnabled{
		"false": HybridAksIPamEnabledFalse,
		"true":  HybridAksIPamEnabledTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HybridAksIPamEnabled(input)
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

type IPAllocationType string

const (
	IPAllocationTypeDualStack IPAllocationType = "DualStack"
	IPAllocationTypeIPVFour   IPAllocationType = "IPV4"
	IPAllocationTypeIPVSix    IPAllocationType = "IPV6"
)

func PossibleValuesForIPAllocationType() []string {
	return []string{
		string(IPAllocationTypeDualStack),
		string(IPAllocationTypeIPVFour),
		string(IPAllocationTypeIPVSix),
	}
}

func (s *IPAllocationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIPAllocationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIPAllocationType(input string) (*IPAllocationType, error) {
	vals := map[string]IPAllocationType{
		"dualstack": IPAllocationTypeDualStack,
		"ipv4":      IPAllocationTypeIPVFour,
		"ipv6":      IPAllocationTypeIPVSix,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IPAllocationType(input)
	return &out, nil
}

type L3NetworkDetailedStatus string

const (
	L3NetworkDetailedStatusAvailable    L3NetworkDetailedStatus = "Available"
	L3NetworkDetailedStatusError        L3NetworkDetailedStatus = "Error"
	L3NetworkDetailedStatusProvisioning L3NetworkDetailedStatus = "Provisioning"
)

func PossibleValuesForL3NetworkDetailedStatus() []string {
	return []string{
		string(L3NetworkDetailedStatusAvailable),
		string(L3NetworkDetailedStatusError),
		string(L3NetworkDetailedStatusProvisioning),
	}
}

func (s *L3NetworkDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseL3NetworkDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseL3NetworkDetailedStatus(input string) (*L3NetworkDetailedStatus, error) {
	vals := map[string]L3NetworkDetailedStatus{
		"available":    L3NetworkDetailedStatusAvailable,
		"error":        L3NetworkDetailedStatusError,
		"provisioning": L3NetworkDetailedStatusProvisioning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := L3NetworkDetailedStatus(input)
	return &out, nil
}

type L3NetworkProvisioningState string

const (
	L3NetworkProvisioningStateAccepted     L3NetworkProvisioningState = "Accepted"
	L3NetworkProvisioningStateCanceled     L3NetworkProvisioningState = "Canceled"
	L3NetworkProvisioningStateFailed       L3NetworkProvisioningState = "Failed"
	L3NetworkProvisioningStateProvisioning L3NetworkProvisioningState = "Provisioning"
	L3NetworkProvisioningStateSucceeded    L3NetworkProvisioningState = "Succeeded"
)

func PossibleValuesForL3NetworkProvisioningState() []string {
	return []string{
		string(L3NetworkProvisioningStateAccepted),
		string(L3NetworkProvisioningStateCanceled),
		string(L3NetworkProvisioningStateFailed),
		string(L3NetworkProvisioningStateProvisioning),
		string(L3NetworkProvisioningStateSucceeded),
	}
}

func (s *L3NetworkProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseL3NetworkProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseL3NetworkProvisioningState(input string) (*L3NetworkProvisioningState, error) {
	vals := map[string]L3NetworkProvisioningState{
		"accepted":     L3NetworkProvisioningStateAccepted,
		"canceled":     L3NetworkProvisioningStateCanceled,
		"failed":       L3NetworkProvisioningStateFailed,
		"provisioning": L3NetworkProvisioningStateProvisioning,
		"succeeded":    L3NetworkProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := L3NetworkProvisioningState(input)
	return &out, nil
}
