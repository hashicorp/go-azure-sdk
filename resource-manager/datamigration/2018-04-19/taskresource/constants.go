package taskresource

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
