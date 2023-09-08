package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreatAssessmentRequestStatus string

const (
	ThreatAssessmentRequestStatuscompleted ThreatAssessmentRequestStatus = "Completed"
	ThreatAssessmentRequestStatuspending   ThreatAssessmentRequestStatus = "Pending"
)

func PossibleValuesForThreatAssessmentRequestStatus() []string {
	return []string{
		string(ThreatAssessmentRequestStatuscompleted),
		string(ThreatAssessmentRequestStatuspending),
	}
}

func parseThreatAssessmentRequestStatus(input string) (*ThreatAssessmentRequestStatus, error) {
	vals := map[string]ThreatAssessmentRequestStatus{
		"completed": ThreatAssessmentRequestStatuscompleted,
		"pending":   ThreatAssessmentRequestStatuspending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ThreatAssessmentRequestStatus(input)
	return &out, nil
}
