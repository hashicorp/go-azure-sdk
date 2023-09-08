package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailAssessmentRequestCategory string

const (
	MailAssessmentRequestCategorymalware   MailAssessmentRequestCategory = "Malware"
	MailAssessmentRequestCategoryphishing  MailAssessmentRequestCategory = "Phishing"
	MailAssessmentRequestCategoryspam      MailAssessmentRequestCategory = "Spam"
	MailAssessmentRequestCategoryundefined MailAssessmentRequestCategory = "Undefined"
)

func PossibleValuesForMailAssessmentRequestCategory() []string {
	return []string{
		string(MailAssessmentRequestCategorymalware),
		string(MailAssessmentRequestCategoryphishing),
		string(MailAssessmentRequestCategoryspam),
		string(MailAssessmentRequestCategoryundefined),
	}
}

func parseMailAssessmentRequestCategory(input string) (*MailAssessmentRequestCategory, error) {
	vals := map[string]MailAssessmentRequestCategory{
		"malware":   MailAssessmentRequestCategorymalware,
		"phishing":  MailAssessmentRequestCategoryphishing,
		"spam":      MailAssessmentRequestCategoryspam,
		"undefined": MailAssessmentRequestCategoryundefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MailAssessmentRequestCategory(input)
	return &out, nil
}
