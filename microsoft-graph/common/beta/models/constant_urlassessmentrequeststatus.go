package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UrlAssessmentRequestStatus string

const (
	UrlAssessmentRequestStatuscompleted UrlAssessmentRequestStatus = "Completed"
	UrlAssessmentRequestStatuspending   UrlAssessmentRequestStatus = "Pending"
)

func PossibleValuesForUrlAssessmentRequestStatus() []string {
	return []string{
		string(UrlAssessmentRequestStatuscompleted),
		string(UrlAssessmentRequestStatuspending),
	}
}

func parseUrlAssessmentRequestStatus(input string) (*UrlAssessmentRequestStatus, error) {
	vals := map[string]UrlAssessmentRequestStatus{
		"completed": UrlAssessmentRequestStatuscompleted,
		"pending":   UrlAssessmentRequestStatuspending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UrlAssessmentRequestStatus(input)
	return &out, nil
}
