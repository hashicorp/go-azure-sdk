package onlinedeployment

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineDeployment interface {
	OnlineDeployment() BaseOnlineDeploymentImpl
}

var _ OnlineDeployment = BaseOnlineDeploymentImpl{}

type BaseOnlineDeploymentImpl struct {
	AppInsightsEnabled        *bool                          `json:"appInsightsEnabled,omitempty"`
	CodeConfiguration         *CodeConfiguration             `json:"codeConfiguration,omitempty"`
	Description               *string                        `json:"description,omitempty"`
	EgressPublicNetworkAccess *EgressPublicNetworkAccessType `json:"egressPublicNetworkAccess,omitempty"`
	EndpointComputeType       EndpointComputeType            `json:"endpointComputeType"`
	EnvironmentId             *string                        `json:"environmentId,omitempty"`
	EnvironmentVariables      *map[string]string             `json:"environmentVariables,omitempty"`
	InstanceType              *string                        `json:"instanceType,omitempty"`
	LivenessProbe             *ProbeSettings                 `json:"livenessProbe,omitempty"`
	Model                     *string                        `json:"model,omitempty"`
	ModelMountPath            *string                        `json:"modelMountPath,omitempty"`
	Properties                *map[string]string             `json:"properties,omitempty"`
	ProvisioningState         *DeploymentProvisioningState   `json:"provisioningState,omitempty"`
	ReadinessProbe            *ProbeSettings                 `json:"readinessProbe,omitempty"`
	RequestSettings           *OnlineRequestSettings         `json:"requestSettings,omitempty"`
	ScaleSettings             OnlineScaleSettings            `json:"scaleSettings"`
}

func (s BaseOnlineDeploymentImpl) OnlineDeployment() BaseOnlineDeploymentImpl {
	return s
}

var _ OnlineDeployment = RawOnlineDeploymentImpl{}

// RawOnlineDeploymentImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawOnlineDeploymentImpl struct {
	onlineDeployment BaseOnlineDeploymentImpl
	Type             string
	Values           map[string]interface{}
}

func (s RawOnlineDeploymentImpl) OnlineDeployment() BaseOnlineDeploymentImpl {
	return s.onlineDeployment
}

var _ json.Unmarshaler = &BaseOnlineDeploymentImpl{}

func (s *BaseOnlineDeploymentImpl) UnmarshalJSON(bytes []byte) error {
	type alias BaseOnlineDeploymentImpl
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into BaseOnlineDeploymentImpl: %+v", err)
	}

	s.AppInsightsEnabled = decoded.AppInsightsEnabled
	s.CodeConfiguration = decoded.CodeConfiguration
	s.Description = decoded.Description
	s.EgressPublicNetworkAccess = decoded.EgressPublicNetworkAccess
	s.EndpointComputeType = decoded.EndpointComputeType
	s.EnvironmentId = decoded.EnvironmentId
	s.EnvironmentVariables = decoded.EnvironmentVariables
	s.InstanceType = decoded.InstanceType
	s.LivenessProbe = decoded.LivenessProbe
	s.Model = decoded.Model
	s.ModelMountPath = decoded.ModelMountPath
	s.Properties = decoded.Properties
	s.ProvisioningState = decoded.ProvisioningState
	s.ReadinessProbe = decoded.ReadinessProbe
	s.RequestSettings = decoded.RequestSettings

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseOnlineDeploymentImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["scaleSettings"]; ok {
		impl, err := UnmarshalOnlineScaleSettingsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ScaleSettings' for 'BaseOnlineDeploymentImpl': %+v", err)
		}
		s.ScaleSettings = impl
	}
	return nil
}

func UnmarshalOnlineDeploymentImplementation(input []byte) (OnlineDeployment, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling OnlineDeployment into map[string]interface: %+v", err)
	}

	value, ok := temp["endpointComputeType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Kubernetes") {
		var out KubernetesOnlineDeployment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into KubernetesOnlineDeployment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Managed") {
		var out ManagedOnlineDeployment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedOnlineDeployment: %+v", err)
		}
		return out, nil
	}

	var parent BaseOnlineDeploymentImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseOnlineDeploymentImpl: %+v", err)
	}

	return RawOnlineDeploymentImpl{
		onlineDeployment: parent,
		Type:             value,
		Values:           temp,
	}, nil

}
