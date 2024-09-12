package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementSettingCategory = DeviceManagementIntentSettingCategory{}

type DeviceManagementIntentSettingCategory struct {
	// The settings this category contains
	Settings *[]DeviceManagementSettingInstance `json:"settings,omitempty"`

	// Fields inherited from DeviceManagementSettingCategory

	// The category name
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The category contains top level required setting
	HasRequiredSetting nullable.Type[bool] `json:"hasRequiredSetting,omitempty"`

	// The setting definitions this category contains
	SettingDefinitions *[]DeviceManagementSettingDefinition `json:"settingDefinitions,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s DeviceManagementIntentSettingCategory) DeviceManagementSettingCategory() BaseDeviceManagementSettingCategoryImpl {
	return BaseDeviceManagementSettingCategoryImpl{
		DisplayName:        s.DisplayName,
		HasRequiredSetting: s.HasRequiredSetting,
		SettingDefinitions: s.SettingDefinitions,
		Id:                 s.Id,
		ODataId:            s.ODataId,
		ODataType:          s.ODataType,
	}
}

func (s DeviceManagementIntentSettingCategory) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementIntentSettingCategory{}

func (s DeviceManagementIntentSettingCategory) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementIntentSettingCategory
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementIntentSettingCategory: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementIntentSettingCategory: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.deviceManagementIntentSettingCategory"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementIntentSettingCategory: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceManagementIntentSettingCategory{}

func (s *DeviceManagementIntentSettingCategory) UnmarshalJSON(bytes []byte) error {
	type alias DeviceManagementIntentSettingCategory
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into DeviceManagementIntentSettingCategory: %+v", err)
	}

	s.DisplayName = decoded.DisplayName
	s.HasRequiredSetting = decoded.HasRequiredSetting
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceManagementIntentSettingCategory into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["settingDefinitions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling SettingDefinitions into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceManagementSettingDefinition, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceManagementSettingDefinitionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'SettingDefinitions' for 'DeviceManagementIntentSettingCategory': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.SettingDefinitions = &output
	}

	if v, ok := temp["settings"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Settings into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceManagementSettingInstance, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceManagementSettingInstanceImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Settings' for 'DeviceManagementIntentSettingCategory': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Settings = &output
	}
	return nil
}
