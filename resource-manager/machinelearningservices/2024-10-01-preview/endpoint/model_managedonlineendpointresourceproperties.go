package endpoint

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EndpointResourceProperties = ManagedOnlineEndpointResourceProperties{}

type ManagedOnlineEndpointResourceProperties struct {
	AuthMode      *EndpointAuthMode `json:"authMode,omitempty"`
	Compute       *string           `json:"compute,omitempty"`
	Description   *string           `json:"description,omitempty"`
	MirrorTraffic *map[string]int64 `json:"mirrorTraffic,omitempty"`
	ScoringUri    *string           `json:"scoringUri,omitempty"`
	Traffic       *map[string]int64 `json:"traffic,omitempty"`

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

func (s ManagedOnlineEndpointResourceProperties) EndpointResourceProperties() BaseEndpointResourcePropertiesImpl {
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

var _ json.Marshaler = ManagedOnlineEndpointResourceProperties{}

func (s ManagedOnlineEndpointResourceProperties) MarshalJSON() ([]byte, error) {
	type wrapper ManagedOnlineEndpointResourceProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedOnlineEndpointResourceProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedOnlineEndpointResourceProperties: %+v", err)
	}

	decoded["endpointType"] = "managedOnlineEndpoint"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedOnlineEndpointResourceProperties: %+v", err)
	}

	return encoded, nil
}
