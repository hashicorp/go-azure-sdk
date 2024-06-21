package softwareupdateconfiguration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SoftwareUpdateConfigurationCollectionItemProperties struct {
	CreationTime        *string                           `json:"creationTime,omitempty"`
	Frequency           *ScheduleFrequency                `json:"frequency,omitempty"`
	LastModifiedTime    *string                           `json:"lastModifiedTime,omitempty"`
	NextRun             *string                           `json:"nextRun,omitempty"`
	ProvisioningState   *string                           `json:"provisioningState,omitempty"`
	StartTime           *string                           `json:"startTime,omitempty"`
	Tasks               *SoftwareUpdateConfigurationTasks `json:"tasks,omitempty"`
	UpdateConfiguration *UpdateConfiguration              `json:"updateConfiguration,omitempty"`
}
