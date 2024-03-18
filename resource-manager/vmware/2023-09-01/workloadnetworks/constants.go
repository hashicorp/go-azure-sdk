package workloadnetworks

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkloadNetworkProvisioningState string

const (
	WorkloadNetworkProvisioningStateBuilding  WorkloadNetworkProvisioningState = "Building"
	WorkloadNetworkProvisioningStateCanceled  WorkloadNetworkProvisioningState = "Canceled"
	WorkloadNetworkProvisioningStateDeleting  WorkloadNetworkProvisioningState = "Deleting"
	WorkloadNetworkProvisioningStateFailed    WorkloadNetworkProvisioningState = "Failed"
	WorkloadNetworkProvisioningStateSucceeded WorkloadNetworkProvisioningState = "Succeeded"
	WorkloadNetworkProvisioningStateUpdating  WorkloadNetworkProvisioningState = "Updating"
)

func PossibleValuesForWorkloadNetworkProvisioningState() []string {
	return []string{
		string(WorkloadNetworkProvisioningStateBuilding),
		string(WorkloadNetworkProvisioningStateCanceled),
		string(WorkloadNetworkProvisioningStateDeleting),
		string(WorkloadNetworkProvisioningStateFailed),
		string(WorkloadNetworkProvisioningStateSucceeded),
		string(WorkloadNetworkProvisioningStateUpdating),
	}
}

func (s *WorkloadNetworkProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWorkloadNetworkProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWorkloadNetworkProvisioningState(input string) (*WorkloadNetworkProvisioningState, error) {
	vals := map[string]WorkloadNetworkProvisioningState{
		"building":  WorkloadNetworkProvisioningStateBuilding,
		"canceled":  WorkloadNetworkProvisioningStateCanceled,
		"deleting":  WorkloadNetworkProvisioningStateDeleting,
		"failed":    WorkloadNetworkProvisioningStateFailed,
		"succeeded": WorkloadNetworkProvisioningStateSucceeded,
		"updating":  WorkloadNetworkProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkloadNetworkProvisioningState(input)
	return &out, nil
}
