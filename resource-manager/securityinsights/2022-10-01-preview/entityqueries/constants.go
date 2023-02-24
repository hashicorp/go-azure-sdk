package entityqueries

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomEntityQueryKind string

const (
	CustomEntityQueryKindActivity CustomEntityQueryKind = "Activity"
)

func PossibleValuesForCustomEntityQueryKind() []string {
	return []string{
		string(CustomEntityQueryKindActivity),
	}
}

func parseCustomEntityQueryKind(input string) (*CustomEntityQueryKind, error) {
	vals := map[string]CustomEntityQueryKind{
		"activity": CustomEntityQueryKindActivity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomEntityQueryKind(input)
	return &out, nil
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

func parseEntityQueryKind(input string) (*EntityQueryKind, error) {
	vals := map[string]EntityQueryKind{
		"activity":  EntityQueryKindActivity,
		"expansion": EntityQueryKindExpansion,
		"insight":   EntityQueryKindInsight,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EntityQueryKind(input)
	return &out, nil
}

type EntityQueryTemplateKind string

const (
	EntityQueryTemplateKindActivity EntityQueryTemplateKind = "Activity"
)

func PossibleValuesForEntityQueryTemplateKind() []string {
	return []string{
		string(EntityQueryTemplateKindActivity),
	}
}

func parseEntityQueryTemplateKind(input string) (*EntityQueryTemplateKind, error) {
	vals := map[string]EntityQueryTemplateKind{
		"activity": EntityQueryTemplateKindActivity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EntityQueryTemplateKind(input)
	return &out, nil
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

func parseEntityType(input string) (*EntityType, error) {
	vals := map[string]EntityType{
		"account":          EntityTypeAccount,
		"azureresource":    EntityTypeAzureResource,
		"cloudapplication": EntityTypeCloudApplication,
		"dns":              EntityTypeDNS,
		"file":             EntityTypeFile,
		"filehash":         EntityTypeFileHash,
		"host":             EntityTypeHost,
		"huntingbookmark":  EntityTypeHuntingBookmark,
		"ip":               EntityTypeIP,
		"iotdevice":        EntityTypeIoTDevice,
		"mailcluster":      EntityTypeMailCluster,
		"mailmessage":      EntityTypeMailMessage,
		"mailbox":          EntityTypeMailbox,
		"malware":          EntityTypeMalware,
		"nic":              EntityTypeNic,
		"process":          EntityTypeProcess,
		"registrykey":      EntityTypeRegistryKey,
		"registryvalue":    EntityTypeRegistryValue,
		"securityalert":    EntityTypeSecurityAlert,
		"securitygroup":    EntityTypeSecurityGroup,
		"submissionmail":   EntityTypeSubmissionMail,
		"url":              EntityTypeURL,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EntityType(input)
	return &out, nil
}

type Kind string

const (
	KindActivity  Kind = "Activity"
	KindExpansion Kind = "Expansion"
)

func PossibleValuesForKind() []string {
	return []string{
		string(KindActivity),
		string(KindExpansion),
	}
}

func parseKind(input string) (*Kind, error) {
	vals := map[string]Kind{
		"activity":  KindActivity,
		"expansion": KindExpansion,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Kind(input)
	return &out, nil
}
