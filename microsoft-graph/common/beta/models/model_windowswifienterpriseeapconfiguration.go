package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsWifiEnterpriseEAPConfiguration struct {
	Assignments                                  *[]DeviceConfigurationAssignment                                            `json:"assignments,omitempty"`
	AuthenticationMethod                         *WindowsWifiEnterpriseEAPConfigurationAuthenticationMethod                  `json:"authenticationMethod,omitempty"`
	AuthenticationPeriodInSeconds                *int64                                                                      `json:"authenticationPeriodInSeconds,omitempty"`
	AuthenticationRetryDelayPeriodInSeconds      *int64                                                                      `json:"authenticationRetryDelayPeriodInSeconds,omitempty"`
	AuthenticationType                           *WindowsWifiEnterpriseEAPConfigurationAuthenticationType                    `json:"authenticationType,omitempty"`
	CacheCredentials                             *bool                                                                       `json:"cacheCredentials,omitempty"`
	ConnectAutomatically                         *bool                                                                       `json:"connectAutomatically,omitempty"`
	ConnectToPreferredNetwork                    *bool                                                                       `json:"connectToPreferredNetwork,omitempty"`
	ConnectWhenNetworkNameIsHidden               *bool                                                                       `json:"connectWhenNetworkNameIsHidden,omitempty"`
	CreatedDateTime                              *string                                                                     `json:"createdDateTime,omitempty"`
	Description                                  *string                                                                     `json:"description,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode  *DeviceManagementApplicabilityRuleDeviceMode                                `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	DeviceManagementApplicabilityRuleOsEdition   *DeviceManagementApplicabilityRuleOsEdition                                 `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOsVersion   *DeviceManagementApplicabilityRuleOsVersion                                 `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceSettingStateSummaries                  *[]SettingStateDeviceSummary                                                `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                         *DeviceConfigurationDeviceOverview                                          `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                               *[]DeviceConfigurationDeviceStatus                                          `json:"deviceStatuses,omitempty"`
	DisableUserPromptForServerValidation         *bool                                                                       `json:"disableUserPromptForServerValidation,omitempty"`
	DisplayName                                  *string                                                                     `json:"displayName,omitempty"`
	EapType                                      *WindowsWifiEnterpriseEAPConfigurationEapType                               `json:"eapType,omitempty"`
	EapolStartPeriodInSeconds                    *int64                                                                      `json:"eapolStartPeriodInSeconds,omitempty"`
	EnablePairwiseMasterKeyCaching               *bool                                                                       `json:"enablePairwiseMasterKeyCaching,omitempty"`
	EnablePreAuthentication                      *bool                                                                       `json:"enablePreAuthentication,omitempty"`
	ForceFIPSCompliance                          *bool                                                                       `json:"forceFIPSCompliance,omitempty"`
	GroupAssignments                             *[]DeviceConfigurationGroupAssignment                                       `json:"groupAssignments,omitempty"`
	Id                                           *string                                                                     `json:"id,omitempty"`
	IdentityCertificateForClientAuthentication   *WindowsCertificateProfileBase                                              `json:"identityCertificateForClientAuthentication,omitempty"`
	InnerAuthenticationProtocolForEAPTTLS        *WindowsWifiEnterpriseEAPConfigurationInnerAuthenticationProtocolForEAPTTLS `json:"innerAuthenticationProtocolForEAPTTLS,omitempty"`
	LastModifiedDateTime                         *string                                                                     `json:"lastModifiedDateTime,omitempty"`
	MaximumAuthenticationFailures                *int64                                                                      `json:"maximumAuthenticationFailures,omitempty"`
	MaximumAuthenticationTimeoutInSeconds        *int64                                                                      `json:"maximumAuthenticationTimeoutInSeconds,omitempty"`
	MaximumEAPOLStartMessages                    *int64                                                                      `json:"maximumEAPOLStartMessages,omitempty"`
	MaximumNumberOfPairwiseMasterKeysInCache     *int64                                                                      `json:"maximumNumberOfPairwiseMasterKeysInCache,omitempty"`
	MaximumPairwiseMasterKeyCacheTimeInMinutes   *int64                                                                      `json:"maximumPairwiseMasterKeyCacheTimeInMinutes,omitempty"`
	MaximumPreAuthenticationAttempts             *int64                                                                      `json:"maximumPreAuthenticationAttempts,omitempty"`
	MeteredConnectionLimit                       *WindowsWifiEnterpriseEAPConfigurationMeteredConnectionLimit                `json:"meteredConnectionLimit,omitempty"`
	NetworkName                                  *string                                                                     `json:"networkName,omitempty"`
	NetworkSingleSignOn                          *WindowsWifiEnterpriseEAPConfigurationNetworkSingleSignOn                   `json:"networkSingleSignOn,omitempty"`
	ODataType                                    *string                                                                     `json:"@odata.type,omitempty"`
	OuterIdentityPrivacyTemporaryValue           *string                                                                     `json:"outerIdentityPrivacyTemporaryValue,omitempty"`
	PerformServerValidation                      *bool                                                                       `json:"performServerValidation,omitempty"`
	PreSharedKey                                 *string                                                                     `json:"preSharedKey,omitempty"`
	PromptForAdditionalAuthenticationCredentials *bool                                                                       `json:"promptForAdditionalAuthenticationCredentials,omitempty"`
	ProxyAutomaticConfigurationUrl               *string                                                                     `json:"proxyAutomaticConfigurationUrl,omitempty"`
	ProxyManualAddress                           *string                                                                     `json:"proxyManualAddress,omitempty"`
	ProxyManualPort                              *int64                                                                      `json:"proxyManualPort,omitempty"`
	ProxySetting                                 *WindowsWifiEnterpriseEAPConfigurationProxySetting                          `json:"proxySetting,omitempty"`
	RequireCryptographicBinding                  *bool                                                                       `json:"requireCryptographicBinding,omitempty"`
	RoleScopeTagIds                              *[]string                                                                   `json:"roleScopeTagIds,omitempty"`
	RootCertificateForClientValidation           *Windows81TrustedRootCertificate                                            `json:"rootCertificateForClientValidation,omitempty"`
	RootCertificatesForServerValidation          *[]Windows81TrustedRootCertificate                                          `json:"rootCertificatesForServerValidation,omitempty"`
	Ssid                                         *string                                                                     `json:"ssid,omitempty"`
	SupportsScopeTags                            *bool                                                                       `json:"supportsScopeTags,omitempty"`
	TrustedServerCertificateNames                *[]string                                                                   `json:"trustedServerCertificateNames,omitempty"`
	UserBasedVirtualLan                          *bool                                                                       `json:"userBasedVirtualLan,omitempty"`
	UserStatusOverview                           *DeviceConfigurationUserOverview                                            `json:"userStatusOverview,omitempty"`
	UserStatuses                                 *[]DeviceConfigurationUserStatus                                            `json:"userStatuses,omitempty"`
	Version                                      *int64                                                                      `json:"version,omitempty"`
	WifiSecurityType                             *WindowsWifiEnterpriseEAPConfigurationWifiSecurityType                      `json:"wifiSecurityType,omitempty"`
}
