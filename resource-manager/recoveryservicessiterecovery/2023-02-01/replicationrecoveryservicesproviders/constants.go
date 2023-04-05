package replicationrecoveryservicesproviders

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgentVersionStatus string

const (
	AgentVersionStatusDeprecated             AgentVersionStatus = "Deprecated"
	AgentVersionStatusNotSupported           AgentVersionStatus = "NotSupported"
	AgentVersionStatusSecurityUpdateRequired AgentVersionStatus = "SecurityUpdateRequired"
	AgentVersionStatusSupported              AgentVersionStatus = "Supported"
	AgentVersionStatusUpdateRequired         AgentVersionStatus = "UpdateRequired"
)

func PossibleValuesForAgentVersionStatus() []string {
	return []string{
		string(AgentVersionStatusDeprecated),
		string(AgentVersionStatusNotSupported),
		string(AgentVersionStatusSecurityUpdateRequired),
		string(AgentVersionStatusSupported),
		string(AgentVersionStatusUpdateRequired),
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
