package assessmentsmetadata

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssessmentType string

const (
	AssessmentTypeBuiltIn         AssessmentType = "BuiltIn"
	AssessmentTypeCustomPolicy    AssessmentType = "CustomPolicy"
	AssessmentTypeCustomerManaged AssessmentType = "CustomerManaged"
)

func PossibleValuesForAssessmentType() []string {
	return []string{
		string(AssessmentTypeBuiltIn),
		string(AssessmentTypeCustomPolicy),
		string(AssessmentTypeCustomerManaged),
	}
}

func (s *AssessmentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAssessmentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAssessmentType(input string) (*AssessmentType, error) {
	vals := map[string]AssessmentType{
		"builtin":         AssessmentTypeBuiltIn,
		"custompolicy":    AssessmentTypeCustomPolicy,
		"customermanaged": AssessmentTypeCustomerManaged,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssessmentType(input)
	return &out, nil
}

type Categories string

const (
	CategoriesCompute           Categories = "Compute"
	CategoriesData              Categories = "Data"
	CategoriesIdentityAndAccess Categories = "IdentityAndAccess"
	CategoriesIoT               Categories = "IoT"
	CategoriesNetworking        Categories = "Networking"
)

func PossibleValuesForCategories() []string {
	return []string{
		string(CategoriesCompute),
		string(CategoriesData),
		string(CategoriesIdentityAndAccess),
		string(CategoriesIoT),
		string(CategoriesNetworking),
	}
}

func (s *Categories) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCategories(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCategories(input string) (*Categories, error) {
	vals := map[string]Categories{
		"compute":           CategoriesCompute,
		"data":              CategoriesData,
		"identityandaccess": CategoriesIdentityAndAccess,
		"iot":               CategoriesIoT,
		"networking":        CategoriesNetworking,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Categories(input)
	return &out, nil
}

type ImplementationEffort string

const (
	ImplementationEffortHigh     ImplementationEffort = "High"
	ImplementationEffortLow      ImplementationEffort = "Low"
	ImplementationEffortModerate ImplementationEffort = "Moderate"
)

func PossibleValuesForImplementationEffort() []string {
	return []string{
		string(ImplementationEffortHigh),
		string(ImplementationEffortLow),
		string(ImplementationEffortModerate),
	}
}

func (s *ImplementationEffort) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseImplementationEffort(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseImplementationEffort(input string) (*ImplementationEffort, error) {
	vals := map[string]ImplementationEffort{
		"high":     ImplementationEffortHigh,
		"low":      ImplementationEffortLow,
		"moderate": ImplementationEffortModerate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImplementationEffort(input)
	return &out, nil
}

type Severity string

const (
	SeverityHigh   Severity = "High"
	SeverityLow    Severity = "Low"
	SeverityMedium Severity = "Medium"
)

func PossibleValuesForSeverity() []string {
	return []string{
		string(SeverityHigh),
		string(SeverityLow),
		string(SeverityMedium),
	}
}

func (s *Severity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSeverity(input string) (*Severity, error) {
	vals := map[string]Severity{
		"high":   SeverityHigh,
		"low":    SeverityLow,
		"medium": SeverityMedium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Severity(input)
	return &out, nil
}

type Threats string

const (
	ThreatsAccountBreach        Threats = "accountBreach"
	ThreatsDataExfiltration     Threats = "dataExfiltration"
	ThreatsDataSpillage         Threats = "dataSpillage"
	ThreatsDenialOfService      Threats = "denialOfService"
	ThreatsElevationOfPrivilege Threats = "elevationOfPrivilege"
	ThreatsMaliciousInsider     Threats = "maliciousInsider"
	ThreatsMissingCoverage      Threats = "missingCoverage"
	ThreatsThreatResistance     Threats = "threatResistance"
)

func PossibleValuesForThreats() []string {
	return []string{
		string(ThreatsAccountBreach),
		string(ThreatsDataExfiltration),
		string(ThreatsDataSpillage),
		string(ThreatsDenialOfService),
		string(ThreatsElevationOfPrivilege),
		string(ThreatsMaliciousInsider),
		string(ThreatsMissingCoverage),
		string(ThreatsThreatResistance),
	}
}

func (s *Threats) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseThreats(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseThreats(input string) (*Threats, error) {
	vals := map[string]Threats{
		"accountbreach":        ThreatsAccountBreach,
		"dataexfiltration":     ThreatsDataExfiltration,
		"dataspillage":         ThreatsDataSpillage,
		"denialofservice":      ThreatsDenialOfService,
		"elevationofprivilege": ThreatsElevationOfPrivilege,
		"maliciousinsider":     ThreatsMaliciousInsider,
		"missingcoverage":      ThreatsMissingCoverage,
		"threatresistance":     ThreatsThreatResistance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Threats(input)
	return &out, nil
}

type UserImpact string

const (
	UserImpactHigh     UserImpact = "High"
	UserImpactLow      UserImpact = "Low"
	UserImpactModerate UserImpact = "Moderate"
)

func PossibleValuesForUserImpact() []string {
	return []string{
		string(UserImpactHigh),
		string(UserImpactLow),
		string(UserImpactModerate),
	}
}

func (s *UserImpact) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserImpact(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserImpact(input string) (*UserImpact, error) {
	vals := map[string]UserImpact{
		"high":     UserImpactHigh,
		"low":      UserImpactLow,
		"moderate": UserImpactModerate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserImpact(input)
	return &out, nil
}
