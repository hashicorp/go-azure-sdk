package hosts

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HostKind string

const (
	HostKindGeneral     HostKind = "General"
	HostKindSpecialized HostKind = "Specialized"
)

func PossibleValuesForHostKind() []string {
	return []string{
		string(HostKindGeneral),
		string(HostKindSpecialized),
	}
}

func (s *HostKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHostKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHostKind(input string) (*HostKind, error) {
	vals := map[string]HostKind{
		"general":     HostKindGeneral,
		"specialized": HostKindSpecialized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HostKind(input)
	return &out, nil
}

type HostMaintenance string

const (
	HostMaintenanceReplacement HostMaintenance = "Replacement"
	HostMaintenanceUpgrade     HostMaintenance = "Upgrade"
)

func PossibleValuesForHostMaintenance() []string {
	return []string{
		string(HostMaintenanceReplacement),
		string(HostMaintenanceUpgrade),
	}
}

func (s *HostMaintenance) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHostMaintenance(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHostMaintenance(input string) (*HostMaintenance, error) {
	vals := map[string]HostMaintenance{
		"replacement": HostMaintenanceReplacement,
		"upgrade":     HostMaintenanceUpgrade,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HostMaintenance(input)
	return &out, nil
}

type HostProvisioningState string

const (
	HostProvisioningStateCanceled  HostProvisioningState = "Canceled"
	HostProvisioningStateFailed    HostProvisioningState = "Failed"
	HostProvisioningStateSucceeded HostProvisioningState = "Succeeded"
)

func PossibleValuesForHostProvisioningState() []string {
	return []string{
		string(HostProvisioningStateCanceled),
		string(HostProvisioningStateFailed),
		string(HostProvisioningStateSucceeded),
	}
}

func (s *HostProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHostProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHostProvisioningState(input string) (*HostProvisioningState, error) {
	vals := map[string]HostProvisioningState{
		"canceled":  HostProvisioningStateCanceled,
		"failed":    HostProvisioningStateFailed,
		"succeeded": HostProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HostProvisioningState(input)
	return &out, nil
}

type SkuTier string

const (
	SkuTierBasic    SkuTier = "Basic"
	SkuTierFree     SkuTier = "Free"
	SkuTierPremium  SkuTier = "Premium"
	SkuTierStandard SkuTier = "Standard"
)

func PossibleValuesForSkuTier() []string {
	return []string{
		string(SkuTierBasic),
		string(SkuTierFree),
		string(SkuTierPremium),
		string(SkuTierStandard),
	}
}

func (s *SkuTier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSkuTier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSkuTier(input string) (*SkuTier, error) {
	vals := map[string]SkuTier{
		"basic":    SkuTierBasic,
		"free":     SkuTierFree,
		"premium":  SkuTierPremium,
		"standard": SkuTierStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuTier(input)
	return &out, nil
}
