package v2workspaceconnectionresource

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EndpointDeploymentResourceProperties = ContentSafetyEndpointDeploymentResourceProperties{}

type ContentSafetyEndpointDeploymentResourceProperties struct {
	Model                EndpointDeploymentModel              `json:"model"`
	RaiPolicyName        *string                              `json:"raiPolicyName,omitempty"`
	Sku                  *CognitiveServicesSku                `json:"sku,omitempty"`
	VersionUpgradeOption *DeploymentModelVersionUpgradeOption `json:"versionUpgradeOption,omitempty"`

	// Fields inherited from EndpointDeploymentResourceProperties

	FailureReason     *string                           `json:"failureReason,omitempty"`
	ProvisioningState *DefaultResourceProvisioningState `json:"provisioningState,omitempty"`
	Type              string                            `json:"type"`
}

func (s ContentSafetyEndpointDeploymentResourceProperties) EndpointDeploymentResourceProperties() BaseEndpointDeploymentResourcePropertiesImpl {
	return BaseEndpointDeploymentResourcePropertiesImpl{
		FailureReason:     s.FailureReason,
		ProvisioningState: s.ProvisioningState,
		Type:              s.Type,
	}
}

var _ json.Marshaler = ContentSafetyEndpointDeploymentResourceProperties{}

func (s ContentSafetyEndpointDeploymentResourceProperties) MarshalJSON() ([]byte, error) {
	type wrapper ContentSafetyEndpointDeploymentResourceProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ContentSafetyEndpointDeploymentResourceProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ContentSafetyEndpointDeploymentResourceProperties: %+v", err)
	}

	decoded["type"] = "Azure.ContentSafety"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ContentSafetyEndpointDeploymentResourceProperties: %+v", err)
	}

	return encoded, nil
}
