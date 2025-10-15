package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActionState struct {
	ActionType    *string            `json:"actionType,omitempty"`
	CorrelationId *string            `json:"correlationId,omitempty"`
	EndTime       *string            `json:"endTime,omitempty"`
	Message       *string            `json:"message,omitempty"`
	StartTime     *string            `json:"startTime,omitempty"`
	Status        *ActionStateStatus `json:"status,omitempty"`
	StepStates    *[]StepState       `json:"stepStates,omitempty"`
}
