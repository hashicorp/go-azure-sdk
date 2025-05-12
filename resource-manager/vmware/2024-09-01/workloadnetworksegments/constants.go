package workloadnetworksegments

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SegmentStatusEnum string

const (
	SegmentStatusEnumFAILURE SegmentStatusEnum = "FAILURE"
	SegmentStatusEnumSUCCESS SegmentStatusEnum = "SUCCESS"
)

func PossibleValuesForSegmentStatusEnum() []string {
	return []string{
		string(SegmentStatusEnumFAILURE),
		string(SegmentStatusEnumSUCCESS),
	}
}

func (s *SegmentStatusEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSegmentStatusEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSegmentStatusEnum(input string) (*SegmentStatusEnum, error) {
	vals := map[string]SegmentStatusEnum{
		"failure": SegmentStatusEnumFAILURE,
		"success": SegmentStatusEnumSUCCESS,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SegmentStatusEnum(input)
	return &out, nil
}

type WorkloadNetworkSegmentProvisioningState string

const (
	WorkloadNetworkSegmentProvisioningStateBuilding  WorkloadNetworkSegmentProvisioningState = "Building"
	WorkloadNetworkSegmentProvisioningStateCanceled  WorkloadNetworkSegmentProvisioningState = "Canceled"
	WorkloadNetworkSegmentProvisioningStateDeleting  WorkloadNetworkSegmentProvisioningState = "Deleting"
	WorkloadNetworkSegmentProvisioningStateFailed    WorkloadNetworkSegmentProvisioningState = "Failed"
	WorkloadNetworkSegmentProvisioningStateSucceeded WorkloadNetworkSegmentProvisioningState = "Succeeded"
	WorkloadNetworkSegmentProvisioningStateUpdating  WorkloadNetworkSegmentProvisioningState = "Updating"
)

func PossibleValuesForWorkloadNetworkSegmentProvisioningState() []string {
	return []string{
		string(WorkloadNetworkSegmentProvisioningStateBuilding),
		string(WorkloadNetworkSegmentProvisioningStateCanceled),
		string(WorkloadNetworkSegmentProvisioningStateDeleting),
		string(WorkloadNetworkSegmentProvisioningStateFailed),
		string(WorkloadNetworkSegmentProvisioningStateSucceeded),
		string(WorkloadNetworkSegmentProvisioningStateUpdating),
	}
}

func (s *WorkloadNetworkSegmentProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWorkloadNetworkSegmentProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWorkloadNetworkSegmentProvisioningState(input string) (*WorkloadNetworkSegmentProvisioningState, error) {
	vals := map[string]WorkloadNetworkSegmentProvisioningState{
		"building":  WorkloadNetworkSegmentProvisioningStateBuilding,
		"canceled":  WorkloadNetworkSegmentProvisioningStateCanceled,
		"deleting":  WorkloadNetworkSegmentProvisioningStateDeleting,
		"failed":    WorkloadNetworkSegmentProvisioningStateFailed,
		"succeeded": WorkloadNetworkSegmentProvisioningStateSucceeded,
		"updating":  WorkloadNetworkSegmentProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkloadNetworkSegmentProvisioningState(input)
	return &out, nil
}
