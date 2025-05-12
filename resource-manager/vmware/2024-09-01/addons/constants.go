package addons

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddonProvisioningState string

const (
	AddonProvisioningStateBuilding  AddonProvisioningState = "Building"
	AddonProvisioningStateCanceled  AddonProvisioningState = "Canceled"
	AddonProvisioningStateCancelled AddonProvisioningState = "Cancelled"
	AddonProvisioningStateDeleting  AddonProvisioningState = "Deleting"
	AddonProvisioningStateFailed    AddonProvisioningState = "Failed"
	AddonProvisioningStateSucceeded AddonProvisioningState = "Succeeded"
	AddonProvisioningStateUpdating  AddonProvisioningState = "Updating"
)

func PossibleValuesForAddonProvisioningState() []string {
	return []string{
		string(AddonProvisioningStateBuilding),
		string(AddonProvisioningStateCanceled),
		string(AddonProvisioningStateCancelled),
		string(AddonProvisioningStateDeleting),
		string(AddonProvisioningStateFailed),
		string(AddonProvisioningStateSucceeded),
		string(AddonProvisioningStateUpdating),
	}
}

func (s *AddonProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAddonProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAddonProvisioningState(input string) (*AddonProvisioningState, error) {
	vals := map[string]AddonProvisioningState{
		"building":  AddonProvisioningStateBuilding,
		"canceled":  AddonProvisioningStateCanceled,
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
	AddonTypeArc AddonType = "Arc"
	AddonTypeHCX AddonType = "HCX"
	AddonTypeSRM AddonType = "SRM"
	AddonTypeVR  AddonType = "VR"
)

func PossibleValuesForAddonType() []string {
	return []string{
		string(AddonTypeArc),
		string(AddonTypeHCX),
		string(AddonTypeSRM),
		string(AddonTypeVR),
	}
}

func (s *AddonType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAddonType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAddonType(input string) (*AddonType, error) {
	vals := map[string]AddonType{
		"arc": AddonTypeArc,
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
