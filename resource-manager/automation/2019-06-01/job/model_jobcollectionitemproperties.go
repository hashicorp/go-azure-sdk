package job

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobCollectionItemProperties struct {
	CreationTime      *string                     `json:"creationTime,omitempty"`
	EndTime           *string                     `json:"endTime,omitempty"`
	JobId             *string                     `json:"jobId,omitempty"`
	LastModifiedTime  *string                     `json:"lastModifiedTime,omitempty"`
	ProvisioningState *string                     `json:"provisioningState,omitempty"`
	RunOn             *string                     `json:"runOn,omitempty"`
	Runbook           *RunbookAssociationProperty `json:"runbook,omitempty"`
	StartTime         *string                     `json:"startTime,omitempty"`
	Status            *JobStatus                  `json:"status,omitempty"`
}
