package jobtargetexecutions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobExecutionProperties struct {
	CreateTime              *string                `json:"createTime,omitempty"`
	CurrentAttemptStartTime *string                `json:"currentAttemptStartTime,omitempty"`
	CurrentAttempts         *int64                 `json:"currentAttempts,omitempty"`
	EndTime                 *string                `json:"endTime,omitempty"`
	JobExecutionId          *string                `json:"jobExecutionId,omitempty"`
	JobVersion              *int64                 `json:"jobVersion,omitempty"`
	LastMessage             *string                `json:"lastMessage,omitempty"`
	Lifecycle               *JobExecutionLifecycle `json:"lifecycle,omitempty"`
	ProvisioningState       *ProvisioningState     `json:"provisioningState,omitempty"`
	StartTime               *string                `json:"startTime,omitempty"`
	StepId                  *int64                 `json:"stepId,omitempty"`
	StepName                *string                `json:"stepName,omitempty"`
	Target                  *JobExecutionTarget    `json:"target,omitempty"`
}
