package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAmazonResourceEvidenceVerdict string

const (
	SecurityAmazonResourceEvidenceVerdictmalicious      SecurityAmazonResourceEvidenceVerdict = "Malicious"
	SecurityAmazonResourceEvidenceVerdictnoThreatsFound SecurityAmazonResourceEvidenceVerdict = "NoThreatsFound"
	SecurityAmazonResourceEvidenceVerdictsuspicious     SecurityAmazonResourceEvidenceVerdict = "Suspicious"
	SecurityAmazonResourceEvidenceVerdictunknown        SecurityAmazonResourceEvidenceVerdict = "Unknown"
)

func PossibleValuesForSecurityAmazonResourceEvidenceVerdict() []string {
	return []string{
		string(SecurityAmazonResourceEvidenceVerdictmalicious),
		string(SecurityAmazonResourceEvidenceVerdictnoThreatsFound),
		string(SecurityAmazonResourceEvidenceVerdictsuspicious),
		string(SecurityAmazonResourceEvidenceVerdictunknown),
	}
}

func parseSecurityAmazonResourceEvidenceVerdict(input string) (*SecurityAmazonResourceEvidenceVerdict, error) {
	vals := map[string]SecurityAmazonResourceEvidenceVerdict{
		"malicious":      SecurityAmazonResourceEvidenceVerdictmalicious,
		"nothreatsfound": SecurityAmazonResourceEvidenceVerdictnoThreatsFound,
		"suspicious":     SecurityAmazonResourceEvidenceVerdictsuspicious,
		"unknown":        SecurityAmazonResourceEvidenceVerdictunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAmazonResourceEvidenceVerdict(input)
	return &out, nil
}
