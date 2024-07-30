package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EasEmailProfileConfigurationBase struct {
	Assignments                                 *[]DeviceConfigurationAssignment                      `json:"assignments,omitempty"`
	CreatedDateTime                             *string                                               `json:"createdDateTime,omitempty"`
	CustomDomainName                            *string                                               `json:"customDomainName,omitempty"`
	Description                                 *string                                               `json:"description,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode          `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	DeviceManagementApplicabilityRuleOsEdition  *DeviceManagementApplicabilityRuleOsEdition           `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOsVersion  *DeviceManagementApplicabilityRuleOsVersion           `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceSettingStateSummaries                 *[]SettingStateDeviceSummary                          `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                        *DeviceConfigurationDeviceOverview                    `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                              *[]DeviceConfigurationDeviceStatus                    `json:"deviceStatuses,omitempty"`
	DisplayName                                 *string                                               `json:"displayName,omitempty"`
	GroupAssignments                            *[]DeviceConfigurationGroupAssignment                 `json:"groupAssignments,omitempty"`
	Id                                          *string                                               `json:"id,omitempty"`
	LastModifiedDateTime                        *string                                               `json:"lastModifiedDateTime,omitempty"`
	ODataType                                   *string                                               `json:"@odata.type,omitempty"`
	RoleScopeTagIds                             *[]string                                             `json:"roleScopeTagIds,omitempty"`
	SupportsScopeTags                           *bool                                                 `json:"supportsScopeTags,omitempty"`
	UserDomainNameSource                        *EasEmailProfileConfigurationBaseUserDomainNameSource `json:"userDomainNameSource,omitempty"`
	UserStatusOverview                          *DeviceConfigurationUserOverview                      `json:"userStatusOverview,omitempty"`
	UserStatuses                                *[]DeviceConfigurationUserStatus                      `json:"userStatuses,omitempty"`
	UsernameAADSource                           *EasEmailProfileConfigurationBaseUsernameAADSource    `json:"usernameAADSource,omitempty"`
	UsernameSource                              *EasEmailProfileConfigurationBaseUsernameSource       `json:"usernameSource,omitempty"`
	Version                                     *int64                                                `json:"version,omitempty"`
}
