package replicationvaulthealth

import "strings"

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

func parseHealthErrorCategory(input string) (*HealthErrorCategory, error) {
	vals := map[string]HealthErrorCategory{
		"agentautoupdateartifactdeleted":     HealthErrorCategoryAgentAutoUpdateArtifactDeleted,
		"agentautoupdateinfra":               HealthErrorCategoryAgentAutoUpdateInfra,
		"agentautoupdaterunasaccount":        HealthErrorCategoryAgentAutoUpdateRunAsAccount,
		"agentautoupdaterunasaccountexpired": HealthErrorCategoryAgentAutoUpdateRunAsAccountExpired,
		"agentautoupdaterunasaccountexpiry":  HealthErrorCategoryAgentAutoUpdateRunAsAccountExpiry,
		"configuration":                      HealthErrorCategoryConfiguration,
		"fabricinfrastructure":               HealthErrorCategoryFabricInfrastructure,
		"none":                               HealthErrorCategoryNone,
		"replication":                        HealthErrorCategoryReplication,
		"testfailover":                       HealthErrorCategoryTestFailover,
		"versionexpiry":                      HealthErrorCategoryVersionExpiry,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HealthErrorCategory(input)
	return &out, nil
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

func parseHealthErrorCustomerResolvability(input string) (*HealthErrorCustomerResolvability, error) {
	vals := map[string]HealthErrorCustomerResolvability{
		"allowed":    HealthErrorCustomerResolvabilityAllowed,
		"notallowed": HealthErrorCustomerResolvabilityNotAllowed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HealthErrorCustomerResolvability(input)
	return &out, nil
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

func parseSeverity(input string) (*Severity, error) {
	vals := map[string]Severity{
		"error":   SeverityError,
		"info":    SeverityInfo,
		"none":    SeverityNONE,
		"warning": SeverityWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Severity(input)
	return &out, nil
}
