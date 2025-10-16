package akriconnectortemplate

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AkriConnectorTemplateManagedConfigurationSettings interface {
	AkriConnectorTemplateManagedConfigurationSettings() BaseAkriConnectorTemplateManagedConfigurationSettingsImpl
}

var _ AkriConnectorTemplateManagedConfigurationSettings = BaseAkriConnectorTemplateManagedConfigurationSettingsImpl{}

type BaseAkriConnectorTemplateManagedConfigurationSettingsImpl struct {
	AdditionalConfiguration        *map[string]string                            `json:"additionalConfiguration,omitempty"`
	Allocation                     AkriConnectorTemplateAllocation               `json:"allocation"`
	ManagedConfigurationType       AkriConnectorTemplateManagedConfigurationType `json:"managedConfigurationType"`
	PersistentVolumeClaimTemplates *[]map[string]interface{}                     `json:"persistentVolumeClaimTemplates,omitempty"`
	PersistentVolumeClaims         *[]AkriConnectorTemplatePersistentVolumeClaim `json:"persistentVolumeClaims,omitempty"`
	Secrets                        *[]AkriConnectorsSecret                       `json:"secrets,omitempty"`
	TrustSettings                  *AkriConnectorTemplateTrustList               `json:"trustSettings,omitempty"`
}

func (s BaseAkriConnectorTemplateManagedConfigurationSettingsImpl) AkriConnectorTemplateManagedConfigurationSettings() BaseAkriConnectorTemplateManagedConfigurationSettingsImpl {
	return s
}

var _ AkriConnectorTemplateManagedConfigurationSettings = RawAkriConnectorTemplateManagedConfigurationSettingsImpl{}

// RawAkriConnectorTemplateManagedConfigurationSettingsImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAkriConnectorTemplateManagedConfigurationSettingsImpl struct {
	akriConnectorTemplateManagedConfigurationSettings BaseAkriConnectorTemplateManagedConfigurationSettingsImpl
	Type                                              string
	Values                                            map[string]interface{}
}

func (s RawAkriConnectorTemplateManagedConfigurationSettingsImpl) AkriConnectorTemplateManagedConfigurationSettings() BaseAkriConnectorTemplateManagedConfigurationSettingsImpl {
	return s.akriConnectorTemplateManagedConfigurationSettings
}

var _ json.Unmarshaler = &BaseAkriConnectorTemplateManagedConfigurationSettingsImpl{}

func (s *BaseAkriConnectorTemplateManagedConfigurationSettingsImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AdditionalConfiguration        *map[string]string                            `json:"additionalConfiguration,omitempty"`
		ManagedConfigurationType       AkriConnectorTemplateManagedConfigurationType `json:"managedConfigurationType"`
		PersistentVolumeClaimTemplates *[]map[string]interface{}                     `json:"persistentVolumeClaimTemplates,omitempty"`
		PersistentVolumeClaims         *[]AkriConnectorTemplatePersistentVolumeClaim `json:"persistentVolumeClaims,omitempty"`
		Secrets                        *[]AkriConnectorsSecret                       `json:"secrets,omitempty"`
		TrustSettings                  *AkriConnectorTemplateTrustList               `json:"trustSettings,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AdditionalConfiguration = decoded.AdditionalConfiguration
	s.ManagedConfigurationType = decoded.ManagedConfigurationType
	s.PersistentVolumeClaimTemplates = decoded.PersistentVolumeClaimTemplates
	s.PersistentVolumeClaims = decoded.PersistentVolumeClaims
	s.Secrets = decoded.Secrets
	s.TrustSettings = decoded.TrustSettings

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseAkriConnectorTemplateManagedConfigurationSettingsImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["allocation"]; ok {
		impl, err := UnmarshalAkriConnectorTemplateAllocationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Allocation' for 'BaseAkriConnectorTemplateManagedConfigurationSettingsImpl': %+v", err)
		}
		s.Allocation = impl
	}

	return nil
}

func UnmarshalAkriConnectorTemplateManagedConfigurationSettingsImplementation(input []byte) (AkriConnectorTemplateManagedConfigurationSettings, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AkriConnectorTemplateManagedConfigurationSettings into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["managedConfigurationType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "ImageConfiguration") {
		var out AkriConnectorTemplateRuntimeImageConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AkriConnectorTemplateRuntimeImageConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "StatefulSetConfiguration") {
		var out AkriConnectorTemplateRuntimeStatefulSetConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AkriConnectorTemplateRuntimeStatefulSetConfiguration: %+v", err)
		}
		return out, nil
	}

	var parent BaseAkriConnectorTemplateManagedConfigurationSettingsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAkriConnectorTemplateManagedConfigurationSettingsImpl: %+v", err)
	}

	return RawAkriConnectorTemplateManagedConfigurationSettingsImpl{
		akriConnectorTemplateManagedConfigurationSettings: parent,
		Type:   value,
		Values: temp,
	}, nil

}
