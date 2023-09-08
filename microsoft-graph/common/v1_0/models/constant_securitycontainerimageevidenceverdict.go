package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContainerImageEvidenceVerdict string

const (
	SecurityContainerImageEvidenceVerdictmalicious      SecurityContainerImageEvidenceVerdict = "Malicious"
	SecurityContainerImageEvidenceVerdictnoThreatsFound SecurityContainerImageEvidenceVerdict = "NoThreatsFound"
	SecurityContainerImageEvidenceVerdictsuspicious     SecurityContainerImageEvidenceVerdict = "Suspicious"
	SecurityContainerImageEvidenceVerdictunknown        SecurityContainerImageEvidenceVerdict = "Unknown"
)

func PossibleValuesForSecurityContainerImageEvidenceVerdict() []string {
	return []string{
		string(SecurityContainerImageEvidenceVerdictmalicious),
		string(SecurityContainerImageEvidenceVerdictnoThreatsFound),
		string(SecurityContainerImageEvidenceVerdictsuspicious),
		string(SecurityContainerImageEvidenceVerdictunknown),
	}
}

func parseSecurityContainerImageEvidenceVerdict(input string) (*SecurityContainerImageEvidenceVerdict, error) {
	vals := map[string]SecurityContainerImageEvidenceVerdict{
		"malicious":      SecurityContainerImageEvidenceVerdictmalicious,
		"nothreatsfound": SecurityContainerImageEvidenceVerdictnoThreatsFound,
		"suspicious":     SecurityContainerImageEvidenceVerdictsuspicious,
		"unknown":        SecurityContainerImageEvidenceVerdictunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityContainerImageEvidenceVerdict(input)
	return &out, nil
}
