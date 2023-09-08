package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVppAppAssignmentSettings struct {
	IsRemovable              *bool   `json:"isRemovable,omitempty"`
	ODataType                *string `json:"@odata.type,omitempty"`
	PreventAutoAppUpdate     *bool   `json:"preventAutoAppUpdate,omitempty"`
	PreventManagedAppBackup  *bool   `json:"preventManagedAppBackup,omitempty"`
	UninstallOnDeviceRemoval *bool   `json:"uninstallOnDeviceRemoval,omitempty"`
	UseDeviceLicensing       *bool   `json:"useDeviceLicensing,omitempty"`
	VpnConfigurationId       *string `json:"vpnConfigurationId,omitempty"`
}
