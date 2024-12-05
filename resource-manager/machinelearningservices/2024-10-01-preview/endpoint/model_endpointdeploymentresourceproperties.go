package endpoint

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndpointDeploymentResourceProperties interface {
	EndpointDeploymentResourceProperties() BaseEndpointDeploymentResourcePropertiesImpl
}

var _ EndpointDeploymentResourceProperties = BaseEndpointDeploymentResourcePropertiesImpl{}

type BaseEndpointDeploymentResourcePropertiesImpl struct {
	FailureReason     *string                           `json:"failureReason,omitempty"`
	ProvisioningState *DefaultResourceProvisioningState `json:"provisioningState,omitempty"`
	Type              string                            `json:"type"`
}

func (s BaseEndpointDeploymentResourcePropertiesImpl) EndpointDeploymentResourceProperties() BaseEndpointDeploymentResourcePropertiesImpl {
	return s
}

var _ EndpointDeploymentResourceProperties = RawEndpointDeploymentResourcePropertiesImpl{}

// RawEndpointDeploymentResourcePropertiesImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawEndpointDeploymentResourcePropertiesImpl struct {
	endpointDeploymentResourceProperties BaseEndpointDeploymentResourcePropertiesImpl
	Type                                 string
	Values                               map[string]interface{}
}

func (s RawEndpointDeploymentResourcePropertiesImpl) EndpointDeploymentResourceProperties() BaseEndpointDeploymentResourcePropertiesImpl {
	return s.endpointDeploymentResourceProperties
}

func UnmarshalEndpointDeploymentResourcePropertiesImplementation(input []byte) (EndpointDeploymentResourceProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EndpointDeploymentResourceProperties into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Azure.ContentSafety") {
		var out ContentSafetyEndpointDeploymentResourceProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ContentSafetyEndpointDeploymentResourceProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "managedOnlineEndpoint") {
		var out ManagedOnlineEndpointDeploymentResourceProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedOnlineEndpointDeploymentResourceProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Azure.OpenAI") {
		var out OpenAIEndpointDeploymentResourceProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OpenAIEndpointDeploymentResourceProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Azure.Speech") {
		var out SpeechEndpointDeploymentResourceProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SpeechEndpointDeploymentResourceProperties: %+v", err)
		}
		return out, nil
	}

	var parent BaseEndpointDeploymentResourcePropertiesImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEndpointDeploymentResourcePropertiesImpl: %+v", err)
	}

	return RawEndpointDeploymentResourcePropertiesImpl{
		endpointDeploymentResourceProperties: parent,
		Type:                                 value,
		Values:                               temp,
	}, nil

}
