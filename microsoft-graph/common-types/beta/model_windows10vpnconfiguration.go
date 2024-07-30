package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10VpnConfiguration struct {
	Assignments                                 *[]DeviceConfigurationAssignment               `json:"assignments,omitempty"`
	AssociatedApps                              *[]Windows10AssociatedApps                     `json:"associatedApps,omitempty"`
	AuthenticationMethod                        *Windows10VpnConfigurationAuthenticationMethod `json:"authenticationMethod,omitempty"`
	ConnectionName                              *string                                        `json:"connectionName,omitempty"`
	ConnectionType                              *Windows10VpnConfigurationConnectionType       `json:"connectionType,omitempty"`
	CreatedDateTime                             *string                                        `json:"createdDateTime,omitempty"`
	CryptographySuite                           *CryptographySuite                             `json:"cryptographySuite,omitempty"`
	CustomXml                                   *string                                        `json:"customXml,omitempty"`
	Description                                 *string                                        `json:"description,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode   `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	DeviceManagementApplicabilityRuleOsEdition  *DeviceManagementApplicabilityRuleOsEdition    `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOsVersion  *DeviceManagementApplicabilityRuleOsVersion    `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceSettingStateSummaries                 *[]SettingStateDeviceSummary                   `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                        *DeviceConfigurationDeviceOverview             `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                              *[]DeviceConfigurationDeviceStatus             `json:"deviceStatuses,omitempty"`
	DisplayName                                 *string                                        `json:"displayName,omitempty"`
	DnsRules                                    *[]VpnDnsRule                                  `json:"dnsRules,omitempty"`
	DnsSuffixes                                 *[]string                                      `json:"dnsSuffixes,omitempty"`
	EapXml                                      *string                                        `json:"eapXml,omitempty"`
	EnableAlwaysOn                              *bool                                          `json:"enableAlwaysOn,omitempty"`
	EnableConditionalAccess                     *bool                                          `json:"enableConditionalAccess,omitempty"`
	EnableDeviceTunnel                          *bool                                          `json:"enableDeviceTunnel,omitempty"`
	EnableDnsRegistration                       *bool                                          `json:"enableDnsRegistration,omitempty"`
	EnableSingleSignOnWithAlternateCertificate  *bool                                          `json:"enableSingleSignOnWithAlternateCertificate,omitempty"`
	EnableSplitTunneling                        *bool                                          `json:"enableSplitTunneling,omitempty"`
	GroupAssignments                            *[]DeviceConfigurationGroupAssignment          `json:"groupAssignments,omitempty"`
	Id                                          *string                                        `json:"id,omitempty"`
	IdentityCertificate                         *WindowsCertificateProfileBase                 `json:"identityCertificate,omitempty"`
	LastModifiedDateTime                        *string                                        `json:"lastModifiedDateTime,omitempty"`
	MicrosoftTunnelSiteId                       *string                                        `json:"microsoftTunnelSiteId,omitempty"`
	ODataType                                   *string                                        `json:"@odata.type,omitempty"`
	OnlyAssociatedAppsCanUseConnection          *bool                                          `json:"onlyAssociatedAppsCanUseConnection,omitempty"`
	ProfileTarget                               *Windows10VpnConfigurationProfileTarget        `json:"profileTarget,omitempty"`
	ProxyServer                                 *Windows10VpnProxyServer                       `json:"proxyServer,omitempty"`
	RememberUserCredentials                     *bool                                          `json:"rememberUserCredentials,omitempty"`
	RoleScopeTagIds                             *[]string                                      `json:"roleScopeTagIds,omitempty"`
	Routes                                      *[]VpnRoute                                    `json:"routes,omitempty"`
	Servers                                     *[]VpnServer                                   `json:"servers,omitempty"`
	SingleSignOnEku                             *ExtendedKeyUsage                              `json:"singleSignOnEku,omitempty"`
	SingleSignOnIssuerHash                      *string                                        `json:"singleSignOnIssuerHash,omitempty"`
	SupportsScopeTags                           *bool                                          `json:"supportsScopeTags,omitempty"`
	TrafficRules                                *[]VpnTrafficRule                              `json:"trafficRules,omitempty"`
	TrustedNetworkDomains                       *[]string                                      `json:"trustedNetworkDomains,omitempty"`
	UserStatusOverview                          *DeviceConfigurationUserOverview               `json:"userStatusOverview,omitempty"`
	UserStatuses                                *[]DeviceConfigurationUserStatus               `json:"userStatuses,omitempty"`
	Version                                     *int64                                         `json:"version,omitempty"`
	WindowsInformationProtectionDomain          *string                                        `json:"windowsInformationProtectionDomain,omitempty"`
}
