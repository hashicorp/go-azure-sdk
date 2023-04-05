package alertrules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertDetail string

const (
	AlertDetailDisplayName AlertDetail = "DisplayName"
	AlertDetailSeverity    AlertDetail = "Severity"
)

func PossibleValuesForAlertDetail() []string {
	return []string{
		string(AlertDetailDisplayName),
		string(AlertDetailSeverity),
	}
}

type AlertRuleKind string

const (
	AlertRuleKindFusion                            AlertRuleKind = "Fusion"
	AlertRuleKindMicrosoftSecurityIncidentCreation AlertRuleKind = "MicrosoftSecurityIncidentCreation"
	AlertRuleKindScheduled                         AlertRuleKind = "Scheduled"
)

func PossibleValuesForAlertRuleKind() []string {
	return []string{
		string(AlertRuleKindFusion),
		string(AlertRuleKindMicrosoftSecurityIncidentCreation),
		string(AlertRuleKindScheduled),
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

type MatchingMethod string

const (
	MatchingMethodAllEntities MatchingMethod = "AllEntities"
	MatchingMethodAnyAlert    MatchingMethod = "AnyAlert"
	MatchingMethodSelected    MatchingMethod = "Selected"
)

func PossibleValuesForMatchingMethod() []string {
	return []string{
		string(MatchingMethodAllEntities),
		string(MatchingMethodAnyAlert),
		string(MatchingMethodSelected),
	}
}

type MicrosoftSecurityProductName string

const (
	MicrosoftSecurityProductNameAzureActiveDirectoryIdentityProtection MicrosoftSecurityProductName = "Azure Active Directory Identity Protection"
	MicrosoftSecurityProductNameAzureAdvancedThreatProtection          MicrosoftSecurityProductName = "Azure Advanced Threat Protection"
	MicrosoftSecurityProductNameAzureSecurityCenter                    MicrosoftSecurityProductName = "Azure Security Center"
	MicrosoftSecurityProductNameAzureSecurityCenterForIoT              MicrosoftSecurityProductName = "Azure Security Center for IoT"
	MicrosoftSecurityProductNameMicrosoftCloudAppSecurity              MicrosoftSecurityProductName = "Microsoft Cloud App Security"
)

func PossibleValuesForMicrosoftSecurityProductName() []string {
	return []string{
		string(MicrosoftSecurityProductNameAzureActiveDirectoryIdentityProtection),
		string(MicrosoftSecurityProductNameAzureAdvancedThreatProtection),
		string(MicrosoftSecurityProductNameAzureSecurityCenter),
		string(MicrosoftSecurityProductNameAzureSecurityCenterForIoT),
		string(MicrosoftSecurityProductNameMicrosoftCloudAppSecurity),
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
