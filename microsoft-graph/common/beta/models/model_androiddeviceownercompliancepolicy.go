package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerCompliancePolicy struct {
	AdvancedThreatProtectionRequiredSecurityLevel      *AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel `json:"advancedThreatProtectionRequiredSecurityLevel,omitempty"`
	Assignments                                        *[]DeviceCompliancePolicyAssignment                                              `json:"assignments,omitempty"`
	CreatedDateTime                                    *string                                                                          `json:"createdDateTime,omitempty"`
	Description                                        *string                                                                          `json:"description,omitempty"`
	DeviceSettingStateSummaries                        *[]SettingStateDeviceSummary                                                     `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                               *DeviceComplianceDeviceOverview                                                  `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                                     *[]DeviceComplianceDeviceStatus                                                  `json:"deviceStatuses,omitempty"`
	DeviceThreatProtectionEnabled                      *bool                                                                            `json:"deviceThreatProtectionEnabled,omitempty"`
	DeviceThreatProtectionRequiredSecurityLevel        *AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel   `json:"deviceThreatProtectionRequiredSecurityLevel,omitempty"`
	DisplayName                                        *string                                                                          `json:"displayName,omitempty"`
	Id                                                 *string                                                                          `json:"id,omitempty"`
	LastModifiedDateTime                               *string                                                                          `json:"lastModifiedDateTime,omitempty"`
	MinAndroidSecurityPatchLevel                       *string                                                                          `json:"minAndroidSecurityPatchLevel,omitempty"`
	ODataType                                          *string                                                                          `json:"@odata.type,omitempty"`
	OsMaximumVersion                                   *string                                                                          `json:"osMaximumVersion,omitempty"`
	OsMinimumVersion                                   *string                                                                          `json:"osMinimumVersion,omitempty"`
	PasswordExpirationDays                             *int64                                                                           `json:"passwordExpirationDays,omitempty"`
	PasswordMinimumLength                              *int64                                                                           `json:"passwordMinimumLength,omitempty"`
	PasswordMinimumLetterCharacters                    *int64                                                                           `json:"passwordMinimumLetterCharacters,omitempty"`
	PasswordMinimumLowerCaseCharacters                 *int64                                                                           `json:"passwordMinimumLowerCaseCharacters,omitempty"`
	PasswordMinimumNonLetterCharacters                 *int64                                                                           `json:"passwordMinimumNonLetterCharacters,omitempty"`
	PasswordMinimumNumericCharacters                   *int64                                                                           `json:"passwordMinimumNumericCharacters,omitempty"`
	PasswordMinimumSymbolCharacters                    *int64                                                                           `json:"passwordMinimumSymbolCharacters,omitempty"`
	PasswordMinimumUpperCaseCharacters                 *int64                                                                           `json:"passwordMinimumUpperCaseCharacters,omitempty"`
	PasswordMinutesOfInactivityBeforeLock              *int64                                                                           `json:"passwordMinutesOfInactivityBeforeLock,omitempty"`
	PasswordPreviousPasswordCountToBlock               *int64                                                                           `json:"passwordPreviousPasswordCountToBlock,omitempty"`
	PasswordRequired                                   *bool                                                                            `json:"passwordRequired,omitempty"`
	PasswordRequiredType                               *AndroidDeviceOwnerCompliancePolicyPasswordRequiredType                          `json:"passwordRequiredType,omitempty"`
	RoleScopeTagIds                                    *[]string                                                                        `json:"roleScopeTagIds,omitempty"`
	ScheduledActionsForRule                            *[]DeviceComplianceScheduledActionForRule                                        `json:"scheduledActionsForRule,omitempty"`
	SecurityRequireIntuneAppIntegrity                  *bool                                                                            `json:"securityRequireIntuneAppIntegrity,omitempty"`
	SecurityRequireSafetyNetAttestationBasicIntegrity  *bool                                                                            `json:"securityRequireSafetyNetAttestationBasicIntegrity,omitempty"`
	SecurityRequireSafetyNetAttestationCertifiedDevice *bool                                                                            `json:"securityRequireSafetyNetAttestationCertifiedDevice,omitempty"`
	StorageRequireEncryption                           *bool                                                                            `json:"storageRequireEncryption,omitempty"`
	UserStatusOverview                                 *DeviceComplianceUserOverview                                                    `json:"userStatusOverview,omitempty"`
	UserStatuses                                       *[]DeviceComplianceUserStatus                                                    `json:"userStatuses,omitempty"`
	Version                                            *int64                                                                           `json:"version,omitempty"`
}
