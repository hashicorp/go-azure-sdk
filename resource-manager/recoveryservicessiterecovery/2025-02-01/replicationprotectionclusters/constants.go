package replicationprotectionclusters

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutoProtectionOfDataDisk string

const (
	AutoProtectionOfDataDiskDisabled AutoProtectionOfDataDisk = "Disabled"
	AutoProtectionOfDataDiskEnabled  AutoProtectionOfDataDisk = "Enabled"
)

func PossibleValuesForAutoProtectionOfDataDisk() []string {
	return []string{
		string(AutoProtectionOfDataDiskDisabled),
		string(AutoProtectionOfDataDiskEnabled),
	}
}

func (s *AutoProtectionOfDataDisk) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAutoProtectionOfDataDisk(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAutoProtectionOfDataDisk(input string) (*AutoProtectionOfDataDisk, error) {
	vals := map[string]AutoProtectionOfDataDisk{
		"disabled": AutoProtectionOfDataDiskDisabled,
		"enabled":  AutoProtectionOfDataDiskEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutoProtectionOfDataDisk(input)
	return &out, nil
}

type FailoverDirection string

const (
	FailoverDirectionPrimaryToRecovery FailoverDirection = "PrimaryToRecovery"
	FailoverDirectionRecoveryToPrimary FailoverDirection = "RecoveryToPrimary"
)

func PossibleValuesForFailoverDirection() []string {
	return []string{
		string(FailoverDirectionPrimaryToRecovery),
		string(FailoverDirectionRecoveryToPrimary),
	}
}

func (s *FailoverDirection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFailoverDirection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFailoverDirection(input string) (*FailoverDirection, error) {
	vals := map[string]FailoverDirection{
		"primarytorecovery": FailoverDirectionPrimaryToRecovery,
		"recoverytoprimary": FailoverDirectionRecoveryToPrimary,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FailoverDirection(input)
	return &out, nil
}

type HealthErrorCustomerResolvability string

const (
	HealthErrorCustomerResolvabilityAllowed    HealthErrorCustomerResolvability = "Allowed"
	HealthErrorCustomerResolvabilityNotAllowed HealthErrorCustomerResolvability = "NotAllowed"
)

func PossibleValuesForHealthErrorCustomerResolvability() []string {
	return []string{
		string(HealthErrorCustomerResolvabilityAllowed),
		string(HealthErrorCustomerResolvabilityNotAllowed),
	}
}

func (s *HealthErrorCustomerResolvability) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHealthErrorCustomerResolvability(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHealthErrorCustomerResolvability(input string) (*HealthErrorCustomerResolvability, error) {
	vals := map[string]HealthErrorCustomerResolvability{
		"allowed":    HealthErrorCustomerResolvabilityAllowed,
		"notallowed": HealthErrorCustomerResolvabilityNotAllowed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HealthErrorCustomerResolvability(input)
	return &out, nil
}

type MultiVMGroupCreateOption string

const (
	MultiVMGroupCreateOptionAutoCreated   MultiVMGroupCreateOption = "AutoCreated"
	MultiVMGroupCreateOptionUserSpecified MultiVMGroupCreateOption = "UserSpecified"
)

func PossibleValuesForMultiVMGroupCreateOption() []string {
	return []string{
		string(MultiVMGroupCreateOptionAutoCreated),
		string(MultiVMGroupCreateOptionUserSpecified),
	}
}

func (s *MultiVMGroupCreateOption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMultiVMGroupCreateOption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMultiVMGroupCreateOption(input string) (*MultiVMGroupCreateOption, error) {
	vals := map[string]MultiVMGroupCreateOption{
		"autocreated":   MultiVMGroupCreateOptionAutoCreated,
		"userspecified": MultiVMGroupCreateOptionUserSpecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MultiVMGroupCreateOption(input)
	return &out, nil
}
