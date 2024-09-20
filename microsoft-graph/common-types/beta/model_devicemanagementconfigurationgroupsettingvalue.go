package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementConfigurationSettingValue = DeviceManagementConfigurationGroupSettingValue{}

type DeviceManagementConfigurationGroupSettingValue struct {
	// Collection of child setting instances contained within this GroupSetting
	Children *[]DeviceManagementConfigurationSettingInstance `json:"children,omitempty"`

	// Fields inherited from DeviceManagementConfigurationSettingValue

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Setting value template reference
	SettingValueTemplateReference *DeviceManagementConfigurationSettingValueTemplateReference `json:"settingValueTemplateReference,omitempty"`
}

func (s DeviceManagementConfigurationGroupSettingValue) DeviceManagementConfigurationSettingValue() BaseDeviceManagementConfigurationSettingValueImpl {
	return BaseDeviceManagementConfigurationSettingValueImpl{
		ODataId:                       s.ODataId,
		ODataType:                     s.ODataType,
		SettingValueTemplateReference: s.SettingValueTemplateReference,
	}
}

var _ json.Marshaler = DeviceManagementConfigurationGroupSettingValue{}

func (s DeviceManagementConfigurationGroupSettingValue) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementConfigurationGroupSettingValue
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementConfigurationGroupSettingValue: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationGroupSettingValue: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.deviceManagementConfigurationGroupSettingValue"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementConfigurationGroupSettingValue: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceManagementConfigurationGroupSettingValue{}

func (s *DeviceManagementConfigurationGroupSettingValue) UnmarshalJSON(bytes []byte) error {

	var decoded struct {
		Children                      *[]DeviceManagementConfigurationSettingInstance             `json:"children,omitempty"`
		ODataId                       *string                                                     `json:"@odata.id,omitempty"`
		ODataType                     *string                                                     `json:"@odata.type,omitempty"`
		SettingValueTemplateReference *DeviceManagementConfigurationSettingValueTemplateReference `json:"settingValueTemplateReference,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.SettingValueTemplateReference = decoded.SettingValueTemplateReference

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceManagementConfigurationGroupSettingValue into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["children"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Children into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceManagementConfigurationSettingInstance, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceManagementConfigurationSettingInstanceImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Children' for 'DeviceManagementConfigurationGroupSettingValue': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Children = &output
	}

	return nil
}
