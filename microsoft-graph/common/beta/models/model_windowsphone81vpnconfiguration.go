package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81VpnConfiguration struct {
	ApplyOnlyToWindows81                        *bool                                               `json:"applyOnlyToWindows81,omitempty"`
	Assignments                                 *[]DeviceConfigurationAssignment                    `json:"assignments,omitempty"`
	AuthenticationMethod                        *WindowsPhone81VpnConfigurationAuthenticationMethod `json:"authenticationMethod,omitempty"`
	BypassVpnOnCompanyWifi                      *bool                                               `json:"bypassVpnOnCompanyWifi,omitempty"`
	BypassVpnOnHomeWifi                         *bool                                               `json:"bypassVpnOnHomeWifi,omitempty"`
	ConnectionName                              *string                                             `json:"connectionName,omitempty"`
	ConnectionType                              *WindowsPhone81VpnConfigurationConnectionType       `json:"connectionType,omitempty"`
	CreatedDateTime                             *string                                             `json:"createdDateTime,omitempty"`
	CustomXml                                   *string                                             `json:"customXml,omitempty"`
	Description                                 *string                                             `json:"description,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode        `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	DeviceManagementApplicabilityRuleOsEdition  *DeviceManagementApplicabilityRuleOsEdition         `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOsVersion  *DeviceManagementApplicabilityRuleOsVersion         `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceSettingStateSummaries                 *[]SettingStateDeviceSummary                        `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                        *DeviceConfigurationDeviceOverview                  `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                              *[]DeviceConfigurationDeviceStatus                  `json:"deviceStatuses,omitempty"`
	DisplayName                                 *string                                             `json:"displayName,omitempty"`
	DnsSuffixSearchList                         *[]string                                           `json:"dnsSuffixSearchList,omitempty"`
	EnableSplitTunneling                        *bool                                               `json:"enableSplitTunneling,omitempty"`
	GroupAssignments                            *[]DeviceConfigurationGroupAssignment               `json:"groupAssignments,omitempty"`
	Id                                          *string                                             `json:"id,omitempty"`
	IdentityCertificate                         *WindowsPhone81CertificateProfileBase               `json:"identityCertificate,omitempty"`
	LastModifiedDateTime                        *string                                             `json:"lastModifiedDateTime,omitempty"`
	LoginGroupOrDomain                          *string                                             `json:"loginGroupOrDomain,omitempty"`
	ODataType                                   *string                                             `json:"@odata.type,omitempty"`
	ProxyServer                                 *Windows81VpnProxyServer                            `json:"proxyServer,omitempty"`
	RememberUserCredentials                     *bool                                               `json:"rememberUserCredentials,omitempty"`
	RoleScopeTagIds                             *[]string                                           `json:"roleScopeTagIds,omitempty"`
	Servers                                     *[]VpnServer                                        `json:"servers,omitempty"`
	SupportsScopeTags                           *bool                                               `json:"supportsScopeTags,omitempty"`
	UserStatusOverview                          *DeviceConfigurationUserOverview                    `json:"userStatusOverview,omitempty"`
	UserStatuses                                *[]DeviceConfigurationUserStatus                    `json:"userStatuses,omitempty"`
	Version                                     *int64                                              `json:"version,omitempty"`
}
