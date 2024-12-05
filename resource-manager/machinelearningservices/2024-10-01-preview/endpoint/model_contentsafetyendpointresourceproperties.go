package endpoint

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EndpointResourceProperties = ContentSafetyEndpointResourceProperties{}

type ContentSafetyEndpointResourceProperties struct {

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

func (s ContentSafetyEndpointResourceProperties) EndpointResourceProperties() BaseEndpointResourcePropertiesImpl {
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

var _ json.Marshaler = ContentSafetyEndpointResourceProperties{}

func (s ContentSafetyEndpointResourceProperties) MarshalJSON() ([]byte, error) {
	type wrapper ContentSafetyEndpointResourceProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ContentSafetyEndpointResourceProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ContentSafetyEndpointResourceProperties: %+v", err)
	}

	decoded["endpointType"] = "Azure.ContentSafety"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ContentSafetyEndpointResourceProperties: %+v", err)
	}

	return encoded, nil
}
