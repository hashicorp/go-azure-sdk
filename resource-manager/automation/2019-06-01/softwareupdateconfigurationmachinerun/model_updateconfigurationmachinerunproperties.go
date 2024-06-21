package softwareupdateconfigurationmachinerun

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateConfigurationMachineRunProperties struct {
	ConfiguredDuration          *string                        `json:"configuredDuration,omitempty"`
	CorrelationId               *string                        `json:"correlationId,omitempty"`
	CreatedBy                   *string                        `json:"createdBy,omitempty"`
	CreationTime                *string                        `json:"creationTime,omitempty"`
	EndTime                     *string                        `json:"endTime,omitempty"`
	Error                       *ErrorResponse                 `json:"error,omitempty"`
	Job                         *JobNavigation                 `json:"job,omitempty"`
	LastModifiedBy              *string                        `json:"lastModifiedBy,omitempty"`
	LastModifiedTime            *string                        `json:"lastModifiedTime,omitempty"`
	OsType                      *string                        `json:"osType,omitempty"`
	SoftwareUpdateConfiguration *UpdateConfigurationNavigation `json:"softwareUpdateConfiguration,omitempty"`
	SourceComputerId            *string                        `json:"sourceComputerId,omitempty"`
	StartTime                   *string                        `json:"startTime,omitempty"`
	Status                      *string                        `json:"status,omitempty"`
	TargetComputer              *string                        `json:"targetComputer,omitempty"`
	TargetComputerType          *string                        `json:"targetComputerType,omitempty"`
}
