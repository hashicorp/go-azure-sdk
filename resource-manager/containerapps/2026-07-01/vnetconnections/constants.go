package vnetconnections

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VnetConnectionProvisioningState string

const (
	VnetConnectionProvisioningStateCanceled   VnetConnectionProvisioningState = "Canceled"
	VnetConnectionProvisioningStateDeleting   VnetConnectionProvisioningState = "Deleting"
	VnetConnectionProvisioningStateFailed     VnetConnectionProvisioningState = "Failed"
	VnetConnectionProvisioningStateInProgress VnetConnectionProvisioningState = "InProgress"
	VnetConnectionProvisioningStateSucceeded  VnetConnectionProvisioningState = "Succeeded"
	VnetConnectionProvisioningStateUpdating   VnetConnectionProvisioningState = "Updating"
)

func PossibleValuesForVnetConnectionProvisioningState() []string {
	return []string{
		string(VnetConnectionProvisioningStateCanceled),
		string(VnetConnectionProvisioningStateDeleting),
		string(VnetConnectionProvisioningStateFailed),
		string(VnetConnectionProvisioningStateInProgress),
		string(VnetConnectionProvisioningStateSucceeded),
		string(VnetConnectionProvisioningStateUpdating),
	}
}

func (s *VnetConnectionProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVnetConnectionProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVnetConnectionProvisioningState(input string) (*VnetConnectionProvisioningState, error) {
	vals := map[string]VnetConnectionProvisioningState{
		"canceled":   VnetConnectionProvisioningStateCanceled,
		"deleting":   VnetConnectionProvisioningStateDeleting,
		"failed":     VnetConnectionProvisioningStateFailed,
		"inprogress": VnetConnectionProvisioningStateInProgress,
		"succeeded":  VnetConnectionProvisioningStateSucceeded,
		"updating":   VnetConnectionProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VnetConnectionProvisioningState(input)
	return &out, nil
}
