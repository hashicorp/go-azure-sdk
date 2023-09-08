package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationPolicy struct {
	Assignments          *[]DeviceManagementConfigurationPolicyAssignment      `json:"assignments,omitempty"`
	CreatedDateTime      *string                                               `json:"createdDateTime,omitempty"`
	CreationSource       *string                                               `json:"creationSource,omitempty"`
	Description          *string                                               `json:"description,omitempty"`
	Id                   *string                                               `json:"id,omitempty"`
	IsAssigned           *bool                                                 `json:"isAssigned,omitempty"`
	LastModifiedDateTime *string                                               `json:"lastModifiedDateTime,omitempty"`
	Name                 *string                                               `json:"name,omitempty"`
	ODataType            *string                                               `json:"@odata.type,omitempty"`
	Platforms            *DeviceManagementConfigurationPolicyPlatforms         `json:"platforms,omitempty"`
	PriorityMetaData     *DeviceManagementPriorityMetaData                     `json:"priorityMetaData,omitempty"`
	RoleScopeTagIds      *[]string                                             `json:"roleScopeTagIds,omitempty"`
	SettingCount         *int64                                                `json:"settingCount,omitempty"`
	Settings             *[]DeviceManagementConfigurationSetting               `json:"settings,omitempty"`
	Technologies         *DeviceManagementConfigurationPolicyTechnologies      `json:"technologies,omitempty"`
	TemplateReference    *DeviceManagementConfigurationPolicyTemplateReference `json:"templateReference,omitempty"`
}
