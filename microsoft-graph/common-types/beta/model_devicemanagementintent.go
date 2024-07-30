package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementIntent struct {
	Assignments                      *[]DeviceManagementIntentAssignment                `json:"assignments,omitempty"`
	Categories                       *[]DeviceManagementIntentSettingCategory           `json:"categories,omitempty"`
	Description                      *string                                            `json:"description,omitempty"`
	DeviceSettingStateSummaries      *[]DeviceManagementIntentDeviceSettingStateSummary `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStateSummary               *DeviceManagementIntentDeviceStateSummary          `json:"deviceStateSummary,omitempty"`
	DeviceStates                     *[]DeviceManagementIntentDeviceState               `json:"deviceStates,omitempty"`
	DisplayName                      *string                                            `json:"displayName,omitempty"`
	Id                               *string                                            `json:"id,omitempty"`
	IsAssigned                       *bool                                              `json:"isAssigned,omitempty"`
	IsMigratingToConfigurationPolicy *bool                                              `json:"isMigratingToConfigurationPolicy,omitempty"`
	LastModifiedDateTime             *string                                            `json:"lastModifiedDateTime,omitempty"`
	ODataType                        *string                                            `json:"@odata.type,omitempty"`
	RoleScopeTagIds                  *[]string                                          `json:"roleScopeTagIds,omitempty"`
	Settings                         *[]DeviceManagementSettingInstance                 `json:"settings,omitempty"`
	TemplateId                       *string                                            `json:"templateId,omitempty"`
	UserStateSummary                 *DeviceManagementIntentUserStateSummary            `json:"userStateSummary,omitempty"`
	UserStates                       *[]DeviceManagementIntentUserState                 `json:"userStates,omitempty"`
}
