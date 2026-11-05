package dotnetcomponents

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DotNetComponentProvisioningState string

const (
	DotNetComponentProvisioningStateCanceled   DotNetComponentProvisioningState = "Canceled"
	DotNetComponentProvisioningStateDeleting   DotNetComponentProvisioningState = "Deleting"
	DotNetComponentProvisioningStateFailed     DotNetComponentProvisioningState = "Failed"
	DotNetComponentProvisioningStateInProgress DotNetComponentProvisioningState = "InProgress"
	DotNetComponentProvisioningStateSucceeded  DotNetComponentProvisioningState = "Succeeded"
)

func PossibleValuesForDotNetComponentProvisioningState() []string {
	return []string{
		string(DotNetComponentProvisioningStateCanceled),
		string(DotNetComponentProvisioningStateDeleting),
		string(DotNetComponentProvisioningStateFailed),
		string(DotNetComponentProvisioningStateInProgress),
		string(DotNetComponentProvisioningStateSucceeded),
	}
}

func (s *DotNetComponentProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDotNetComponentProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDotNetComponentProvisioningState(input string) (*DotNetComponentProvisioningState, error) {
	vals := map[string]DotNetComponentProvisioningState{
		"canceled":   DotNetComponentProvisioningStateCanceled,
		"deleting":   DotNetComponentProvisioningStateDeleting,
		"failed":     DotNetComponentProvisioningStateFailed,
		"inprogress": DotNetComponentProvisioningStateInProgress,
		"succeeded":  DotNetComponentProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DotNetComponentProvisioningState(input)
	return &out, nil
}

type DotNetComponentType string

const (
	DotNetComponentTypeAspireDashboard DotNetComponentType = "AspireDashboard"
)

func PossibleValuesForDotNetComponentType() []string {
	return []string{
		string(DotNetComponentTypeAspireDashboard),
	}
}

func (s *DotNetComponentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDotNetComponentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDotNetComponentType(input string) (*DotNetComponentType, error) {
	vals := map[string]DotNetComponentType{
		"aspiredashboard": DotNetComponentTypeAspireDashboard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DotNetComponentType(input)
	return &out, nil
}
