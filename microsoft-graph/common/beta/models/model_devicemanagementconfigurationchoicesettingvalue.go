package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationChoiceSettingValue struct {
	Children                      *[]DeviceManagementConfigurationSettingInstance             `json:"children,omitempty"`
	ODataType                     *string                                                     `json:"@odata.type,omitempty"`
	SettingValueTemplateReference *DeviceManagementConfigurationSettingValueTemplateReference `json:"settingValueTemplateReference,omitempty"`
	Value                         *string                                                     `json:"value,omitempty"`
}
