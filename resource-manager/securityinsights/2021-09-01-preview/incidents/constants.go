package incidents

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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
		"collection":          AttackTacticCollection,
		"commandandcontrol":   AttackTacticCommandAndControl,
		"credentialaccess":    AttackTacticCredentialAccess,
		"defenseevasion":      AttackTacticDefenseEvasion,
		"discovery":           AttackTacticDiscovery,
		"execution":           AttackTacticExecution,
		"exfiltration":        AttackTacticExfiltration,
		"impact":              AttackTacticImpact,
		"initialaccess":       AttackTacticInitialAccess,
		"lateralmovement":     AttackTacticLateralMovement,
		"persistence":         AttackTacticPersistence,
		"preattack":           AttackTacticPreAttack,
		"privilegeescalation": AttackTacticPrivilegeEscalation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttackTactic(input)
	return &out, nil
}

type IncidentClassification string

const (
	IncidentClassificationBenignPositive IncidentClassification = "BenignPositive"
	IncidentClassificationFalsePositive  IncidentClassification = "FalsePositive"
	IncidentClassificationTruePositive   IncidentClassification = "TruePositive"
	IncidentClassificationUndetermined   IncidentClassification = "Undetermined"
)

func PossibleValuesForIncidentClassification() []string {
	return []string{
		string(IncidentClassificationBenignPositive),
		string(IncidentClassificationFalsePositive),
		string(IncidentClassificationTruePositive),
		string(IncidentClassificationUndetermined),
	}
}

func (s *IncidentClassification) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIncidentClassification(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIncidentClassification(input string) (*IncidentClassification, error) {
	vals := map[string]IncidentClassification{
		"benignpositive": IncidentClassificationBenignPositive,
		"falsepositive":  IncidentClassificationFalsePositive,
		"truepositive":   IncidentClassificationTruePositive,
		"undetermined":   IncidentClassificationUndetermined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IncidentClassification(input)
	return &out, nil
}

type IncidentClassificationReason string

const (
	IncidentClassificationReasonInaccurateData        IncidentClassificationReason = "InaccurateData"
	IncidentClassificationReasonIncorrectAlertLogic   IncidentClassificationReason = "IncorrectAlertLogic"
	IncidentClassificationReasonSuspiciousActivity    IncidentClassificationReason = "SuspiciousActivity"
	IncidentClassificationReasonSuspiciousButExpected IncidentClassificationReason = "SuspiciousButExpected"
)

func PossibleValuesForIncidentClassificationReason() []string {
	return []string{
		string(IncidentClassificationReasonInaccurateData),
		string(IncidentClassificationReasonIncorrectAlertLogic),
		string(IncidentClassificationReasonSuspiciousActivity),
		string(IncidentClassificationReasonSuspiciousButExpected),
	}
}

func (s *IncidentClassificationReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIncidentClassificationReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIncidentClassificationReason(input string) (*IncidentClassificationReason, error) {
	vals := map[string]IncidentClassificationReason{
		"inaccuratedata":        IncidentClassificationReasonInaccurateData,
		"incorrectalertlogic":   IncidentClassificationReasonIncorrectAlertLogic,
		"suspiciousactivity":    IncidentClassificationReasonSuspiciousActivity,
		"suspiciousbutexpected": IncidentClassificationReasonSuspiciousButExpected,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IncidentClassificationReason(input)
	return &out, nil
}

type IncidentLabelType string

const (
	IncidentLabelTypeSystem IncidentLabelType = "System"
	IncidentLabelTypeUser   IncidentLabelType = "User"
)

func PossibleValuesForIncidentLabelType() []string {
	return []string{
		string(IncidentLabelTypeSystem),
		string(IncidentLabelTypeUser),
	}
}

func (s *IncidentLabelType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIncidentLabelType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIncidentLabelType(input string) (*IncidentLabelType, error) {
	vals := map[string]IncidentLabelType{
		"system": IncidentLabelTypeSystem,
		"user":   IncidentLabelTypeUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IncidentLabelType(input)
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

type IncidentStatus string

const (
	IncidentStatusActive IncidentStatus = "Active"
	IncidentStatusClosed IncidentStatus = "Closed"
	IncidentStatusNew    IncidentStatus = "New"
)

func PossibleValuesForIncidentStatus() []string {
	return []string{
		string(IncidentStatusActive),
		string(IncidentStatusClosed),
		string(IncidentStatusNew),
	}
}

func (s *IncidentStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIncidentStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIncidentStatus(input string) (*IncidentStatus, error) {
	vals := map[string]IncidentStatus{
		"active": IncidentStatusActive,
		"closed": IncidentStatusClosed,
		"new":    IncidentStatusNew,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IncidentStatus(input)
	return &out, nil
}

type OwnerType string

const (
	OwnerTypeGroup   OwnerType = "Group"
	OwnerTypeUnknown OwnerType = "Unknown"
	OwnerTypeUser    OwnerType = "User"
)

func PossibleValuesForOwnerType() []string {
	return []string{
		string(OwnerTypeGroup),
		string(OwnerTypeUnknown),
		string(OwnerTypeUser),
	}
}

func (s *OwnerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOwnerType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOwnerType(input string) (*OwnerType, error) {
	vals := map[string]OwnerType{
		"group":   OwnerTypeGroup,
		"unknown": OwnerTypeUnknown,
		"user":    OwnerTypeUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OwnerType(input)
	return &out, nil
}
