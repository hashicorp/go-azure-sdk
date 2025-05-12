package purestoragepolicies

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PureStoragePolicyProvisioningState string

const (
	PureStoragePolicyProvisioningStateCanceled  PureStoragePolicyProvisioningState = "Canceled"
	PureStoragePolicyProvisioningStateDeleting  PureStoragePolicyProvisioningState = "Deleting"
	PureStoragePolicyProvisioningStateFailed    PureStoragePolicyProvisioningState = "Failed"
	PureStoragePolicyProvisioningStateSucceeded PureStoragePolicyProvisioningState = "Succeeded"
	PureStoragePolicyProvisioningStateUpdating  PureStoragePolicyProvisioningState = "Updating"
)

func PossibleValuesForPureStoragePolicyProvisioningState() []string {
	return []string{
		string(PureStoragePolicyProvisioningStateCanceled),
		string(PureStoragePolicyProvisioningStateDeleting),
		string(PureStoragePolicyProvisioningStateFailed),
		string(PureStoragePolicyProvisioningStateSucceeded),
		string(PureStoragePolicyProvisioningStateUpdating),
	}
}

func (s *PureStoragePolicyProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePureStoragePolicyProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePureStoragePolicyProvisioningState(input string) (*PureStoragePolicyProvisioningState, error) {
	vals := map[string]PureStoragePolicyProvisioningState{
		"canceled":  PureStoragePolicyProvisioningStateCanceled,
		"deleting":  PureStoragePolicyProvisioningStateDeleting,
		"failed":    PureStoragePolicyProvisioningStateFailed,
		"succeeded": PureStoragePolicyProvisioningStateSucceeded,
		"updating":  PureStoragePolicyProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PureStoragePolicyProvisioningState(input)
	return &out, nil
}
