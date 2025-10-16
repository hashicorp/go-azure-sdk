package akriconnectortemplate

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AkriConnectorTemplateProperties struct {
	AioMetadata                 *AkriConnectorTemplateAioMetadata                `json:"aioMetadata,omitempty"`
	ConnectorMetadataRef        *string                                          `json:"connectorMetadataRef,omitempty"`
	DeviceInboundEndpointTypes  []AkriConnectorTemplateDeviceInboundEndpointType `json:"deviceInboundEndpointTypes"`
	Diagnostics                 *AkriConnectorTemplateDiagnostics                `json:"diagnostics,omitempty"`
	HealthState                 *ResourceHealthState                             `json:"healthState,omitempty"`
	MqttConnectionConfiguration *AkriConnectorsMqttConnectionConfiguration       `json:"mqttConnectionConfiguration,omitempty"`
	ProvisioningState           *ProvisioningState                               `json:"provisioningState,omitempty"`
	RuntimeConfiguration        AkriConnectorTemplateRuntimeConfiguration        `json:"runtimeConfiguration"`
}

var _ json.Unmarshaler = &AkriConnectorTemplateProperties{}

func (s *AkriConnectorTemplateProperties) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AioMetadata                 *AkriConnectorTemplateAioMetadata                `json:"aioMetadata,omitempty"`
		ConnectorMetadataRef        *string                                          `json:"connectorMetadataRef,omitempty"`
		DeviceInboundEndpointTypes  []AkriConnectorTemplateDeviceInboundEndpointType `json:"deviceInboundEndpointTypes"`
		Diagnostics                 *AkriConnectorTemplateDiagnostics                `json:"diagnostics,omitempty"`
		HealthState                 *ResourceHealthState                             `json:"healthState,omitempty"`
		MqttConnectionConfiguration *AkriConnectorsMqttConnectionConfiguration       `json:"mqttConnectionConfiguration,omitempty"`
		ProvisioningState           *ProvisioningState                               `json:"provisioningState,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AioMetadata = decoded.AioMetadata
	s.ConnectorMetadataRef = decoded.ConnectorMetadataRef
	s.DeviceInboundEndpointTypes = decoded.DeviceInboundEndpointTypes
	s.Diagnostics = decoded.Diagnostics
	s.HealthState = decoded.HealthState
	s.MqttConnectionConfiguration = decoded.MqttConnectionConfiguration
	s.ProvisioningState = decoded.ProvisioningState

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AkriConnectorTemplateProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["runtimeConfiguration"]; ok {
		impl, err := UnmarshalAkriConnectorTemplateRuntimeConfigurationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'RuntimeConfiguration' for 'AkriConnectorTemplateProperties': %+v", err)
		}
		s.RuntimeConfiguration = impl
	}

	return nil
}
