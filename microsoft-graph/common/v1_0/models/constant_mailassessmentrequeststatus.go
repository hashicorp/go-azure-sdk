package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailAssessmentRequestStatus string

const (
	MailAssessmentRequestStatuscompleted MailAssessmentRequestStatus = "Completed"
	MailAssessmentRequestStatuspending   MailAssessmentRequestStatus = "Pending"
)

func PossibleValuesForMailAssessmentRequestStatus() []string {
	return []string{
		string(MailAssessmentRequestStatuscompleted),
		string(MailAssessmentRequestStatuspending),
	}
}

func parseMailAssessmentRequestStatus(input string) (*MailAssessmentRequestStatus, error) {
	vals := map[string]MailAssessmentRequestStatus{
		"completed": MailAssessmentRequestStatuscompleted,
		"pending":   MailAssessmentRequestStatuspending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MailAssessmentRequestStatus(input)
	return &out, nil
}
