package agentdeployment

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AgentDeploymentProperties = ManagedAgentDeployment{}

type ManagedAgentDeployment struct {

	// Fields inherited from AgentDeploymentProperties

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

func (s ManagedAgentDeployment) AgentDeploymentProperties() BaseAgentDeploymentPropertiesImpl {
	return BaseAgentDeploymentPropertiesImpl{
		Agents:            s.Agents,
		DeploymentId:      s.DeploymentId,
		DeploymentType:    s.DeploymentType,
		Description:       s.Description,
		DisplayName:       s.DisplayName,
		Protocols:         s.Protocols,
		ProvisioningState: s.ProvisioningState,
		State:             s.State,
		Tags:              s.Tags,
	}
}

var _ json.Marshaler = ManagedAgentDeployment{}

func (s ManagedAgentDeployment) MarshalJSON() ([]byte, error) {
	type wrapper ManagedAgentDeployment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedAgentDeployment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedAgentDeployment: %+v", err)
	}

	decoded["deploymentType"] = "Managed"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedAgentDeployment: %+v", err)
	}

	return encoded, nil
}
