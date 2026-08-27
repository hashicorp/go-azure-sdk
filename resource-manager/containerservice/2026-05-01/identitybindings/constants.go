package identitybindings

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityBindingProvisioningState string

const (
	IdentityBindingProvisioningStateCanceled  IdentityBindingProvisioningState = "Canceled"
	IdentityBindingProvisioningStateCreating  IdentityBindingProvisioningState = "Creating"
	IdentityBindingProvisioningStateDeleting  IdentityBindingProvisioningState = "Deleting"
	IdentityBindingProvisioningStateFailed    IdentityBindingProvisioningState = "Failed"
	IdentityBindingProvisioningStateSucceeded IdentityBindingProvisioningState = "Succeeded"
	IdentityBindingProvisioningStateUpdating  IdentityBindingProvisioningState = "Updating"
)

func PossibleValuesForIdentityBindingProvisioningState() []string {
	return []string{
		string(IdentityBindingProvisioningStateCanceled),
		string(IdentityBindingProvisioningStateCreating),
		string(IdentityBindingProvisioningStateDeleting),
		string(IdentityBindingProvisioningStateFailed),
		string(IdentityBindingProvisioningStateSucceeded),
		string(IdentityBindingProvisioningStateUpdating),
	}
}

func (s *IdentityBindingProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityBindingProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityBindingProvisioningState(input string) (*IdentityBindingProvisioningState, error) {
	vals := map[string]IdentityBindingProvisioningState{
		"canceled":  IdentityBindingProvisioningStateCanceled,
		"creating":  IdentityBindingProvisioningStateCreating,
		"deleting":  IdentityBindingProvisioningStateDeleting,
		"failed":    IdentityBindingProvisioningStateFailed,
		"succeeded": IdentityBindingProvisioningStateSucceeded,
		"updating":  IdentityBindingProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityBindingProvisioningState(input)
	return &out, nil
}
