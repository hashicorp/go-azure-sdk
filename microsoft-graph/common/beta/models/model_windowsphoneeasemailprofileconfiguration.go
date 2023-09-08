package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhoneEASEmailProfileConfiguration struct {
	AccountName                                 *string                                                        `json:"accountName,omitempty"`
	ApplyOnlyToWindowsPhone81                   *bool                                                          `json:"applyOnlyToWindowsPhone81,omitempty"`
	Assignments                                 *[]DeviceConfigurationAssignment                               `json:"assignments,omitempty"`
	CreatedDateTime                             *string                                                        `json:"createdDateTime,omitempty"`
	CustomDomainName                            *string                                                        `json:"customDomainName,omitempty"`
	Description                                 *string                                                        `json:"description,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode                   `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	DeviceManagementApplicabilityRuleOsEdition  *DeviceManagementApplicabilityRuleOsEdition                    `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOsVersion  *DeviceManagementApplicabilityRuleOsVersion                    `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceSettingStateSummaries                 *[]SettingStateDeviceSummary                                   `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                        *DeviceConfigurationDeviceOverview                             `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                              *[]DeviceConfigurationDeviceStatus                             `json:"deviceStatuses,omitempty"`
	DisplayName                                 *string                                                        `json:"displayName,omitempty"`
	DurationOfEmailToSync                       *WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync `json:"durationOfEmailToSync,omitempty"`
	EmailAddressSource                          *WindowsPhoneEASEmailProfileConfigurationEmailAddressSource    `json:"emailAddressSource,omitempty"`
	EmailSyncSchedule                           *WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule     `json:"emailSyncSchedule,omitempty"`
	GroupAssignments                            *[]DeviceConfigurationGroupAssignment                          `json:"groupAssignments,omitempty"`
	HostName                                    *string                                                        `json:"hostName,omitempty"`
	Id                                          *string                                                        `json:"id,omitempty"`
	LastModifiedDateTime                        *string                                                        `json:"lastModifiedDateTime,omitempty"`
	ODataType                                   *string                                                        `json:"@odata.type,omitempty"`
	RequireSsl                                  *bool                                                          `json:"requireSsl,omitempty"`
	RoleScopeTagIds                             *[]string                                                      `json:"roleScopeTagIds,omitempty"`
	SupportsScopeTags                           *bool                                                          `json:"supportsScopeTags,omitempty"`
	SyncCalendar                                *bool                                                          `json:"syncCalendar,omitempty"`
	SyncContacts                                *bool                                                          `json:"syncContacts,omitempty"`
	SyncTasks                                   *bool                                                          `json:"syncTasks,omitempty"`
	UserDomainNameSource                        *WindowsPhoneEASEmailProfileConfigurationUserDomainNameSource  `json:"userDomainNameSource,omitempty"`
	UserStatusOverview                          *DeviceConfigurationUserOverview                               `json:"userStatusOverview,omitempty"`
	UserStatuses                                *[]DeviceConfigurationUserStatus                               `json:"userStatuses,omitempty"`
	UsernameAADSource                           *WindowsPhoneEASEmailProfileConfigurationUsernameAADSource     `json:"usernameAADSource,omitempty"`
	UsernameSource                              *WindowsPhoneEASEmailProfileConfigurationUsernameSource        `json:"usernameSource,omitempty"`
	Version                                     *int64                                                         `json:"version,omitempty"`
}
