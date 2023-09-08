package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UrlAssessmentRequestCategory string

const (
	UrlAssessmentRequestCategorymalware   UrlAssessmentRequestCategory = "Malware"
	UrlAssessmentRequestCategoryphishing  UrlAssessmentRequestCategory = "Phishing"
	UrlAssessmentRequestCategoryspam      UrlAssessmentRequestCategory = "Spam"
	UrlAssessmentRequestCategoryundefined UrlAssessmentRequestCategory = "Undefined"
)

func PossibleValuesForUrlAssessmentRequestCategory() []string {
	return []string{
		string(UrlAssessmentRequestCategorymalware),
		string(UrlAssessmentRequestCategoryphishing),
		string(UrlAssessmentRequestCategoryspam),
		string(UrlAssessmentRequestCategoryundefined),
	}
}

func parseUrlAssessmentRequestCategory(input string) (*UrlAssessmentRequestCategory, error) {
	vals := map[string]UrlAssessmentRequestCategory{
		"malware":   UrlAssessmentRequestCategorymalware,
		"phishing":  UrlAssessmentRequestCategoryphishing,
		"spam":      UrlAssessmentRequestCategoryspam,
		"undefined": UrlAssessmentRequestCategoryundefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UrlAssessmentRequestCategory(input)
	return &out, nil
}
