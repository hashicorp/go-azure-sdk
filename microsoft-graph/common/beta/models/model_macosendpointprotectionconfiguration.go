package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSEndpointProtectionConfiguration struct {
	AdvancedThreatProtectionAutomaticSampleSubmission    *MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmission `json:"advancedThreatProtectionAutomaticSampleSubmission,omitempty"`
	AdvancedThreatProtectionCloudDelivered               *MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivered            `json:"advancedThreatProtectionCloudDelivered,omitempty"`
	AdvancedThreatProtectionDiagnosticDataCollection     *MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollection  `json:"advancedThreatProtectionDiagnosticDataCollection,omitempty"`
	AdvancedThreatProtectionExcludedExtensions           *[]string                                                                              `json:"advancedThreatProtectionExcludedExtensions,omitempty"`
	AdvancedThreatProtectionExcludedFiles                *[]string                                                                              `json:"advancedThreatProtectionExcludedFiles,omitempty"`
	AdvancedThreatProtectionExcludedFolders              *[]string                                                                              `json:"advancedThreatProtectionExcludedFolders,omitempty"`
	AdvancedThreatProtectionExcludedProcesses            *[]string                                                                              `json:"advancedThreatProtectionExcludedProcesses,omitempty"`
	AdvancedThreatProtectionRealTime                     *MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTime                  `json:"advancedThreatProtectionRealTime,omitempty"`
	Assignments                                          *[]DeviceConfigurationAssignment                                                       `json:"assignments,omitempty"`
	CreatedDateTime                                      *string                                                                                `json:"createdDateTime,omitempty"`
	Description                                          *string                                                                                `json:"description,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode          *DeviceManagementApplicabilityRuleDeviceMode                                           `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	DeviceManagementApplicabilityRuleOsEdition           *DeviceManagementApplicabilityRuleOsEdition                                            `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOsVersion           *DeviceManagementApplicabilityRuleOsVersion                                            `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceSettingStateSummaries                          *[]SettingStateDeviceSummary                                                           `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                                 *DeviceConfigurationDeviceOverview                                                     `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                                       *[]DeviceConfigurationDeviceStatus                                                     `json:"deviceStatuses,omitempty"`
	DisplayName                                          *string                                                                                `json:"displayName,omitempty"`
	FileVaultAllowDeferralUntilSignOut                   *bool                                                                                  `json:"fileVaultAllowDeferralUntilSignOut,omitempty"`
	FileVaultDisablePromptAtSignOut                      *bool                                                                                  `json:"fileVaultDisablePromptAtSignOut,omitempty"`
	FileVaultEnabled                                     *bool                                                                                  `json:"fileVaultEnabled,omitempty"`
	FileVaultHidePersonalRecoveryKey                     *bool                                                                                  `json:"fileVaultHidePersonalRecoveryKey,omitempty"`
	FileVaultInstitutionalRecoveryKeyCertificate         *string                                                                                `json:"fileVaultInstitutionalRecoveryKeyCertificate,omitempty"`
	FileVaultInstitutionalRecoveryKeyCertificateFileName *string                                                                                `json:"fileVaultInstitutionalRecoveryKeyCertificateFileName,omitempty"`
	FileVaultNumberOfTimesUserCanIgnore                  *int64                                                                                 `json:"fileVaultNumberOfTimesUserCanIgnore,omitempty"`
	FileVaultPersonalRecoveryKeyHelpMessage              *string                                                                                `json:"fileVaultPersonalRecoveryKeyHelpMessage,omitempty"`
	FileVaultPersonalRecoveryKeyRotationInMonths         *int64                                                                                 `json:"fileVaultPersonalRecoveryKeyRotationInMonths,omitempty"`
	FileVaultSelectedRecoveryKeyTypes                    *MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypes                 `json:"fileVaultSelectedRecoveryKeyTypes,omitempty"`
	FirewallApplications                                 *[]MacOSFirewallApplication                                                            `json:"firewallApplications,omitempty"`
	FirewallBlockAllIncoming                             *bool                                                                                  `json:"firewallBlockAllIncoming,omitempty"`
	FirewallEnableStealthMode                            *bool                                                                                  `json:"firewallEnableStealthMode,omitempty"`
	FirewallEnabled                                      *bool                                                                                  `json:"firewallEnabled,omitempty"`
	GatekeeperAllowedAppSource                           *MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource                        `json:"gatekeeperAllowedAppSource,omitempty"`
	GatekeeperBlockOverride                              *bool                                                                                  `json:"gatekeeperBlockOverride,omitempty"`
	GroupAssignments                                     *[]DeviceConfigurationGroupAssignment                                                  `json:"groupAssignments,omitempty"`
	Id                                                   *string                                                                                `json:"id,omitempty"`
	LastModifiedDateTime                                 *string                                                                                `json:"lastModifiedDateTime,omitempty"`
	ODataType                                            *string                                                                                `json:"@odata.type,omitempty"`
	RoleScopeTagIds                                      *[]string                                                                              `json:"roleScopeTagIds,omitempty"`
	SupportsScopeTags                                    *bool                                                                                  `json:"supportsScopeTags,omitempty"`
	UserStatusOverview                                   *DeviceConfigurationUserOverview                                                       `json:"userStatusOverview,omitempty"`
	UserStatuses                                         *[]DeviceConfigurationUserStatus                                                       `json:"userStatuses,omitempty"`
	Version                                              *int64                                                                                 `json:"version,omitempty"`
}
