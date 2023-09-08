package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementCompliancePolicy struct {
	Assignments             *[]DeviceManagementConfigurationPolicyAssignment    `json:"assignments,omitempty"`
	CreatedDateTime         *string                                             `json:"createdDateTime,omitempty"`
	CreationSource          *string                                             `json:"creationSource,omitempty"`
	Description             *string                                             `json:"description,omitempty"`
	Id                      *string                                             `json:"id,omitempty"`
	IsAssigned              *bool                                               `json:"isAssigned,omitempty"`
	LastModifiedDateTime    *string                                             `json:"lastModifiedDateTime,omitempty"`
	Name                    *string                                             `json:"name,omitempty"`
	ODataType               *string                                             `json:"@odata.type,omitempty"`
	Platforms               *DeviceManagementCompliancePolicyPlatforms          `json:"platforms,omitempty"`
	RoleScopeTagIds         *[]string                                           `json:"roleScopeTagIds,omitempty"`
	ScheduledActionsForRule *[]DeviceManagementComplianceScheduledActionForRule `json:"scheduledActionsForRule,omitempty"`
	SettingCount            *int64                                              `json:"settingCount,omitempty"`
	Settings                *[]DeviceManagementConfigurationSetting             `json:"settings,omitempty"`
	Technologies            *DeviceManagementCompliancePolicyTechnologies       `json:"technologies,omitempty"`
}
