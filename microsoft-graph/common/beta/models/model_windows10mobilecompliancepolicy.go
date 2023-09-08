package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10MobileCompliancePolicy struct {
	ActiveFirewallRequired                *bool                                                `json:"activeFirewallRequired,omitempty"`
	Assignments                           *[]DeviceCompliancePolicyAssignment                  `json:"assignments,omitempty"`
	BitLockerEnabled                      *bool                                                `json:"bitLockerEnabled,omitempty"`
	CodeIntegrityEnabled                  *bool                                                `json:"codeIntegrityEnabled,omitempty"`
	CreatedDateTime                       *string                                              `json:"createdDateTime,omitempty"`
	Description                           *string                                              `json:"description,omitempty"`
	DeviceSettingStateSummaries           *[]SettingStateDeviceSummary                         `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                  *DeviceComplianceDeviceOverview                      `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                        *[]DeviceComplianceDeviceStatus                      `json:"deviceStatuses,omitempty"`
	DisplayName                           *string                                              `json:"displayName,omitempty"`
	EarlyLaunchAntiMalwareDriverEnabled   *bool                                                `json:"earlyLaunchAntiMalwareDriverEnabled,omitempty"`
	Id                                    *string                                              `json:"id,omitempty"`
	LastModifiedDateTime                  *string                                              `json:"lastModifiedDateTime,omitempty"`
	ODataType                             *string                                              `json:"@odata.type,omitempty"`
	OsMaximumVersion                      *string                                              `json:"osMaximumVersion,omitempty"`
	OsMinimumVersion                      *string                                              `json:"osMinimumVersion,omitempty"`
	PasswordBlockSimple                   *bool                                                `json:"passwordBlockSimple,omitempty"`
	PasswordExpirationDays                *int64                                               `json:"passwordExpirationDays,omitempty"`
	PasswordMinimumCharacterSetCount      *int64                                               `json:"passwordMinimumCharacterSetCount,omitempty"`
	PasswordMinimumLength                 *int64                                               `json:"passwordMinimumLength,omitempty"`
	PasswordMinutesOfInactivityBeforeLock *int64                                               `json:"passwordMinutesOfInactivityBeforeLock,omitempty"`
	PasswordPreviousPasswordBlockCount    *int64                                               `json:"passwordPreviousPasswordBlockCount,omitempty"`
	PasswordRequireToUnlockFromIdle       *bool                                                `json:"passwordRequireToUnlockFromIdle,omitempty"`
	PasswordRequired                      *bool                                                `json:"passwordRequired,omitempty"`
	PasswordRequiredType                  *Windows10MobileCompliancePolicyPasswordRequiredType `json:"passwordRequiredType,omitempty"`
	RoleScopeTagIds                       *[]string                                            `json:"roleScopeTagIds,omitempty"`
	ScheduledActionsForRule               *[]DeviceComplianceScheduledActionForRule            `json:"scheduledActionsForRule,omitempty"`
	SecureBootEnabled                     *bool                                                `json:"secureBootEnabled,omitempty"`
	StorageRequireEncryption              *bool                                                `json:"storageRequireEncryption,omitempty"`
	UserStatusOverview                    *DeviceComplianceUserOverview                        `json:"userStatusOverview,omitempty"`
	UserStatuses                          *[]DeviceComplianceUserStatus                        `json:"userStatuses,omitempty"`
	ValidOperatingSystemBuildRanges       *[]OperatingSystemVersionRange                       `json:"validOperatingSystemBuildRanges,omitempty"`
	Version                               *int64                                               `json:"version,omitempty"`
}
