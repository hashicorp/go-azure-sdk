package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailAssessmentRequestRequestSource string

const (
	MailAssessmentRequestRequestSourceadministrator MailAssessmentRequestRequestSource = "Administrator"
	MailAssessmentRequestRequestSourceundefined     MailAssessmentRequestRequestSource = "Undefined"
	MailAssessmentRequestRequestSourceuser          MailAssessmentRequestRequestSource = "User"
)

func PossibleValuesForMailAssessmentRequestRequestSource() []string {
	return []string{
		string(MailAssessmentRequestRequestSourceadministrator),
		string(MailAssessmentRequestRequestSourceundefined),
		string(MailAssessmentRequestRequestSourceuser),
	}
}

func parseMailAssessmentRequestRequestSource(input string) (*MailAssessmentRequestRequestSource, error) {
	vals := map[string]MailAssessmentRequestRequestSource{
		"administrator": MailAssessmentRequestRequestSourceadministrator,
		"undefined":     MailAssessmentRequestRequestSourceundefined,
		"user":          MailAssessmentRequestRequestSourceuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MailAssessmentRequestRequestSource(input)
	return &out, nil
}
