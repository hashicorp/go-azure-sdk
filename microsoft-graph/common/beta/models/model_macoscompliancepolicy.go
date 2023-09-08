package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSCompliancePolicy struct {
	AdvancedThreatProtectionRequiredSecurityLevel *MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel `json:"advancedThreatProtectionRequiredSecurityLevel,omitempty"`
	Assignments                                   *[]DeviceCompliancePolicyAssignment                                 `json:"assignments,omitempty"`
	CreatedDateTime                               *string                                                             `json:"createdDateTime,omitempty"`
	Description                                   *string                                                             `json:"description,omitempty"`
	DeviceSettingStateSummaries                   *[]SettingStateDeviceSummary                                        `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                          *DeviceComplianceDeviceOverview                                     `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                                *[]DeviceComplianceDeviceStatus                                     `json:"deviceStatuses,omitempty"`
	DeviceThreatProtectionEnabled                 *bool                                                               `json:"deviceThreatProtectionEnabled,omitempty"`
	DeviceThreatProtectionRequiredSecurityLevel   *MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel   `json:"deviceThreatProtectionRequiredSecurityLevel,omitempty"`
	DisplayName                                   *string                                                             `json:"displayName,omitempty"`
	FirewallBlockAllIncoming                      *bool                                                               `json:"firewallBlockAllIncoming,omitempty"`
	FirewallEnableStealthMode                     *bool                                                               `json:"firewallEnableStealthMode,omitempty"`
	FirewallEnabled                               *bool                                                               `json:"firewallEnabled,omitempty"`
	GatekeeperAllowedAppSource                    *MacOSCompliancePolicyGatekeeperAllowedAppSource                    `json:"gatekeeperAllowedAppSource,omitempty"`
	Id                                            *string                                                             `json:"id,omitempty"`
	LastModifiedDateTime                          *string                                                             `json:"lastModifiedDateTime,omitempty"`
	ODataType                                     *string                                                             `json:"@odata.type,omitempty"`
	OsMaximumBuildVersion                         *string                                                             `json:"osMaximumBuildVersion,omitempty"`
	OsMaximumVersion                              *string                                                             `json:"osMaximumVersion,omitempty"`
	OsMinimumBuildVersion                         *string                                                             `json:"osMinimumBuildVersion,omitempty"`
	OsMinimumVersion                              *string                                                             `json:"osMinimumVersion,omitempty"`
	PasswordBlockSimple                           *bool                                                               `json:"passwordBlockSimple,omitempty"`
	PasswordExpirationDays                        *int64                                                              `json:"passwordExpirationDays,omitempty"`
	PasswordMinimumCharacterSetCount              *int64                                                              `json:"passwordMinimumCharacterSetCount,omitempty"`
	PasswordMinimumLength                         *int64                                                              `json:"passwordMinimumLength,omitempty"`
	PasswordMinutesOfInactivityBeforeLock         *int64                                                              `json:"passwordMinutesOfInactivityBeforeLock,omitempty"`
	PasswordPreviousPasswordBlockCount            *int64                                                              `json:"passwordPreviousPasswordBlockCount,omitempty"`
	PasswordRequired                              *bool                                                               `json:"passwordRequired,omitempty"`
	PasswordRequiredType                          *MacOSCompliancePolicyPasswordRequiredType                          `json:"passwordRequiredType,omitempty"`
	RoleScopeTagIds                               *[]string                                                           `json:"roleScopeTagIds,omitempty"`
	ScheduledActionsForRule                       *[]DeviceComplianceScheduledActionForRule                           `json:"scheduledActionsForRule,omitempty"`
	StorageRequireEncryption                      *bool                                                               `json:"storageRequireEncryption,omitempty"`
	SystemIntegrityProtectionEnabled              *bool                                                               `json:"systemIntegrityProtectionEnabled,omitempty"`
	UserStatusOverview                            *DeviceComplianceUserOverview                                       `json:"userStatusOverview,omitempty"`
	UserStatuses                                  *[]DeviceComplianceUserStatus                                       `json:"userStatuses,omitempty"`
	Version                                       *int64                                                              `json:"version,omitempty"`
}
