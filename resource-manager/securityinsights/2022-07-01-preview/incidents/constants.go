package incidents

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttackTactic string

const (
	AttackTacticCollection              AttackTactic = "Collection"
	AttackTacticCommandAndControl       AttackTactic = "CommandAndControl"
	AttackTacticCredentialAccess        AttackTactic = "CredentialAccess"
	AttackTacticDefenseEvasion          AttackTactic = "DefenseEvasion"
	AttackTacticDiscovery               AttackTactic = "Discovery"
	AttackTacticExecution               AttackTactic = "Execution"
	AttackTacticExfiltration            AttackTactic = "Exfiltration"
	AttackTacticImpact                  AttackTactic = "Impact"
	AttackTacticImpairProcessControl    AttackTactic = "ImpairProcessControl"
	AttackTacticInhibitResponseFunction AttackTactic = "InhibitResponseFunction"
	AttackTacticInitialAccess           AttackTactic = "InitialAccess"
	AttackTacticLateralMovement         AttackTactic = "LateralMovement"
	AttackTacticPersistence             AttackTactic = "Persistence"
	AttackTacticPreAttack               AttackTactic = "PreAttack"
	AttackTacticPrivilegeEscalation     AttackTactic = "PrivilegeEscalation"
	AttackTacticReconnaissance          AttackTactic = "Reconnaissance"
	AttackTacticResourceDevelopment     AttackTactic = "ResourceDevelopment"
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
		string(AttackTacticImpairProcessControl),
		string(AttackTacticInhibitResponseFunction),
		string(AttackTacticInitialAccess),
		string(AttackTacticLateralMovement),
		string(AttackTacticPersistence),
		string(AttackTacticPreAttack),
		string(AttackTacticPrivilegeEscalation),
		string(AttackTacticReconnaissance),
		string(AttackTacticResourceDevelopment),
	}
}

type IncidentClassification string

const (
	IncidentClassificationBenignPositive IncidentClassification = "BenignPositive"
	IncidentClassificationFalsePositive  IncidentClassification = "FalsePositive"
	IncidentClassificationTruePositive   IncidentClassification = "TruePositive"
	IncidentClassificationUndetermined   IncidentClassification = "Undetermined"
)

func PossibleValuesForIncidentClassification() []string {
	return []string{
		string(IncidentClassificationBenignPositive),
		string(IncidentClassificationFalsePositive),
		string(IncidentClassificationTruePositive),
		string(IncidentClassificationUndetermined),
	}
}

type IncidentClassificationReason string

const (
	IncidentClassificationReasonInaccurateData        IncidentClassificationReason = "InaccurateData"
	IncidentClassificationReasonIncorrectAlertLogic   IncidentClassificationReason = "IncorrectAlertLogic"
	IncidentClassificationReasonSuspiciousActivity    IncidentClassificationReason = "SuspiciousActivity"
	IncidentClassificationReasonSuspiciousButExpected IncidentClassificationReason = "SuspiciousButExpected"
)

func PossibleValuesForIncidentClassificationReason() []string {
	return []string{
		string(IncidentClassificationReasonInaccurateData),
		string(IncidentClassificationReasonIncorrectAlertLogic),
		string(IncidentClassificationReasonSuspiciousActivity),
		string(IncidentClassificationReasonSuspiciousButExpected),
	}
}

type IncidentLabelType string

const (
	IncidentLabelTypeAutoAssigned IncidentLabelType = "AutoAssigned"
	IncidentLabelTypeUser         IncidentLabelType = "User"
)

func PossibleValuesForIncidentLabelType() []string {
	return []string{
		string(IncidentLabelTypeAutoAssigned),
		string(IncidentLabelTypeUser),
	}
}

type IncidentSeverity string

const (
	IncidentSeverityHigh          IncidentSeverity = "High"
	IncidentSeverityInformational IncidentSeverity = "Informational"
	IncidentSeverityLow           IncidentSeverity = "Low"
	IncidentSeverityMedium        IncidentSeverity = "Medium"
)

func PossibleValuesForIncidentSeverity() []string {
	return []string{
		string(IncidentSeverityHigh),
		string(IncidentSeverityInformational),
		string(IncidentSeverityLow),
		string(IncidentSeverityMedium),
	}
}

type IncidentStatus string

const (
	IncidentStatusActive IncidentStatus = "Active"
	IncidentStatusClosed IncidentStatus = "Closed"
	IncidentStatusNew    IncidentStatus = "New"
)

func PossibleValuesForIncidentStatus() []string {
	return []string{
		string(IncidentStatusActive),
		string(IncidentStatusClosed),
		string(IncidentStatusNew),
	}
}

type OwnerType string

const (
	OwnerTypeGroup   OwnerType = "Group"
	OwnerTypeUnknown OwnerType = "Unknown"
	OwnerTypeUser    OwnerType = "User"
)

func PossibleValuesForOwnerType() []string {
	return []string{
		string(OwnerTypeGroup),
		string(OwnerTypeUnknown),
		string(OwnerTypeUser),
	}
}
