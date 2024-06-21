package softwareupdateconfiguration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SoftwareUpdateConfigurationProperties struct {
	CreatedBy           *string                           `json:"createdBy,omitempty"`
	CreationTime        *string                           `json:"creationTime,omitempty"`
	Error               *ErrorResponse                    `json:"error,omitempty"`
	LastModifiedBy      *string                           `json:"lastModifiedBy,omitempty"`
	LastModifiedTime    *string                           `json:"lastModifiedTime,omitempty"`
	ProvisioningState   *string                           `json:"provisioningState,omitempty"`
	ScheduleInfo        SUCScheduleProperties             `json:"scheduleInfo"`
	Tasks               *SoftwareUpdateConfigurationTasks `json:"tasks,omitempty"`
	UpdateConfiguration UpdateConfiguration               `json:"updateConfiguration"`
}
