package subassessments

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssessedResourceType string

const (
	AssessedResourceTypeContainerRegistryVulnerability AssessedResourceType = "ContainerRegistryVulnerability"
	AssessedResourceTypeServerVulnerability            AssessedResourceType = "ServerVulnerability"
	AssessedResourceTypeSqlServerVulnerability         AssessedResourceType = "SqlServerVulnerability"
)

func PossibleValuesForAssessedResourceType() []string {
	return []string{
		string(AssessedResourceTypeContainerRegistryVulnerability),
		string(AssessedResourceTypeServerVulnerability),
		string(AssessedResourceTypeSqlServerVulnerability),
	}
}

func parseAssessedResourceType(input string) (*AssessedResourceType, error) {
	vals := map[string]AssessedResourceType{
		"containerregistryvulnerability": AssessedResourceTypeContainerRegistryVulnerability,
		"servervulnerability":            AssessedResourceTypeServerVulnerability,
		"sqlservervulnerability":         AssessedResourceTypeSqlServerVulnerability,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssessedResourceType(input)
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

type Source string

const (
	SourceAzure        Source = "Azure"
	SourceOnPremise    Source = "OnPremise"
	SourceOnPremiseSql Source = "OnPremiseSql"
)

func PossibleValuesForSource() []string {
	return []string{
		string(SourceAzure),
		string(SourceOnPremise),
		string(SourceOnPremiseSql),
	}
}

func parseSource(input string) (*Source, error) {
	vals := map[string]Source{
		"azure":        SourceAzure,
		"onpremise":    SourceOnPremise,
		"onpremisesql": SourceOnPremiseSql,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Source(input)
	return &out, nil
}

type SubAssessmentStatusCode string

const (
	SubAssessmentStatusCodeHealthy       SubAssessmentStatusCode = "Healthy"
	SubAssessmentStatusCodeNotApplicable SubAssessmentStatusCode = "NotApplicable"
	SubAssessmentStatusCodeUnhealthy     SubAssessmentStatusCode = "Unhealthy"
)

func PossibleValuesForSubAssessmentStatusCode() []string {
	return []string{
		string(SubAssessmentStatusCodeHealthy),
		string(SubAssessmentStatusCodeNotApplicable),
		string(SubAssessmentStatusCodeUnhealthy),
	}
}

func parseSubAssessmentStatusCode(input string) (*SubAssessmentStatusCode, error) {
	vals := map[string]SubAssessmentStatusCode{
		"healthy":       SubAssessmentStatusCodeHealthy,
		"notapplicable": SubAssessmentStatusCodeNotApplicable,
		"unhealthy":     SubAssessmentStatusCodeUnhealthy,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubAssessmentStatusCode(input)
	return &out, nil
}
