package cloudservicesnetworks

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudServicesNetworkDetailedStatus string

const (
	CloudServicesNetworkDetailedStatusAvailable    CloudServicesNetworkDetailedStatus = "Available"
	CloudServicesNetworkDetailedStatusError        CloudServicesNetworkDetailedStatus = "Error"
	CloudServicesNetworkDetailedStatusProvisioning CloudServicesNetworkDetailedStatus = "Provisioning"
)

func PossibleValuesForCloudServicesNetworkDetailedStatus() []string {
	return []string{
		string(CloudServicesNetworkDetailedStatusAvailable),
		string(CloudServicesNetworkDetailedStatusError),
		string(CloudServicesNetworkDetailedStatusProvisioning),
	}
}

func (s *CloudServicesNetworkDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudServicesNetworkDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudServicesNetworkDetailedStatus(input string) (*CloudServicesNetworkDetailedStatus, error) {
	vals := map[string]CloudServicesNetworkDetailedStatus{
		"available":    CloudServicesNetworkDetailedStatusAvailable,
		"error":        CloudServicesNetworkDetailedStatusError,
		"provisioning": CloudServicesNetworkDetailedStatusProvisioning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudServicesNetworkDetailedStatus(input)
	return &out, nil
}

type CloudServicesNetworkEnableDefaultEgressEndpoints string

const (
	CloudServicesNetworkEnableDefaultEgressEndpointsFalse CloudServicesNetworkEnableDefaultEgressEndpoints = "False"
	CloudServicesNetworkEnableDefaultEgressEndpointsTrue  CloudServicesNetworkEnableDefaultEgressEndpoints = "True"
)

func PossibleValuesForCloudServicesNetworkEnableDefaultEgressEndpoints() []string {
	return []string{
		string(CloudServicesNetworkEnableDefaultEgressEndpointsFalse),
		string(CloudServicesNetworkEnableDefaultEgressEndpointsTrue),
	}
}

func (s *CloudServicesNetworkEnableDefaultEgressEndpoints) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudServicesNetworkEnableDefaultEgressEndpoints(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudServicesNetworkEnableDefaultEgressEndpoints(input string) (*CloudServicesNetworkEnableDefaultEgressEndpoints, error) {
	vals := map[string]CloudServicesNetworkEnableDefaultEgressEndpoints{
		"false": CloudServicesNetworkEnableDefaultEgressEndpointsFalse,
		"true":  CloudServicesNetworkEnableDefaultEgressEndpointsTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudServicesNetworkEnableDefaultEgressEndpoints(input)
	return &out, nil
}

type CloudServicesNetworkProvisioningState string

const (
	CloudServicesNetworkProvisioningStateAccepted     CloudServicesNetworkProvisioningState = "Accepted"
	CloudServicesNetworkProvisioningStateCanceled     CloudServicesNetworkProvisioningState = "Canceled"
	CloudServicesNetworkProvisioningStateFailed       CloudServicesNetworkProvisioningState = "Failed"
	CloudServicesNetworkProvisioningStateProvisioning CloudServicesNetworkProvisioningState = "Provisioning"
	CloudServicesNetworkProvisioningStateSucceeded    CloudServicesNetworkProvisioningState = "Succeeded"
)

func PossibleValuesForCloudServicesNetworkProvisioningState() []string {
	return []string{
		string(CloudServicesNetworkProvisioningStateAccepted),
		string(CloudServicesNetworkProvisioningStateCanceled),
		string(CloudServicesNetworkProvisioningStateFailed),
		string(CloudServicesNetworkProvisioningStateProvisioning),
		string(CloudServicesNetworkProvisioningStateSucceeded),
	}
}

func (s *CloudServicesNetworkProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudServicesNetworkProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudServicesNetworkProvisioningState(input string) (*CloudServicesNetworkProvisioningState, error) {
	vals := map[string]CloudServicesNetworkProvisioningState{
		"accepted":     CloudServicesNetworkProvisioningStateAccepted,
		"canceled":     CloudServicesNetworkProvisioningStateCanceled,
		"failed":       CloudServicesNetworkProvisioningStateFailed,
		"provisioning": CloudServicesNetworkProvisioningStateProvisioning,
		"succeeded":    CloudServicesNetworkProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudServicesNetworkProvisioningState(input)
	return &out, nil
}

type CloudServicesNetworkStorageMode string

const (
	CloudServicesNetworkStorageModeNone     CloudServicesNetworkStorageMode = "None"
	CloudServicesNetworkStorageModeStandard CloudServicesNetworkStorageMode = "Standard"
)

func PossibleValuesForCloudServicesNetworkStorageMode() []string {
	return []string{
		string(CloudServicesNetworkStorageModeNone),
		string(CloudServicesNetworkStorageModeStandard),
	}
}

func (s *CloudServicesNetworkStorageMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudServicesNetworkStorageMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudServicesNetworkStorageMode(input string) (*CloudServicesNetworkStorageMode, error) {
	vals := map[string]CloudServicesNetworkStorageMode{
		"none":     CloudServicesNetworkStorageModeNone,
		"standard": CloudServicesNetworkStorageModeStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudServicesNetworkStorageMode(input)
	return &out, nil
}

type CloudServicesNetworkStorageStatusStatus string

const (
	CloudServicesNetworkStorageStatusStatusAvailable       CloudServicesNetworkStorageStatusStatus = "Available"
	CloudServicesNetworkStorageStatusStatusExpandingVolume CloudServicesNetworkStorageStatusStatus = "ExpandingVolume"
	CloudServicesNetworkStorageStatusStatusExpansionFailed CloudServicesNetworkStorageStatusStatus = "ExpansionFailed"
)

func PossibleValuesForCloudServicesNetworkStorageStatusStatus() []string {
	return []string{
		string(CloudServicesNetworkStorageStatusStatusAvailable),
		string(CloudServicesNetworkStorageStatusStatusExpandingVolume),
		string(CloudServicesNetworkStorageStatusStatusExpansionFailed),
	}
}

func (s *CloudServicesNetworkStorageStatusStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudServicesNetworkStorageStatusStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudServicesNetworkStorageStatusStatus(input string) (*CloudServicesNetworkStorageStatusStatus, error) {
	vals := map[string]CloudServicesNetworkStorageStatusStatus{
		"available":       CloudServicesNetworkStorageStatusStatusAvailable,
		"expandingvolume": CloudServicesNetworkStorageStatusStatusExpandingVolume,
		"expansionfailed": CloudServicesNetworkStorageStatusStatusExpansionFailed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudServicesNetworkStorageStatusStatus(input)
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
