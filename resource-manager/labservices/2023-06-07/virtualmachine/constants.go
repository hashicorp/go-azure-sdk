package virtualmachine

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningState string

const (
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateLocked    ProvisioningState = "Locked"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateLocked),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

func (s *ProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"creating":  ProvisioningStateCreating,
		"deleting":  ProvisioningStateDeleting,
		"failed":    ProvisioningStateFailed,
		"locked":    ProvisioningStateLocked,
		"succeeded": ProvisioningStateSucceeded,
		"updating":  ProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}

type VirtualMachineState string

const (
	VirtualMachineStateRedeploying       VirtualMachineState = "Redeploying"
	VirtualMachineStateReimaging         VirtualMachineState = "Reimaging"
	VirtualMachineStateResettingPassword VirtualMachineState = "ResettingPassword"
	VirtualMachineStateRunning           VirtualMachineState = "Running"
	VirtualMachineStateStarting          VirtualMachineState = "Starting"
	VirtualMachineStateStopped           VirtualMachineState = "Stopped"
	VirtualMachineStateStopping          VirtualMachineState = "Stopping"
)

func PossibleValuesForVirtualMachineState() []string {
	return []string{
		string(VirtualMachineStateRedeploying),
		string(VirtualMachineStateReimaging),
		string(VirtualMachineStateResettingPassword),
		string(VirtualMachineStateRunning),
		string(VirtualMachineStateStarting),
		string(VirtualMachineStateStopped),
		string(VirtualMachineStateStopping),
	}
}

func (s *VirtualMachineState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualMachineState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualMachineState(input string) (*VirtualMachineState, error) {
	vals := map[string]VirtualMachineState{
		"redeploying":       VirtualMachineStateRedeploying,
		"reimaging":         VirtualMachineStateReimaging,
		"resettingpassword": VirtualMachineStateResettingPassword,
		"running":           VirtualMachineStateRunning,
		"starting":          VirtualMachineStateStarting,
		"stopped":           VirtualMachineStateStopped,
		"stopping":          VirtualMachineStateStopping,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachineState(input)
	return &out, nil
}

type VirtualMachineType string

const (
	VirtualMachineTypeTemplate VirtualMachineType = "Template"
	VirtualMachineTypeUser     VirtualMachineType = "User"
)

func PossibleValuesForVirtualMachineType() []string {
	return []string{
		string(VirtualMachineTypeTemplate),
		string(VirtualMachineTypeUser),
	}
}

func (s *VirtualMachineType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualMachineType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualMachineType(input string) (*VirtualMachineType, error) {
	vals := map[string]VirtualMachineType{
		"template": VirtualMachineTypeTemplate,
		"user":     VirtualMachineTypeUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachineType(input)
	return &out, nil
}
