package replicationvaulthealth

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HealthErrorCategory string

const (
	HealthErrorCategoryAgentAutoUpdateArtifactDeleted     HealthErrorCategory = "AgentAutoUpdateArtifactDeleted"
	HealthErrorCategoryAgentAutoUpdateInfra               HealthErrorCategory = "AgentAutoUpdateInfra"
	HealthErrorCategoryAgentAutoUpdateRunAsAccount        HealthErrorCategory = "AgentAutoUpdateRunAsAccount"
	HealthErrorCategoryAgentAutoUpdateRunAsAccountExpired HealthErrorCategory = "AgentAutoUpdateRunAsAccountExpired"
	HealthErrorCategoryAgentAutoUpdateRunAsAccountExpiry  HealthErrorCategory = "AgentAutoUpdateRunAsAccountExpiry"
	HealthErrorCategoryConfiguration                      HealthErrorCategory = "Configuration"
	HealthErrorCategoryFabricInfrastructure               HealthErrorCategory = "FabricInfrastructure"
	HealthErrorCategoryNone                               HealthErrorCategory = "None"
	HealthErrorCategoryReplication                        HealthErrorCategory = "Replication"
	HealthErrorCategoryTestFailover                       HealthErrorCategory = "TestFailover"
	HealthErrorCategoryVersionExpiry                      HealthErrorCategory = "VersionExpiry"
)

func PossibleValuesForHealthErrorCategory() []string {
	return []string{
		string(HealthErrorCategoryAgentAutoUpdateArtifactDeleted),
		string(HealthErrorCategoryAgentAutoUpdateInfra),
		string(HealthErrorCategoryAgentAutoUpdateRunAsAccount),
		string(HealthErrorCategoryAgentAutoUpdateRunAsAccountExpired),
		string(HealthErrorCategoryAgentAutoUpdateRunAsAccountExpiry),
		string(HealthErrorCategoryConfiguration),
		string(HealthErrorCategoryFabricInfrastructure),
		string(HealthErrorCategoryNone),
		string(HealthErrorCategoryReplication),
		string(HealthErrorCategoryTestFailover),
		string(HealthErrorCategoryVersionExpiry),
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

type Severity string

const (
	SeverityError   Severity = "Error"
	SeverityInfo    Severity = "Info"
	SeverityNONE    Severity = "NONE"
	SeverityWarning Severity = "Warning"
)

func PossibleValuesForSeverity() []string {
	return []string{
		string(SeverityError),
		string(SeverityInfo),
		string(SeverityNONE),
		string(SeverityWarning),
	}
}
