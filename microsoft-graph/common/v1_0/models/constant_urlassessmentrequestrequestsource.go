package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UrlAssessmentRequestRequestSource string

const (
	UrlAssessmentRequestRequestSourceadministrator UrlAssessmentRequestRequestSource = "Administrator"
	UrlAssessmentRequestRequestSourceundefined     UrlAssessmentRequestRequestSource = "Undefined"
	UrlAssessmentRequestRequestSourceuser          UrlAssessmentRequestRequestSource = "User"
)

func PossibleValuesForUrlAssessmentRequestRequestSource() []string {
	return []string{
		string(UrlAssessmentRequestRequestSourceadministrator),
		string(UrlAssessmentRequestRequestSourceundefined),
		string(UrlAssessmentRequestRequestSourceuser),
	}
}

func parseUrlAssessmentRequestRequestSource(input string) (*UrlAssessmentRequestRequestSource, error) {
	vals := map[string]UrlAssessmentRequestRequestSource{
		"administrator": UrlAssessmentRequestRequestSourceadministrator,
		"undefined":     UrlAssessmentRequestRequestSourceundefined,
		"user":          UrlAssessmentRequestRequestSourceuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UrlAssessmentRequestRequestSource(input)
	return &out, nil
}
