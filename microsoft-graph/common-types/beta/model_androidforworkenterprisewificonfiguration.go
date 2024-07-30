package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkEnterpriseWiFiConfiguration struct {
	Assignments                                 *[]DeviceConfigurationAssignment                                                `json:"assignments,omitempty"`
	AuthenticationMethod                        *AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethod                  `json:"authenticationMethod,omitempty"`
	ConnectAutomatically                        *bool                                                                           `json:"connectAutomatically,omitempty"`
	ConnectWhenNetworkNameIsHidden              *bool                                                                           `json:"connectWhenNetworkNameIsHidden,omitempty"`
	CreatedDateTime                             *string                                                                         `json:"createdDateTime,omitempty"`
	Description                                 *string                                                                         `json:"description,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode                                    `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	DeviceManagementApplicabilityRuleOsEdition  *DeviceManagementApplicabilityRuleOsEdition                                     `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOsVersion  *DeviceManagementApplicabilityRuleOsVersion                                     `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceSettingStateSummaries                 *[]SettingStateDeviceSummary                                                    `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                        *DeviceConfigurationDeviceOverview                                              `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                              *[]DeviceConfigurationDeviceStatus                                              `json:"deviceStatuses,omitempty"`
	DisplayName                                 *string                                                                         `json:"displayName,omitempty"`
	EapType                                     *AndroidForWorkEnterpriseWiFiConfigurationEapType                               `json:"eapType,omitempty"`
	GroupAssignments                            *[]DeviceConfigurationGroupAssignment                                           `json:"groupAssignments,omitempty"`
	Id                                          *string                                                                         `json:"id,omitempty"`
	IdentityCertificateForClientAuthentication  *AndroidForWorkCertificateProfileBase                                           `json:"identityCertificateForClientAuthentication,omitempty"`
	InnerAuthenticationProtocolForEapTtls       *AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls `json:"innerAuthenticationProtocolForEapTtls,omitempty"`
	InnerAuthenticationProtocolForPeap          *AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap    `json:"innerAuthenticationProtocolForPeap,omitempty"`
	LastModifiedDateTime                        *string                                                                         `json:"lastModifiedDateTime,omitempty"`
	NetworkName                                 *string                                                                         `json:"networkName,omitempty"`
	ODataType                                   *string                                                                         `json:"@odata.type,omitempty"`
	OuterIdentityPrivacyTemporaryValue          *string                                                                         `json:"outerIdentityPrivacyTemporaryValue,omitempty"`
	RoleScopeTagIds                             *[]string                                                                       `json:"roleScopeTagIds,omitempty"`
	RootCertificateForServerValidation          *AndroidForWorkTrustedRootCertificate                                           `json:"rootCertificateForServerValidation,omitempty"`
	Ssid                                        *string                                                                         `json:"ssid,omitempty"`
	SupportsScopeTags                           *bool                                                                           `json:"supportsScopeTags,omitempty"`
	TrustedServerCertificateNames               *[]string                                                                       `json:"trustedServerCertificateNames,omitempty"`
	UserStatusOverview                          *DeviceConfigurationUserOverview                                                `json:"userStatusOverview,omitempty"`
	UserStatuses                                *[]DeviceConfigurationUserStatus                                                `json:"userStatuses,omitempty"`
	Version                                     *int64                                                                          `json:"version,omitempty"`
	WiFiSecurityType                            *AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityType                      `json:"wiFiSecurityType,omitempty"`
}
