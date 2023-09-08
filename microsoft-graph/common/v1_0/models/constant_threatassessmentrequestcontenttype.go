package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreatAssessmentRequestContentType string

const (
	ThreatAssessmentRequestContentTypefile ThreatAssessmentRequestContentType = "File"
	ThreatAssessmentRequestContentTypemail ThreatAssessmentRequestContentType = "Mail"
	ThreatAssessmentRequestContentTypeurl  ThreatAssessmentRequestContentType = "Url"
)

func PossibleValuesForThreatAssessmentRequestContentType() []string {
	return []string{
		string(ThreatAssessmentRequestContentTypefile),
		string(ThreatAssessmentRequestContentTypemail),
		string(ThreatAssessmentRequestContentTypeurl),
	}
}

func parseThreatAssessmentRequestContentType(input string) (*ThreatAssessmentRequestContentType, error) {
	vals := map[string]ThreatAssessmentRequestContentType{
		"file": ThreatAssessmentRequestContentTypefile,
		"mail": ThreatAssessmentRequestContentTypemail,
		"url":  ThreatAssessmentRequestContentTypeurl,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ThreatAssessmentRequestContentType(input)
	return &out, nil
}
