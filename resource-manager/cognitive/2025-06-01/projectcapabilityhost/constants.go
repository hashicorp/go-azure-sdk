package projectcapabilityhost

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CapabilityHostProvisioningState string

const (
	CapabilityHostProvisioningStateCanceled  CapabilityHostProvisioningState = "Canceled"
	CapabilityHostProvisioningStateCreating  CapabilityHostProvisioningState = "Creating"
	CapabilityHostProvisioningStateDeleting  CapabilityHostProvisioningState = "Deleting"
	CapabilityHostProvisioningStateFailed    CapabilityHostProvisioningState = "Failed"
	CapabilityHostProvisioningStateSucceeded CapabilityHostProvisioningState = "Succeeded"
	CapabilityHostProvisioningStateUpdating  CapabilityHostProvisioningState = "Updating"
)

func PossibleValuesForCapabilityHostProvisioningState() []string {
	return []string{
		string(CapabilityHostProvisioningStateCanceled),
		string(CapabilityHostProvisioningStateCreating),
		string(CapabilityHostProvisioningStateDeleting),
		string(CapabilityHostProvisioningStateFailed),
		string(CapabilityHostProvisioningStateSucceeded),
		string(CapabilityHostProvisioningStateUpdating),
	}
}

func (s *CapabilityHostProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCapabilityHostProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCapabilityHostProvisioningState(input string) (*CapabilityHostProvisioningState, error) {
	vals := map[string]CapabilityHostProvisioningState{
		"canceled":  CapabilityHostProvisioningStateCanceled,
		"creating":  CapabilityHostProvisioningStateCreating,
		"deleting":  CapabilityHostProvisioningStateDeleting,
		"failed":    CapabilityHostProvisioningStateFailed,
		"succeeded": CapabilityHostProvisioningStateSucceeded,
		"updating":  CapabilityHostProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CapabilityHostProvisioningState(input)
	return &out, nil
}
