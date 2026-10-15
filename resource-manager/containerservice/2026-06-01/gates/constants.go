package gates

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GateProvisioningState string

const (
	GateProvisioningStateCanceled  GateProvisioningState = "Canceled"
	GateProvisioningStateFailed    GateProvisioningState = "Failed"
	GateProvisioningStateSucceeded GateProvisioningState = "Succeeded"
)

func PossibleValuesForGateProvisioningState() []string {
	return []string{
		string(GateProvisioningStateCanceled),
		string(GateProvisioningStateFailed),
		string(GateProvisioningStateSucceeded),
	}
}

func (s *GateProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGateProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGateProvisioningState(input string) (*GateProvisioningState, error) {
	vals := map[string]GateProvisioningState{
		"canceled":  GateProvisioningStateCanceled,
		"failed":    GateProvisioningStateFailed,
		"succeeded": GateProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GateProvisioningState(input)
	return &out, nil
}

type GateState string

const (
	GateStateCompleted GateState = "Completed"
	GateStatePending   GateState = "Pending"
	GateStateSkipped   GateState = "Skipped"
)

func PossibleValuesForGateState() []string {
	return []string{
		string(GateStateCompleted),
		string(GateStatePending),
		string(GateStateSkipped),
	}
}

func (s *GateState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGateState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGateState(input string) (*GateState, error) {
	vals := map[string]GateState{
		"completed": GateStateCompleted,
		"pending":   GateStatePending,
		"skipped":   GateStateSkipped,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GateState(input)
	return &out, nil
}

type GateType string

const (
	GateTypeApproval GateType = "Approval"
)

func PossibleValuesForGateType() []string {
	return []string{
		string(GateTypeApproval),
	}
}

func (s *GateType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGateType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGateType(input string) (*GateType, error) {
	vals := map[string]GateType{
		"approval": GateTypeApproval,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GateType(input)
	return &out, nil
}

type Timing string

const (
	TimingAfter  Timing = "After"
	TimingBefore Timing = "Before"
)

func PossibleValuesForTiming() []string {
	return []string{
		string(TimingAfter),
		string(TimingBefore),
	}
}

func (s *Timing) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTiming(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTiming(input string) (*Timing, error) {
	vals := map[string]Timing{
		"after":  TimingAfter,
		"before": TimingBefore,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Timing(input)
	return &out, nil
}
