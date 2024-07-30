package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVpnConfiguration struct {
	Assignments                                 *[]DeviceConfigurationAssignment             `json:"assignments,omitempty"`
	AssociatedDomains                           *[]string                                    `json:"associatedDomains,omitempty"`
	AuthenticationMethod                        *IosVpnConfigurationAuthenticationMethod     `json:"authenticationMethod,omitempty"`
	CloudName                                   *string                                      `json:"cloudName,omitempty"`
	ConnectionName                              *string                                      `json:"connectionName,omitempty"`
	ConnectionType                              *IosVpnConfigurationConnectionType           `json:"connectionType,omitempty"`
	CreatedDateTime                             *string                                      `json:"createdDateTime,omitempty"`
	CustomData                                  *[]KeyValue                                  `json:"customData,omitempty"`
	CustomKeyValueData                          *[]KeyValuePair                              `json:"customKeyValueData,omitempty"`
	DerivedCredentialSettings                   *DeviceManagementDerivedCredentialSettings   `json:"derivedCredentialSettings,omitempty"`
	Description                                 *string                                      `json:"description,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	DeviceManagementApplicabilityRuleOsEdition  *DeviceManagementApplicabilityRuleOsEdition  `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOsVersion  *DeviceManagementApplicabilityRuleOsVersion  `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceSettingStateSummaries                 *[]SettingStateDeviceSummary                 `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                        *DeviceConfigurationDeviceOverview           `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                              *[]DeviceConfigurationDeviceStatus           `json:"deviceStatuses,omitempty"`
	DisableOnDemandUserOverride                 *bool                                        `json:"disableOnDemandUserOverride,omitempty"`
	DisconnectOnIdle                            *bool                                        `json:"disconnectOnIdle,omitempty"`
	DisconnectOnIdleTimerInSeconds              *int64                                       `json:"disconnectOnIdleTimerInSeconds,omitempty"`
	DisplayName                                 *string                                      `json:"displayName,omitempty"`
	EnablePerApp                                *bool                                        `json:"enablePerApp,omitempty"`
	EnableSplitTunneling                        *bool                                        `json:"enableSplitTunneling,omitempty"`
	ExcludeList                                 *[]string                                    `json:"excludeList,omitempty"`
	ExcludedDomains                             *[]string                                    `json:"excludedDomains,omitempty"`
	GroupAssignments                            *[]DeviceConfigurationGroupAssignment        `json:"groupAssignments,omitempty"`
	Id                                          *string                                      `json:"id,omitempty"`
	Identifier                                  *string                                      `json:"identifier,omitempty"`
	IdentityCertificate                         *IosCertificateProfileBase                   `json:"identityCertificate,omitempty"`
	LastModifiedDateTime                        *string                                      `json:"lastModifiedDateTime,omitempty"`
	LoginGroupOrDomain                          *string                                      `json:"loginGroupOrDomain,omitempty"`
	MicrosoftTunnelSiteId                       *string                                      `json:"microsoftTunnelSiteId,omitempty"`
	ODataType                                   *string                                      `json:"@odata.type,omitempty"`
	OnDemandRules                               *[]VpnOnDemandRule                           `json:"onDemandRules,omitempty"`
	OptInToDeviceIdSharing                      *bool                                        `json:"optInToDeviceIdSharing,omitempty"`
	ProviderType                                *IosVpnConfigurationProviderType             `json:"providerType,omitempty"`
	ProxyServer                                 *VpnProxyServer                              `json:"proxyServer,omitempty"`
	Realm                                       *string                                      `json:"realm,omitempty"`
	Role                                        *string                                      `json:"role,omitempty"`
	RoleScopeTagIds                             *[]string                                    `json:"roleScopeTagIds,omitempty"`
	SafariDomains                               *[]string                                    `json:"safariDomains,omitempty"`
	Server                                      *VpnServer                                   `json:"server,omitempty"`
	StrictEnforcement                           *bool                                        `json:"strictEnforcement,omitempty"`
	SupportsScopeTags                           *bool                                        `json:"supportsScopeTags,omitempty"`
	TargetedMobileApps                          *[]AppListItem                               `json:"targetedMobileApps,omitempty"`
	UserDomain                                  *string                                      `json:"userDomain,omitempty"`
	UserStatusOverview                          *DeviceConfigurationUserOverview             `json:"userStatusOverview,omitempty"`
	UserStatuses                                *[]DeviceConfigurationUserStatus             `json:"userStatuses,omitempty"`
	Version                                     *int64                                       `json:"version,omitempty"`
}
