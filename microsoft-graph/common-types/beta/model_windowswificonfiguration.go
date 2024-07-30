package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsWifiConfiguration struct {
	Assignments                                 *[]DeviceConfigurationAssignment                `json:"assignments,omitempty"`
	ConnectAutomatically                        *bool                                           `json:"connectAutomatically,omitempty"`
	ConnectToPreferredNetwork                   *bool                                           `json:"connectToPreferredNetwork,omitempty"`
	ConnectWhenNetworkNameIsHidden              *bool                                           `json:"connectWhenNetworkNameIsHidden,omitempty"`
	CreatedDateTime                             *string                                         `json:"createdDateTime,omitempty"`
	Description                                 *string                                         `json:"description,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode    `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	DeviceManagementApplicabilityRuleOsEdition  *DeviceManagementApplicabilityRuleOsEdition     `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOsVersion  *DeviceManagementApplicabilityRuleOsVersion     `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceSettingStateSummaries                 *[]SettingStateDeviceSummary                    `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                        *DeviceConfigurationDeviceOverview              `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                              *[]DeviceConfigurationDeviceStatus              `json:"deviceStatuses,omitempty"`
	DisplayName                                 *string                                         `json:"displayName,omitempty"`
	ForceFIPSCompliance                         *bool                                           `json:"forceFIPSCompliance,omitempty"`
	GroupAssignments                            *[]DeviceConfigurationGroupAssignment           `json:"groupAssignments,omitempty"`
	Id                                          *string                                         `json:"id,omitempty"`
	LastModifiedDateTime                        *string                                         `json:"lastModifiedDateTime,omitempty"`
	MeteredConnectionLimit                      *WindowsWifiConfigurationMeteredConnectionLimit `json:"meteredConnectionLimit,omitempty"`
	NetworkName                                 *string                                         `json:"networkName,omitempty"`
	ODataType                                   *string                                         `json:"@odata.type,omitempty"`
	PreSharedKey                                *string                                         `json:"preSharedKey,omitempty"`
	ProxyAutomaticConfigurationUrl              *string                                         `json:"proxyAutomaticConfigurationUrl,omitempty"`
	ProxyManualAddress                          *string                                         `json:"proxyManualAddress,omitempty"`
	ProxyManualPort                             *int64                                          `json:"proxyManualPort,omitempty"`
	ProxySetting                                *WindowsWifiConfigurationProxySetting           `json:"proxySetting,omitempty"`
	RoleScopeTagIds                             *[]string                                       `json:"roleScopeTagIds,omitempty"`
	Ssid                                        *string                                         `json:"ssid,omitempty"`
	SupportsScopeTags                           *bool                                           `json:"supportsScopeTags,omitempty"`
	UserStatusOverview                          *DeviceConfigurationUserOverview                `json:"userStatusOverview,omitempty"`
	UserStatuses                                *[]DeviceConfigurationUserStatus                `json:"userStatuses,omitempty"`
	Version                                     *int64                                          `json:"version,omitempty"`
	WifiSecurityType                            *WindowsWifiConfigurationWifiSecurityType       `json:"wifiSecurityType,omitempty"`
}
