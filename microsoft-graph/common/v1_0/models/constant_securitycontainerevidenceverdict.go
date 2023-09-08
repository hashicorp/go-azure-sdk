package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContainerEvidenceVerdict string

const (
	SecurityContainerEvidenceVerdictmalicious      SecurityContainerEvidenceVerdict = "Malicious"
	SecurityContainerEvidenceVerdictnoThreatsFound SecurityContainerEvidenceVerdict = "NoThreatsFound"
	SecurityContainerEvidenceVerdictsuspicious     SecurityContainerEvidenceVerdict = "Suspicious"
	SecurityContainerEvidenceVerdictunknown        SecurityContainerEvidenceVerdict = "Unknown"
)

func PossibleValuesForSecurityContainerEvidenceVerdict() []string {
	return []string{
		string(SecurityContainerEvidenceVerdictmalicious),
		string(SecurityContainerEvidenceVerdictnoThreatsFound),
		string(SecurityContainerEvidenceVerdictsuspicious),
		string(SecurityContainerEvidenceVerdictunknown),
	}
}

func parseSecurityContainerEvidenceVerdict(input string) (*SecurityContainerEvidenceVerdict, error) {
	vals := map[string]SecurityContainerEvidenceVerdict{
		"malicious":      SecurityContainerEvidenceVerdictmalicious,
		"nothreatsfound": SecurityContainerEvidenceVerdictnoThreatsFound,
		"suspicious":     SecurityContainerEvidenceVerdictsuspicious,
		"unknown":        SecurityContainerEvidenceVerdictunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityContainerEvidenceVerdict(input)
	return &out, nil
}
