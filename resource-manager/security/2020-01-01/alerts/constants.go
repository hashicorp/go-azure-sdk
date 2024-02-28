package alerts

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
	AlertStatusActive    AlertStatus = "Active"
	AlertStatusDismissed AlertStatus = "Dismissed"
	AlertStatusResolved  AlertStatus = "Resolved"
)

func PossibleValuesForAlertStatus() []string {
	return []string{
		string(AlertStatusActive),
		string(AlertStatusDismissed),
		string(AlertStatusResolved),
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
		"active":    AlertStatusActive,
		"dismissed": AlertStatusDismissed,
		"resolved":  AlertStatusResolved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertStatus(input)
	return &out, nil
}

type Intent string

const (
	IntentCollection          Intent = "Collection"
	IntentCommandAndControl   Intent = "CommandAndControl"
	IntentCredentialAccess    Intent = "CredentialAccess"
	IntentDefenseEvasion      Intent = "DefenseEvasion"
	IntentDiscovery           Intent = "Discovery"
	IntentExecution           Intent = "Execution"
	IntentExfiltration        Intent = "Exfiltration"
	IntentExploitation        Intent = "Exploitation"
	IntentImpact              Intent = "Impact"
	IntentInitialAccess       Intent = "InitialAccess"
	IntentLateralMovement     Intent = "LateralMovement"
	IntentPersistence         Intent = "Persistence"
	IntentPreAttack           Intent = "PreAttack"
	IntentPrivilegeEscalation Intent = "PrivilegeEscalation"
	IntentProbing             Intent = "Probing"
	IntentUnknown             Intent = "Unknown"
)

func PossibleValuesForIntent() []string {
	return []string{
		string(IntentCollection),
		string(IntentCommandAndControl),
		string(IntentCredentialAccess),
		string(IntentDefenseEvasion),
		string(IntentDiscovery),
		string(IntentExecution),
		string(IntentExfiltration),
		string(IntentExploitation),
		string(IntentImpact),
		string(IntentInitialAccess),
		string(IntentLateralMovement),
		string(IntentPersistence),
		string(IntentPreAttack),
		string(IntentPrivilegeEscalation),
		string(IntentProbing),
		string(IntentUnknown),
	}
}

func (s *Intent) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIntent(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIntent(input string) (*Intent, error) {
	vals := map[string]Intent{
		"collection":          IntentCollection,
		"commandandcontrol":   IntentCommandAndControl,
		"credentialaccess":    IntentCredentialAccess,
		"defenseevasion":      IntentDefenseEvasion,
		"discovery":           IntentDiscovery,
		"execution":           IntentExecution,
		"exfiltration":        IntentExfiltration,
		"exploitation":        IntentExploitation,
		"impact":              IntentImpact,
		"initialaccess":       IntentInitialAccess,
		"lateralmovement":     IntentLateralMovement,
		"persistence":         IntentPersistence,
		"preattack":           IntentPreAttack,
		"privilegeescalation": IntentPrivilegeEscalation,
		"probing":             IntentProbing,
		"unknown":             IntentUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Intent(input)
	return &out, nil
}

type ResourceIdentifierType string

const (
	ResourceIdentifierTypeAzureResource ResourceIdentifierType = "AzureResource"
	ResourceIdentifierTypeLogAnalytics  ResourceIdentifierType = "LogAnalytics"
)

func PossibleValuesForResourceIdentifierType() []string {
	return []string{
		string(ResourceIdentifierTypeAzureResource),
		string(ResourceIdentifierTypeLogAnalytics),
	}
}

func (s *ResourceIdentifierType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResourceIdentifierType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResourceIdentifierType(input string) (*ResourceIdentifierType, error) {
	vals := map[string]ResourceIdentifierType{
		"azureresource": ResourceIdentifierTypeAzureResource,
		"loganalytics":  ResourceIdentifierTypeLogAnalytics,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceIdentifierType(input)
	return &out, nil
}
