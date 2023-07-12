package patch

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommandState string

const (
	CommandStateAccepted  CommandState = "Accepted"
	CommandStateFailed    CommandState = "Failed"
	CommandStateRunning   CommandState = "Running"
	CommandStateSucceeded CommandState = "Succeeded"
	CommandStateUnknown   CommandState = "Unknown"
)

func PossibleValuesForCommandState() []string {
	return []string{
		string(CommandStateAccepted),
		string(CommandStateFailed),
		string(CommandStateRunning),
		string(CommandStateSucceeded),
		string(CommandStateUnknown),
	}
}

func (s *CommandState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCommandState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCommandState(input string) (*CommandState, error) {
	vals := map[string]CommandState{
		"accepted":  CommandStateAccepted,
		"failed":    CommandStateFailed,
		"running":   CommandStateRunning,
		"succeeded": CommandStateSucceeded,
		"unknown":   CommandStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CommandState(input)
	return &out, nil
}

type ProjectProvisioningState string

const (
	ProjectProvisioningStateDeleting  ProjectProvisioningState = "Deleting"
	ProjectProvisioningStateSucceeded ProjectProvisioningState = "Succeeded"
)

func PossibleValuesForProjectProvisioningState() []string {
	return []string{
		string(ProjectProvisioningStateDeleting),
		string(ProjectProvisioningStateSucceeded),
	}
}

func (s *ProjectProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProjectProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProjectProvisioningState(input string) (*ProjectProvisioningState, error) {
	vals := map[string]ProjectProvisioningState{
		"deleting":  ProjectProvisioningStateDeleting,
		"succeeded": ProjectProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProjectProvisioningState(input)
	return &out, nil
}

type ProjectSourcePlatform string

const (
	ProjectSourcePlatformSQL     ProjectSourcePlatform = "SQL"
	ProjectSourcePlatformUnknown ProjectSourcePlatform = "Unknown"
)

func PossibleValuesForProjectSourcePlatform() []string {
	return []string{
		string(ProjectSourcePlatformSQL),
		string(ProjectSourcePlatformUnknown),
	}
}

func (s *ProjectSourcePlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProjectSourcePlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProjectSourcePlatform(input string) (*ProjectSourcePlatform, error) {
	vals := map[string]ProjectSourcePlatform{
		"sql":     ProjectSourcePlatformSQL,
		"unknown": ProjectSourcePlatformUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProjectSourcePlatform(input)
	return &out, nil
}

type ProjectTargetPlatform string

const (
	ProjectTargetPlatformSQLDB   ProjectTargetPlatform = "SQLDB"
	ProjectTargetPlatformUnknown ProjectTargetPlatform = "Unknown"
)

func PossibleValuesForProjectTargetPlatform() []string {
	return []string{
		string(ProjectTargetPlatformSQLDB),
		string(ProjectTargetPlatformUnknown),
	}
}

func (s *ProjectTargetPlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProjectTargetPlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProjectTargetPlatform(input string) (*ProjectTargetPlatform, error) {
	vals := map[string]ProjectTargetPlatform{
		"sqldb":   ProjectTargetPlatformSQLDB,
		"unknown": ProjectTargetPlatformUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProjectTargetPlatform(input)
	return &out, nil
}

type ServiceProvisioningState string

const (
	ServiceProvisioningStateAccepted      ServiceProvisioningState = "Accepted"
	ServiceProvisioningStateDeleting      ServiceProvisioningState = "Deleting"
	ServiceProvisioningStateDeploying     ServiceProvisioningState = "Deploying"
	ServiceProvisioningStateFailed        ServiceProvisioningState = "Failed"
	ServiceProvisioningStateFailedToStart ServiceProvisioningState = "FailedToStart"
	ServiceProvisioningStateFailedToStop  ServiceProvisioningState = "FailedToStop"
	ServiceProvisioningStateStarting      ServiceProvisioningState = "Starting"
	ServiceProvisioningStateStopped       ServiceProvisioningState = "Stopped"
	ServiceProvisioningStateStopping      ServiceProvisioningState = "Stopping"
	ServiceProvisioningStateSucceeded     ServiceProvisioningState = "Succeeded"
)

func PossibleValuesForServiceProvisioningState() []string {
	return []string{
		string(ServiceProvisioningStateAccepted),
		string(ServiceProvisioningStateDeleting),
		string(ServiceProvisioningStateDeploying),
		string(ServiceProvisioningStateFailed),
		string(ServiceProvisioningStateFailedToStart),
		string(ServiceProvisioningStateFailedToStop),
		string(ServiceProvisioningStateStarting),
		string(ServiceProvisioningStateStopped),
		string(ServiceProvisioningStateStopping),
		string(ServiceProvisioningStateSucceeded),
	}
}

func (s *ServiceProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServiceProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServiceProvisioningState(input string) (*ServiceProvisioningState, error) {
	vals := map[string]ServiceProvisioningState{
		"accepted":      ServiceProvisioningStateAccepted,
		"deleting":      ServiceProvisioningStateDeleting,
		"deploying":     ServiceProvisioningStateDeploying,
		"failed":        ServiceProvisioningStateFailed,
		"failedtostart": ServiceProvisioningStateFailedToStart,
		"failedtostop":  ServiceProvisioningStateFailedToStop,
		"starting":      ServiceProvisioningStateStarting,
		"stopped":       ServiceProvisioningStateStopped,
		"stopping":      ServiceProvisioningStateStopping,
		"succeeded":     ServiceProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceProvisioningState(input)
	return &out, nil
}

type TaskState string

const (
	TaskStateCanceled              TaskState = "Canceled"
	TaskStateFailed                TaskState = "Failed"
	TaskStateFailedInputValidation TaskState = "FailedInputValidation"
	TaskStateFaulted               TaskState = "Faulted"
	TaskStateQueued                TaskState = "Queued"
	TaskStateRunning               TaskState = "Running"
	TaskStateSucceeded             TaskState = "Succeeded"
	TaskStateUnknown               TaskState = "Unknown"
)

func PossibleValuesForTaskState() []string {
	return []string{
		string(TaskStateCanceled),
		string(TaskStateFailed),
		string(TaskStateFailedInputValidation),
		string(TaskStateFaulted),
		string(TaskStateQueued),
		string(TaskStateRunning),
		string(TaskStateSucceeded),
		string(TaskStateUnknown),
	}
}

func (s *TaskState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTaskState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTaskState(input string) (*TaskState, error) {
	vals := map[string]TaskState{
		"canceled":              TaskStateCanceled,
		"failed":                TaskStateFailed,
		"failedinputvalidation": TaskStateFailedInputValidation,
		"faulted":               TaskStateFaulted,
		"queued":                TaskStateQueued,
		"running":               TaskStateRunning,
		"succeeded":             TaskStateSucceeded,
		"unknown":               TaskStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TaskState(input)
	return &out, nil
}
