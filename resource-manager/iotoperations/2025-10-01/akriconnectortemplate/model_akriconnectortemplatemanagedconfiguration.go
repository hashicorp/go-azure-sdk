package akriconnectortemplate

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AkriConnectorTemplateRuntimeConfiguration = AkriConnectorTemplateManagedConfiguration{}

type AkriConnectorTemplateManagedConfiguration struct {
	ManagedConfigurationSettings AkriConnectorTemplateManagedConfigurationSettings `json:"managedConfigurationSettings"`

	// Fields inherited from AkriConnectorTemplateRuntimeConfiguration

	RuntimeConfigurationType AkriConnectorTemplateRuntimeConfigurationType `json:"runtimeConfigurationType"`
}

func (s AkriConnectorTemplateManagedConfiguration) AkriConnectorTemplateRuntimeConfiguration() BaseAkriConnectorTemplateRuntimeConfigurationImpl {
	return BaseAkriConnectorTemplateRuntimeConfigurationImpl{
		RuntimeConfigurationType: s.RuntimeConfigurationType,
	}
}

var _ json.Marshaler = AkriConnectorTemplateManagedConfiguration{}

func (s AkriConnectorTemplateManagedConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper AkriConnectorTemplateManagedConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AkriConnectorTemplateManagedConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AkriConnectorTemplateManagedConfiguration: %+v", err)
	}

	decoded["runtimeConfigurationType"] = "ManagedConfiguration"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AkriConnectorTemplateManagedConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AkriConnectorTemplateManagedConfiguration{}

func (s *AkriConnectorTemplateManagedConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		RuntimeConfigurationType AkriConnectorTemplateRuntimeConfigurationType `json:"runtimeConfigurationType"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.RuntimeConfigurationType = decoded.RuntimeConfigurationType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AkriConnectorTemplateManagedConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["managedConfigurationSettings"]; ok {
		impl, err := UnmarshalAkriConnectorTemplateManagedConfigurationSettingsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ManagedConfigurationSettings' for 'AkriConnectorTemplateManagedConfiguration': %+v", err)
		}
		s.ManagedConfigurationSettings = impl
	}

	return nil
}
