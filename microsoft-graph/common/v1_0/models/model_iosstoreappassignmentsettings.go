package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosStoreAppAssignmentSettings struct {
	IsRemovable              *bool   `json:"isRemovable,omitempty"`
	ODataType                *string `json:"@odata.type,omitempty"`
	UninstallOnDeviceRemoval *bool   `json:"uninstallOnDeviceRemoval,omitempty"`
	VpnConfigurationId       *string `json:"vpnConfigurationId,omitempty"`
}
