package softwareupdateconfigurationrun

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SoftwareUpdateConfigurationRunProperties struct {
	ComputerCount               *int64                               `json:"computerCount,omitempty"`
	ConfiguredDuration          *string                              `json:"configuredDuration,omitempty"`
	CreatedBy                   *string                              `json:"createdBy,omitempty"`
	CreationTime                *string                              `json:"creationTime,omitempty"`
	EndTime                     *string                              `json:"endTime,omitempty"`
	FailedCount                 *int64                               `json:"failedCount,omitempty"`
	LastModifiedBy              *string                              `json:"lastModifiedBy,omitempty"`
	LastModifiedTime            *string                              `json:"lastModifiedTime,omitempty"`
	OsType                      *string                              `json:"osType,omitempty"`
	SoftwareUpdateConfiguration *UpdateConfigurationNavigation       `json:"softwareUpdateConfiguration,omitempty"`
	StartTime                   *string                              `json:"startTime,omitempty"`
	Status                      *string                              `json:"status,omitempty"`
	Tasks                       *SoftwareUpdateConfigurationRunTasks `json:"tasks,omitempty"`
}
