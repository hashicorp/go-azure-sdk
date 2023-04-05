package agents

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgentStatus string

const (
	AgentStatusExecuting         AgentStatus = "Executing"
	AgentStatusOffline           AgentStatus = "Offline"
	AgentStatusOnline            AgentStatus = "Online"
	AgentStatusRegistering       AgentStatus = "Registering"
	AgentStatusRequiresAttention AgentStatus = "RequiresAttention"
	AgentStatusUnregistering     AgentStatus = "Unregistering"
)

func PossibleValuesForAgentStatus() []string {
	return []string{
		string(AgentStatusExecuting),
		string(AgentStatusOffline),
		string(AgentStatusOnline),
		string(AgentStatusRegistering),
		string(AgentStatusRequiresAttention),
		string(AgentStatusUnregistering),
	}
}

func (s *AgentStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForAgentStatus() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = AgentStatus(decoded)
	return nil
}

type ProvisioningState string

const (
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateSucceeded),
	}
}

func (s *ProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForProvisioningState() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ProvisioningState(decoded)
	return nil
}
