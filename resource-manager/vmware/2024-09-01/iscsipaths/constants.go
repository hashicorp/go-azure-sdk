package iscsipaths

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IscsiPathProvisioningState string

const (
	IscsiPathProvisioningStateBuilding  IscsiPathProvisioningState = "Building"
	IscsiPathProvisioningStateCanceled  IscsiPathProvisioningState = "Canceled"
	IscsiPathProvisioningStateDeleting  IscsiPathProvisioningState = "Deleting"
	IscsiPathProvisioningStateFailed    IscsiPathProvisioningState = "Failed"
	IscsiPathProvisioningStatePending   IscsiPathProvisioningState = "Pending"
	IscsiPathProvisioningStateSucceeded IscsiPathProvisioningState = "Succeeded"
	IscsiPathProvisioningStateUpdating  IscsiPathProvisioningState = "Updating"
)

func PossibleValuesForIscsiPathProvisioningState() []string {
	return []string{
		string(IscsiPathProvisioningStateBuilding),
		string(IscsiPathProvisioningStateCanceled),
		string(IscsiPathProvisioningStateDeleting),
		string(IscsiPathProvisioningStateFailed),
		string(IscsiPathProvisioningStatePending),
		string(IscsiPathProvisioningStateSucceeded),
		string(IscsiPathProvisioningStateUpdating),
	}
}

func (s *IscsiPathProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIscsiPathProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIscsiPathProvisioningState(input string) (*IscsiPathProvisioningState, error) {
	vals := map[string]IscsiPathProvisioningState{
		"building":  IscsiPathProvisioningStateBuilding,
		"canceled":  IscsiPathProvisioningStateCanceled,
		"deleting":  IscsiPathProvisioningStateDeleting,
		"failed":    IscsiPathProvisioningStateFailed,
		"pending":   IscsiPathProvisioningStatePending,
		"succeeded": IscsiPathProvisioningStateSucceeded,
		"updating":  IscsiPathProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IscsiPathProvisioningState(input)
	return &out, nil
}
