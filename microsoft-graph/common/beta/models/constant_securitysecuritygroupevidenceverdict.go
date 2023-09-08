package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySecurityGroupEvidenceVerdict string

const (
	SecuritySecurityGroupEvidenceVerdictmalicious      SecuritySecurityGroupEvidenceVerdict = "Malicious"
	SecuritySecurityGroupEvidenceVerdictnoThreatsFound SecuritySecurityGroupEvidenceVerdict = "NoThreatsFound"
	SecuritySecurityGroupEvidenceVerdictsuspicious     SecuritySecurityGroupEvidenceVerdict = "Suspicious"
	SecuritySecurityGroupEvidenceVerdictunknown        SecuritySecurityGroupEvidenceVerdict = "Unknown"
)

func PossibleValuesForSecuritySecurityGroupEvidenceVerdict() []string {
	return []string{
		string(SecuritySecurityGroupEvidenceVerdictmalicious),
		string(SecuritySecurityGroupEvidenceVerdictnoThreatsFound),
		string(SecuritySecurityGroupEvidenceVerdictsuspicious),
		string(SecuritySecurityGroupEvidenceVerdictunknown),
	}
}

func parseSecuritySecurityGroupEvidenceVerdict(input string) (*SecuritySecurityGroupEvidenceVerdict, error) {
	vals := map[string]SecuritySecurityGroupEvidenceVerdict{
		"malicious":      SecuritySecurityGroupEvidenceVerdictmalicious,
		"nothreatsfound": SecuritySecurityGroupEvidenceVerdictnoThreatsFound,
		"suspicious":     SecuritySecurityGroupEvidenceVerdictsuspicious,
		"unknown":        SecuritySecurityGroupEvidenceVerdictunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecuritySecurityGroupEvidenceVerdict(input)
	return &out, nil
}
