package agentregistrationinformation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgentRegistrationKeyName string

const (
	AgentRegistrationKeyNamePrimary   AgentRegistrationKeyName = "primary"
	AgentRegistrationKeyNameSecondary AgentRegistrationKeyName = "secondary"
)

func PossibleValuesForAgentRegistrationKeyName() []string {
	return []string{
		string(AgentRegistrationKeyNamePrimary),
		string(AgentRegistrationKeyNameSecondary),
	}
}
