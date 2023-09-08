package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailFileAssessmentRequestCategory string

const (
	EmailFileAssessmentRequestCategorymalware   EmailFileAssessmentRequestCategory = "Malware"
	EmailFileAssessmentRequestCategoryphishing  EmailFileAssessmentRequestCategory = "Phishing"
	EmailFileAssessmentRequestCategoryspam      EmailFileAssessmentRequestCategory = "Spam"
	EmailFileAssessmentRequestCategoryundefined EmailFileAssessmentRequestCategory = "Undefined"
)

func PossibleValuesForEmailFileAssessmentRequestCategory() []string {
	return []string{
		string(EmailFileAssessmentRequestCategorymalware),
		string(EmailFileAssessmentRequestCategoryphishing),
		string(EmailFileAssessmentRequestCategoryspam),
		string(EmailFileAssessmentRequestCategoryundefined),
	}
}

func parseEmailFileAssessmentRequestCategory(input string) (*EmailFileAssessmentRequestCategory, error) {
	vals := map[string]EmailFileAssessmentRequestCategory{
		"malware":   EmailFileAssessmentRequestCategorymalware,
		"phishing":  EmailFileAssessmentRequestCategoryphishing,
		"spam":      EmailFileAssessmentRequestCategoryspam,
		"undefined": EmailFileAssessmentRequestCategoryundefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EmailFileAssessmentRequestCategory(input)
	return &out, nil
}
