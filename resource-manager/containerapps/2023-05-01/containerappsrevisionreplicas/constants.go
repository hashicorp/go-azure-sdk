package containerappsrevisionreplicas

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerAppContainerRunningState string

const (
	ContainerAppContainerRunningStateRunning    ContainerAppContainerRunningState = "Running"
	ContainerAppContainerRunningStateTerminated ContainerAppContainerRunningState = "Terminated"
	ContainerAppContainerRunningStateWaiting    ContainerAppContainerRunningState = "Waiting"
)

func PossibleValuesForContainerAppContainerRunningState() []string {
	return []string{
		string(ContainerAppContainerRunningStateRunning),
		string(ContainerAppContainerRunningStateTerminated),
		string(ContainerAppContainerRunningStateWaiting),
	}
}

func parseContainerAppContainerRunningState(input string) (*ContainerAppContainerRunningState, error) {
	vals := map[string]ContainerAppContainerRunningState{
		"running":    ContainerAppContainerRunningStateRunning,
		"terminated": ContainerAppContainerRunningStateTerminated,
		"waiting":    ContainerAppContainerRunningStateWaiting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContainerAppContainerRunningState(input)
	return &out, nil
}

type ContainerAppReplicaRunningState string

const (
	ContainerAppReplicaRunningStateNotRunning ContainerAppReplicaRunningState = "NotRunning"
	ContainerAppReplicaRunningStateRunning    ContainerAppReplicaRunningState = "Running"
	ContainerAppReplicaRunningStateUnknown    ContainerAppReplicaRunningState = "Unknown"
)

func PossibleValuesForContainerAppReplicaRunningState() []string {
	return []string{
		string(ContainerAppReplicaRunningStateNotRunning),
		string(ContainerAppReplicaRunningStateRunning),
		string(ContainerAppReplicaRunningStateUnknown),
	}
}

func parseContainerAppReplicaRunningState(input string) (*ContainerAppReplicaRunningState, error) {
	vals := map[string]ContainerAppReplicaRunningState{
		"notrunning": ContainerAppReplicaRunningStateNotRunning,
		"running":    ContainerAppReplicaRunningStateRunning,
		"unknown":    ContainerAppReplicaRunningStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContainerAppReplicaRunningState(input)
	return &out, nil
}
