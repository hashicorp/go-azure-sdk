package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StepState struct {
	EndTime   *string          `json:"endTime,omitempty"`
	Message   *string          `json:"message,omitempty"`
	StartTime *string          `json:"startTime,omitempty"`
	Status    *StepStateStatus `json:"status,omitempty"`
	StepName  *string          `json:"stepName,omitempty"`
}
