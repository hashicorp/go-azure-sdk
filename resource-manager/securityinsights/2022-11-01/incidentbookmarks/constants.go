package incidentbookmarks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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

type AlertStatus string

const (
	AlertStatusDismissed  AlertStatus = "Dismissed"
	AlertStatusInProgress AlertStatus = "InProgress"
	AlertStatusNew        AlertStatus = "New"
	AlertStatusResolved   AlertStatus = "Resolved"
	AlertStatusUnknown    AlertStatus = "Unknown"
)

func PossibleValuesForAlertStatus() []string {
	return []string{
		string(AlertStatusDismissed),
		string(AlertStatusInProgress),
		string(AlertStatusNew),
		string(AlertStatusResolved),
		string(AlertStatusUnknown),
	}
}

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

type ConfidenceLevel string

const (
	ConfidenceLevelHigh    ConfidenceLevel = "High"
	ConfidenceLevelLow     ConfidenceLevel = "Low"
	ConfidenceLevelUnknown ConfidenceLevel = "Unknown"
)

func PossibleValuesForConfidenceLevel() []string {
	return []string{
		string(ConfidenceLevelHigh),
		string(ConfidenceLevelLow),
		string(ConfidenceLevelUnknown),
	}
}

type ConfidenceScoreStatus string

const (
	ConfidenceScoreStatusFinal         ConfidenceScoreStatus = "Final"
	ConfidenceScoreStatusInProcess     ConfidenceScoreStatus = "InProcess"
	ConfidenceScoreStatusNotApplicable ConfidenceScoreStatus = "NotApplicable"
	ConfidenceScoreStatusNotFinal      ConfidenceScoreStatus = "NotFinal"
)

func PossibleValuesForConfidenceScoreStatus() []string {
	return []string{
		string(ConfidenceScoreStatusFinal),
		string(ConfidenceScoreStatusInProcess),
		string(ConfidenceScoreStatusNotApplicable),
		string(ConfidenceScoreStatusNotFinal),
	}
}

type EntityKindEnum string

const (
	EntityKindEnumAccount          EntityKindEnum = "Account"
	EntityKindEnumAzureResource    EntityKindEnum = "AzureResource"
	EntityKindEnumBookmark         EntityKindEnum = "Bookmark"
	EntityKindEnumCloudApplication EntityKindEnum = "CloudApplication"
	EntityKindEnumDnsResolution    EntityKindEnum = "DnsResolution"
	EntityKindEnumFile             EntityKindEnum = "File"
	EntityKindEnumFileHash         EntityKindEnum = "FileHash"
	EntityKindEnumHost             EntityKindEnum = "Host"
	EntityKindEnumIP               EntityKindEnum = "Ip"
	EntityKindEnumIoTDevice        EntityKindEnum = "IoTDevice"
	EntityKindEnumMailCluster      EntityKindEnum = "MailCluster"
	EntityKindEnumMailMessage      EntityKindEnum = "MailMessage"
	EntityKindEnumMailbox          EntityKindEnum = "Mailbox"
	EntityKindEnumMalware          EntityKindEnum = "Malware"
	EntityKindEnumProcess          EntityKindEnum = "Process"
	EntityKindEnumRegistryKey      EntityKindEnum = "RegistryKey"
	EntityKindEnumRegistryValue    EntityKindEnum = "RegistryValue"
	EntityKindEnumSecurityAlert    EntityKindEnum = "SecurityAlert"
	EntityKindEnumSecurityGroup    EntityKindEnum = "SecurityGroup"
	EntityKindEnumSubmissionMail   EntityKindEnum = "SubmissionMail"
	EntityKindEnumUrl              EntityKindEnum = "Url"
)

func PossibleValuesForEntityKindEnum() []string {
	return []string{
		string(EntityKindEnumAccount),
		string(EntityKindEnumAzureResource),
		string(EntityKindEnumBookmark),
		string(EntityKindEnumCloudApplication),
		string(EntityKindEnumDnsResolution),
		string(EntityKindEnumFile),
		string(EntityKindEnumFileHash),
		string(EntityKindEnumHost),
		string(EntityKindEnumIP),
		string(EntityKindEnumIoTDevice),
		string(EntityKindEnumMailCluster),
		string(EntityKindEnumMailMessage),
		string(EntityKindEnumMailbox),
		string(EntityKindEnumMalware),
		string(EntityKindEnumProcess),
		string(EntityKindEnumRegistryKey),
		string(EntityKindEnumRegistryValue),
		string(EntityKindEnumSecurityAlert),
		string(EntityKindEnumSecurityGroup),
		string(EntityKindEnumSubmissionMail),
		string(EntityKindEnumUrl),
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

type KillChainIntent string

const (
	KillChainIntentCollection          KillChainIntent = "Collection"
	KillChainIntentCommandAndControl   KillChainIntent = "CommandAndControl"
	KillChainIntentCredentialAccess    KillChainIntent = "CredentialAccess"
	KillChainIntentDefenseEvasion      KillChainIntent = "DefenseEvasion"
	KillChainIntentDiscovery           KillChainIntent = "Discovery"
	KillChainIntentExecution           KillChainIntent = "Execution"
	KillChainIntentExfiltration        KillChainIntent = "Exfiltration"
	KillChainIntentExploitation        KillChainIntent = "Exploitation"
	KillChainIntentImpact              KillChainIntent = "Impact"
	KillChainIntentLateralMovement     KillChainIntent = "LateralMovement"
	KillChainIntentPersistence         KillChainIntent = "Persistence"
	KillChainIntentPrivilegeEscalation KillChainIntent = "PrivilegeEscalation"
	KillChainIntentProbing             KillChainIntent = "Probing"
	KillChainIntentUnknown             KillChainIntent = "Unknown"
)

func PossibleValuesForKillChainIntent() []string {
	return []string{
		string(KillChainIntentCollection),
		string(KillChainIntentCommandAndControl),
		string(KillChainIntentCredentialAccess),
		string(KillChainIntentDefenseEvasion),
		string(KillChainIntentDiscovery),
		string(KillChainIntentExecution),
		string(KillChainIntentExfiltration),
		string(KillChainIntentExploitation),
		string(KillChainIntentImpact),
		string(KillChainIntentLateralMovement),
		string(KillChainIntentPersistence),
		string(KillChainIntentPrivilegeEscalation),
		string(KillChainIntentProbing),
		string(KillChainIntentUnknown),
	}
}
