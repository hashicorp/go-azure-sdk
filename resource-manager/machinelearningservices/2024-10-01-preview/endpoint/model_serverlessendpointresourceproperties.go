package endpoint

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EndpointResourceProperties = ServerlessEndpointResourceProperties{}

type ServerlessEndpointResourceProperties struct {
	AuthMode                  *ServerlessInferenceEndpointAuthMode   `json:"authMode,omitempty"`
	CapacityReservation       *ServerlessEndpointCapacityReservation `json:"capacityReservation,omitempty"`
	ContentSafety             *ServerlessEndpointContentSafety       `json:"contentSafety,omitempty"`
	EndpointState             *ServerlessEndpointState               `json:"endpointState,omitempty"`
	InferenceEndpoint         *ServerlessEndpointInferenceEndpoint   `json:"inferenceEndpoint,omitempty"`
	MarketplaceSubscriptionId *string                                `json:"marketplaceSubscriptionId,omitempty"`
	Metadata                  *interface{}                           `json:"metadata,omitempty"`
	ModelSettings             *ServerlessEndpointModelSettings       `json:"modelSettings,omitempty"`
	Offer                     *ServerlessOffer                       `json:"offer,omitempty"`

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

func (s ServerlessEndpointResourceProperties) EndpointResourceProperties() BaseEndpointResourcePropertiesImpl {
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

var _ json.Marshaler = ServerlessEndpointResourceProperties{}

func (s ServerlessEndpointResourceProperties) MarshalJSON() ([]byte, error) {
	type wrapper ServerlessEndpointResourceProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ServerlessEndpointResourceProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ServerlessEndpointResourceProperties: %+v", err)
	}

	decoded["endpointType"] = "serverlessEndpoint"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ServerlessEndpointResourceProperties: %+v", err)
	}

	return encoded, nil
}
