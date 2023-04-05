package entities

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

type EntityItemQueryKind string

const (
	EntityItemQueryKindInsight EntityItemQueryKind = "Insight"
)

func PossibleValuesForEntityItemQueryKind() []string {
	return []string{
		string(EntityItemQueryKindInsight),
	}
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
	EntityKindIP               EntityKind = "Ip"
	EntityKindIoTDevice        EntityKind = "IoTDevice"
	EntityKindMailCluster      EntityKind = "MailCluster"
	EntityKindMailMessage      EntityKind = "MailMessage"
	EntityKindMailbox          EntityKind = "Mailbox"
	EntityKindMalware          EntityKind = "Malware"
	EntityKindNic              EntityKind = "Nic"
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
		string(EntityKindIP),
		string(EntityKindIoTDevice),
		string(EntityKindMailCluster),
		string(EntityKindMailMessage),
		string(EntityKindMailbox),
		string(EntityKindMalware),
		string(EntityKindNic),
		string(EntityKindProcess),
		string(EntityKindRegistryKey),
		string(EntityKindRegistryValue),
		string(EntityKindSecurityAlert),
		string(EntityKindSecurityGroup),
		string(EntityKindSubmissionMail),
		string(EntityKindUrl),
	}
}

type EntityQueryKind string

const (
	EntityQueryKindActivity  EntityQueryKind = "Activity"
	EntityQueryKindExpansion EntityQueryKind = "Expansion"
	EntityQueryKindInsight   EntityQueryKind = "Insight"
)

func PossibleValuesForEntityQueryKind() []string {
	return []string{
		string(EntityQueryKindActivity),
		string(EntityQueryKindExpansion),
		string(EntityQueryKindInsight),
	}
}

type EntityTimelineKind string

const (
	EntityTimelineKindActivity      EntityTimelineKind = "Activity"
	EntityTimelineKindAnomaly       EntityTimelineKind = "Anomaly"
	EntityTimelineKindBookmark      EntityTimelineKind = "Bookmark"
	EntityTimelineKindSecurityAlert EntityTimelineKind = "SecurityAlert"
)

func PossibleValuesForEntityTimelineKind() []string {
	return []string{
		string(EntityTimelineKindActivity),
		string(EntityTimelineKindAnomaly),
		string(EntityTimelineKindBookmark),
		string(EntityTimelineKindSecurityAlert),
	}
}

type EntityType string

const (
	EntityTypeAccount          EntityType = "Account"
	EntityTypeAzureResource    EntityType = "AzureResource"
	EntityTypeCloudApplication EntityType = "CloudApplication"
	EntityTypeDNS              EntityType = "DNS"
	EntityTypeFile             EntityType = "File"
	EntityTypeFileHash         EntityType = "FileHash"
	EntityTypeHost             EntityType = "Host"
	EntityTypeHuntingBookmark  EntityType = "HuntingBookmark"
	EntityTypeIP               EntityType = "IP"
	EntityTypeIoTDevice        EntityType = "IoTDevice"
	EntityTypeMailCluster      EntityType = "MailCluster"
	EntityTypeMailMessage      EntityType = "MailMessage"
	EntityTypeMailbox          EntityType = "Mailbox"
	EntityTypeMalware          EntityType = "Malware"
	EntityTypeNic              EntityType = "Nic"
	EntityTypeProcess          EntityType = "Process"
	EntityTypeRegistryKey      EntityType = "RegistryKey"
	EntityTypeRegistryValue    EntityType = "RegistryValue"
	EntityTypeSecurityAlert    EntityType = "SecurityAlert"
	EntityTypeSecurityGroup    EntityType = "SecurityGroup"
	EntityTypeSubmissionMail   EntityType = "SubmissionMail"
	EntityTypeURL              EntityType = "URL"
)

func PossibleValuesForEntityType() []string {
	return []string{
		string(EntityTypeAccount),
		string(EntityTypeAzureResource),
		string(EntityTypeCloudApplication),
		string(EntityTypeDNS),
		string(EntityTypeFile),
		string(EntityTypeFileHash),
		string(EntityTypeHost),
		string(EntityTypeHuntingBookmark),
		string(EntityTypeIP),
		string(EntityTypeIoTDevice),
		string(EntityTypeMailCluster),
		string(EntityTypeMailMessage),
		string(EntityTypeMailbox),
		string(EntityTypeMalware),
		string(EntityTypeNic),
		string(EntityTypeProcess),
		string(EntityTypeRegistryKey),
		string(EntityTypeRegistryValue),
		string(EntityTypeSecurityAlert),
		string(EntityTypeSecurityGroup),
		string(EntityTypeSubmissionMail),
		string(EntityTypeURL),
	}
}

type GetInsightsError string

const (
	GetInsightsErrorInsight GetInsightsError = "Insight"
)

func PossibleValuesForGetInsightsError() []string {
	return []string{
		string(GetInsightsErrorInsight),
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

type OutputType string

const (
	OutputTypeDate   OutputType = "Date"
	OutputTypeEntity OutputType = "Entity"
	OutputTypeNumber OutputType = "Number"
	OutputTypeString OutputType = "String"
)

func PossibleValuesForOutputType() []string {
	return []string{
		string(OutputTypeDate),
		string(OutputTypeEntity),
		string(OutputTypeNumber),
		string(OutputTypeString),
	}
}
