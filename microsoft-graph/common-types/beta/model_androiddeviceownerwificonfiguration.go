package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerWiFiConfiguration interface {
	Entity
	DeviceConfiguration
	AndroidDeviceOwnerWiFiConfiguration() BaseAndroidDeviceOwnerWiFiConfigurationImpl
}

var _ AndroidDeviceOwnerWiFiConfiguration = BaseAndroidDeviceOwnerWiFiConfigurationImpl{}

type BaseAndroidDeviceOwnerWiFiConfigurationImpl struct {
	// Connect automatically when this network is in range. Setting this to true will skip the user prompt and automatically
	// connect the device to Wi-Fi network.
	ConnectAutomatically nullable.Type[bool] `json:"connectAutomatically,omitempty"`

	// When set to true, this profile forces the device to connect to a network that doesn't broadcast its SSID to all
	// devices.
	ConnectWhenNetworkNameIsHidden nullable.Type[bool] `json:"connectWhenNetworkNameIsHidden,omitempty"`

	// The MAC address randomization mode for Android device Wi-Fi configuration. Possible values include automatic and
	// hardware. Default value is automatic. Possible values are: automatic, hardware, unknownFutureValue.
	MacAddressRandomizationMode *MacAddressRandomizationMode `json:"macAddressRandomizationMode,omitempty"`

	// Network Name
	NetworkName *string `json:"networkName,omitempty"`

	// This is the pre-shared key for WPA Personal Wi-Fi network.
	PreSharedKey nullable.Type[string] `json:"preSharedKey,omitempty"`

	// This is the pre-shared key for WPA Personal Wi-Fi network.
	PreSharedKeyIsSet *bool `json:"preSharedKeyIsSet,omitempty"`

	// Specify the proxy server configuration script URL.
	ProxyAutomaticConfigurationUrl nullable.Type[string] `json:"proxyAutomaticConfigurationUrl,omitempty"`

	// List of hosts to exclude using the proxy on connections for. These hosts can use wildcards such as .example.com.
	ProxyExclusionList nullable.Type[string] `json:"proxyExclusionList,omitempty"`

	// Specify the proxy server IP address. Android documentation does not specify IPv4 or IPv6. For example: 192.168.1.1.
	ProxyManualAddress nullable.Type[string] `json:"proxyManualAddress,omitempty"`

	// Specify the proxy server port.
	ProxyManualPort nullable.Type[int64] `json:"proxyManualPort,omitempty"`

	// Wi-Fi Proxy Settings.
	ProxySettings *WiFiProxySetting `json:"proxySettings,omitempty"`

	// This is the name of the Wi-Fi network that is broadcast to all devices.
	Ssid *string `json:"ssid,omitempty"`

	// Wi-Fi Security Types for Android Device Owner.
	WiFiSecurityType *AndroidDeviceOwnerWiFiSecurityType `json:"wiFiSecurityType,omitempty"`

	// Fields inherited from DeviceConfiguration

	// The list of assignments for the device configuration profile.
	Assignments *[]DeviceConfigurationAssignment `json:"assignments,omitempty"`

	// DateTime the object was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Admin provided description of the Device Configuration.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The device mode applicability rule for this Policy.
	DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`

	// The OS edition applicability for this Policy.
	DeviceManagementApplicabilityRuleOsEdition *DeviceManagementApplicabilityRuleOsEdition `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`

	// The OS version applicability rule for this Policy.
	DeviceManagementApplicabilityRuleOsVersion *DeviceManagementApplicabilityRuleOsVersion `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`

	// Device Configuration Setting State Device Summary
	DeviceSettingStateSummaries *[]SettingStateDeviceSummary `json:"deviceSettingStateSummaries,omitempty"`

	// Device Configuration devices status overview
	DeviceStatusOverview *DeviceConfigurationDeviceOverview `json:"deviceStatusOverview,omitempty"`

	// Device configuration installation status by device.
	DeviceStatuses *[]DeviceConfigurationDeviceStatus `json:"deviceStatuses,omitempty"`

	// Admin provided name of the device configuration.
	DisplayName *string `json:"displayName,omitempty"`

	// The list of group assignments for the device configuration profile.
	GroupAssignments *[]DeviceConfigurationGroupAssignment `json:"groupAssignments,omitempty"`

	// DateTime the object was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// List of Scope Tags for this Entity instance.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// Indicates whether or not the underlying Device Configuration supports the assignment of scope tags. Assigning to the
	// ScopeTags property is not allowed when this value is false and entities will not be visible to scoped users. This
	// occurs for Legacy policies created in Silverlight and can be resolved by deleting and recreating the policy in the
	// Azure Portal. This property is read-only.
	SupportsScopeTags *bool `json:"supportsScopeTags,omitempty"`

	// Device Configuration users status overview
	UserStatusOverview *DeviceConfigurationUserOverview `json:"userStatusOverview,omitempty"`

	// Device configuration installation status by user.
	UserStatuses *[]DeviceConfigurationUserStatus `json:"userStatuses,omitempty"`

	// Version of the device configuration.
	Version *int64 `json:"version,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s BaseAndroidDeviceOwnerWiFiConfigurationImpl) AndroidDeviceOwnerWiFiConfiguration() BaseAndroidDeviceOwnerWiFiConfigurationImpl {
	return s
}

func (s BaseAndroidDeviceOwnerWiFiConfigurationImpl) DeviceConfiguration() BaseDeviceConfigurationImpl {
	return BaseDeviceConfigurationImpl{
		Assignments:     s.Assignments,
		CreatedDateTime: s.CreatedDateTime,
		Description:     s.Description,
		DeviceManagementApplicabilityRuleDeviceMode: s.DeviceManagementApplicabilityRuleDeviceMode,
		DeviceManagementApplicabilityRuleOsEdition:  s.DeviceManagementApplicabilityRuleOsEdition,
		DeviceManagementApplicabilityRuleOsVersion:  s.DeviceManagementApplicabilityRuleOsVersion,
		DeviceSettingStateSummaries:                 s.DeviceSettingStateSummaries,
		DeviceStatusOverview:                        s.DeviceStatusOverview,
		DeviceStatuses:                              s.DeviceStatuses,
		DisplayName:                                 s.DisplayName,
		GroupAssignments:                            s.GroupAssignments,
		LastModifiedDateTime:                        s.LastModifiedDateTime,
		RoleScopeTagIds:                             s.RoleScopeTagIds,
		SupportsScopeTags:                           s.SupportsScopeTags,
		UserStatusOverview:                          s.UserStatusOverview,
		UserStatuses:                                s.UserStatuses,
		Version:                                     s.Version,
		Id:                                          s.Id,
		ODataId:                                     s.ODataId,
		ODataType:                                   s.ODataType,
	}
}

func (s BaseAndroidDeviceOwnerWiFiConfigurationImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ AndroidDeviceOwnerWiFiConfiguration = RawAndroidDeviceOwnerWiFiConfigurationImpl{}

// RawAndroidDeviceOwnerWiFiConfigurationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAndroidDeviceOwnerWiFiConfigurationImpl struct {
	androidDeviceOwnerWiFiConfiguration BaseAndroidDeviceOwnerWiFiConfigurationImpl
	Type                                string
	Values                              map[string]interface{}
}

func (s RawAndroidDeviceOwnerWiFiConfigurationImpl) AndroidDeviceOwnerWiFiConfiguration() BaseAndroidDeviceOwnerWiFiConfigurationImpl {
	return s.androidDeviceOwnerWiFiConfiguration
}

func (s RawAndroidDeviceOwnerWiFiConfigurationImpl) DeviceConfiguration() BaseDeviceConfigurationImpl {
	return s.androidDeviceOwnerWiFiConfiguration.DeviceConfiguration()
}

func (s RawAndroidDeviceOwnerWiFiConfigurationImpl) Entity() BaseEntityImpl {
	return s.androidDeviceOwnerWiFiConfiguration.Entity()
}

var _ json.Marshaler = BaseAndroidDeviceOwnerWiFiConfigurationImpl{}

func (s BaseAndroidDeviceOwnerWiFiConfigurationImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAndroidDeviceOwnerWiFiConfigurationImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAndroidDeviceOwnerWiFiConfigurationImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAndroidDeviceOwnerWiFiConfigurationImpl: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.androidDeviceOwnerWiFiConfiguration"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAndroidDeviceOwnerWiFiConfigurationImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalAndroidDeviceOwnerWiFiConfigurationImplementation(input []byte) (AndroidDeviceOwnerWiFiConfiguration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidDeviceOwnerWiFiConfiguration into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.androidDeviceOwnerEnterpriseWiFiConfiguration") {
		var out AndroidDeviceOwnerEnterpriseWiFiConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidDeviceOwnerEnterpriseWiFiConfiguration: %+v", err)
		}
		return out, nil
	}

	var parent BaseAndroidDeviceOwnerWiFiConfigurationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAndroidDeviceOwnerWiFiConfigurationImpl: %+v", err)
	}

	return RawAndroidDeviceOwnerWiFiConfigurationImpl{
		androidDeviceOwnerWiFiConfiguration: parent,
		Type:                                value,
		Values:                              temp,
	}, nil

}
