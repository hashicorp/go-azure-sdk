package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsWiredNetworkConfiguration struct {
	Assignments                                         *[]DeviceConfigurationAssignment                                       `json:"assignments,omitempty"`
	AuthenticationBlockPeriodInMinutes                  *int64                                                                 `json:"authenticationBlockPeriodInMinutes,omitempty"`
	AuthenticationMethod                                *WindowsWiredNetworkConfigurationAuthenticationMethod                  `json:"authenticationMethod,omitempty"`
	AuthenticationPeriodInSeconds                       *int64                                                                 `json:"authenticationPeriodInSeconds,omitempty"`
	AuthenticationRetryDelayPeriodInSeconds             *int64                                                                 `json:"authenticationRetryDelayPeriodInSeconds,omitempty"`
	AuthenticationType                                  *WindowsWiredNetworkConfigurationAuthenticationType                    `json:"authenticationType,omitempty"`
	CacheCredentials                                    *bool                                                                  `json:"cacheCredentials,omitempty"`
	CreatedDateTime                                     *string                                                                `json:"createdDateTime,omitempty"`
	Description                                         *string                                                                `json:"description,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode         *DeviceManagementApplicabilityRuleDeviceMode                           `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	DeviceManagementApplicabilityRuleOsEdition          *DeviceManagementApplicabilityRuleOsEdition                            `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOsVersion          *DeviceManagementApplicabilityRuleOsVersion                            `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceSettingStateSummaries                         *[]SettingStateDeviceSummary                                           `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                                *DeviceConfigurationDeviceOverview                                     `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                                      *[]DeviceConfigurationDeviceStatus                                     `json:"deviceStatuses,omitempty"`
	DisableUserPromptForServerValidation                *bool                                                                  `json:"disableUserPromptForServerValidation,omitempty"`
	DisplayName                                         *string                                                                `json:"displayName,omitempty"`
	EapType                                             *WindowsWiredNetworkConfigurationEapType                               `json:"eapType,omitempty"`
	EapolStartPeriodInSeconds                           *int64                                                                 `json:"eapolStartPeriodInSeconds,omitempty"`
	Enforce8021X                                        *bool                                                                  `json:"enforce8021X,omitempty"`
	ForceFIPSCompliance                                 *bool                                                                  `json:"forceFIPSCompliance,omitempty"`
	GroupAssignments                                    *[]DeviceConfigurationGroupAssignment                                  `json:"groupAssignments,omitempty"`
	Id                                                  *string                                                                `json:"id,omitempty"`
	IdentityCertificateForClientAuthentication          *WindowsCertificateProfileBase                                         `json:"identityCertificateForClientAuthentication,omitempty"`
	InnerAuthenticationProtocolForEAPTTLS               *WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS `json:"innerAuthenticationProtocolForEAPTTLS,omitempty"`
	LastModifiedDateTime                                *string                                                                `json:"lastModifiedDateTime,omitempty"`
	MaximumAuthenticationFailures                       *int64                                                                 `json:"maximumAuthenticationFailures,omitempty"`
	MaximumEAPOLStartMessages                           *int64                                                                 `json:"maximumEAPOLStartMessages,omitempty"`
	ODataType                                           *string                                                                `json:"@odata.type,omitempty"`
	OuterIdentityPrivacyTemporaryValue                  *string                                                                `json:"outerIdentityPrivacyTemporaryValue,omitempty"`
	PerformServerValidation                             *bool                                                                  `json:"performServerValidation,omitempty"`
	RequireCryptographicBinding                         *bool                                                                  `json:"requireCryptographicBinding,omitempty"`
	RoleScopeTagIds                                     *[]string                                                              `json:"roleScopeTagIds,omitempty"`
	RootCertificateForClientValidation                  *Windows81TrustedRootCertificate                                       `json:"rootCertificateForClientValidation,omitempty"`
	RootCertificatesForServerValidation                 *[]Windows81TrustedRootCertificate                                     `json:"rootCertificatesForServerValidation,omitempty"`
	SecondaryAuthenticationMethod                       *WindowsWiredNetworkConfigurationSecondaryAuthenticationMethod         `json:"secondaryAuthenticationMethod,omitempty"`
	SecondaryIdentityCertificateForClientAuthentication *WindowsCertificateProfileBase                                         `json:"secondaryIdentityCertificateForClientAuthentication,omitempty"`
	SecondaryRootCertificateForClientValidation         *Windows81TrustedRootCertificate                                       `json:"secondaryRootCertificateForClientValidation,omitempty"`
	SupportsScopeTags                                   *bool                                                                  `json:"supportsScopeTags,omitempty"`
	TrustedServerCertificateNames                       *[]string                                                              `json:"trustedServerCertificateNames,omitempty"`
	UserStatusOverview                                  *DeviceConfigurationUserOverview                                       `json:"userStatusOverview,omitempty"`
	UserStatuses                                        *[]DeviceConfigurationUserStatus                                       `json:"userStatuses,omitempty"`
	Version                                             *int64                                                                 `json:"version,omitempty"`
}
