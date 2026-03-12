package csmdeploymentstatuses

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentBuildStatus string

const (
	DeploymentBuildStatusBuildAborted             DeploymentBuildStatus = "BuildAborted"
	DeploymentBuildStatusBuildFailed              DeploymentBuildStatus = "BuildFailed"
	DeploymentBuildStatusBuildInProgress          DeploymentBuildStatus = "BuildInProgress"
	DeploymentBuildStatusBuildPending             DeploymentBuildStatus = "BuildPending"
	DeploymentBuildStatusBuildRequestReceived     DeploymentBuildStatus = "BuildRequestReceived"
	DeploymentBuildStatusBuildSuccessful          DeploymentBuildStatus = "BuildSuccessful"
	DeploymentBuildStatusPostBuildRestartRequired DeploymentBuildStatus = "PostBuildRestartRequired"
	DeploymentBuildStatusRuntimeFailed            DeploymentBuildStatus = "RuntimeFailed"
	DeploymentBuildStatusRuntimeStarting          DeploymentBuildStatus = "RuntimeStarting"
	DeploymentBuildStatusRuntimeSuccessful        DeploymentBuildStatus = "RuntimeSuccessful"
	DeploymentBuildStatusStartPolling             DeploymentBuildStatus = "StartPolling"
	DeploymentBuildStatusStartPollingWithRestart  DeploymentBuildStatus = "StartPollingWithRestart"
	DeploymentBuildStatusTimedOut                 DeploymentBuildStatus = "TimedOut"
)

func PossibleValuesForDeploymentBuildStatus() []string {
	return []string{
		string(DeploymentBuildStatusBuildAborted),
		string(DeploymentBuildStatusBuildFailed),
		string(DeploymentBuildStatusBuildInProgress),
		string(DeploymentBuildStatusBuildPending),
		string(DeploymentBuildStatusBuildRequestReceived),
		string(DeploymentBuildStatusBuildSuccessful),
		string(DeploymentBuildStatusPostBuildRestartRequired),
		string(DeploymentBuildStatusRuntimeFailed),
		string(DeploymentBuildStatusRuntimeStarting),
		string(DeploymentBuildStatusRuntimeSuccessful),
		string(DeploymentBuildStatusStartPolling),
		string(DeploymentBuildStatusStartPollingWithRestart),
		string(DeploymentBuildStatusTimedOut),
	}
}

func (s *DeploymentBuildStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeploymentBuildStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeploymentBuildStatus(input string) (*DeploymentBuildStatus, error) {
	vals := map[string]DeploymentBuildStatus{
		"buildaborted":             DeploymentBuildStatusBuildAborted,
		"buildfailed":              DeploymentBuildStatusBuildFailed,
		"buildinprogress":          DeploymentBuildStatusBuildInProgress,
		"buildpending":             DeploymentBuildStatusBuildPending,
		"buildrequestreceived":     DeploymentBuildStatusBuildRequestReceived,
		"buildsuccessful":          DeploymentBuildStatusBuildSuccessful,
		"postbuildrestartrequired": DeploymentBuildStatusPostBuildRestartRequired,
		"runtimefailed":            DeploymentBuildStatusRuntimeFailed,
		"runtimestarting":          DeploymentBuildStatusRuntimeStarting,
		"runtimesuccessful":        DeploymentBuildStatusRuntimeSuccessful,
		"startpolling":             DeploymentBuildStatusStartPolling,
		"startpollingwithrestart":  DeploymentBuildStatusStartPollingWithRestart,
		"timedout":                 DeploymentBuildStatusTimedOut,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeploymentBuildStatus(input)
	return &out, nil
}
