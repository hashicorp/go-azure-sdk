package agentdeployment

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgentDeploymentProperties interface {
	AgentDeploymentProperties() BaseAgentDeploymentPropertiesImpl
}

var _ AgentDeploymentProperties = BaseAgentDeploymentPropertiesImpl{}

type BaseAgentDeploymentPropertiesImpl struct {
	Agents            *[]VersionedAgentReference        `json:"agents,omitempty"`
	DeploymentId      *string                           `json:"deploymentId,omitempty"`
	DeploymentType    AgentDeploymentType               `json:"deploymentType"`
	Description       *string                           `json:"description,omitempty"`
	DisplayName       *string                           `json:"displayName,omitempty"`
	Protocols         *[]AgentProtocolVersion           `json:"protocols,omitempty"`
	ProvisioningState *AgentDeploymentProvisioningState `json:"provisioningState,omitempty"`
	State             *AgentDeploymentState             `json:"state,omitempty"`
	Tags              *map[string]string                `json:"tags,omitempty"`
}

func (s BaseAgentDeploymentPropertiesImpl) AgentDeploymentProperties() BaseAgentDeploymentPropertiesImpl {
	return s
}

var _ AgentDeploymentProperties = RawAgentDeploymentPropertiesImpl{}

// RawAgentDeploymentPropertiesImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAgentDeploymentPropertiesImpl struct {
	agentDeploymentProperties BaseAgentDeploymentPropertiesImpl
	Type                      string
	Values                    map[string]interface{}
}

func (s RawAgentDeploymentPropertiesImpl) AgentDeploymentProperties() BaseAgentDeploymentPropertiesImpl {
	return s.agentDeploymentProperties
}

func UnmarshalAgentDeploymentPropertiesImplementation(input []byte) (AgentDeploymentProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AgentDeploymentProperties into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["deploymentType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Hosted") {
		var out HostedAgentDeployment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HostedAgentDeployment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Managed") {
		var out ManagedAgentDeployment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedAgentDeployment: %+v", err)
		}
		return out, nil
	}

	var parent BaseAgentDeploymentPropertiesImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAgentDeploymentPropertiesImpl: %+v", err)
	}

	return RawAgentDeploymentPropertiesImpl{
		agentDeploymentProperties: parent,
		Type:                      value,
		Values:                    temp,
	}, nil

}
