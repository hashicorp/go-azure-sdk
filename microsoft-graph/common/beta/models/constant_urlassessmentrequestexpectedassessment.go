package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UrlAssessmentRequestExpectedAssessment string

const (
	UrlAssessmentRequestExpectedAssessmentblock   UrlAssessmentRequestExpectedAssessment = "Block"
	UrlAssessmentRequestExpectedAssessmentunblock UrlAssessmentRequestExpectedAssessment = "Unblock"
)

func PossibleValuesForUrlAssessmentRequestExpectedAssessment() []string {
	return []string{
		string(UrlAssessmentRequestExpectedAssessmentblock),
		string(UrlAssessmentRequestExpectedAssessmentunblock),
	}
}

func parseUrlAssessmentRequestExpectedAssessment(input string) (*UrlAssessmentRequestExpectedAssessment, error) {
	vals := map[string]UrlAssessmentRequestExpectedAssessment{
		"block":   UrlAssessmentRequestExpectedAssessmentblock,
		"unblock": UrlAssessmentRequestExpectedAssessmentunblock,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UrlAssessmentRequestExpectedAssessment(input)
	return &out, nil
}
