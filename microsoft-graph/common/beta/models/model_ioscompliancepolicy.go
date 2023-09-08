package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosCompliancePolicy struct {
	AdvancedThreatProtectionRequiredSecurityLevel  *IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel `json:"advancedThreatProtectionRequiredSecurityLevel,omitempty"`
	Assignments                                    *[]DeviceCompliancePolicyAssignment                               `json:"assignments,omitempty"`
	CreatedDateTime                                *string                                                           `json:"createdDateTime,omitempty"`
	Description                                    *string                                                           `json:"description,omitempty"`
	DeviceSettingStateSummaries                    *[]SettingStateDeviceSummary                                      `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                           *DeviceComplianceDeviceOverview                                   `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                                 *[]DeviceComplianceDeviceStatus                                   `json:"deviceStatuses,omitempty"`
	DeviceThreatProtectionEnabled                  *bool                                                             `json:"deviceThreatProtectionEnabled,omitempty"`
	DeviceThreatProtectionRequiredSecurityLevel    *IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel   `json:"deviceThreatProtectionRequiredSecurityLevel,omitempty"`
	DisplayName                                    *string                                                           `json:"displayName,omitempty"`
	Id                                             *string                                                           `json:"id,omitempty"`
	LastModifiedDateTime                           *string                                                           `json:"lastModifiedDateTime,omitempty"`
	ManagedEmailProfileRequired                    *bool                                                             `json:"managedEmailProfileRequired,omitempty"`
	ODataType                                      *string                                                           `json:"@odata.type,omitempty"`
	OsMaximumBuildVersion                          *string                                                           `json:"osMaximumBuildVersion,omitempty"`
	OsMaximumVersion                               *string                                                           `json:"osMaximumVersion,omitempty"`
	OsMinimumBuildVersion                          *string                                                           `json:"osMinimumBuildVersion,omitempty"`
	OsMinimumVersion                               *string                                                           `json:"osMinimumVersion,omitempty"`
	PasscodeBlockSimple                            *bool                                                             `json:"passcodeBlockSimple,omitempty"`
	PasscodeExpirationDays                         *int64                                                            `json:"passcodeExpirationDays,omitempty"`
	PasscodeMinimumCharacterSetCount               *int64                                                            `json:"passcodeMinimumCharacterSetCount,omitempty"`
	PasscodeMinimumLength                          *int64                                                            `json:"passcodeMinimumLength,omitempty"`
	PasscodeMinutesOfInactivityBeforeLock          *int64                                                            `json:"passcodeMinutesOfInactivityBeforeLock,omitempty"`
	PasscodeMinutesOfInactivityBeforeScreenTimeout *int64                                                            `json:"passcodeMinutesOfInactivityBeforeScreenTimeout,omitempty"`
	PasscodePreviousPasscodeBlockCount             *int64                                                            `json:"passcodePreviousPasscodeBlockCount,omitempty"`
	PasscodeRequired                               *bool                                                             `json:"passcodeRequired,omitempty"`
	PasscodeRequiredType                           *IosCompliancePolicyPasscodeRequiredType                          `json:"passcodeRequiredType,omitempty"`
	RestrictedApps                                 *[]AppListItem                                                    `json:"restrictedApps,omitempty"`
	RoleScopeTagIds                                *[]string                                                         `json:"roleScopeTagIds,omitempty"`
	ScheduledActionsForRule                        *[]DeviceComplianceScheduledActionForRule                         `json:"scheduledActionsForRule,omitempty"`
	SecurityBlockJailbrokenDevices                 *bool                                                             `json:"securityBlockJailbrokenDevices,omitempty"`
	UserStatusOverview                             *DeviceComplianceUserOverview                                     `json:"userStatusOverview,omitempty"`
	UserStatuses                                   *[]DeviceComplianceUserStatus                                     `json:"userStatuses,omitempty"`
	Version                                        *int64                                                            `json:"version,omitempty"`
}
