package replicationfabrics

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

type ProtectionHealth string

const (
	ProtectionHealthCritical ProtectionHealth = "Critical"
	ProtectionHealthNone     ProtectionHealth = "None"
	ProtectionHealthNormal   ProtectionHealth = "Normal"
	ProtectionHealthWarning  ProtectionHealth = "Warning"
)

func PossibleValuesForProtectionHealth() []string {
	return []string{
		string(ProtectionHealthCritical),
		string(ProtectionHealthNone),
		string(ProtectionHealthNormal),
		string(ProtectionHealthWarning),
	}
}

type RcmComponentStatus string

const (
	RcmComponentStatusCritical RcmComponentStatus = "Critical"
	RcmComponentStatusHealthy  RcmComponentStatus = "Healthy"
	RcmComponentStatusUnknown  RcmComponentStatus = "Unknown"
	RcmComponentStatusWarning  RcmComponentStatus = "Warning"
)

func PossibleValuesForRcmComponentStatus() []string {
	return []string{
		string(RcmComponentStatusCritical),
		string(RcmComponentStatusHealthy),
		string(RcmComponentStatusUnknown),
		string(RcmComponentStatusWarning),
	}
}
