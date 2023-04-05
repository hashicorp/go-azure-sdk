package addons

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddonProvisioningState string

const (
	AddonProvisioningStateBuilding  AddonProvisioningState = "Building"
	AddonProvisioningStateCancelled AddonProvisioningState = "Cancelled"
	AddonProvisioningStateDeleting  AddonProvisioningState = "Deleting"
	AddonProvisioningStateFailed    AddonProvisioningState = "Failed"
	AddonProvisioningStateSucceeded AddonProvisioningState = "Succeeded"
	AddonProvisioningStateUpdating  AddonProvisioningState = "Updating"
)

func PossibleValuesForAddonProvisioningState() []string {
	return []string{
		string(AddonProvisioningStateBuilding),
		string(AddonProvisioningStateCancelled),
		string(AddonProvisioningStateDeleting),
		string(AddonProvisioningStateFailed),
		string(AddonProvisioningStateSucceeded),
		string(AddonProvisioningStateUpdating),
	}
}

type AddonType string

const (
	AddonTypeHCX AddonType = "HCX"
	AddonTypeSRM AddonType = "SRM"
	AddonTypeVR  AddonType = "VR"
)

func PossibleValuesForAddonType() []string {
	return []string{
		string(AddonTypeHCX),
		string(AddonTypeSRM),
		string(AddonTypeVR),
	}
}
