package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUrlEvidenceVerdict string

const (
	SecurityUrlEvidenceVerdictmalicious      SecurityUrlEvidenceVerdict = "Malicious"
	SecurityUrlEvidenceVerdictnoThreatsFound SecurityUrlEvidenceVerdict = "NoThreatsFound"
	SecurityUrlEvidenceVerdictsuspicious     SecurityUrlEvidenceVerdict = "Suspicious"
	SecurityUrlEvidenceVerdictunknown        SecurityUrlEvidenceVerdict = "Unknown"
)

func PossibleValuesForSecurityUrlEvidenceVerdict() []string {
	return []string{
		string(SecurityUrlEvidenceVerdictmalicious),
		string(SecurityUrlEvidenceVerdictnoThreatsFound),
		string(SecurityUrlEvidenceVerdictsuspicious),
		string(SecurityUrlEvidenceVerdictunknown),
	}
}

func parseSecurityUrlEvidenceVerdict(input string) (*SecurityUrlEvidenceVerdict, error) {
	vals := map[string]SecurityUrlEvidenceVerdict{
		"malicious":      SecurityUrlEvidenceVerdictmalicious,
		"nothreatsfound": SecurityUrlEvidenceVerdictnoThreatsFound,
		"suspicious":     SecurityUrlEvidenceVerdictsuspicious,
		"unknown":        SecurityUrlEvidenceVerdictunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityUrlEvidenceVerdict(input)
	return &out, nil
}
