package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreatAssessmentResultResultType string

const (
	ThreatAssessmentResultResultTypecheckPolicy ThreatAssessmentResultResultType = "CheckPolicy"
	ThreatAssessmentResultResultTyperescan      ThreatAssessmentResultResultType = "Rescan"
)

func PossibleValuesForThreatAssessmentResultResultType() []string {
	return []string{
		string(ThreatAssessmentResultResultTypecheckPolicy),
		string(ThreatAssessmentResultResultTyperescan),
	}
}

func parseThreatAssessmentResultResultType(input string) (*ThreatAssessmentResultResultType, error) {
	vals := map[string]ThreatAssessmentResultResultType{
		"checkpolicy": ThreatAssessmentResultResultTypecheckPolicy,
		"rescan":      ThreatAssessmentResultResultTyperescan,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ThreatAssessmentResultResultType(input)
	return &out, nil
}
