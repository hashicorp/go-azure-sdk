package replicationprotectioncontainermappings

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgentAutoUpdateStatus string

const (
	AgentAutoUpdateStatusDisabled AgentAutoUpdateStatus = "Disabled"
	AgentAutoUpdateStatusEnabled  AgentAutoUpdateStatus = "Enabled"
)

func PossibleValuesForAgentAutoUpdateStatus() []string {
	return []string{
		string(AgentAutoUpdateStatusDisabled),
		string(AgentAutoUpdateStatusEnabled),
	}
}

type AutomationAccountAuthenticationType string

const (
	AutomationAccountAuthenticationTypeRunAsAccount           AutomationAccountAuthenticationType = "RunAsAccount"
	AutomationAccountAuthenticationTypeSystemAssignedIdentity AutomationAccountAuthenticationType = "SystemAssignedIdentity"
)

func PossibleValuesForAutomationAccountAuthenticationType() []string {
	return []string{
		string(AutomationAccountAuthenticationTypeRunAsAccount),
		string(AutomationAccountAuthenticationTypeSystemAssignedIdentity),
	}
}

type HealthErrorCustomerResolvability string

const (
	HealthErrorCustomerResolvabilityAllowed    HealthErrorCustomerResolvability = "Allowed"
	HealthErrorCustomerResolvabilityNotAllowed HealthErrorCustomerResolvability = "NotAllowed"
)

func PossibleValuesForHealthErrorCustomerResolvability() []string {
	return []string{
		string(HealthErrorCustomerResolvabilityAllowed),
		string(HealthErrorCustomerResolvabilityNotAllowed),
	}
}
