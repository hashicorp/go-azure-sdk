package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerEnterpriseWiFiConfiguration struct {
	Assignments                                 *[]DeviceConfigurationAssignment                                                    `json:"assignments,omitempty"`
	AuthenticationMethod                        *AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod                  `json:"authenticationMethod,omitempty"`
	ConnectAutomatically                        *bool                                                                               `json:"connectAutomatically,omitempty"`
	ConnectWhenNetworkNameIsHidden              *bool                                                                               `json:"connectWhenNetworkNameIsHidden,omitempty"`
	CreatedDateTime                             *string                                                                             `json:"createdDateTime,omitempty"`
	DerivedCredentialSettings                   *DeviceManagementDerivedCredentialSettings                                          `json:"derivedCredentialSettings,omitempty"`
	Description                                 *string                                                                             `json:"description,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode                                        `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	DeviceManagementApplicabilityRuleOsEdition  *DeviceManagementApplicabilityRuleOsEdition                                         `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOsVersion  *DeviceManagementApplicabilityRuleOsVersion                                         `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceSettingStateSummaries                 *[]SettingStateDeviceSummary                                                        `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                        *DeviceConfigurationDeviceOverview                                                  `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                              *[]DeviceConfigurationDeviceStatus                                                  `json:"deviceStatuses,omitempty"`
	DisplayName                                 *string                                                                             `json:"displayName,omitempty"`
	EapType                                     *AndroidDeviceOwnerEnterpriseWiFiConfigurationEapType                               `json:"eapType,omitempty"`
	GroupAssignments                            *[]DeviceConfigurationGroupAssignment                                               `json:"groupAssignments,omitempty"`
	Id                                          *string                                                                             `json:"id,omitempty"`
	IdentityCertificateForClientAuthentication  *AndroidDeviceOwnerCertificateProfileBase                                           `json:"identityCertificateForClientAuthentication,omitempty"`
	InnerAuthenticationProtocolForEapTtls       *AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls `json:"innerAuthenticationProtocolForEapTtls,omitempty"`
	InnerAuthenticationProtocolForPeap          *AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap    `json:"innerAuthenticationProtocolForPeap,omitempty"`
	LastModifiedDateTime                        *string                                                                             `json:"lastModifiedDateTime,omitempty"`
	NetworkName                                 *string                                                                             `json:"networkName,omitempty"`
	ODataType                                   *string                                                                             `json:"@odata.type,omitempty"`
	OuterIdentityPrivacyTemporaryValue          *string                                                                             `json:"outerIdentityPrivacyTemporaryValue,omitempty"`
	PreSharedKey                                *string                                                                             `json:"preSharedKey,omitempty"`
	PreSharedKeyIsSet                           *bool                                                                               `json:"preSharedKeyIsSet,omitempty"`
	ProxyAutomaticConfigurationUrl              *string                                                                             `json:"proxyAutomaticConfigurationUrl,omitempty"`
	ProxyExclusionList                          *string                                                                             `json:"proxyExclusionList,omitempty"`
	ProxyManualAddress                          *string                                                                             `json:"proxyManualAddress,omitempty"`
	ProxyManualPort                             *int64                                                                              `json:"proxyManualPort,omitempty"`
	ProxySettings                               *AndroidDeviceOwnerEnterpriseWiFiConfigurationProxySettings                         `json:"proxySettings,omitempty"`
	RoleScopeTagIds                             *[]string                                                                           `json:"roleScopeTagIds,omitempty"`
	RootCertificateForServerValidation          *AndroidDeviceOwnerTrustedRootCertificate                                           `json:"rootCertificateForServerValidation,omitempty"`
	Ssid                                        *string                                                                             `json:"ssid,omitempty"`
	SupportsScopeTags                           *bool                                                                               `json:"supportsScopeTags,omitempty"`
	TrustedServerCertificateNames               *[]string                                                                           `json:"trustedServerCertificateNames,omitempty"`
	UserStatusOverview                          *DeviceConfigurationUserOverview                                                    `json:"userStatusOverview,omitempty"`
	UserStatuses                                *[]DeviceConfigurationUserStatus                                                    `json:"userStatuses,omitempty"`
	Version                                     *int64                                                                              `json:"version,omitempty"`
	WiFiSecurityType                            *AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType                      `json:"wiFiSecurityType,omitempty"`
}
