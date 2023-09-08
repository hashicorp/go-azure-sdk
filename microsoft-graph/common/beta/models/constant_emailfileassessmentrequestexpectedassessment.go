package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailFileAssessmentRequestExpectedAssessment string

const (
	EmailFileAssessmentRequestExpectedAssessmentblock   EmailFileAssessmentRequestExpectedAssessment = "Block"
	EmailFileAssessmentRequestExpectedAssessmentunblock EmailFileAssessmentRequestExpectedAssessment = "Unblock"
)

func PossibleValuesForEmailFileAssessmentRequestExpectedAssessment() []string {
	return []string{
		string(EmailFileAssessmentRequestExpectedAssessmentblock),
		string(EmailFileAssessmentRequestExpectedAssessmentunblock),
	}
}

func parseEmailFileAssessmentRequestExpectedAssessment(input string) (*EmailFileAssessmentRequestExpectedAssessment, error) {
	vals := map[string]EmailFileAssessmentRequestExpectedAssessment{
		"block":   EmailFileAssessmentRequestExpectedAssessmentblock,
		"unblock": EmailFileAssessmentRequestExpectedAssessmentunblock,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EmailFileAssessmentRequestExpectedAssessment(input)
	return &out, nil
}
