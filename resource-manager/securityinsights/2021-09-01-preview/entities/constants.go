package entities

import (
	"encoding/json"
	"fmt"
	"strings"
)

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

func (s *AlertSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAlertSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
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

type EntityItemQueryKind string

const (
	EntityItemQueryKindInsight EntityItemQueryKind = "Insight"
)

func PossibleValuesForEntityItemQueryKind() []string {
	return []string{
		string(EntityItemQueryKindInsight),
	}
}

func (s *EntityItemQueryKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEntityItemQueryKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEntityItemQueryKind(input string) (*EntityItemQueryKind, error) {
	vals := map[string]EntityItemQueryKind{
		"insight": EntityItemQueryKindInsight,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EntityItemQueryKind(input)
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

func (s *EntityKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEntityKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
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

func (s *EntityQueryKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEntityQueryKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
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

type EntityTimelineKind string

const (
	EntityTimelineKindActivity      EntityTimelineKind = "Activity"
	EntityTimelineKindBookmark      EntityTimelineKind = "Bookmark"
	EntityTimelineKindSecurityAlert EntityTimelineKind = "SecurityAlert"
)

func PossibleValuesForEntityTimelineKind() []string {
	return []string{
		string(EntityTimelineKindActivity),
		string(EntityTimelineKindBookmark),
		string(EntityTimelineKindSecurityAlert),
	}
}

func (s *EntityTimelineKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEntityTimelineKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEntityTimelineKind(input string) (*EntityTimelineKind, error) {
	vals := map[string]EntityTimelineKind{
		"activity":      EntityTimelineKindActivity,
		"bookmark":      EntityTimelineKindBookmark,
		"securityalert": EntityTimelineKindSecurityAlert,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EntityTimelineKind(input)
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
		string(EntityTypeProcess),
		string(EntityTypeRegistryKey),
		string(EntityTypeRegistryValue),
		string(EntityTypeSecurityAlert),
		string(EntityTypeSecurityGroup),
		string(EntityTypeSubmissionMail),
		string(EntityTypeURL),
	}
}

func (s *EntityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEntityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
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
	KindInsight Kind = "Insight"
)

func PossibleValuesForKind() []string {
	return []string{
		string(KindInsight),
	}
}

func (s *Kind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKind(input string) (*Kind, error) {
	vals := map[string]Kind{
		"insight": KindInsight,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Kind(input)
	return &out, nil
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

func (s *OutputType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOutputType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOutputType(input string) (*OutputType, error) {
	vals := map[string]OutputType{
		"date":   OutputTypeDate,
		"entity": OutputTypeEntity,
		"number": OutputTypeNumber,
		"string": OutputTypeString,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OutputType(input)
	return &out, nil
}
