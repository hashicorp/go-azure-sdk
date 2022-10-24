package bookmark

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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
		"ip":               EntityKindIP,
		"iotdevice":        EntityKindIoTDevice,
		"mailcluster":      EntityKindMailCluster,
		"mailmessage":      EntityKindMailMessage,
		"mailbox":          EntityKindMailbox,
		"malware":          EntityKindMalware,
		"nic":              EntityKindNic,
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
