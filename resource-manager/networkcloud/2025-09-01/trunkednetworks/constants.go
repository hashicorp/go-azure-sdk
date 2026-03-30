package trunkednetworks

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

type TrunkedNetworkDetailedStatus string

const (
	TrunkedNetworkDetailedStatusAvailable    TrunkedNetworkDetailedStatus = "Available"
	TrunkedNetworkDetailedStatusError        TrunkedNetworkDetailedStatus = "Error"
	TrunkedNetworkDetailedStatusProvisioning TrunkedNetworkDetailedStatus = "Provisioning"
)

func PossibleValuesForTrunkedNetworkDetailedStatus() []string {
	return []string{
		string(TrunkedNetworkDetailedStatusAvailable),
		string(TrunkedNetworkDetailedStatusError),
		string(TrunkedNetworkDetailedStatusProvisioning),
	}
}

func (s *TrunkedNetworkDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTrunkedNetworkDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTrunkedNetworkDetailedStatus(input string) (*TrunkedNetworkDetailedStatus, error) {
	vals := map[string]TrunkedNetworkDetailedStatus{
		"available":    TrunkedNetworkDetailedStatusAvailable,
		"error":        TrunkedNetworkDetailedStatusError,
		"provisioning": TrunkedNetworkDetailedStatusProvisioning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TrunkedNetworkDetailedStatus(input)
	return &out, nil
}

type TrunkedNetworkProvisioningState string

const (
	TrunkedNetworkProvisioningStateAccepted     TrunkedNetworkProvisioningState = "Accepted"
	TrunkedNetworkProvisioningStateCanceled     TrunkedNetworkProvisioningState = "Canceled"
	TrunkedNetworkProvisioningStateFailed       TrunkedNetworkProvisioningState = "Failed"
	TrunkedNetworkProvisioningStateProvisioning TrunkedNetworkProvisioningState = "Provisioning"
	TrunkedNetworkProvisioningStateSucceeded    TrunkedNetworkProvisioningState = "Succeeded"
)

func PossibleValuesForTrunkedNetworkProvisioningState() []string {
	return []string{
		string(TrunkedNetworkProvisioningStateAccepted),
		string(TrunkedNetworkProvisioningStateCanceled),
		string(TrunkedNetworkProvisioningStateFailed),
		string(TrunkedNetworkProvisioningStateProvisioning),
		string(TrunkedNetworkProvisioningStateSucceeded),
	}
}

func (s *TrunkedNetworkProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTrunkedNetworkProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTrunkedNetworkProvisioningState(input string) (*TrunkedNetworkProvisioningState, error) {
	vals := map[string]TrunkedNetworkProvisioningState{
		"accepted":     TrunkedNetworkProvisioningStateAccepted,
		"canceled":     TrunkedNetworkProvisioningStateCanceled,
		"failed":       TrunkedNetworkProvisioningStateFailed,
		"provisioning": TrunkedNetworkProvisioningStateProvisioning,
		"succeeded":    TrunkedNetworkProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TrunkedNetworkProvisioningState(input)
	return &out, nil
}
