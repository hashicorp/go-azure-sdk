package baremetalmachinekeysets

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BareMetalMachineKeySetDetailedStatus string

const (
	BareMetalMachineKeySetDetailedStatusAllActive   BareMetalMachineKeySetDetailedStatus = "AllActive"
	BareMetalMachineKeySetDetailedStatusAllInvalid  BareMetalMachineKeySetDetailedStatus = "AllInvalid"
	BareMetalMachineKeySetDetailedStatusSomeInvalid BareMetalMachineKeySetDetailedStatus = "SomeInvalid"
	BareMetalMachineKeySetDetailedStatusValidating  BareMetalMachineKeySetDetailedStatus = "Validating"
)

func PossibleValuesForBareMetalMachineKeySetDetailedStatus() []string {
	return []string{
		string(BareMetalMachineKeySetDetailedStatusAllActive),
		string(BareMetalMachineKeySetDetailedStatusAllInvalid),
		string(BareMetalMachineKeySetDetailedStatusSomeInvalid),
		string(BareMetalMachineKeySetDetailedStatusValidating),
	}
}

func (s *BareMetalMachineKeySetDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBareMetalMachineKeySetDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBareMetalMachineKeySetDetailedStatus(input string) (*BareMetalMachineKeySetDetailedStatus, error) {
	vals := map[string]BareMetalMachineKeySetDetailedStatus{
		"allactive":   BareMetalMachineKeySetDetailedStatusAllActive,
		"allinvalid":  BareMetalMachineKeySetDetailedStatusAllInvalid,
		"someinvalid": BareMetalMachineKeySetDetailedStatusSomeInvalid,
		"validating":  BareMetalMachineKeySetDetailedStatusValidating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BareMetalMachineKeySetDetailedStatus(input)
	return &out, nil
}

type BareMetalMachineKeySetPrivilegeLevel string

const (
	BareMetalMachineKeySetPrivilegeLevelOther     BareMetalMachineKeySetPrivilegeLevel = "Other"
	BareMetalMachineKeySetPrivilegeLevelStandard  BareMetalMachineKeySetPrivilegeLevel = "Standard"
	BareMetalMachineKeySetPrivilegeLevelSuperuser BareMetalMachineKeySetPrivilegeLevel = "Superuser"
)

func PossibleValuesForBareMetalMachineKeySetPrivilegeLevel() []string {
	return []string{
		string(BareMetalMachineKeySetPrivilegeLevelOther),
		string(BareMetalMachineKeySetPrivilegeLevelStandard),
		string(BareMetalMachineKeySetPrivilegeLevelSuperuser),
	}
}

func (s *BareMetalMachineKeySetPrivilegeLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBareMetalMachineKeySetPrivilegeLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBareMetalMachineKeySetPrivilegeLevel(input string) (*BareMetalMachineKeySetPrivilegeLevel, error) {
	vals := map[string]BareMetalMachineKeySetPrivilegeLevel{
		"other":     BareMetalMachineKeySetPrivilegeLevelOther,
		"standard":  BareMetalMachineKeySetPrivilegeLevelStandard,
		"superuser": BareMetalMachineKeySetPrivilegeLevelSuperuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BareMetalMachineKeySetPrivilegeLevel(input)
	return &out, nil
}

type BareMetalMachineKeySetProvisioningState string

const (
	BareMetalMachineKeySetProvisioningStateAccepted     BareMetalMachineKeySetProvisioningState = "Accepted"
	BareMetalMachineKeySetProvisioningStateCanceled     BareMetalMachineKeySetProvisioningState = "Canceled"
	BareMetalMachineKeySetProvisioningStateFailed       BareMetalMachineKeySetProvisioningState = "Failed"
	BareMetalMachineKeySetProvisioningStateProvisioning BareMetalMachineKeySetProvisioningState = "Provisioning"
	BareMetalMachineKeySetProvisioningStateSucceeded    BareMetalMachineKeySetProvisioningState = "Succeeded"
)

func PossibleValuesForBareMetalMachineKeySetProvisioningState() []string {
	return []string{
		string(BareMetalMachineKeySetProvisioningStateAccepted),
		string(BareMetalMachineKeySetProvisioningStateCanceled),
		string(BareMetalMachineKeySetProvisioningStateFailed),
		string(BareMetalMachineKeySetProvisioningStateProvisioning),
		string(BareMetalMachineKeySetProvisioningStateSucceeded),
	}
}

func (s *BareMetalMachineKeySetProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBareMetalMachineKeySetProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBareMetalMachineKeySetProvisioningState(input string) (*BareMetalMachineKeySetProvisioningState, error) {
	vals := map[string]BareMetalMachineKeySetProvisioningState{
		"accepted":     BareMetalMachineKeySetProvisioningStateAccepted,
		"canceled":     BareMetalMachineKeySetProvisioningStateCanceled,
		"failed":       BareMetalMachineKeySetProvisioningStateFailed,
		"provisioning": BareMetalMachineKeySetProvisioningStateProvisioning,
		"succeeded":    BareMetalMachineKeySetProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BareMetalMachineKeySetProvisioningState(input)
	return &out, nil
}

type BareMetalMachineKeySetUserSetupStatus string

const (
	BareMetalMachineKeySetUserSetupStatusActive  BareMetalMachineKeySetUserSetupStatus = "Active"
	BareMetalMachineKeySetUserSetupStatusInvalid BareMetalMachineKeySetUserSetupStatus = "Invalid"
)

func PossibleValuesForBareMetalMachineKeySetUserSetupStatus() []string {
	return []string{
		string(BareMetalMachineKeySetUserSetupStatusActive),
		string(BareMetalMachineKeySetUserSetupStatusInvalid),
	}
}

func (s *BareMetalMachineKeySetUserSetupStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBareMetalMachineKeySetUserSetupStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBareMetalMachineKeySetUserSetupStatus(input string) (*BareMetalMachineKeySetUserSetupStatus, error) {
	vals := map[string]BareMetalMachineKeySetUserSetupStatus{
		"active":  BareMetalMachineKeySetUserSetupStatusActive,
		"invalid": BareMetalMachineKeySetUserSetupStatusInvalid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BareMetalMachineKeySetUserSetupStatus(input)
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
