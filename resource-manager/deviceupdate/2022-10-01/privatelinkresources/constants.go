package privatelinkresources

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupIdProvisioningState string

const (
	GroupIdProvisioningStateCanceled  GroupIdProvisioningState = "Canceled"
	GroupIdProvisioningStateFailed    GroupIdProvisioningState = "Failed"
	GroupIdProvisioningStateSucceeded GroupIdProvisioningState = "Succeeded"
)

func PossibleValuesForGroupIdProvisioningState() []string {
	return []string{
		string(GroupIdProvisioningStateCanceled),
		string(GroupIdProvisioningStateFailed),
		string(GroupIdProvisioningStateSucceeded),
	}
}

func (s *GroupIdProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupIdProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupIdProvisioningState(input string) (*GroupIdProvisioningState, error) {
	vals := map[string]GroupIdProvisioningState{
		"canceled":  GroupIdProvisioningStateCanceled,
		"failed":    GroupIdProvisioningStateFailed,
		"succeeded": GroupIdProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupIdProvisioningState(input)
	return &out, nil
}
