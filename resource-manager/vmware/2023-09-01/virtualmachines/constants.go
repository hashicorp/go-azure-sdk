package virtualmachines

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineProvisioningState string

const (
	VirtualMachineProvisioningStateCanceled  VirtualMachineProvisioningState = "Canceled"
	VirtualMachineProvisioningStateFailed    VirtualMachineProvisioningState = "Failed"
	VirtualMachineProvisioningStateSucceeded VirtualMachineProvisioningState = "Succeeded"
)

func PossibleValuesForVirtualMachineProvisioningState() []string {
	return []string{
		string(VirtualMachineProvisioningStateCanceled),
		string(VirtualMachineProvisioningStateFailed),
		string(VirtualMachineProvisioningStateSucceeded),
	}
}

func (s *VirtualMachineProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualMachineProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualMachineProvisioningState(input string) (*VirtualMachineProvisioningState, error) {
	vals := map[string]VirtualMachineProvisioningState{
		"canceled":  VirtualMachineProvisioningStateCanceled,
		"failed":    VirtualMachineProvisioningStateFailed,
		"succeeded": VirtualMachineProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachineProvisioningState(input)
	return &out, nil
}

type VirtualMachineRestrictMovementState string

const (
	VirtualMachineRestrictMovementStateDisabled VirtualMachineRestrictMovementState = "Disabled"
	VirtualMachineRestrictMovementStateEnabled  VirtualMachineRestrictMovementState = "Enabled"
)

func PossibleValuesForVirtualMachineRestrictMovementState() []string {
	return []string{
		string(VirtualMachineRestrictMovementStateDisabled),
		string(VirtualMachineRestrictMovementStateEnabled),
	}
}

func (s *VirtualMachineRestrictMovementState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualMachineRestrictMovementState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualMachineRestrictMovementState(input string) (*VirtualMachineRestrictMovementState, error) {
	vals := map[string]VirtualMachineRestrictMovementState{
		"disabled": VirtualMachineRestrictMovementStateDisabled,
		"enabled":  VirtualMachineRestrictMovementStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachineRestrictMovementState(input)
	return &out, nil
}
