package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEasEmailProfileConfiguration struct {
	AccountName                                   *string                                                   `json:"accountName,omitempty"`
	Assignments                                   *[]DeviceConfigurationAssignment                          `json:"assignments,omitempty"`
	AuthenticationMethod                          *IosEasEmailProfileConfigurationAuthenticationMethod      `json:"authenticationMethod,omitempty"`
	BlockMovingMessagesToOtherEmailAccounts       *bool                                                     `json:"blockMovingMessagesToOtherEmailAccounts,omitempty"`
	BlockSendingEmailFromThirdPartyApps           *bool                                                     `json:"blockSendingEmailFromThirdPartyApps,omitempty"`
	BlockSyncingRecentlyUsedEmailAddresses        *bool                                                     `json:"blockSyncingRecentlyUsedEmailAddresses,omitempty"`
	CreatedDateTime                               *string                                                   `json:"createdDateTime,omitempty"`
	CustomDomainName                              *string                                                   `json:"customDomainName,omitempty"`
	DerivedCredentialSettings                     *DeviceManagementDerivedCredentialSettings                `json:"derivedCredentialSettings,omitempty"`
	Description                                   *string                                                   `json:"description,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode   *DeviceManagementApplicabilityRuleDeviceMode              `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	DeviceManagementApplicabilityRuleOsEdition    *DeviceManagementApplicabilityRuleOsEdition               `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOsVersion    *DeviceManagementApplicabilityRuleOsVersion               `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceSettingStateSummaries                   *[]SettingStateDeviceSummary                              `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                          *DeviceConfigurationDeviceOverview                        `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                                *[]DeviceConfigurationDeviceStatus                        `json:"deviceStatuses,omitempty"`
	DisplayName                                   *string                                                   `json:"displayName,omitempty"`
	DurationOfEmailToSync                         *IosEasEmailProfileConfigurationDurationOfEmailToSync     `json:"durationOfEmailToSync,omitempty"`
	EasServices                                   *IosEasEmailProfileConfigurationEasServices               `json:"easServices,omitempty"`
	EasServicesUserOverrideEnabled                *bool                                                     `json:"easServicesUserOverrideEnabled,omitempty"`
	EmailAddressSource                            *IosEasEmailProfileConfigurationEmailAddressSource        `json:"emailAddressSource,omitempty"`
	EncryptionCertificateType                     *IosEasEmailProfileConfigurationEncryptionCertificateType `json:"encryptionCertificateType,omitempty"`
	GroupAssignments                              *[]DeviceConfigurationGroupAssignment                     `json:"groupAssignments,omitempty"`
	HostName                                      *string                                                   `json:"hostName,omitempty"`
	Id                                            *string                                                   `json:"id,omitempty"`
	IdentityCertificate                           *IosCertificateProfileBase                                `json:"identityCertificate,omitempty"`
	LastModifiedDateTime                          *string                                                   `json:"lastModifiedDateTime,omitempty"`
	ODataType                                     *string                                                   `json:"@odata.type,omitempty"`
	PerAppVPNProfileId                            *string                                                   `json:"perAppVPNProfileId,omitempty"`
	RequireSmime                                  *bool                                                     `json:"requireSmime,omitempty"`
	RequireSsl                                    *bool                                                     `json:"requireSsl,omitempty"`
	RoleScopeTagIds                               *[]string                                                 `json:"roleScopeTagIds,omitempty"`
	SigningCertificateType                        *IosEasEmailProfileConfigurationSigningCertificateType    `json:"signingCertificateType,omitempty"`
	SmimeEnablePerMessageSwitch                   *bool                                                     `json:"smimeEnablePerMessageSwitch,omitempty"`
	SmimeEncryptByDefaultEnabled                  *bool                                                     `json:"smimeEncryptByDefaultEnabled,omitempty"`
	SmimeEncryptByDefaultUserOverrideEnabled      *bool                                                     `json:"smimeEncryptByDefaultUserOverrideEnabled,omitempty"`
	SmimeEncryptionCertificate                    *IosCertificateProfile                                    `json:"smimeEncryptionCertificate,omitempty"`
	SmimeEncryptionCertificateUserOverrideEnabled *bool                                                     `json:"smimeEncryptionCertificateUserOverrideEnabled,omitempty"`
	SmimeSigningCertificate                       *IosCertificateProfile                                    `json:"smimeSigningCertificate,omitempty"`
	SmimeSigningCertificateUserOverrideEnabled    *bool                                                     `json:"smimeSigningCertificateUserOverrideEnabled,omitempty"`
	SmimeSigningEnabled                           *bool                                                     `json:"smimeSigningEnabled,omitempty"`
	SmimeSigningUserOverrideEnabled               *bool                                                     `json:"smimeSigningUserOverrideEnabled,omitempty"`
	SupportsScopeTags                             *bool                                                     `json:"supportsScopeTags,omitempty"`
	UseOAuth                                      *bool                                                     `json:"useOAuth,omitempty"`
	UserDomainNameSource                          *IosEasEmailProfileConfigurationUserDomainNameSource      `json:"userDomainNameSource,omitempty"`
	UserStatusOverview                            *DeviceConfigurationUserOverview                          `json:"userStatusOverview,omitempty"`
	UserStatuses                                  *[]DeviceConfigurationUserStatus                          `json:"userStatuses,omitempty"`
	UsernameAADSource                             *IosEasEmailProfileConfigurationUsernameAADSource         `json:"usernameAADSource,omitempty"`
	UsernameSource                                *IosEasEmailProfileConfigurationUsernameSource            `json:"usernameSource,omitempty"`
	Version                                       *int64                                                    `json:"version,omitempty"`
}
