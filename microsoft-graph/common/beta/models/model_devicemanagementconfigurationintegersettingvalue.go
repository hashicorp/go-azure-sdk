package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationIntegerSettingValue struct {
	ODataType                     *string                                                     `json:"@odata.type,omitempty"`
	SettingValueTemplateReference *DeviceManagementConfigurationSettingValueTemplateReference `json:"settingValueTemplateReference,omitempty"`
	Value                         *int64                                                      `json:"value,omitempty"`
}
