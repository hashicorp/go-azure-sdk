package placementpolicies

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AffinityStrength string

const (
	AffinityStrengthMust   AffinityStrength = "Must"
	AffinityStrengthShould AffinityStrength = "Should"
)

func PossibleValuesForAffinityStrength() []string {
	return []string{
		string(AffinityStrengthMust),
		string(AffinityStrengthShould),
	}
}

func (s *AffinityStrength) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAffinityStrength(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAffinityStrength(input string) (*AffinityStrength, error) {
	vals := map[string]AffinityStrength{
		"must":   AffinityStrengthMust,
		"should": AffinityStrengthShould,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AffinityStrength(input)
	return &out, nil
}

type AffinityType string

const (
	AffinityTypeAffinity     AffinityType = "Affinity"
	AffinityTypeAntiAffinity AffinityType = "AntiAffinity"
)

func PossibleValuesForAffinityType() []string {
	return []string{
		string(AffinityTypeAffinity),
		string(AffinityTypeAntiAffinity),
	}
}

func (s *AffinityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAffinityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAffinityType(input string) (*AffinityType, error) {
	vals := map[string]AffinityType{
		"affinity":     AffinityTypeAffinity,
		"antiaffinity": AffinityTypeAntiAffinity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AffinityType(input)
	return &out, nil
}

type AzureHybridBenefitType string

const (
	AzureHybridBenefitTypeNone    AzureHybridBenefitType = "None"
	AzureHybridBenefitTypeSqlHost AzureHybridBenefitType = "SqlHost"
)

func PossibleValuesForAzureHybridBenefitType() []string {
	return []string{
		string(AzureHybridBenefitTypeNone),
		string(AzureHybridBenefitTypeSqlHost),
	}
}

func (s *AzureHybridBenefitType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAzureHybridBenefitType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAzureHybridBenefitType(input string) (*AzureHybridBenefitType, error) {
	vals := map[string]AzureHybridBenefitType{
		"none":    AzureHybridBenefitTypeNone,
		"sqlhost": AzureHybridBenefitTypeSqlHost,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AzureHybridBenefitType(input)
	return &out, nil
}

type PlacementPolicyProvisioningState string

const (
	PlacementPolicyProvisioningStateBuilding  PlacementPolicyProvisioningState = "Building"
	PlacementPolicyProvisioningStateCanceled  PlacementPolicyProvisioningState = "Canceled"
	PlacementPolicyProvisioningStateDeleting  PlacementPolicyProvisioningState = "Deleting"
	PlacementPolicyProvisioningStateFailed    PlacementPolicyProvisioningState = "Failed"
	PlacementPolicyProvisioningStateSucceeded PlacementPolicyProvisioningState = "Succeeded"
	PlacementPolicyProvisioningStateUpdating  PlacementPolicyProvisioningState = "Updating"
)

func PossibleValuesForPlacementPolicyProvisioningState() []string {
	return []string{
		string(PlacementPolicyProvisioningStateBuilding),
		string(PlacementPolicyProvisioningStateCanceled),
		string(PlacementPolicyProvisioningStateDeleting),
		string(PlacementPolicyProvisioningStateFailed),
		string(PlacementPolicyProvisioningStateSucceeded),
		string(PlacementPolicyProvisioningStateUpdating),
	}
}

func (s *PlacementPolicyProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlacementPolicyProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlacementPolicyProvisioningState(input string) (*PlacementPolicyProvisioningState, error) {
	vals := map[string]PlacementPolicyProvisioningState{
		"building":  PlacementPolicyProvisioningStateBuilding,
		"canceled":  PlacementPolicyProvisioningStateCanceled,
		"deleting":  PlacementPolicyProvisioningStateDeleting,
		"failed":    PlacementPolicyProvisioningStateFailed,
		"succeeded": PlacementPolicyProvisioningStateSucceeded,
		"updating":  PlacementPolicyProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlacementPolicyProvisioningState(input)
	return &out, nil
}

type PlacementPolicyState string

const (
	PlacementPolicyStateDisabled PlacementPolicyState = "Disabled"
	PlacementPolicyStateEnabled  PlacementPolicyState = "Enabled"
)

func PossibleValuesForPlacementPolicyState() []string {
	return []string{
		string(PlacementPolicyStateDisabled),
		string(PlacementPolicyStateEnabled),
	}
}

func (s *PlacementPolicyState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlacementPolicyState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlacementPolicyState(input string) (*PlacementPolicyState, error) {
	vals := map[string]PlacementPolicyState{
		"disabled": PlacementPolicyStateDisabled,
		"enabled":  PlacementPolicyStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlacementPolicyState(input)
	return &out, nil
}

type PlacementPolicyType string

const (
	PlacementPolicyTypeVMHost PlacementPolicyType = "VmHost"
	PlacementPolicyTypeVMVM   PlacementPolicyType = "VmVm"
)

func PossibleValuesForPlacementPolicyType() []string {
	return []string{
		string(PlacementPolicyTypeVMHost),
		string(PlacementPolicyTypeVMVM),
	}
}

func (s *PlacementPolicyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlacementPolicyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlacementPolicyType(input string) (*PlacementPolicyType, error) {
	vals := map[string]PlacementPolicyType{
		"vmhost": PlacementPolicyTypeVMHost,
		"vmvm":   PlacementPolicyTypeVMVM,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlacementPolicyType(input)
	return &out, nil
}

type VirtualMachineRestrictMovementState string

const (
	VirtualMachineRestrictMovementStateDisabled VirtualMachineRestrictMovementState = "Disabled"
	VirtualMachineRestrictMovementStateEnabled  VirtualMachineRestrictMovementState = "Enabled"
)

func PossibleValuesForVirtualMachineRestrictMovementState() []string {
	return []string{
		string(VirtualMachineRestrictMovementStateDisabled),
		string(VirtualMachineRestrictMovementStateEnabled),
	}
}

func (s *VirtualMachineRestrictMovementState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualMachineRestrictMovementState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualMachineRestrictMovementState(input string) (*VirtualMachineRestrictMovementState, error) {
	vals := map[string]VirtualMachineRestrictMovementState{
		"disabled": VirtualMachineRestrictMovementStateDisabled,
		"enabled":  VirtualMachineRestrictMovementStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachineRestrictMovementState(input)
	return &out, nil
}
