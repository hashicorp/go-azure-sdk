package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileVpnConfiguration struct {
	AlwaysOn                                    *bool                                                   `json:"alwaysOn,omitempty"`
	AlwaysOnLockdown                            *bool                                                   `json:"alwaysOnLockdown,omitempty"`
	Assignments                                 *[]DeviceConfigurationAssignment                        `json:"assignments,omitempty"`
	AuthenticationMethod                        *AndroidWorkProfileVpnConfigurationAuthenticationMethod `json:"authenticationMethod,omitempty"`
	ConnectionName                              *string                                                 `json:"connectionName,omitempty"`
	ConnectionType                              *AndroidWorkProfileVpnConfigurationConnectionType       `json:"connectionType,omitempty"`
	CreatedDateTime                             *string                                                 `json:"createdDateTime,omitempty"`
	CustomData                                  *[]KeyValue                                             `json:"customData,omitempty"`
	CustomKeyValueData                          *[]KeyValuePair                                         `json:"customKeyValueData,omitempty"`
	Description                                 *string                                                 `json:"description,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode            `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	DeviceManagementApplicabilityRuleOsEdition  *DeviceManagementApplicabilityRuleOsEdition             `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOsVersion  *DeviceManagementApplicabilityRuleOsVersion             `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceSettingStateSummaries                 *[]SettingStateDeviceSummary                            `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                        *DeviceConfigurationDeviceOverview                      `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                              *[]DeviceConfigurationDeviceStatus                      `json:"deviceStatuses,omitempty"`
	DisplayName                                 *string                                                 `json:"displayName,omitempty"`
	Fingerprint                                 *string                                                 `json:"fingerprint,omitempty"`
	GroupAssignments                            *[]DeviceConfigurationGroupAssignment                   `json:"groupAssignments,omitempty"`
	Id                                          *string                                                 `json:"id,omitempty"`
	IdentityCertificate                         *AndroidWorkProfileCertificateProfileBase               `json:"identityCertificate,omitempty"`
	LastModifiedDateTime                        *string                                                 `json:"lastModifiedDateTime,omitempty"`
	MicrosoftTunnelSiteId                       *string                                                 `json:"microsoftTunnelSiteId,omitempty"`
	ODataType                                   *string                                                 `json:"@odata.type,omitempty"`
	ProxyExclusionList                          *[]string                                               `json:"proxyExclusionList,omitempty"`
	ProxyServer                                 *VpnProxyServer                                         `json:"proxyServer,omitempty"`
	Realm                                       *string                                                 `json:"realm,omitempty"`
	Role                                        *string                                                 `json:"role,omitempty"`
	RoleScopeTagIds                             *[]string                                               `json:"roleScopeTagIds,omitempty"`
	Servers                                     *[]VpnServer                                            `json:"servers,omitempty"`
	SupportsScopeTags                           *bool                                                   `json:"supportsScopeTags,omitempty"`
	TargetedMobileApps                          *[]AppListItem                                          `json:"targetedMobileApps,omitempty"`
	TargetedPackageIds                          *[]string                                               `json:"targetedPackageIds,omitempty"`
	UserStatusOverview                          *DeviceConfigurationUserOverview                        `json:"userStatusOverview,omitempty"`
	UserStatuses                                *[]DeviceConfigurationUserStatus                        `json:"userStatuses,omitempty"`
	Version                                     *int64                                                  `json:"version,omitempty"`
}
