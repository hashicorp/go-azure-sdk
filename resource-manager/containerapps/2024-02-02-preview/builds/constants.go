package builds

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BuildProvisioningState string

const (
	BuildProvisioningStateCanceled  BuildProvisioningState = "Canceled"
	BuildProvisioningStateCreating  BuildProvisioningState = "Creating"
	BuildProvisioningStateDeleting  BuildProvisioningState = "Deleting"
	BuildProvisioningStateFailed    BuildProvisioningState = "Failed"
	BuildProvisioningStateSucceeded BuildProvisioningState = "Succeeded"
	BuildProvisioningStateUpdating  BuildProvisioningState = "Updating"
)

func PossibleValuesForBuildProvisioningState() []string {
	return []string{
		string(BuildProvisioningStateCanceled),
		string(BuildProvisioningStateCreating),
		string(BuildProvisioningStateDeleting),
		string(BuildProvisioningStateFailed),
		string(BuildProvisioningStateSucceeded),
		string(BuildProvisioningStateUpdating),
	}
}

func (s *BuildProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBuildProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBuildProvisioningState(input string) (*BuildProvisioningState, error) {
	vals := map[string]BuildProvisioningState{
		"canceled":  BuildProvisioningStateCanceled,
		"creating":  BuildProvisioningStateCreating,
		"deleting":  BuildProvisioningStateDeleting,
		"failed":    BuildProvisioningStateFailed,
		"succeeded": BuildProvisioningStateSucceeded,
		"updating":  BuildProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BuildProvisioningState(input)
	return &out, nil
}

type BuildStatus string

const (
	BuildStatusCanceled   BuildStatus = "Canceled"
	BuildStatusFailed     BuildStatus = "Failed"
	BuildStatusInProgress BuildStatus = "InProgress"
	BuildStatusNotStarted BuildStatus = "NotStarted"
	BuildStatusSucceeded  BuildStatus = "Succeeded"
)

func PossibleValuesForBuildStatus() []string {
	return []string{
		string(BuildStatusCanceled),
		string(BuildStatusFailed),
		string(BuildStatusInProgress),
		string(BuildStatusNotStarted),
		string(BuildStatusSucceeded),
	}
}

func (s *BuildStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBuildStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBuildStatus(input string) (*BuildStatus, error) {
	vals := map[string]BuildStatus{
		"canceled":   BuildStatusCanceled,
		"failed":     BuildStatusFailed,
		"inprogress": BuildStatusInProgress,
		"notstarted": BuildStatusNotStarted,
		"succeeded":  BuildStatusSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BuildStatus(input)
	return &out, nil
}
