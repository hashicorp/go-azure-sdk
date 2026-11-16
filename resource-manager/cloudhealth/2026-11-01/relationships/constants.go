package relationships

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HealthModelProvisioningState string

const (
	HealthModelProvisioningStateCanceled  HealthModelProvisioningState = "Canceled"
	HealthModelProvisioningStateCreating  HealthModelProvisioningState = "Creating"
	HealthModelProvisioningStateDeleting  HealthModelProvisioningState = "Deleting"
	HealthModelProvisioningStateFailed    HealthModelProvisioningState = "Failed"
	HealthModelProvisioningStateSucceeded HealthModelProvisioningState = "Succeeded"
)

func PossibleValuesForHealthModelProvisioningState() []string {
	return []string{
		string(HealthModelProvisioningStateCanceled),
		string(HealthModelProvisioningStateCreating),
		string(HealthModelProvisioningStateDeleting),
		string(HealthModelProvisioningStateFailed),
		string(HealthModelProvisioningStateSucceeded),
	}
}

func (s *HealthModelProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHealthModelProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHealthModelProvisioningState(input string) (*HealthModelProvisioningState, error) {
	vals := map[string]HealthModelProvisioningState{
		"canceled":  HealthModelProvisioningStateCanceled,
		"creating":  HealthModelProvisioningStateCreating,
		"deleting":  HealthModelProvisioningStateDeleting,
		"failed":    HealthModelProvisioningStateFailed,
		"succeeded": HealthModelProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HealthModelProvisioningState(input)
	return &out, nil
}
