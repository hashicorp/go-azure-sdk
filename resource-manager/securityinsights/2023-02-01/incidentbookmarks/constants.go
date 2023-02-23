package incidentbookmarks

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

func parseAttackTactic(input string) (*AttackTactic, error) {
	vals := map[string]AttackTactic{
		"collection":              AttackTacticCollection,
		"commandandcontrol":       AttackTacticCommandAndControl,
		"credentialaccess":        AttackTacticCredentialAccess,
		"defenseevasion":          AttackTacticDefenseEvasion,
		"discovery":               AttackTacticDiscovery,
		"execution":               AttackTacticExecution,
		"exfiltration":            AttackTacticExfiltration,
		"impact":                  AttackTacticImpact,
		"impairprocesscontrol":    AttackTacticImpairProcessControl,
		"inhibitresponsefunction": AttackTacticInhibitResponseFunction,
		"initialaccess":           AttackTacticInitialAccess,
		"lateralmovement":         AttackTacticLateralMovement,
		"persistence":             AttackTacticPersistence,
		"preattack":               AttackTacticPreAttack,
		"privilegeescalation":     AttackTacticPrivilegeEscalation,
		"reconnaissance":          AttackTacticReconnaissance,
		"resourcedevelopment":     AttackTacticResourceDevelopment,
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

func parseEntityKindEnum(input string) (*EntityKindEnum, error) {
	vals := map[string]EntityKindEnum{
		"account":          EntityKindEnumAccount,
		"azureresource":    EntityKindEnumAzureResource,
		"bookmark":         EntityKindEnumBookmark,
		"cloudapplication": EntityKindEnumCloudApplication,
		"dnsresolution":    EntityKindEnumDnsResolution,
		"file":             EntityKindEnumFile,
		"filehash":         EntityKindEnumFileHash,
		"host":             EntityKindEnumHost,
		"ip":               EntityKindEnumIP,
		"iotdevice":        EntityKindEnumIoTDevice,
		"mailcluster":      EntityKindEnumMailCluster,
		"mailmessage":      EntityKindEnumMailMessage,
		"mailbox":          EntityKindEnumMailbox,
		"malware":          EntityKindEnumMalware,
		"process":          EntityKindEnumProcess,
		"registrykey":      EntityKindEnumRegistryKey,
		"registryvalue":    EntityKindEnumRegistryValue,
		"securityalert":    EntityKindEnumSecurityAlert,
		"securitygroup":    EntityKindEnumSecurityGroup,
		"submissionmail":   EntityKindEnumSubmissionMail,
		"url":              EntityKindEnumUrl,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EntityKindEnum(input)
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
