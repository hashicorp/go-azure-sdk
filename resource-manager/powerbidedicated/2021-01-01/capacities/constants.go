package capacities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CapacityProvisioningState string

const (
	CapacityProvisioningStateDeleting     CapacityProvisioningState = "Deleting"
	CapacityProvisioningStateFailed       CapacityProvisioningState = "Failed"
	CapacityProvisioningStatePaused       CapacityProvisioningState = "Paused"
	CapacityProvisioningStatePausing      CapacityProvisioningState = "Pausing"
	CapacityProvisioningStatePreparing    CapacityProvisioningState = "Preparing"
	CapacityProvisioningStateProvisioning CapacityProvisioningState = "Provisioning"
	CapacityProvisioningStateResuming     CapacityProvisioningState = "Resuming"
	CapacityProvisioningStateScaling      CapacityProvisioningState = "Scaling"
	CapacityProvisioningStateSucceeded    CapacityProvisioningState = "Succeeded"
	CapacityProvisioningStateSuspended    CapacityProvisioningState = "Suspended"
	CapacityProvisioningStateSuspending   CapacityProvisioningState = "Suspending"
	CapacityProvisioningStateUpdating     CapacityProvisioningState = "Updating"
)

func PossibleValuesForCapacityProvisioningState() []string {
	return []string{
		string(CapacityProvisioningStateDeleting),
		string(CapacityProvisioningStateFailed),
		string(CapacityProvisioningStatePaused),
		string(CapacityProvisioningStatePausing),
		string(CapacityProvisioningStatePreparing),
		string(CapacityProvisioningStateProvisioning),
		string(CapacityProvisioningStateResuming),
		string(CapacityProvisioningStateScaling),
		string(CapacityProvisioningStateSucceeded),
		string(CapacityProvisioningStateSuspended),
		string(CapacityProvisioningStateSuspending),
		string(CapacityProvisioningStateUpdating),
	}
}

type CapacitySkuTier string

const (
	CapacitySkuTierAutoPremiumHost CapacitySkuTier = "AutoPremiumHost"
	CapacitySkuTierPBIEAzure       CapacitySkuTier = "PBIE_Azure"
	CapacitySkuTierPremium         CapacitySkuTier = "Premium"
)

func PossibleValuesForCapacitySkuTier() []string {
	return []string{
		string(CapacitySkuTierAutoPremiumHost),
		string(CapacitySkuTierPBIEAzure),
		string(CapacitySkuTierPremium),
	}
}

type Mode string

const (
	ModeGenOne Mode = "Gen1"
	ModeGenTwo Mode = "Gen2"
)

func PossibleValuesForMode() []string {
	return []string{
		string(ModeGenOne),
		string(ModeGenTwo),
	}
}

type State string

const (
	StateDeleting     State = "Deleting"
	StateFailed       State = "Failed"
	StatePaused       State = "Paused"
	StatePausing      State = "Pausing"
	StatePreparing    State = "Preparing"
	StateProvisioning State = "Provisioning"
	StateResuming     State = "Resuming"
	StateScaling      State = "Scaling"
	StateSucceeded    State = "Succeeded"
	StateSuspended    State = "Suspended"
	StateSuspending   State = "Suspending"
	StateUpdating     State = "Updating"
)

func PossibleValuesForState() []string {
	return []string{
		string(StateDeleting),
		string(StateFailed),
		string(StatePaused),
		string(StatePausing),
		string(StatePreparing),
		string(StateProvisioning),
		string(StateResuming),
		string(StateScaling),
		string(StateSucceeded),
		string(StateSuspended),
		string(StateSuspending),
		string(StateUpdating),
	}
}
