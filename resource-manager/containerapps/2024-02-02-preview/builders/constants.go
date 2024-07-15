package builders

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BuilderProvisioningState string

const (
	BuilderProvisioningStateCanceled  BuilderProvisioningState = "Canceled"
	BuilderProvisioningStateCreating  BuilderProvisioningState = "Creating"
	BuilderProvisioningStateDeleting  BuilderProvisioningState = "Deleting"
	BuilderProvisioningStateFailed    BuilderProvisioningState = "Failed"
	BuilderProvisioningStateSucceeded BuilderProvisioningState = "Succeeded"
	BuilderProvisioningStateUpdating  BuilderProvisioningState = "Updating"
)

func PossibleValuesForBuilderProvisioningState() []string {
	return []string{
		string(BuilderProvisioningStateCanceled),
		string(BuilderProvisioningStateCreating),
		string(BuilderProvisioningStateDeleting),
		string(BuilderProvisioningStateFailed),
		string(BuilderProvisioningStateSucceeded),
		string(BuilderProvisioningStateUpdating),
	}
}

func (s *BuilderProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBuilderProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBuilderProvisioningState(input string) (*BuilderProvisioningState, error) {
	vals := map[string]BuilderProvisioningState{
		"canceled":  BuilderProvisioningStateCanceled,
		"creating":  BuilderProvisioningStateCreating,
		"deleting":  BuilderProvisioningStateDeleting,
		"failed":    BuilderProvisioningStateFailed,
		"succeeded": BuilderProvisioningStateSucceeded,
		"updating":  BuilderProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BuilderProvisioningState(input)
	return &out, nil
}
