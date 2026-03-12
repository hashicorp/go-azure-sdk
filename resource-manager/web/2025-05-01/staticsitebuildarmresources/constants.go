package staticsitebuildarmresources

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BuildStatus string

const (
	BuildStatusDeleting             BuildStatus = "Deleting"
	BuildStatusDeploying            BuildStatus = "Deploying"
	BuildStatusDetached             BuildStatus = "Detached"
	BuildStatusFailed               BuildStatus = "Failed"
	BuildStatusReady                BuildStatus = "Ready"
	BuildStatusUploading            BuildStatus = "Uploading"
	BuildStatusWaitingForDeployment BuildStatus = "WaitingForDeployment"
)

func PossibleValuesForBuildStatus() []string {
	return []string{
		string(BuildStatusDeleting),
		string(BuildStatusDeploying),
		string(BuildStatusDetached),
		string(BuildStatusFailed),
		string(BuildStatusReady),
		string(BuildStatusUploading),
		string(BuildStatusWaitingForDeployment),
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
		"deleting":             BuildStatusDeleting,
		"deploying":            BuildStatusDeploying,
		"detached":             BuildStatusDetached,
		"failed":               BuildStatusFailed,
		"ready":                BuildStatusReady,
		"uploading":            BuildStatusUploading,
		"waitingfordeployment": BuildStatusWaitingForDeployment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BuildStatus(input)
	return &out, nil
}

type TriggerTypes string

const (
	TriggerTypesHTTPTrigger TriggerTypes = "HttpTrigger"
	TriggerTypesUnknown     TriggerTypes = "Unknown"
)

func PossibleValuesForTriggerTypes() []string {
	return []string{
		string(TriggerTypesHTTPTrigger),
		string(TriggerTypesUnknown),
	}
}

func (s *TriggerTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTriggerTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTriggerTypes(input string) (*TriggerTypes, error) {
	vals := map[string]TriggerTypes{
		"httptrigger": TriggerTypesHTTPTrigger,
		"unknown":     TriggerTypesUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TriggerTypes(input)
	return &out, nil
}
