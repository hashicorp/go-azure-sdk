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

func (s *AlertStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAlertStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
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

type AntispamMailDirection string

const (
	AntispamMailDirectionInbound  AntispamMailDirection = "Inbound"
	AntispamMailDirectionIntraorg AntispamMailDirection = "Intraorg"
	AntispamMailDirectionOutbound AntispamMailDirection = "Outbound"
	AntispamMailDirectionUnknown  AntispamMailDirection = "Unknown"
)

func PossibleValuesForAntispamMailDirection() []string {
	return []string{
		string(AntispamMailDirectionInbound),
		string(AntispamMailDirectionIntraorg),
		string(AntispamMailDirectionOutbound),
		string(AntispamMailDirectionUnknown),
	}
}

func (s *AntispamMailDirection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAntispamMailDirection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAntispamMailDirection(input string) (*AntispamMailDirection, error) {
	vals := map[string]AntispamMailDirection{
		"inbound":  AntispamMailDirectionInbound,
		"intraorg": AntispamMailDirectionIntraorg,
		"outbound": AntispamMailDirectionOutbound,
		"unknown":  AntispamMailDirectionUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AntispamMailDirection(input)
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

func (s *AttackTactic) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttackTactic(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
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

func (s *ConfidenceLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConfidenceLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
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

func (s *ConfidenceScoreStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConfidenceScoreStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
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

type DeliveryAction string

const (
	DeliveryActionBlocked         DeliveryAction = "Blocked"
	DeliveryActionDelivered       DeliveryAction = "Delivered"
	DeliveryActionDeliveredAsSpam DeliveryAction = "DeliveredAsSpam"
	DeliveryActionReplaced        DeliveryAction = "Replaced"
	DeliveryActionUnknown         DeliveryAction = "Unknown"
)

func PossibleValuesForDeliveryAction() []string {
	return []string{
		string(DeliveryActionBlocked),
		string(DeliveryActionDelivered),
		string(DeliveryActionDeliveredAsSpam),
		string(DeliveryActionReplaced),
		string(DeliveryActionUnknown),
	}
}

func (s *DeliveryAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeliveryAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeliveryAction(input string) (*DeliveryAction, error) {
	vals := map[string]DeliveryAction{
		"blocked":         DeliveryActionBlocked,
		"delivered":       DeliveryActionDelivered,
		"deliveredasspam": DeliveryActionDeliveredAsSpam,
		"replaced":        DeliveryActionReplaced,
		"unknown":         DeliveryActionUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeliveryAction(input)
	return &out, nil
}

type DeliveryLocation string

const (
	DeliveryLocationDeletedFolder DeliveryLocation = "DeletedFolder"
	DeliveryLocationDropped       DeliveryLocation = "Dropped"
	DeliveryLocationExternal      DeliveryLocation = "External"
	DeliveryLocationFailed        DeliveryLocation = "Failed"
	DeliveryLocationForwarded     DeliveryLocation = "Forwarded"
	DeliveryLocationInbox         DeliveryLocation = "Inbox"
	DeliveryLocationJunkFolder    DeliveryLocation = "JunkFolder"
	DeliveryLocationQuarantine    DeliveryLocation = "Quarantine"
	DeliveryLocationUnknown       DeliveryLocation = "Unknown"
)

func PossibleValuesForDeliveryLocation() []string {
	return []string{
		string(DeliveryLocationDeletedFolder),
		string(DeliveryLocationDropped),
		string(DeliveryLocationExternal),
		string(DeliveryLocationFailed),
		string(DeliveryLocationForwarded),
		string(DeliveryLocationInbox),
		string(DeliveryLocationJunkFolder),
		string(DeliveryLocationQuarantine),
		string(DeliveryLocationUnknown),
	}
}

func (s *DeliveryLocation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeliveryLocation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeliveryLocation(input string) (*DeliveryLocation, error) {
	vals := map[string]DeliveryLocation{
		"deletedfolder": DeliveryLocationDeletedFolder,
		"dropped":       DeliveryLocationDropped,
		"external":      DeliveryLocationExternal,
		"failed":        DeliveryLocationFailed,
		"forwarded":     DeliveryLocationForwarded,
		"inbox":         DeliveryLocationInbox,
		"junkfolder":    DeliveryLocationJunkFolder,
		"quarantine":    DeliveryLocationQuarantine,
		"unknown":       DeliveryLocationUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeliveryLocation(input)
	return &out, nil
}

type DeviceImportance string

const (
	DeviceImportanceHigh    DeviceImportance = "High"
	DeviceImportanceLow     DeviceImportance = "Low"
	DeviceImportanceNormal  DeviceImportance = "Normal"
	DeviceImportanceUnknown DeviceImportance = "Unknown"
)

func PossibleValuesForDeviceImportance() []string {
	return []string{
		string(DeviceImportanceHigh),
		string(DeviceImportanceLow),
		string(DeviceImportanceNormal),
		string(DeviceImportanceUnknown),
	}
}

func (s *DeviceImportance) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceImportance(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceImportance(input string) (*DeviceImportance, error) {
	vals := map[string]DeviceImportance{
		"high":    DeviceImportanceHigh,
		"low":     DeviceImportanceLow,
		"normal":  DeviceImportanceNormal,
		"unknown": DeviceImportanceUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceImportance(input)
	return &out, nil
}

type ElevationToken string

const (
	ElevationTokenDefault ElevationToken = "Default"
	ElevationTokenFull    ElevationToken = "Full"
	ElevationTokenLimited ElevationToken = "Limited"
)

func PossibleValuesForElevationToken() []string {
	return []string{
		string(ElevationTokenDefault),
		string(ElevationTokenFull),
		string(ElevationTokenLimited),
	}
}

func (s *ElevationToken) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseElevationToken(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseElevationToken(input string) (*ElevationToken, error) {
	vals := map[string]ElevationToken{
		"default": ElevationTokenDefault,
		"full":    ElevationTokenFull,
		"limited": ElevationTokenLimited,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ElevationToken(input)
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
	EntityKindNic              EntityKind = "Nic"
	EntityKindProcess          EntityKind = "Process"
	EntityKindRegistryKey      EntityKind = "RegistryKey"
	EntityKindRegistryValue    EntityKind = "RegistryValue"
	EntityKindSecurityAlert    EntityKind = "SecurityAlert"
	EntityKindSecurityGroup    EntityKind = "SecurityGroup"
	EntityKindSubmissionMail   EntityKind = "SubmissionMail"
	EntityKindURL              EntityKind = "Url"
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
		string(EntityKindURL),
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
		"nic":              EntityKindNic,
		"process":          EntityKindProcess,
		"registrykey":      EntityKindRegistryKey,
		"registryvalue":    EntityKindRegistryValue,
		"securityalert":    EntityKindSecurityAlert,
		"securitygroup":    EntityKindSecurityGroup,
		"submissionmail":   EntityKindSubmissionMail,
		"url":              EntityKindURL,
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
		"anomaly":       EntityTimelineKindAnomaly,
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

type FileHashAlgorithm string

const (
	FileHashAlgorithmMDFive          FileHashAlgorithm = "MD5"
	FileHashAlgorithmSHAOne          FileHashAlgorithm = "SHA1"
	FileHashAlgorithmSHATwoFiveSix   FileHashAlgorithm = "SHA256"
	FileHashAlgorithmSHATwoFiveSixAC FileHashAlgorithm = "SHA256AC"
	FileHashAlgorithmUnknown         FileHashAlgorithm = "Unknown"
)

func PossibleValuesForFileHashAlgorithm() []string {
	return []string{
		string(FileHashAlgorithmMDFive),
		string(FileHashAlgorithmSHAOne),
		string(FileHashAlgorithmSHATwoFiveSix),
		string(FileHashAlgorithmSHATwoFiveSixAC),
		string(FileHashAlgorithmUnknown),
	}
}

func (s *FileHashAlgorithm) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFileHashAlgorithm(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFileHashAlgorithm(input string) (*FileHashAlgorithm, error) {
	vals := map[string]FileHashAlgorithm{
		"md5":      FileHashAlgorithmMDFive,
		"sha1":     FileHashAlgorithmSHAOne,
		"sha256":   FileHashAlgorithmSHATwoFiveSix,
		"sha256ac": FileHashAlgorithmSHATwoFiveSixAC,
		"unknown":  FileHashAlgorithmUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FileHashAlgorithm(input)
	return &out, nil
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

func (s *GetInsightsError) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGetInsightsError(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGetInsightsError(input string) (*GetInsightsError, error) {
	vals := map[string]GetInsightsError{
		"insight": GetInsightsErrorInsight,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GetInsightsError(input)
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

func (s *IncidentSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIncidentSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
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

func (s *KillChainIntent) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKillChainIntent(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
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

type OSFamily string

const (
	OSFamilyAndroid OSFamily = "Android"
	OSFamilyIOS     OSFamily = "IOS"
	OSFamilyLinux   OSFamily = "Linux"
	OSFamilyUnknown OSFamily = "Unknown"
	OSFamilyWindows OSFamily = "Windows"
)

func PossibleValuesForOSFamily() []string {
	return []string{
		string(OSFamilyAndroid),
		string(OSFamilyIOS),
		string(OSFamilyLinux),
		string(OSFamilyUnknown),
		string(OSFamilyWindows),
	}
}

func (s *OSFamily) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOSFamily(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOSFamily(input string) (*OSFamily, error) {
	vals := map[string]OSFamily{
		"android": OSFamilyAndroid,
		"ios":     OSFamilyIOS,
		"linux":   OSFamilyLinux,
		"unknown": OSFamilyUnknown,
		"windows": OSFamilyWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OSFamily(input)
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

type RegistryHive string

const (
	RegistryHiveHKEYA                        RegistryHive = "HKEY_A"
	RegistryHiveHKEYCLASSESROOT              RegistryHive = "HKEY_CLASSES_ROOT"
	RegistryHiveHKEYCURRENTCONFIG            RegistryHive = "HKEY_CURRENT_CONFIG"
	RegistryHiveHKEYCURRENTUSER              RegistryHive = "HKEY_CURRENT_USER"
	RegistryHiveHKEYCURRENTUSERLOCALSETTINGS RegistryHive = "HKEY_CURRENT_USER_LOCAL_SETTINGS"
	RegistryHiveHKEYLOCALMACHINE             RegistryHive = "HKEY_LOCAL_MACHINE"
	RegistryHiveHKEYPERFORMANCEDATA          RegistryHive = "HKEY_PERFORMANCE_DATA"
	RegistryHiveHKEYPERFORMANCENLSTEXT       RegistryHive = "HKEY_PERFORMANCE_NLSTEXT"
	RegistryHiveHKEYPERFORMANCETEXT          RegistryHive = "HKEY_PERFORMANCE_TEXT"
	RegistryHiveHKEYUSERS                    RegistryHive = "HKEY_USERS"
)

func PossibleValuesForRegistryHive() []string {
	return []string{
		string(RegistryHiveHKEYA),
		string(RegistryHiveHKEYCLASSESROOT),
		string(RegistryHiveHKEYCURRENTCONFIG),
		string(RegistryHiveHKEYCURRENTUSER),
		string(RegistryHiveHKEYCURRENTUSERLOCALSETTINGS),
		string(RegistryHiveHKEYLOCALMACHINE),
		string(RegistryHiveHKEYPERFORMANCEDATA),
		string(RegistryHiveHKEYPERFORMANCENLSTEXT),
		string(RegistryHiveHKEYPERFORMANCETEXT),
		string(RegistryHiveHKEYUSERS),
	}
}

func (s *RegistryHive) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRegistryHive(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRegistryHive(input string) (*RegistryHive, error) {
	vals := map[string]RegistryHive{
		"hkey_a":                           RegistryHiveHKEYA,
		"hkey_classes_root":                RegistryHiveHKEYCLASSESROOT,
		"hkey_current_config":              RegistryHiveHKEYCURRENTCONFIG,
		"hkey_current_user":                RegistryHiveHKEYCURRENTUSER,
		"hkey_current_user_local_settings": RegistryHiveHKEYCURRENTUSERLOCALSETTINGS,
		"hkey_local_machine":               RegistryHiveHKEYLOCALMACHINE,
		"hkey_performance_data":            RegistryHiveHKEYPERFORMANCEDATA,
		"hkey_performance_nlstext":         RegistryHiveHKEYPERFORMANCENLSTEXT,
		"hkey_performance_text":            RegistryHiveHKEYPERFORMANCETEXT,
		"hkey_users":                       RegistryHiveHKEYUSERS,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RegistryHive(input)
	return &out, nil
}

type RegistryValueKind string

const (
	RegistryValueKindBinary       RegistryValueKind = "Binary"
	RegistryValueKindDWord        RegistryValueKind = "DWord"
	RegistryValueKindExpandString RegistryValueKind = "ExpandString"
	RegistryValueKindMultiString  RegistryValueKind = "MultiString"
	RegistryValueKindNone         RegistryValueKind = "None"
	RegistryValueKindQWord        RegistryValueKind = "QWord"
	RegistryValueKindString       RegistryValueKind = "String"
	RegistryValueKindUnknown      RegistryValueKind = "Unknown"
)

func PossibleValuesForRegistryValueKind() []string {
	return []string{
		string(RegistryValueKindBinary),
		string(RegistryValueKindDWord),
		string(RegistryValueKindExpandString),
		string(RegistryValueKindMultiString),
		string(RegistryValueKindNone),
		string(RegistryValueKindQWord),
		string(RegistryValueKindString),
		string(RegistryValueKindUnknown),
	}
}

func (s *RegistryValueKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRegistryValueKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRegistryValueKind(input string) (*RegistryValueKind, error) {
	vals := map[string]RegistryValueKind{
		"binary":       RegistryValueKindBinary,
		"dword":        RegistryValueKindDWord,
		"expandstring": RegistryValueKindExpandString,
		"multistring":  RegistryValueKindMultiString,
		"none":         RegistryValueKindNone,
		"qword":        RegistryValueKindQWord,
		"string":       RegistryValueKindString,
		"unknown":      RegistryValueKindUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RegistryValueKind(input)
	return &out, nil
}
