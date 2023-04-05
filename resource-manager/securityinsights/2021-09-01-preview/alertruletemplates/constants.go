package alertruletemplates

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

type EntityMappingType string

const (
	EntityMappingTypeAccount          EntityMappingType = "Account"
	EntityMappingTypeAzureResource    EntityMappingType = "AzureResource"
	EntityMappingTypeCloudApplication EntityMappingType = "CloudApplication"
	EntityMappingTypeDNS              EntityMappingType = "DNS"
	EntityMappingTypeFile             EntityMappingType = "File"
	EntityMappingTypeFileHash         EntityMappingType = "FileHash"
	EntityMappingTypeHost             EntityMappingType = "Host"
	EntityMappingTypeIP               EntityMappingType = "IP"
	EntityMappingTypeMailCluster      EntityMappingType = "MailCluster"
	EntityMappingTypeMailMessage      EntityMappingType = "MailMessage"
	EntityMappingTypeMailbox          EntityMappingType = "Mailbox"
	EntityMappingTypeMalware          EntityMappingType = "Malware"
	EntityMappingTypeProcess          EntityMappingType = "Process"
	EntityMappingTypeRegistryKey      EntityMappingType = "RegistryKey"
	EntityMappingTypeRegistryValue    EntityMappingType = "RegistryValue"
	EntityMappingTypeSecurityGroup    EntityMappingType = "SecurityGroup"
	EntityMappingTypeSubmissionMail   EntityMappingType = "SubmissionMail"
	EntityMappingTypeURL              EntityMappingType = "URL"
)

func PossibleValuesForEntityMappingType() []string {
	return []string{
		string(EntityMappingTypeAccount),
		string(EntityMappingTypeAzureResource),
		string(EntityMappingTypeCloudApplication),
		string(EntityMappingTypeDNS),
		string(EntityMappingTypeFile),
		string(EntityMappingTypeFileHash),
		string(EntityMappingTypeHost),
		string(EntityMappingTypeIP),
		string(EntityMappingTypeMailCluster),
		string(EntityMappingTypeMailMessage),
		string(EntityMappingTypeMailbox),
		string(EntityMappingTypeMalware),
		string(EntityMappingTypeProcess),
		string(EntityMappingTypeRegistryKey),
		string(EntityMappingTypeRegistryValue),
		string(EntityMappingTypeSecurityGroup),
		string(EntityMappingTypeSubmissionMail),
		string(EntityMappingTypeURL),
	}
}

type EventGroupingAggregationKind string

const (
	EventGroupingAggregationKindAlertPerResult EventGroupingAggregationKind = "AlertPerResult"
	EventGroupingAggregationKindSingleAlert    EventGroupingAggregationKind = "SingleAlert"
)

func PossibleValuesForEventGroupingAggregationKind() []string {
	return []string{
		string(EventGroupingAggregationKindAlertPerResult),
		string(EventGroupingAggregationKindSingleAlert),
	}
}

type TemplateStatus string

const (
	TemplateStatusAvailable    TemplateStatus = "Available"
	TemplateStatusInstalled    TemplateStatus = "Installed"
	TemplateStatusNotAvailable TemplateStatus = "NotAvailable"
)

func PossibleValuesForTemplateStatus() []string {
	return []string{
		string(TemplateStatusAvailable),
		string(TemplateStatusInstalled),
		string(TemplateStatusNotAvailable),
	}
}

type TriggerOperator string

const (
	TriggerOperatorEqual       TriggerOperator = "Equal"
	TriggerOperatorGreaterThan TriggerOperator = "GreaterThan"
	TriggerOperatorLessThan    TriggerOperator = "LessThan"
	TriggerOperatorNotEqual    TriggerOperator = "NotEqual"
)

func PossibleValuesForTriggerOperator() []string {
	return []string{
		string(TriggerOperatorEqual),
		string(TriggerOperatorGreaterThan),
		string(TriggerOperatorLessThan),
		string(TriggerOperatorNotEqual),
	}
}
