package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationChoiceSettingValueTemplate struct {
	// Choice Setting Value Default Template.
	DefaultValue DeviceManagementConfigurationChoiceSettingValueDefaultTemplate `json:"defaultValue"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Recommended definition override.
	RecommendedValueDefinition *DeviceManagementConfigurationChoiceSettingValueDefinitionTemplate `json:"recommendedValueDefinition,omitempty"`

	// Required definition override.
	RequiredValueDefinition *DeviceManagementConfigurationChoiceSettingValueDefinitionTemplate `json:"requiredValueDefinition,omitempty"`

	// Setting Value Template Id
	SettingValueTemplateId *string `json:"settingValueTemplateId,omitempty"`
}

var _ json.Unmarshaler = &DeviceManagementConfigurationChoiceSettingValueTemplate{}

func (s *DeviceManagementConfigurationChoiceSettingValueTemplate) UnmarshalJSON(bytes []byte) error {
	type alias DeviceManagementConfigurationChoiceSettingValueTemplate
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into DeviceManagementConfigurationChoiceSettingValueTemplate: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.RecommendedValueDefinition = decoded.RecommendedValueDefinition
	s.RequiredValueDefinition = decoded.RequiredValueDefinition
	s.SettingValueTemplateId = decoded.SettingValueTemplateId

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceManagementConfigurationChoiceSettingValueTemplate into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["defaultValue"]; ok {
		impl, err := UnmarshalDeviceManagementConfigurationChoiceSettingValueDefaultTemplateImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'DefaultValue' for 'DeviceManagementConfigurationChoiceSettingValueTemplate': %+v", err)
		}
		s.DefaultValue = impl
	}
	return nil
}
