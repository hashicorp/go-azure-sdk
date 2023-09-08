package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreatAssessmentRequestExpectedAssessment string

const (
	ThreatAssessmentRequestExpectedAssessmentblock   ThreatAssessmentRequestExpectedAssessment = "Block"
	ThreatAssessmentRequestExpectedAssessmentunblock ThreatAssessmentRequestExpectedAssessment = "Unblock"
)

func PossibleValuesForThreatAssessmentRequestExpectedAssessment() []string {
	return []string{
		string(ThreatAssessmentRequestExpectedAssessmentblock),
		string(ThreatAssessmentRequestExpectedAssessmentunblock),
	}
}

func parseThreatAssessmentRequestExpectedAssessment(input string) (*ThreatAssessmentRequestExpectedAssessment, error) {
	vals := map[string]ThreatAssessmentRequestExpectedAssessment{
		"block":   ThreatAssessmentRequestExpectedAssessmentblock,
		"unblock": ThreatAssessmentRequestExpectedAssessmentunblock,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ThreatAssessmentRequestExpectedAssessment(input)
	return &out, nil
}
