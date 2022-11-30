package placementpolicies

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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

type PlacementPolicyProvisioningState string

const (
	PlacementPolicyProvisioningStateBuilding  PlacementPolicyProvisioningState = "Building"
	PlacementPolicyProvisioningStateDeleting  PlacementPolicyProvisioningState = "Deleting"
	PlacementPolicyProvisioningStateFailed    PlacementPolicyProvisioningState = "Failed"
	PlacementPolicyProvisioningStateSucceeded PlacementPolicyProvisioningState = "Succeeded"
	PlacementPolicyProvisioningStateUpdating  PlacementPolicyProvisioningState = "Updating"
)

func PossibleValuesForPlacementPolicyProvisioningState() []string {
	return []string{
		string(PlacementPolicyProvisioningStateBuilding),
		string(PlacementPolicyProvisioningStateDeleting),
		string(PlacementPolicyProvisioningStateFailed),
		string(PlacementPolicyProvisioningStateSucceeded),
		string(PlacementPolicyProvisioningStateUpdating),
	}
}

func parsePlacementPolicyProvisioningState(input string) (*PlacementPolicyProvisioningState, error) {
	vals := map[string]PlacementPolicyProvisioningState{
		"building":  PlacementPolicyProvisioningStateBuilding,
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
