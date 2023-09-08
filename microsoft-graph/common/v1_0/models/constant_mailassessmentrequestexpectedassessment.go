package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailAssessmentRequestExpectedAssessment string

const (
	MailAssessmentRequestExpectedAssessmentblock   MailAssessmentRequestExpectedAssessment = "Block"
	MailAssessmentRequestExpectedAssessmentunblock MailAssessmentRequestExpectedAssessment = "Unblock"
)

func PossibleValuesForMailAssessmentRequestExpectedAssessment() []string {
	return []string{
		string(MailAssessmentRequestExpectedAssessmentblock),
		string(MailAssessmentRequestExpectedAssessmentunblock),
	}
}

func parseMailAssessmentRequestExpectedAssessment(input string) (*MailAssessmentRequestExpectedAssessment, error) {
	vals := map[string]MailAssessmentRequestExpectedAssessment{
		"block":   MailAssessmentRequestExpectedAssessmentblock,
		"unblock": MailAssessmentRequestExpectedAssessmentunblock,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MailAssessmentRequestExpectedAssessment(input)
	return &out, nil
}
