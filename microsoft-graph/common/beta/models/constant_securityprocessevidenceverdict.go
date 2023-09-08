package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityProcessEvidenceVerdict string

const (
	SecurityProcessEvidenceVerdictmalicious      SecurityProcessEvidenceVerdict = "Malicious"
	SecurityProcessEvidenceVerdictnoThreatsFound SecurityProcessEvidenceVerdict = "NoThreatsFound"
	SecurityProcessEvidenceVerdictsuspicious     SecurityProcessEvidenceVerdict = "Suspicious"
	SecurityProcessEvidenceVerdictunknown        SecurityProcessEvidenceVerdict = "Unknown"
)

func PossibleValuesForSecurityProcessEvidenceVerdict() []string {
	return []string{
		string(SecurityProcessEvidenceVerdictmalicious),
		string(SecurityProcessEvidenceVerdictnoThreatsFound),
		string(SecurityProcessEvidenceVerdictsuspicious),
		string(SecurityProcessEvidenceVerdictunknown),
	}
}

func parseSecurityProcessEvidenceVerdict(input string) (*SecurityProcessEvidenceVerdict, error) {
	vals := map[string]SecurityProcessEvidenceVerdict{
		"malicious":      SecurityProcessEvidenceVerdictmalicious,
		"nothreatsfound": SecurityProcessEvidenceVerdictnoThreatsFound,
		"suspicious":     SecurityProcessEvidenceVerdictsuspicious,
		"unknown":        SecurityProcessEvidenceVerdictunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityProcessEvidenceVerdict(input)
	return &out, nil
}
