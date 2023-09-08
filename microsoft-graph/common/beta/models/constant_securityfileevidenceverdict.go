package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileEvidenceVerdict string

const (
	SecurityFileEvidenceVerdictmalicious      SecurityFileEvidenceVerdict = "Malicious"
	SecurityFileEvidenceVerdictnoThreatsFound SecurityFileEvidenceVerdict = "NoThreatsFound"
	SecurityFileEvidenceVerdictsuspicious     SecurityFileEvidenceVerdict = "Suspicious"
	SecurityFileEvidenceVerdictunknown        SecurityFileEvidenceVerdict = "Unknown"
)

func PossibleValuesForSecurityFileEvidenceVerdict() []string {
	return []string{
		string(SecurityFileEvidenceVerdictmalicious),
		string(SecurityFileEvidenceVerdictnoThreatsFound),
		string(SecurityFileEvidenceVerdictsuspicious),
		string(SecurityFileEvidenceVerdictunknown),
	}
}

func parseSecurityFileEvidenceVerdict(input string) (*SecurityFileEvidenceVerdict, error) {
	vals := map[string]SecurityFileEvidenceVerdict{
		"malicious":      SecurityFileEvidenceVerdictmalicious,
		"nothreatsfound": SecurityFileEvidenceVerdictnoThreatsFound,
		"suspicious":     SecurityFileEvidenceVerdictsuspicious,
		"unknown":        SecurityFileEvidenceVerdictunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileEvidenceVerdict(input)
	return &out, nil
}
