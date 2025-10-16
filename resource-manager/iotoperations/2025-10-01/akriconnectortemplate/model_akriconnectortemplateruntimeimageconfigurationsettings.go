package akriconnectortemplate

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AkriConnectorTemplateRuntimeImageConfigurationSettings struct {
	ImageName         string                          `json:"imageName"`
	ImagePullPolicy   *AkriConnectorsImagePullPolicy  `json:"imagePullPolicy,omitempty"`
	RegistrySettings  AkriConnectorsRegistrySettings  `json:"registrySettings"`
	Replicas          *int64                          `json:"replicas,omitempty"`
	TagDigestSettings AkriConnectorsTagDigestSettings `json:"tagDigestSettings"`
}

var _ json.Unmarshaler = &AkriConnectorTemplateRuntimeImageConfigurationSettings{}

func (s *AkriConnectorTemplateRuntimeImageConfigurationSettings) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ImageName       string                         `json:"imageName"`
		ImagePullPolicy *AkriConnectorsImagePullPolicy `json:"imagePullPolicy,omitempty"`
		Replicas        *int64                         `json:"replicas,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ImageName = decoded.ImageName
	s.ImagePullPolicy = decoded.ImagePullPolicy
	s.Replicas = decoded.Replicas

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AkriConnectorTemplateRuntimeImageConfigurationSettings into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["registrySettings"]; ok {
		impl, err := UnmarshalAkriConnectorsRegistrySettingsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'RegistrySettings' for 'AkriConnectorTemplateRuntimeImageConfigurationSettings': %+v", err)
		}
		s.RegistrySettings = impl
	}

	if v, ok := temp["tagDigestSettings"]; ok {
		impl, err := UnmarshalAkriConnectorsTagDigestSettingsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'TagDigestSettings' for 'AkriConnectorTemplateRuntimeImageConfigurationSettings': %+v", err)
		}
		s.TagDigestSettings = impl
	}

	return nil
}
