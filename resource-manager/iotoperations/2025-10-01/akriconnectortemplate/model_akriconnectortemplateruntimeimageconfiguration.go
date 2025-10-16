package akriconnectortemplate

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AkriConnectorTemplateManagedConfigurationSettings = AkriConnectorTemplateRuntimeImageConfiguration{}

type AkriConnectorTemplateRuntimeImageConfiguration struct {
	ImageConfigurationSettings AkriConnectorTemplateRuntimeImageConfigurationSettings `json:"imageConfigurationSettings"`

	// Fields inherited from AkriConnectorTemplateManagedConfigurationSettings

	AdditionalConfiguration        *map[string]string                            `json:"additionalConfiguration,omitempty"`
	Allocation                     AkriConnectorTemplateAllocation               `json:"allocation"`
	ManagedConfigurationType       AkriConnectorTemplateManagedConfigurationType `json:"managedConfigurationType"`
	PersistentVolumeClaimTemplates *[]map[string]interface{}                     `json:"persistentVolumeClaimTemplates,omitempty"`
	PersistentVolumeClaims         *[]AkriConnectorTemplatePersistentVolumeClaim `json:"persistentVolumeClaims,omitempty"`
	Secrets                        *[]AkriConnectorsSecret                       `json:"secrets,omitempty"`
	TrustSettings                  *AkriConnectorTemplateTrustList               `json:"trustSettings,omitempty"`
}

func (s AkriConnectorTemplateRuntimeImageConfiguration) AkriConnectorTemplateManagedConfigurationSettings() BaseAkriConnectorTemplateManagedConfigurationSettingsImpl {
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

var _ json.Marshaler = AkriConnectorTemplateRuntimeImageConfiguration{}

func (s AkriConnectorTemplateRuntimeImageConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper AkriConnectorTemplateRuntimeImageConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AkriConnectorTemplateRuntimeImageConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AkriConnectorTemplateRuntimeImageConfiguration: %+v", err)
	}

	decoded["managedConfigurationType"] = "ImageConfiguration"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AkriConnectorTemplateRuntimeImageConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AkriConnectorTemplateRuntimeImageConfiguration{}

func (s *AkriConnectorTemplateRuntimeImageConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ImageConfigurationSettings     AkriConnectorTemplateRuntimeImageConfigurationSettings `json:"imageConfigurationSettings"`
		AdditionalConfiguration        *map[string]string                                     `json:"additionalConfiguration,omitempty"`
		ManagedConfigurationType       AkriConnectorTemplateManagedConfigurationType          `json:"managedConfigurationType"`
		PersistentVolumeClaimTemplates *[]map[string]interface{}                              `json:"persistentVolumeClaimTemplates,omitempty"`
		PersistentVolumeClaims         *[]AkriConnectorTemplatePersistentVolumeClaim          `json:"persistentVolumeClaims,omitempty"`
		Secrets                        *[]AkriConnectorsSecret                                `json:"secrets,omitempty"`
		TrustSettings                  *AkriConnectorTemplateTrustList                        `json:"trustSettings,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ImageConfigurationSettings = decoded.ImageConfigurationSettings
	s.AdditionalConfiguration = decoded.AdditionalConfiguration
	s.ManagedConfigurationType = decoded.ManagedConfigurationType
	s.PersistentVolumeClaimTemplates = decoded.PersistentVolumeClaimTemplates
	s.PersistentVolumeClaims = decoded.PersistentVolumeClaims
	s.Secrets = decoded.Secrets
	s.TrustSettings = decoded.TrustSettings

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AkriConnectorTemplateRuntimeImageConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["allocation"]; ok {
		impl, err := UnmarshalAkriConnectorTemplateAllocationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Allocation' for 'AkriConnectorTemplateRuntimeImageConfiguration': %+v", err)
		}
		s.Allocation = impl
	}

	return nil
}
