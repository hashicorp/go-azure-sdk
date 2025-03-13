package endpoint

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EndpointDeploymentResourceProperties = ManagedOnlineEndpointDeploymentResourceProperties{}

type ManagedOnlineEndpointDeploymentResourceProperties struct {
	EndpointComputeType *EndpointComputeType `json:"endpointComputeType,omitempty"`
	Model               *string              `json:"model,omitempty"`

	// Fields inherited from EndpointDeploymentResourceProperties

	FailureReason     *string                           `json:"failureReason,omitempty"`
	ProvisioningState *DefaultResourceProvisioningState `json:"provisioningState,omitempty"`
	Type              string                            `json:"type"`
}

func (s ManagedOnlineEndpointDeploymentResourceProperties) EndpointDeploymentResourceProperties() BaseEndpointDeploymentResourcePropertiesImpl {
	return BaseEndpointDeploymentResourcePropertiesImpl{
		FailureReason:     s.FailureReason,
		ProvisioningState: s.ProvisioningState,
		Type:              s.Type,
	}
}

var _ json.Marshaler = ManagedOnlineEndpointDeploymentResourceProperties{}

func (s ManagedOnlineEndpointDeploymentResourceProperties) MarshalJSON() ([]byte, error) {
	type wrapper ManagedOnlineEndpointDeploymentResourceProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedOnlineEndpointDeploymentResourceProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedOnlineEndpointDeploymentResourceProperties: %+v", err)
	}

	decoded["type"] = "managedOnlineEndpoint"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedOnlineEndpointDeploymentResourceProperties: %+v", err)
	}

	return encoded, nil
}
