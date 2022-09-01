package alertruletemplates

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertRuleKind string

const (
	AlertRuleKindFusion                            AlertRuleKind = "Fusion"
	AlertRuleKindMLBehaviorAnalytics               AlertRuleKind = "MLBehaviorAnalytics"
	AlertRuleKindMicrosoftSecurityIncidentCreation AlertRuleKind = "MicrosoftSecurityIncidentCreation"
	AlertRuleKindNRT                               AlertRuleKind = "NRT"
	AlertRuleKindScheduled                         AlertRuleKind = "Scheduled"
	AlertRuleKindThreatIntelligence                AlertRuleKind = "ThreatIntelligence"
)

func PossibleValuesForAlertRuleKind() []string {
	return []string{
		string(AlertRuleKindFusion),
		string(AlertRuleKindMLBehaviorAnalytics),
		string(AlertRuleKindMicrosoftSecurityIncidentCreation),
		string(AlertRuleKindNRT),
		string(AlertRuleKindScheduled),
		string(AlertRuleKindThreatIntelligence),
	}
}

func parseAlertRuleKind(input string) (*AlertRuleKind, error) {
	vals := map[string]AlertRuleKind{
		"fusion":                            AlertRuleKindFusion,
		"mlbehavioranalytics":               AlertRuleKindMLBehaviorAnalytics,
		"microsoftsecurityincidentcreation": AlertRuleKindMicrosoftSecurityIncidentCreation,
		"nrt":                               AlertRuleKindNRT,
		"scheduled":                         AlertRuleKindScheduled,
		"threatintelligence":                AlertRuleKindThreatIntelligence,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertRuleKind(input)
	return &out, nil
}

type AlertSeverity string

const (
	AlertSeverityHigh          AlertSeverity = "High"
	AlertSeverityInformational AlertSeverity = "Informational"
	AlertSeverityLow           AlertSeverity = "Low"
	AlertSeverityMedium        AlertSeverity = "Medium"
)

func PossibleValuesForAlertSeverity() []string {
	return []string{
		string(AlertSeverityHigh),
		string(AlertSeverityInformational),
		string(AlertSeverityLow),
		string(AlertSeverityMedium),
	}
}

func parseAlertSeverity(input string) (*AlertSeverity, error) {
	vals := map[string]AlertSeverity{
		"high":          AlertSeverityHigh,
		"informational": AlertSeverityInformational,
		"low":           AlertSeverityLow,
		"medium":        AlertSeverityMedium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertSeverity(input)
	return &out, nil
}

type AttackTactic string

const (
	AttackTacticCollection          AttackTactic = "Collection"
	AttackTacticCommandAndControl   AttackTactic = "CommandAndControl"
	AttackTacticCredentialAccess    AttackTactic = "CredentialAccess"
	AttackTacticDefenseEvasion      AttackTactic = "DefenseEvasion"
	AttackTacticDiscovery           AttackTactic = "Discovery"
	AttackTacticExecution           AttackTactic = "Execution"
	AttackTacticExfiltration        AttackTactic = "Exfiltration"
	AttackTacticImpact              AttackTactic = "Impact"
	AttackTacticInitialAccess       AttackTactic = "InitialAccess"
	AttackTacticLateralMovement     AttackTactic = "LateralMovement"
	AttackTacticPersistence         AttackTactic = "Persistence"
	AttackTacticPreAttack           AttackTactic = "PreAttack"
	AttackTacticPrivilegeEscalation AttackTactic = "PrivilegeEscalation"
)

func PossibleValuesForAttackTactic() []string {
	return []string{
		string(AttackTacticCollection),
		string(AttackTacticCommandAndControl),
		string(AttackTacticCredentialAccess),
		string(AttackTacticDefenseEvasion),
		string(AttackTacticDiscovery),
		string(AttackTacticExecution),
		string(AttackTacticExfiltration),
		string(AttackTacticImpact),
		string(AttackTacticInitialAccess),
		string(AttackTacticLateralMovement),
		string(AttackTacticPersistence),
		string(AttackTacticPreAttack),
		string(AttackTacticPrivilegeEscalation),
	}
}

func parseAttackTactic(input string) (*AttackTactic, error) {
	vals := map[string]AttackTactic{
		"collection":          AttackTacticCollection,
		"commandandcontrol":   AttackTacticCommandAndControl,
		"credentialaccess":    AttackTacticCredentialAccess,
		"defenseevasion":      AttackTacticDefenseEvasion,
		"discovery":           AttackTacticDiscovery,
		"execution":           AttackTacticExecution,
		"exfiltration":        AttackTacticExfiltration,
		"impact":              AttackTacticImpact,
		"initialaccess":       AttackTacticInitialAccess,
		"lateralmovement":     AttackTacticLateralMovement,
		"persistence":         AttackTacticPersistence,
		"preattack":           AttackTacticPreAttack,
		"privilegeescalation": AttackTacticPrivilegeEscalation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttackTactic(input)
	return &out, nil
}
