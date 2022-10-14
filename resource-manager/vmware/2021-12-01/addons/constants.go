package addons

import "strings"

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

func parseAddonProvisioningState(input string) (*AddonProvisioningState, error) {
	vals := map[string]AddonProvisioningState{
		"building":  AddonProvisioningStateBuilding,
		"cancelled": AddonProvisioningStateCancelled,
		"deleting":  AddonProvisioningStateDeleting,
		"failed":    AddonProvisioningStateFailed,
		"succeeded": AddonProvisioningStateSucceeded,
		"updating":  AddonProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AddonProvisioningState(input)
	return &out, nil
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

func parseAddonType(input string) (*AddonType, error) {
	vals := map[string]AddonType{
		"hcx": AddonTypeHCX,
		"srm": AddonTypeSRM,
		"vr":  AddonTypeVR,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AddonType(input)
	return &out, nil
}
