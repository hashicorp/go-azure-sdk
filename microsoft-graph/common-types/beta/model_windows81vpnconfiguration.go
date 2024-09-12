package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81VpnConfiguration interface {
	Entity
	DeviceConfiguration
	WindowsVpnConfiguration
	Windows81VpnConfiguration() BaseWindows81VpnConfigurationImpl
}

var _ Windows81VpnConfiguration = BaseWindows81VpnConfigurationImpl{}

type BaseWindows81VpnConfigurationImpl struct {
	// Value indicating whether this policy only applies to Windows 8.1. This property is read-only.
	ApplyOnlyToWindows81 *bool `json:"applyOnlyToWindows81,omitempty"`

	// Windows VPN connection type.
	ConnectionType *WindowsVpnConnectionType `json:"connectionType,omitempty"`

	// Enable split tunneling for the VPN.
	EnableSplitTunneling *bool `json:"enableSplitTunneling,omitempty"`

	// Login group or domain when connection type is set to Dell SonicWALL Mobile Connection.
	LoginGroupOrDomain nullable.Type[string] `json:"loginGroupOrDomain,omitempty"`

	// Proxy Server.
	ProxyServer *Windows81VpnProxyServer `json:"proxyServer,omitempty"`

	// Fields inherited from WindowsVpnConfiguration

	// Connection name displayed to the user.
	ConnectionName *string `json:"connectionName,omitempty"`

	// Custom XML commands that configures the VPN connection. (UTF8 encoded byte array)
	CustomXml nullable.Type[string] `json:"customXml,omitempty"`

	// List of VPN Servers on the network. Make sure end users can access these network locations. This collection can
	// contain a maximum of 500 elements.
	Servers *[]VpnServer `json:"servers,omitempty"`

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

func (s BaseWindows81VpnConfigurationImpl) Windows81VpnConfiguration() BaseWindows81VpnConfigurationImpl {
	return s
}

func (s BaseWindows81VpnConfigurationImpl) WindowsVpnConfiguration() BaseWindowsVpnConfigurationImpl {
	return BaseWindowsVpnConfigurationImpl{
		ConnectionName:  s.ConnectionName,
		CustomXml:       s.CustomXml,
		Servers:         s.Servers,
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

func (s BaseWindows81VpnConfigurationImpl) DeviceConfiguration() BaseDeviceConfigurationImpl {
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

func (s BaseWindows81VpnConfigurationImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ Windows81VpnConfiguration = RawWindows81VpnConfigurationImpl{}

// RawWindows81VpnConfigurationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWindows81VpnConfigurationImpl struct {
	windows81VpnConfiguration BaseWindows81VpnConfigurationImpl
	Type                      string
	Values                    map[string]interface{}
}

func (s RawWindows81VpnConfigurationImpl) Windows81VpnConfiguration() BaseWindows81VpnConfigurationImpl {
	return s.windows81VpnConfiguration
}

func (s RawWindows81VpnConfigurationImpl) WindowsVpnConfiguration() BaseWindowsVpnConfigurationImpl {
	return s.windows81VpnConfiguration.WindowsVpnConfiguration()
}

func (s RawWindows81VpnConfigurationImpl) DeviceConfiguration() BaseDeviceConfigurationImpl {
	return s.windows81VpnConfiguration.DeviceConfiguration()
}

func (s RawWindows81VpnConfigurationImpl) Entity() BaseEntityImpl {
	return s.windows81VpnConfiguration.Entity()
}

var _ json.Marshaler = BaseWindows81VpnConfigurationImpl{}

func (s BaseWindows81VpnConfigurationImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseWindows81VpnConfigurationImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseWindows81VpnConfigurationImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseWindows81VpnConfigurationImpl: %+v", err)
	}

	delete(decoded, "applyOnlyToWindows81")
	decoded["@odata.type"] = "#microsoft.graph.windows81VpnConfiguration"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseWindows81VpnConfigurationImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalWindows81VpnConfigurationImplementation(input []byte) (Windows81VpnConfiguration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Windows81VpnConfiguration into map[string]interface: %+v", err)
	}

	value, ok := temp["@odata.type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsPhone81VpnConfiguration") {
		var out WindowsPhone81VpnConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsPhone81VpnConfiguration: %+v", err)
		}
		return out, nil
	}

	var parent BaseWindows81VpnConfigurationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWindows81VpnConfigurationImpl: %+v", err)
	}

	return RawWindows81VpnConfigurationImpl{
		windows81VpnConfiguration: parent,
		Type:                      value,
		Values:                    temp,
	}, nil

}
