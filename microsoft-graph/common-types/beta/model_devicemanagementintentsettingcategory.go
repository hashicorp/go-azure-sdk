package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementIntentSettingCategory struct {
	DisplayName        *string                              `json:"displayName,omitempty"`
	HasRequiredSetting *bool                                `json:"hasRequiredSetting,omitempty"`
	Id                 *string                              `json:"id,omitempty"`
	ODataType          *string                              `json:"@odata.type,omitempty"`
	SettingDefinitions *[]DeviceManagementSettingDefinition `json:"settingDefinitions,omitempty"`
	Settings           *[]DeviceManagementSettingInstance   `json:"settings,omitempty"`
}
