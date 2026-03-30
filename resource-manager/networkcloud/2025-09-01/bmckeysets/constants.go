package bmckeysets

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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

type BmcKeySetDetailedStatus string

const (
	BmcKeySetDetailedStatusAllActive   BmcKeySetDetailedStatus = "AllActive"
	BmcKeySetDetailedStatusAllInvalid  BmcKeySetDetailedStatus = "AllInvalid"
	BmcKeySetDetailedStatusSomeInvalid BmcKeySetDetailedStatus = "SomeInvalid"
	BmcKeySetDetailedStatusValidating  BmcKeySetDetailedStatus = "Validating"
)

func PossibleValuesForBmcKeySetDetailedStatus() []string {
	return []string{
		string(BmcKeySetDetailedStatusAllActive),
		string(BmcKeySetDetailedStatusAllInvalid),
		string(BmcKeySetDetailedStatusSomeInvalid),
		string(BmcKeySetDetailedStatusValidating),
	}
}

func (s *BmcKeySetDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBmcKeySetDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBmcKeySetDetailedStatus(input string) (*BmcKeySetDetailedStatus, error) {
	vals := map[string]BmcKeySetDetailedStatus{
		"allactive":   BmcKeySetDetailedStatusAllActive,
		"allinvalid":  BmcKeySetDetailedStatusAllInvalid,
		"someinvalid": BmcKeySetDetailedStatusSomeInvalid,
		"validating":  BmcKeySetDetailedStatusValidating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BmcKeySetDetailedStatus(input)
	return &out, nil
}

type BmcKeySetPrivilegeLevel string

const (
	BmcKeySetPrivilegeLevelAdministrator BmcKeySetPrivilegeLevel = "Administrator"
	BmcKeySetPrivilegeLevelReadOnly      BmcKeySetPrivilegeLevel = "ReadOnly"
)

func PossibleValuesForBmcKeySetPrivilegeLevel() []string {
	return []string{
		string(BmcKeySetPrivilegeLevelAdministrator),
		string(BmcKeySetPrivilegeLevelReadOnly),
	}
}

func (s *BmcKeySetPrivilegeLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBmcKeySetPrivilegeLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBmcKeySetPrivilegeLevel(input string) (*BmcKeySetPrivilegeLevel, error) {
	vals := map[string]BmcKeySetPrivilegeLevel{
		"administrator": BmcKeySetPrivilegeLevelAdministrator,
		"readonly":      BmcKeySetPrivilegeLevelReadOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BmcKeySetPrivilegeLevel(input)
	return &out, nil
}

type BmcKeySetProvisioningState string

const (
	BmcKeySetProvisioningStateAccepted     BmcKeySetProvisioningState = "Accepted"
	BmcKeySetProvisioningStateCanceled     BmcKeySetProvisioningState = "Canceled"
	BmcKeySetProvisioningStateFailed       BmcKeySetProvisioningState = "Failed"
	BmcKeySetProvisioningStateProvisioning BmcKeySetProvisioningState = "Provisioning"
	BmcKeySetProvisioningStateSucceeded    BmcKeySetProvisioningState = "Succeeded"
)

func PossibleValuesForBmcKeySetProvisioningState() []string {
	return []string{
		string(BmcKeySetProvisioningStateAccepted),
		string(BmcKeySetProvisioningStateCanceled),
		string(BmcKeySetProvisioningStateFailed),
		string(BmcKeySetProvisioningStateProvisioning),
		string(BmcKeySetProvisioningStateSucceeded),
	}
}

func (s *BmcKeySetProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBmcKeySetProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBmcKeySetProvisioningState(input string) (*BmcKeySetProvisioningState, error) {
	vals := map[string]BmcKeySetProvisioningState{
		"accepted":     BmcKeySetProvisioningStateAccepted,
		"canceled":     BmcKeySetProvisioningStateCanceled,
		"failed":       BmcKeySetProvisioningStateFailed,
		"provisioning": BmcKeySetProvisioningStateProvisioning,
		"succeeded":    BmcKeySetProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BmcKeySetProvisioningState(input)
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
