package incidentalerts

import "strings"

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

func parseAlertStatus(input string) (*AlertStatus, error) {
	vals := map[string]AlertStatus{
		"dismissed":  AlertStatusDismissed,
		"inprogress": AlertStatusInProgress,
		"new":        AlertStatusNew,
		"resolved":   AlertStatusResolved,
		"unknown":    AlertStatusUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertStatus(input)
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

func parseConfidenceLevel(input string) (*ConfidenceLevel, error) {
	vals := map[string]ConfidenceLevel{
		"high":    ConfidenceLevelHigh,
		"low":     ConfidenceLevelLow,
		"unknown": ConfidenceLevelUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConfidenceLevel(input)
	return &out, nil
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

func parseConfidenceScoreStatus(input string) (*ConfidenceScoreStatus, error) {
	vals := map[string]ConfidenceScoreStatus{
		"final":         ConfidenceScoreStatusFinal,
		"inprocess":     ConfidenceScoreStatusInProcess,
		"notapplicable": ConfidenceScoreStatusNotApplicable,
		"notfinal":      ConfidenceScoreStatusNotFinal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConfidenceScoreStatus(input)
	return &out, nil
}

type EntityKind string

const (
	EntityKindAccount          EntityKind = "Account"
	EntityKindAzureResource    EntityKind = "AzureResource"
	EntityKindBookmark         EntityKind = "Bookmark"
	EntityKindCloudApplication EntityKind = "CloudApplication"
	EntityKindDnsResolution    EntityKind = "DnsResolution"
	EntityKindFile             EntityKind = "File"
	EntityKindFileHash         EntityKind = "FileHash"
	EntityKindHost             EntityKind = "Host"
	EntityKindIoTDevice        EntityKind = "IoTDevice"
	EntityKindIp               EntityKind = "Ip"
	EntityKindMailCluster      EntityKind = "MailCluster"
	EntityKindMailMessage      EntityKind = "MailMessage"
	EntityKindMailbox          EntityKind = "Mailbox"
	EntityKindMalware          EntityKind = "Malware"
	EntityKindProcess          EntityKind = "Process"
	EntityKindRegistryKey      EntityKind = "RegistryKey"
	EntityKindRegistryValue    EntityKind = "RegistryValue"
	EntityKindSecurityAlert    EntityKind = "SecurityAlert"
	EntityKindSecurityGroup    EntityKind = "SecurityGroup"
	EntityKindSubmissionMail   EntityKind = "SubmissionMail"
	EntityKindUrl              EntityKind = "Url"
)

func PossibleValuesForEntityKind() []string {
	return []string{
		string(EntityKindAccount),
		string(EntityKindAzureResource),
		string(EntityKindBookmark),
		string(EntityKindCloudApplication),
		string(EntityKindDnsResolution),
		string(EntityKindFile),
		string(EntityKindFileHash),
		string(EntityKindHost),
		string(EntityKindIoTDevice),
		string(EntityKindIp),
		string(EntityKindMailCluster),
		string(EntityKindMailMessage),
		string(EntityKindMailbox),
		string(EntityKindMalware),
		string(EntityKindProcess),
		string(EntityKindRegistryKey),
		string(EntityKindRegistryValue),
		string(EntityKindSecurityAlert),
		string(EntityKindSecurityGroup),
		string(EntityKindSubmissionMail),
		string(EntityKindUrl),
	}
}

func parseEntityKind(input string) (*EntityKind, error) {
	vals := map[string]EntityKind{
		"account":          EntityKindAccount,
		"azureresource":    EntityKindAzureResource,
		"bookmark":         EntityKindBookmark,
		"cloudapplication": EntityKindCloudApplication,
		"dnsresolution":    EntityKindDnsResolution,
		"file":             EntityKindFile,
		"filehash":         EntityKindFileHash,
		"host":             EntityKindHost,
		"iotdevice":        EntityKindIoTDevice,
		"ip":               EntityKindIp,
		"mailcluster":      EntityKindMailCluster,
		"mailmessage":      EntityKindMailMessage,
		"mailbox":          EntityKindMailbox,
		"malware":          EntityKindMalware,
		"process":          EntityKindProcess,
		"registrykey":      EntityKindRegistryKey,
		"registryvalue":    EntityKindRegistryValue,
		"securityalert":    EntityKindSecurityAlert,
		"securitygroup":    EntityKindSecurityGroup,
		"submissionmail":   EntityKindSubmissionMail,
		"url":              EntityKindUrl,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EntityKind(input)
	return &out, nil
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

func parseIncidentSeverity(input string) (*IncidentSeverity, error) {
	vals := map[string]IncidentSeverity{
		"high":          IncidentSeverityHigh,
		"informational": IncidentSeverityInformational,
		"low":           IncidentSeverityLow,
		"medium":        IncidentSeverityMedium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IncidentSeverity(input)
	return &out, nil
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

func parseKillChainIntent(input string) (*KillChainIntent, error) {
	vals := map[string]KillChainIntent{
		"collection":          KillChainIntentCollection,
		"commandandcontrol":   KillChainIntentCommandAndControl,
		"credentialaccess":    KillChainIntentCredentialAccess,
		"defenseevasion":      KillChainIntentDefenseEvasion,
		"discovery":           KillChainIntentDiscovery,
		"execution":           KillChainIntentExecution,
		"exfiltration":        KillChainIntentExfiltration,
		"exploitation":        KillChainIntentExploitation,
		"impact":              KillChainIntentImpact,
		"lateralmovement":     KillChainIntentLateralMovement,
		"persistence":         KillChainIntentPersistence,
		"privilegeescalation": KillChainIntentPrivilegeEscalation,
		"probing":             KillChainIntentProbing,
		"unknown":             KillChainIntentUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KillChainIntent(input)
	return &out, nil
}
