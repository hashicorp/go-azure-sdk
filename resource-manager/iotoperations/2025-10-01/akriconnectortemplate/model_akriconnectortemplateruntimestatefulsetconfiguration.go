package akriconnectortemplate

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AkriConnectorTemplateManagedConfigurationSettings = AkriConnectorTemplateRuntimeStatefulSetConfiguration{}

type AkriConnectorTemplateRuntimeStatefulSetConfiguration struct {
	StatefulSetConfigurationSettings map[string]interface{} `json:"statefulSetConfigurationSettings"`

	// Fields inherited from AkriConnectorTemplateManagedConfigurationSettings

	AdditionalConfiguration        *map[string]string                            `json:"additionalConfiguration,omitempty"`
	Allocation                     AkriConnectorTemplateAllocation               `json:"allocation"`
	ManagedConfigurationType       AkriConnectorTemplateManagedConfigurationType `json:"managedConfigurationType"`
	PersistentVolumeClaimTemplates *[]map[string]interface{}                     `json:"persistentVolumeClaimTemplates,omitempty"`
	PersistentVolumeClaims         *[]AkriConnectorTemplatePersistentVolumeClaim `json:"persistentVolumeClaims,omitempty"`
	Secrets                        *[]AkriConnectorsSecret                       `json:"secrets,omitempty"`
	TrustSettings                  *AkriConnectorTemplateTrustList               `json:"trustSettings,omitempty"`
}

func (s AkriConnectorTemplateRuntimeStatefulSetConfiguration) AkriConnectorTemplateManagedConfigurationSettings() BaseAkriConnectorTemplateManagedConfigurationSettingsImpl {
	return BaseAkriConnectorTemplateManagedConfigurationSettingsImpl{
		AdditionalConfiguration:        s.AdditionalConfiguration,
		Allocation:                     s.Allocation,
		ManagedConfigurationType:       s.ManagedConfigurationType,
		PersistentVolumeClaimTemplates: s.PersistentVolumeClaimTemplates,
		PersistentVolumeClaims:         s.PersistentVolumeClaims,
		Secrets:                        s.Secrets,
		TrustSettings:                  s.TrustSettings,
	}
}

var _ json.Marshaler = AkriConnectorTemplateRuntimeStatefulSetConfiguration{}

func (s AkriConnectorTemplateRuntimeStatefulSetConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper AkriConnectorTemplateRuntimeStatefulSetConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AkriConnectorTemplateRuntimeStatefulSetConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AkriConnectorTemplateRuntimeStatefulSetConfiguration: %+v", err)
	}

	decoded["managedConfigurationType"] = "StatefulSetConfiguration"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AkriConnectorTemplateRuntimeStatefulSetConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AkriConnectorTemplateRuntimeStatefulSetConfiguration{}

func (s *AkriConnectorTemplateRuntimeStatefulSetConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		StatefulSetConfigurationSettings map[string]interface{}                        `json:"statefulSetConfigurationSettings"`
		AdditionalConfiguration          *map[string]string                            `json:"additionalConfiguration,omitempty"`
		ManagedConfigurationType         AkriConnectorTemplateManagedConfigurationType `json:"managedConfigurationType"`
		PersistentVolumeClaimTemplates   *[]map[string]interface{}                     `json:"persistentVolumeClaimTemplates,omitempty"`
		PersistentVolumeClaims           *[]AkriConnectorTemplatePersistentVolumeClaim `json:"persistentVolumeClaims,omitempty"`
		Secrets                          *[]AkriConnectorsSecret                       `json:"secrets,omitempty"`
		TrustSettings                    *AkriConnectorTemplateTrustList               `json:"trustSettings,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.StatefulSetConfigurationSettings = decoded.StatefulSetConfigurationSettings
	s.AdditionalConfiguration = decoded.AdditionalConfiguration
	s.ManagedConfigurationType = decoded.ManagedConfigurationType
	s.PersistentVolumeClaimTemplates = decoded.PersistentVolumeClaimTemplates
	s.PersistentVolumeClaims = decoded.PersistentVolumeClaims
	s.Secrets = decoded.Secrets
	s.TrustSettings = decoded.TrustSettings

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AkriConnectorTemplateRuntimeStatefulSetConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["allocation"]; ok {
		impl, err := UnmarshalAkriConnectorTemplateAllocationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Allocation' for 'AkriConnectorTemplateRuntimeStatefulSetConfiguration': %+v", err)
		}
		s.Allocation = impl
	}

	return nil
}
