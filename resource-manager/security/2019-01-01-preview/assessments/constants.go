package assessments

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssessmentStatusCode string

const (
	AssessmentStatusCodeHealthy       AssessmentStatusCode = "Healthy"
	AssessmentStatusCodeNotApplicable AssessmentStatusCode = "NotApplicable"
	AssessmentStatusCodeUnhealthy     AssessmentStatusCode = "Unhealthy"
)

func PossibleValuesForAssessmentStatusCode() []string {
	return []string{
		string(AssessmentStatusCodeHealthy),
		string(AssessmentStatusCodeNotApplicable),
		string(AssessmentStatusCodeUnhealthy),
	}
}

func parseAssessmentStatusCode(input string) (*AssessmentStatusCode, error) {
	vals := map[string]AssessmentStatusCode{
		"healthy":       AssessmentStatusCodeHealthy,
		"notapplicable": AssessmentStatusCodeNotApplicable,
		"unhealthy":     AssessmentStatusCodeUnhealthy,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssessmentStatusCode(input)
	return &out, nil
}

type ExpandEnum string

const (
	ExpandEnumLinks    ExpandEnum = "links"
	ExpandEnumMetadata ExpandEnum = "metadata"
)

func PossibleValuesForExpandEnum() []string {
	return []string{
		string(ExpandEnumLinks),
		string(ExpandEnumMetadata),
	}
}

func parseExpandEnum(input string) (*ExpandEnum, error) {
	vals := map[string]ExpandEnum{
		"links":    ExpandEnumLinks,
		"metadata": ExpandEnumMetadata,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExpandEnum(input)
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
