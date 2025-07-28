package alertruletemplates

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertProperty string

const (
	AlertPropertyAlertLink            AlertProperty = "AlertLink"
	AlertPropertyConfidenceLevel      AlertProperty = "ConfidenceLevel"
	AlertPropertyConfidenceScore      AlertProperty = "ConfidenceScore"
	AlertPropertyExtendedLinks        AlertProperty = "ExtendedLinks"
	AlertPropertyProductComponentName AlertProperty = "ProductComponentName"
	AlertPropertyProductName          AlertProperty = "ProductName"
	AlertPropertyProviderName         AlertProperty = "ProviderName"
	AlertPropertyRemediationSteps     AlertProperty = "RemediationSteps"
	AlertPropertyTechniques           AlertProperty = "Techniques"
)

func PossibleValuesForAlertProperty() []string {
	return []string{
		string(AlertPropertyAlertLink),
		string(AlertPropertyConfidenceLevel),
		string(AlertPropertyConfidenceScore),
		string(AlertPropertyExtendedLinks),
		string(AlertPropertyProductComponentName),
		string(AlertPropertyProductName),
		string(AlertPropertyProviderName),
		string(AlertPropertyRemediationSteps),
		string(AlertPropertyTechniques),
	}
}

func (s *AlertProperty) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAlertProperty(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAlertProperty(input string) (*AlertProperty, error) {
	vals := map[string]AlertProperty{
		"alertlink":            AlertPropertyAlertLink,
		"confidencelevel":      AlertPropertyConfidenceLevel,
		"confidencescore":      AlertPropertyConfidenceScore,
		"extendedlinks":        AlertPropertyExtendedLinks,
		"productcomponentname": AlertPropertyProductComponentName,
		"productname":          AlertPropertyProductName,
		"providername":         AlertPropertyProviderName,
		"remediationsteps":     AlertPropertyRemediationSteps,
		"techniques":           AlertPropertyTechniques,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertProperty(input)
	return &out, nil
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

func (s *AlertRuleKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAlertRuleKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAlertRuleKind(input string) (*AlertRuleKind, error) {
	vals := map[string]AlertRuleKind{
		"fusion":                            AlertRuleKindFusion,
		"microsoftsecurityincidentcreation": AlertRuleKindMicrosoftSecurityIncidentCreation,
		"scheduled":                         AlertRuleKindScheduled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertRuleKind(input)
	return &out, nil
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

func (s *EntityMappingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEntityMappingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEntityMappingType(input string) (*EntityMappingType, error) {
	vals := map[string]EntityMappingType{
		"account":          EntityMappingTypeAccount,
		"azureresource":    EntityMappingTypeAzureResource,
		"cloudapplication": EntityMappingTypeCloudApplication,
		"dns":              EntityMappingTypeDNS,
		"file":             EntityMappingTypeFile,
		"filehash":         EntityMappingTypeFileHash,
		"host":             EntityMappingTypeHost,
		"ip":               EntityMappingTypeIP,
		"mailcluster":      EntityMappingTypeMailCluster,
		"mailmessage":      EntityMappingTypeMailMessage,
		"mailbox":          EntityMappingTypeMailbox,
		"malware":          EntityMappingTypeMalware,
		"process":          EntityMappingTypeProcess,
		"registrykey":      EntityMappingTypeRegistryKey,
		"registryvalue":    EntityMappingTypeRegistryValue,
		"securitygroup":    EntityMappingTypeSecurityGroup,
		"submissionmail":   EntityMappingTypeSubmissionMail,
		"url":              EntityMappingTypeURL,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EntityMappingType(input)
	return &out, nil
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

func (s *EventGroupingAggregationKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventGroupingAggregationKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventGroupingAggregationKind(input string) (*EventGroupingAggregationKind, error) {
	vals := map[string]EventGroupingAggregationKind{
		"alertperresult": EventGroupingAggregationKindAlertPerResult,
		"singlealert":    EventGroupingAggregationKindSingleAlert,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventGroupingAggregationKind(input)
	return &out, nil
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

func (s *MicrosoftSecurityProductName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMicrosoftSecurityProductName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMicrosoftSecurityProductName(input string) (*MicrosoftSecurityProductName, error) {
	vals := map[string]MicrosoftSecurityProductName{
		"azure active directory identity protection": MicrosoftSecurityProductNameAzureActiveDirectoryIdentityProtection,
		"azure advanced threat protection":           MicrosoftSecurityProductNameAzureAdvancedThreatProtection,
		"azure security center":                      MicrosoftSecurityProductNameAzureSecurityCenter,
		"azure security center for iot":              MicrosoftSecurityProductNameAzureSecurityCenterForIoT,
		"microsoft cloud app security":               MicrosoftSecurityProductNameMicrosoftCloudAppSecurity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftSecurityProductName(input)
	return &out, nil
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

func (s *TemplateStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTemplateStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTemplateStatus(input string) (*TemplateStatus, error) {
	vals := map[string]TemplateStatus{
		"available":    TemplateStatusAvailable,
		"installed":    TemplateStatusInstalled,
		"notavailable": TemplateStatusNotAvailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TemplateStatus(input)
	return &out, nil
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

func (s *TriggerOperator) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTriggerOperator(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTriggerOperator(input string) (*TriggerOperator, error) {
	vals := map[string]TriggerOperator{
		"equal":       TriggerOperatorEqual,
		"greaterthan": TriggerOperatorGreaterThan,
		"lessthan":    TriggerOperatorLessThan,
		"notequal":    TriggerOperatorNotEqual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TriggerOperator(input)
	return &out, nil
}
