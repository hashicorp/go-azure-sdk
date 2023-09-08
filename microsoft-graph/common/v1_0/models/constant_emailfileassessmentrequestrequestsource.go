package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailFileAssessmentRequestRequestSource string

const (
	EmailFileAssessmentRequestRequestSourceadministrator EmailFileAssessmentRequestRequestSource = "Administrator"
	EmailFileAssessmentRequestRequestSourceundefined     EmailFileAssessmentRequestRequestSource = "Undefined"
	EmailFileAssessmentRequestRequestSourceuser          EmailFileAssessmentRequestRequestSource = "User"
)

func PossibleValuesForEmailFileAssessmentRequestRequestSource() []string {
	return []string{
		string(EmailFileAssessmentRequestRequestSourceadministrator),
		string(EmailFileAssessmentRequestRequestSourceundefined),
		string(EmailFileAssessmentRequestRequestSourceuser),
	}
}

func parseEmailFileAssessmentRequestRequestSource(input string) (*EmailFileAssessmentRequestRequestSource, error) {
	vals := map[string]EmailFileAssessmentRequestRequestSource{
		"administrator": EmailFileAssessmentRequestRequestSourceadministrator,
		"undefined":     EmailFileAssessmentRequestRequestSourceundefined,
		"user":          EmailFileAssessmentRequestRequestSourceuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EmailFileAssessmentRequestRequestSource(input)
	return &out, nil
}
