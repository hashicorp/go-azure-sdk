package endpoint

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EndpointResourceProperties = OpenAIEndpointResourceProperties{}

type OpenAIEndpointResourceProperties struct {

	// Fields inherited from EndpointResourceProperties

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

func (s OpenAIEndpointResourceProperties) EndpointResourceProperties() BaseEndpointResourcePropertiesImpl {
	return BaseEndpointResourcePropertiesImpl{
		AssociatedResourceId:           s.AssociatedResourceId,
		Deployments:                    s.Deployments,
		EndpointType:                   s.EndpointType,
		EndpointUri:                    s.EndpointUri,
		FailureReason:                  s.FailureReason,
		Location:                       s.Location,
		Name:                           s.Name,
		ProvisioningState:              s.ProvisioningState,
		ShouldCreateAiServicesEndpoint: s.ShouldCreateAiServicesEndpoint,
	}
}

var _ json.Marshaler = OpenAIEndpointResourceProperties{}

func (s OpenAIEndpointResourceProperties) MarshalJSON() ([]byte, error) {
	type wrapper OpenAIEndpointResourceProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OpenAIEndpointResourceProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OpenAIEndpointResourceProperties: %+v", err)
	}

	decoded["endpointType"] = "Azure.OpenAI"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OpenAIEndpointResourceProperties: %+v", err)
	}

	return encoded, nil
}
