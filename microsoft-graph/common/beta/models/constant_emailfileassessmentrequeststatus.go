package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailFileAssessmentRequestStatus string

const (
	EmailFileAssessmentRequestStatuscompleted EmailFileAssessmentRequestStatus = "Completed"
	EmailFileAssessmentRequestStatuspending   EmailFileAssessmentRequestStatus = "Pending"
)

func PossibleValuesForEmailFileAssessmentRequestStatus() []string {
	return []string{
		string(EmailFileAssessmentRequestStatuscompleted),
		string(EmailFileAssessmentRequestStatuspending),
	}
}

func parseEmailFileAssessmentRequestStatus(input string) (*EmailFileAssessmentRequestStatus, error) {
	vals := map[string]EmailFileAssessmentRequestStatus{
		"completed": EmailFileAssessmentRequestStatuscompleted,
		"pending":   EmailFileAssessmentRequestStatuspending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EmailFileAssessmentRequestStatus(input)
	return &out, nil
}
