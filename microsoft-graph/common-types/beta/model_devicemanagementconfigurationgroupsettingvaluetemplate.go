package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationGroupSettingValueTemplate struct {
	Children               *[]DeviceManagementConfigurationSettingInstanceTemplate `json:"children,omitempty"`
	ODataType              *string                                                 `json:"@odata.type,omitempty"`
	SettingValueTemplateId *string                                                 `json:"settingValueTemplateId,omitempty"`
}
