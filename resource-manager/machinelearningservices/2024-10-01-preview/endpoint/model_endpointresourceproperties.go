package endpoint

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndpointResourceProperties interface {
	EndpointResourceProperties() BaseEndpointResourcePropertiesImpl
}

var _ EndpointResourceProperties = BaseEndpointResourcePropertiesImpl{}

type BaseEndpointResourcePropertiesImpl struct {
	AssociatedResourceId           *string                                              `json:"associatedResourceId,omitempty"`
	Deployments                    *[]EndpointDeploymentResourcePropertiesBasicResource `json:"deployments,omitempty"`
	EndpointType                   EndpointType                                         `json:"endpointType"`
	EndpointUri                    *string                                              `json:"endpointUri,omitempty"`
	FailureReason                  *string                                              `json:"failureReason,omitempty"`
	Location                       *string                                              `json:"location,omitempty"`
	Name                           *string                                              `json:"name,omitempty"`
	ProvisioningState              *DefaultResourceProvisioningState                    `json:"provisioningState,omitempty"`
	ShouldCreateAiServicesEndpoint *bool                                                `json:"shouldCreateAiServicesEndpoint,omitempty"`
}

func (s BaseEndpointResourcePropertiesImpl) EndpointResourceProperties() BaseEndpointResourcePropertiesImpl {
	return s
}

var _ EndpointResourceProperties = RawEndpointResourcePropertiesImpl{}

// RawEndpointResourcePropertiesImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawEndpointResourcePropertiesImpl struct {
	endpointResourceProperties BaseEndpointResourcePropertiesImpl
	Type                       string
	Values                     map[string]interface{}
}

func (s RawEndpointResourcePropertiesImpl) EndpointResourceProperties() BaseEndpointResourcePropertiesImpl {
	return s.endpointResourceProperties
}

func UnmarshalEndpointResourcePropertiesImplementation(input []byte) (EndpointResourceProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EndpointResourceProperties into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["endpointType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Azure.ContentSafety") {
		var out ContentSafetyEndpointResourceProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ContentSafetyEndpointResourceProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "managedOnlineEndpoint") {
		var out ManagedOnlineEndpointResourceProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedOnlineEndpointResourceProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Azure.OpenAI") {
		var out OpenAIEndpointResourceProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OpenAIEndpointResourceProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "serverlessEndpoint") {
		var out ServerlessEndpointResourceProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServerlessEndpointResourceProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Azure.Speech") {
		var out SpeechEndpointResourceProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SpeechEndpointResourceProperties: %+v", err)
		}
		return out, nil
	}

	var parent BaseEndpointResourcePropertiesImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEndpointResourcePropertiesImpl: %+v", err)
	}

	return RawEndpointResourcePropertiesImpl{
		endpointResourceProperties: parent,
		Type:                       value,
		Values:                     temp,
	}, nil

}
