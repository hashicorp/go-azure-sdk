package sourcecontrolsyncjob

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SourceControlSyncJobByIdProperties struct {
	CreationTime           *string            `json:"creationTime,omitempty"`
	EndTime                *string            `json:"endTime,omitempty"`
	Exception              *string            `json:"exception,omitempty"`
	ProvisioningState      *ProvisioningState `json:"provisioningState,omitempty"`
	SourceControlSyncJobId *string            `json:"sourceControlSyncJobId,omitempty"`
	StartTime              *string            `json:"startTime,omitempty"`
	SyncType               *SyncType          `json:"syncType,omitempty"`
}
