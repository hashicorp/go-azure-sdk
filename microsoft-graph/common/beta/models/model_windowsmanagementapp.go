package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagementApp struct {
	AvailableVersion                   *string                               `json:"availableVersion,omitempty"`
	HealthStates                       *[]WindowsManagementAppHealthState    `json:"healthStates,omitempty"`
	Id                                 *string                               `json:"id,omitempty"`
	ManagedInstaller                   *WindowsManagementAppManagedInstaller `json:"managedInstaller,omitempty"`
	ManagedInstallerConfiguredDateTime *string                               `json:"managedInstallerConfiguredDateTime,omitempty"`
	ODataType                          *string                               `json:"@odata.type,omitempty"`
}
