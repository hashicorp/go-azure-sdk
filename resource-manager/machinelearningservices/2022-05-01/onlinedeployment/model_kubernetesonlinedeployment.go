package onlinedeployment

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OnlineDeployment = KubernetesOnlineDeployment{}

type KubernetesOnlineDeployment struct {
	ContainerResourceRequirements *ContainerResourceRequirements `json:"containerResourceRequirements,omitempty"`

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

var _ json.Marshaler = KubernetesOnlineDeployment{}

func (s KubernetesOnlineDeployment) MarshalJSON() ([]byte, error) {
	type wrapper KubernetesOnlineDeployment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling KubernetesOnlineDeployment: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling KubernetesOnlineDeployment: %+v", err)
	}
	decoded["endpointComputeType"] = "Kubernetes"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling KubernetesOnlineDeployment: %+v", err)
	}

	return encoded, nil
}
