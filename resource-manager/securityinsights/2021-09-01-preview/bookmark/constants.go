package bookmark

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
		string(EntityKindProcess),
		string(EntityKindRegistryKey),
		string(EntityKindRegistryValue),
		string(EntityKindSecurityAlert),
		string(EntityKindSecurityGroup),
		string(EntityKindSubmissionMail),
		string(EntityKindUrl),
	}
}
