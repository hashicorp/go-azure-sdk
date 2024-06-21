package virtualmachinescalesetvmruncommands

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineRunCommandInstanceView struct {
	EndTime          *string               `json:"endTime,omitempty"`
	Error            *string               `json:"error,omitempty"`
	ExecutionMessage *string               `json:"executionMessage,omitempty"`
	ExecutionState   *ExecutionState       `json:"executionState,omitempty"`
	ExitCode         *int64                `json:"exitCode,omitempty"`
	Output           *string               `json:"output,omitempty"`
	StartTime        *string               `json:"startTime,omitempty"`
	Statuses         *[]InstanceViewStatus `json:"statuses,omitempty"`
}
