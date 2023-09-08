package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIpEvidenceVerdict string

const (
	SecurityIpEvidenceVerdictmalicious      SecurityIpEvidenceVerdict = "Malicious"
	SecurityIpEvidenceVerdictnoThreatsFound SecurityIpEvidenceVerdict = "NoThreatsFound"
	SecurityIpEvidenceVerdictsuspicious     SecurityIpEvidenceVerdict = "Suspicious"
	SecurityIpEvidenceVerdictunknown        SecurityIpEvidenceVerdict = "Unknown"
)

func PossibleValuesForSecurityIpEvidenceVerdict() []string {
	return []string{
		string(SecurityIpEvidenceVerdictmalicious),
		string(SecurityIpEvidenceVerdictnoThreatsFound),
		string(SecurityIpEvidenceVerdictsuspicious),
		string(SecurityIpEvidenceVerdictunknown),
	}
}

func parseSecurityIpEvidenceVerdict(input string) (*SecurityIpEvidenceVerdict, error) {
	vals := map[string]SecurityIpEvidenceVerdict{
		"malicious":      SecurityIpEvidenceVerdictmalicious,
		"nothreatsfound": SecurityIpEvidenceVerdictnoThreatsFound,
		"suspicious":     SecurityIpEvidenceVerdictsuspicious,
		"unknown":        SecurityIpEvidenceVerdictunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityIpEvidenceVerdict(input)
	return &out, nil
}
