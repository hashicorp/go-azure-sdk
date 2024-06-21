package job

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobProperties struct {
	CreationTime           *string                      `json:"creationTime,omitempty"`
	EndTime                *string                      `json:"endTime,omitempty"`
	Exception              *string                      `json:"exception,omitempty"`
	JobId                  *string                      `json:"jobId,omitempty"`
	JobScheduleId          *string                      `json:"jobScheduleId,omitempty"`
	LastModifiedTime       *string                      `json:"lastModifiedTime,omitempty"`
	LastStatusModifiedTime *string                      `json:"lastStatusModifiedTime,omitempty"`
	Parameters             *map[string]string           `json:"parameters,omitempty"`
	ProvisioningState      *JobProvisioningState        `json:"provisioningState,omitempty"`
	RunOn                  *string                      `json:"runOn,omitempty"`
	Runbook                *RunbookAssociationProperty  `json:"runbook,omitempty"`
	Schedule               *ScheduleAssociationProperty `json:"schedule,omitempty"`
	StartTime              *string                      `json:"startTime,omitempty"`
	StartedBy              *string                      `json:"startedBy,omitempty"`
	Status                 *JobStatus                   `json:"status,omitempty"`
	StatusDetails          *string                      `json:"statusDetails,omitempty"`
}
