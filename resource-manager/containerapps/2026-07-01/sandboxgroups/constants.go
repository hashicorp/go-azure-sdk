package sandboxgroups

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                       ManagedServiceIdentityType = "None"
	ManagedServiceIdentityTypeSystemAssigned             ManagedServiceIdentityType = "SystemAssigned"
	ManagedServiceIdentityTypeSystemAssignedUserAssigned ManagedServiceIdentityType = "SystemAssigned,UserAssigned"
	ManagedServiceIdentityTypeUserAssigned               ManagedServiceIdentityType = "UserAssigned"
)

func PossibleValuesForManagedServiceIdentityType() []string {
	return []string{
		string(ManagedServiceIdentityTypeNone),
		string(ManagedServiceIdentityTypeSystemAssigned),
		string(ManagedServiceIdentityTypeSystemAssignedUserAssigned),
		string(ManagedServiceIdentityTypeUserAssigned),
	}
}

func (s *ManagedServiceIdentityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedServiceIdentityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedServiceIdentityType(input string) (*ManagedServiceIdentityType, error) {
	vals := map[string]ManagedServiceIdentityType{
		"none":                        ManagedServiceIdentityTypeNone,
		"systemassigned":              ManagedServiceIdentityTypeSystemAssigned,
		"systemassigned,userassigned": ManagedServiceIdentityTypeSystemAssignedUserAssigned,
		"userassigned":                ManagedServiceIdentityTypeUserAssigned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedServiceIdentityType(input)
	return &out, nil
}

type SandboxGroupProvisioningState string

const (
	SandboxGroupProvisioningStateCanceled   SandboxGroupProvisioningState = "Canceled"
	SandboxGroupProvisioningStateDeleting   SandboxGroupProvisioningState = "Deleting"
	SandboxGroupProvisioningStateFailed     SandboxGroupProvisioningState = "Failed"
	SandboxGroupProvisioningStateInProgress SandboxGroupProvisioningState = "InProgress"
	SandboxGroupProvisioningStateSucceeded  SandboxGroupProvisioningState = "Succeeded"
	SandboxGroupProvisioningStateUpdating   SandboxGroupProvisioningState = "Updating"
)

func PossibleValuesForSandboxGroupProvisioningState() []string {
	return []string{
		string(SandboxGroupProvisioningStateCanceled),
		string(SandboxGroupProvisioningStateDeleting),
		string(SandboxGroupProvisioningStateFailed),
		string(SandboxGroupProvisioningStateInProgress),
		string(SandboxGroupProvisioningStateSucceeded),
		string(SandboxGroupProvisioningStateUpdating),
	}
}

func (s *SandboxGroupProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSandboxGroupProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSandboxGroupProvisioningState(input string) (*SandboxGroupProvisioningState, error) {
	vals := map[string]SandboxGroupProvisioningState{
		"canceled":   SandboxGroupProvisioningStateCanceled,
		"deleting":   SandboxGroupProvisioningStateDeleting,
		"failed":     SandboxGroupProvisioningStateFailed,
		"inprogress": SandboxGroupProvisioningStateInProgress,
		"succeeded":  SandboxGroupProvisioningStateSucceeded,
		"updating":   SandboxGroupProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SandboxGroupProvisioningState(input)
	return &out, nil
}
