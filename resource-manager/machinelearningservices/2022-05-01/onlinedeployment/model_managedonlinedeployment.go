package onlinedeployment

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OnlineDeployment = ManagedOnlineDeployment{}

type ManagedOnlineDeployment struct {

	// Fields inherited from OnlineDeployment
	AppInsightsEnabled   *bool                        `json:"appInsightsEnabled,omitempty"`
	CodeConfiguration    *CodeConfiguration           `json:"codeConfiguration,omitempty"`
	Description          *string                      `json:"description,omitempty"`
	EnvironmentId        *string                      `json:"environmentId,omitempty"`
	EnvironmentVariables *map[string]string           `json:"environmentVariables,omitempty"`
	InstanceType         *string                      `json:"instanceType,omitempty"`
	LivenessProbe        *ProbeSettings               `json:"livenessProbe,omitempty"`
	Model                *string                      `json:"model,omitempty"`
	ModelMountPath       *string                      `json:"modelMountPath,omitempty"`
	Properties           *map[string]string           `json:"properties,omitempty"`
	ProvisioningState    *DeploymentProvisioningState `json:"provisioningState,omitempty"`
	ReadinessProbe       *ProbeSettings               `json:"readinessProbe,omitempty"`
	RequestSettings      *OnlineRequestSettings       `json:"requestSettings,omitempty"`
	ScaleSettings        OnlineScaleSettings          `json:"scaleSettings"`
}

var _ json.Marshaler = ManagedOnlineDeployment{}

func (s ManagedOnlineDeployment) MarshalJSON() ([]byte, error) {
	type wrapper ManagedOnlineDeployment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedOnlineDeployment: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedOnlineDeployment: %+v", err)
	}
	decoded["endpointComputeType"] = "Managed"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedOnlineDeployment: %+v", err)
	}

	return encoded, nil
}
