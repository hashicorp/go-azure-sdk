package template

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateTemplateCreateInstanceRequest struct {
	Description     *string                            `json:"description,omitempty"`
	DisplayName     *string                            `json:"displayName,omitempty"`
	RoleScopeTagIds *[]string                          `json:"roleScopeTagIds,omitempty"`
	SettingsDelta   *[]DeviceManagementSettingInstance `json:"settingsDelta,omitempty"`
}
