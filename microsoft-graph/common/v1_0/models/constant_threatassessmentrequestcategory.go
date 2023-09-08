package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreatAssessmentRequestCategory string

const (
	ThreatAssessmentRequestCategorymalware   ThreatAssessmentRequestCategory = "Malware"
	ThreatAssessmentRequestCategoryphishing  ThreatAssessmentRequestCategory = "Phishing"
	ThreatAssessmentRequestCategoryspam      ThreatAssessmentRequestCategory = "Spam"
	ThreatAssessmentRequestCategoryundefined ThreatAssessmentRequestCategory = "Undefined"
)

func PossibleValuesForThreatAssessmentRequestCategory() []string {
	return []string{
		string(ThreatAssessmentRequestCategorymalware),
		string(ThreatAssessmentRequestCategoryphishing),
		string(ThreatAssessmentRequestCategoryspam),
		string(ThreatAssessmentRequestCategoryundefined),
	}
}

func parseThreatAssessmentRequestCategory(input string) (*ThreatAssessmentRequestCategory, error) {
	vals := map[string]ThreatAssessmentRequestCategory{
		"malware":   ThreatAssessmentRequestCategorymalware,
		"phishing":  ThreatAssessmentRequestCategoryphishing,
		"spam":      ThreatAssessmentRequestCategoryspam,
		"undefined": ThreatAssessmentRequestCategoryundefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ThreatAssessmentRequestCategory(input)
	return &out, nil
}
