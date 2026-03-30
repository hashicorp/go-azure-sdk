package storageappliances

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

type RemoteVendorManagementFeature string

const (
	RemoteVendorManagementFeatureSupported   RemoteVendorManagementFeature = "Supported"
	RemoteVendorManagementFeatureUnsupported RemoteVendorManagementFeature = "Unsupported"
)

func PossibleValuesForRemoteVendorManagementFeature() []string {
	return []string{
		string(RemoteVendorManagementFeatureSupported),
		string(RemoteVendorManagementFeatureUnsupported),
	}
}

func (s *RemoteVendorManagementFeature) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRemoteVendorManagementFeature(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRemoteVendorManagementFeature(input string) (*RemoteVendorManagementFeature, error) {
	vals := map[string]RemoteVendorManagementFeature{
		"supported":   RemoteVendorManagementFeatureSupported,
		"unsupported": RemoteVendorManagementFeatureUnsupported,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RemoteVendorManagementFeature(input)
	return &out, nil
}

type RemoteVendorManagementStatus string

const (
	RemoteVendorManagementStatusDisabled    RemoteVendorManagementStatus = "Disabled"
	RemoteVendorManagementStatusEnabled     RemoteVendorManagementStatus = "Enabled"
	RemoteVendorManagementStatusUnsupported RemoteVendorManagementStatus = "Unsupported"
)

func PossibleValuesForRemoteVendorManagementStatus() []string {
	return []string{
		string(RemoteVendorManagementStatusDisabled),
		string(RemoteVendorManagementStatusEnabled),
		string(RemoteVendorManagementStatusUnsupported),
	}
}

func (s *RemoteVendorManagementStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRemoteVendorManagementStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRemoteVendorManagementStatus(input string) (*RemoteVendorManagementStatus, error) {
	vals := map[string]RemoteVendorManagementStatus{
		"disabled":    RemoteVendorManagementStatusDisabled,
		"enabled":     RemoteVendorManagementStatusEnabled,
		"unsupported": RemoteVendorManagementStatusUnsupported,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RemoteVendorManagementStatus(input)
	return &out, nil
}

type StorageApplianceDetailedStatus string

const (
	StorageApplianceDetailedStatusAvailable    StorageApplianceDetailedStatus = "Available"
	StorageApplianceDetailedStatusDegraded     StorageApplianceDetailedStatus = "Degraded"
	StorageApplianceDetailedStatusError        StorageApplianceDetailedStatus = "Error"
	StorageApplianceDetailedStatusProvisioning StorageApplianceDetailedStatus = "Provisioning"
)

func PossibleValuesForStorageApplianceDetailedStatus() []string {
	return []string{
		string(StorageApplianceDetailedStatusAvailable),
		string(StorageApplianceDetailedStatusDegraded),
		string(StorageApplianceDetailedStatusError),
		string(StorageApplianceDetailedStatusProvisioning),
	}
}

func (s *StorageApplianceDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStorageApplianceDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStorageApplianceDetailedStatus(input string) (*StorageApplianceDetailedStatus, error) {
	vals := map[string]StorageApplianceDetailedStatus{
		"available":    StorageApplianceDetailedStatusAvailable,
		"degraded":     StorageApplianceDetailedStatusDegraded,
		"error":        StorageApplianceDetailedStatusError,
		"provisioning": StorageApplianceDetailedStatusProvisioning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StorageApplianceDetailedStatus(input)
	return &out, nil
}

type StorageApplianceProvisioningState string

const (
	StorageApplianceProvisioningStateAccepted     StorageApplianceProvisioningState = "Accepted"
	StorageApplianceProvisioningStateCanceled     StorageApplianceProvisioningState = "Canceled"
	StorageApplianceProvisioningStateFailed       StorageApplianceProvisioningState = "Failed"
	StorageApplianceProvisioningStateProvisioning StorageApplianceProvisioningState = "Provisioning"
	StorageApplianceProvisioningStateSucceeded    StorageApplianceProvisioningState = "Succeeded"
)

func PossibleValuesForStorageApplianceProvisioningState() []string {
	return []string{
		string(StorageApplianceProvisioningStateAccepted),
		string(StorageApplianceProvisioningStateCanceled),
		string(StorageApplianceProvisioningStateFailed),
		string(StorageApplianceProvisioningStateProvisioning),
		string(StorageApplianceProvisioningStateSucceeded),
	}
}

func (s *StorageApplianceProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStorageApplianceProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStorageApplianceProvisioningState(input string) (*StorageApplianceProvisioningState, error) {
	vals := map[string]StorageApplianceProvisioningState{
		"accepted":     StorageApplianceProvisioningStateAccepted,
		"canceled":     StorageApplianceProvisioningStateCanceled,
		"failed":       StorageApplianceProvisioningStateFailed,
		"provisioning": StorageApplianceProvisioningStateProvisioning,
		"succeeded":    StorageApplianceProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StorageApplianceProvisioningState(input)
	return &out, nil
}
