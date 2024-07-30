package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10CompliancePolicy struct {
	ActiveFirewallRequired                      *bool                                                                 `json:"activeFirewallRequired,omitempty"`
	AntiSpywareRequired                         *bool                                                                 `json:"antiSpywareRequired,omitempty"`
	AntivirusRequired                           *bool                                                                 `json:"antivirusRequired,omitempty"`
	Assignments                                 *[]DeviceCompliancePolicyAssignment                                   `json:"assignments,omitempty"`
	BitLockerEnabled                            *bool                                                                 `json:"bitLockerEnabled,omitempty"`
	CodeIntegrityEnabled                        *bool                                                                 `json:"codeIntegrityEnabled,omitempty"`
	ConfigurationManagerComplianceRequired      *bool                                                                 `json:"configurationManagerComplianceRequired,omitempty"`
	CreatedDateTime                             *string                                                               `json:"createdDateTime,omitempty"`
	DefenderEnabled                             *bool                                                                 `json:"defenderEnabled,omitempty"`
	DefenderVersion                             *string                                                               `json:"defenderVersion,omitempty"`
	Description                                 *string                                                               `json:"description,omitempty"`
	DeviceCompliancePolicyScript                *DeviceCompliancePolicyScript                                         `json:"deviceCompliancePolicyScript,omitempty"`
	DeviceSettingStateSummaries                 *[]SettingStateDeviceSummary                                          `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                        *DeviceComplianceDeviceOverview                                       `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                              *[]DeviceComplianceDeviceStatus                                       `json:"deviceStatuses,omitempty"`
	DeviceThreatProtectionEnabled               *bool                                                                 `json:"deviceThreatProtectionEnabled,omitempty"`
	DeviceThreatProtectionRequiredSecurityLevel *Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel `json:"deviceThreatProtectionRequiredSecurityLevel,omitempty"`
	DisplayName                                 *string                                                               `json:"displayName,omitempty"`
	EarlyLaunchAntiMalwareDriverEnabled         *bool                                                                 `json:"earlyLaunchAntiMalwareDriverEnabled,omitempty"`
	FirmwareProtectionEnabled                   *bool                                                                 `json:"firmwareProtectionEnabled,omitempty"`
	Id                                          *string                                                               `json:"id,omitempty"`
	KernelDmaProtectionEnabled                  *bool                                                                 `json:"kernelDmaProtectionEnabled,omitempty"`
	LastModifiedDateTime                        *string                                                               `json:"lastModifiedDateTime,omitempty"`
	MemoryIntegrityEnabled                      *bool                                                                 `json:"memoryIntegrityEnabled,omitempty"`
	MobileOsMaximumVersion                      *string                                                               `json:"mobileOsMaximumVersion,omitempty"`
	MobileOsMinimumVersion                      *string                                                               `json:"mobileOsMinimumVersion,omitempty"`
	ODataType                                   *string                                                               `json:"@odata.type,omitempty"`
	OsMaximumVersion                            *string                                                               `json:"osMaximumVersion,omitempty"`
	OsMinimumVersion                            *string                                                               `json:"osMinimumVersion,omitempty"`
	PasswordBlockSimple                         *bool                                                                 `json:"passwordBlockSimple,omitempty"`
	PasswordExpirationDays                      *int64                                                                `json:"passwordExpirationDays,omitempty"`
	PasswordMinimumCharacterSetCount            *int64                                                                `json:"passwordMinimumCharacterSetCount,omitempty"`
	PasswordMinimumLength                       *int64                                                                `json:"passwordMinimumLength,omitempty"`
	PasswordMinutesOfInactivityBeforeLock       *int64                                                                `json:"passwordMinutesOfInactivityBeforeLock,omitempty"`
	PasswordPreviousPasswordBlockCount          *int64                                                                `json:"passwordPreviousPasswordBlockCount,omitempty"`
	PasswordRequired                            *bool                                                                 `json:"passwordRequired,omitempty"`
	PasswordRequiredToUnlockFromIdle            *bool                                                                 `json:"passwordRequiredToUnlockFromIdle,omitempty"`
	PasswordRequiredType                        *Windows10CompliancePolicyPasswordRequiredType                        `json:"passwordRequiredType,omitempty"`
	RequireHealthyDeviceReport                  *bool                                                                 `json:"requireHealthyDeviceReport,omitempty"`
	RoleScopeTagIds                             *[]string                                                             `json:"roleScopeTagIds,omitempty"`
	RtpEnabled                                  *bool                                                                 `json:"rtpEnabled,omitempty"`
	ScheduledActionsForRule                     *[]DeviceComplianceScheduledActionForRule                             `json:"scheduledActionsForRule,omitempty"`
	SecureBootEnabled                           *bool                                                                 `json:"secureBootEnabled,omitempty"`
	SignatureOutOfDate                          *bool                                                                 `json:"signatureOutOfDate,omitempty"`
	StorageRequireEncryption                    *bool                                                                 `json:"storageRequireEncryption,omitempty"`
	TpmRequired                                 *bool                                                                 `json:"tpmRequired,omitempty"`
	UserStatusOverview                          *DeviceComplianceUserOverview                                         `json:"userStatusOverview,omitempty"`
	UserStatuses                                *[]DeviceComplianceUserStatus                                         `json:"userStatuses,omitempty"`
	ValidOperatingSystemBuildRanges             *[]OperatingSystemVersionRange                                        `json:"validOperatingSystemBuildRanges,omitempty"`
	Version                                     *int64                                                                `json:"version,omitempty"`
	VirtualizationBasedSecurityEnabled          *bool                                                                 `json:"virtualizationBasedSecurityEnabled,omitempty"`
}
