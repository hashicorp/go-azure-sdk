package placementpolicies

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
