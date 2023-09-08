package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingOccurrence struct {
	MaxDeviceOccurrence *int64  `json:"maxDeviceOccurrence,omitempty"`
	MinDeviceOccurrence *int64  `json:"minDeviceOccurrence,omitempty"`
	ODataType           *string `json:"@odata.type,omitempty"`
}
