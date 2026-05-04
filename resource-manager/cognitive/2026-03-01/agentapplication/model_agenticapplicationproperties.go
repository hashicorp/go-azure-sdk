package agentapplication

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgenticApplicationProperties struct {
	AgentIdentityBlueprint  *AssignedIdentity                    `json:"agentIdentityBlueprint,omitempty"`
	Agents                  *[]AgentReferenceProperties          `json:"agents,omitempty"`
	AuthorizationPolicy     ApplicationAuthorizationPolicy       `json:"authorizationPolicy"`
	BaseURL                 *string                              `json:"baseUrl,omitempty"`
	DefaultInstanceIdentity *AssignedIdentity                    `json:"defaultInstanceIdentity,omitempty"`
	Description             *string                              `json:"description,omitempty"`
	DisplayName             *string                              `json:"displayName,omitempty"`
	IsEnabled               *bool                                `json:"isEnabled,omitempty"`
	ProvisioningState       *AgenticApplicationProvisioningState `json:"provisioningState,omitempty"`
	Tags                    *map[string]string                   `json:"tags,omitempty"`
	TrafficRoutingPolicy    *ApplicationTrafficRoutingPolicy     `json:"trafficRoutingPolicy,omitempty"`
}

var _ json.Unmarshaler = &AgenticApplicationProperties{}

func (s *AgenticApplicationProperties) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AgentIdentityBlueprint  *AssignedIdentity                    `json:"agentIdentityBlueprint,omitempty"`
		Agents                  *[]AgentReferenceProperties          `json:"agents,omitempty"`
		BaseURL                 *string                              `json:"baseUrl,omitempty"`
		DefaultInstanceIdentity *AssignedIdentity                    `json:"defaultInstanceIdentity,omitempty"`
		Description             *string                              `json:"description,omitempty"`
		DisplayName             *string                              `json:"displayName,omitempty"`
		IsEnabled               *bool                                `json:"isEnabled,omitempty"`
		ProvisioningState       *AgenticApplicationProvisioningState `json:"provisioningState,omitempty"`
		Tags                    *map[string]string                   `json:"tags,omitempty"`
		TrafficRoutingPolicy    *ApplicationTrafficRoutingPolicy     `json:"trafficRoutingPolicy,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AgentIdentityBlueprint = decoded.AgentIdentityBlueprint
	s.Agents = decoded.Agents
	s.BaseURL = decoded.BaseURL
	s.DefaultInstanceIdentity = decoded.DefaultInstanceIdentity
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.IsEnabled = decoded.IsEnabled
	s.ProvisioningState = decoded.ProvisioningState
	s.Tags = decoded.Tags
	s.TrafficRoutingPolicy = decoded.TrafficRoutingPolicy

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AgenticApplicationProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["authorizationPolicy"]; ok {
		impl, err := UnmarshalApplicationAuthorizationPolicyImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AuthorizationPolicy' for 'AgenticApplicationProperties': %+v", err)
		}
		s.AuthorizationPolicy = impl
	}

	return nil
}
