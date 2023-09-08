package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBlobEvidenceVerdict string

const (
	SecurityBlobEvidenceVerdictmalicious      SecurityBlobEvidenceVerdict = "Malicious"
	SecurityBlobEvidenceVerdictnoThreatsFound SecurityBlobEvidenceVerdict = "NoThreatsFound"
	SecurityBlobEvidenceVerdictsuspicious     SecurityBlobEvidenceVerdict = "Suspicious"
	SecurityBlobEvidenceVerdictunknown        SecurityBlobEvidenceVerdict = "Unknown"
)

func PossibleValuesForSecurityBlobEvidenceVerdict() []string {
	return []string{
		string(SecurityBlobEvidenceVerdictmalicious),
		string(SecurityBlobEvidenceVerdictnoThreatsFound),
		string(SecurityBlobEvidenceVerdictsuspicious),
		string(SecurityBlobEvidenceVerdictunknown),
	}
}

func parseSecurityBlobEvidenceVerdict(input string) (*SecurityBlobEvidenceVerdict, error) {
	vals := map[string]SecurityBlobEvidenceVerdict{
		"malicious":      SecurityBlobEvidenceVerdictmalicious,
		"nothreatsfound": SecurityBlobEvidenceVerdictnoThreatsFound,
		"suspicious":     SecurityBlobEvidenceVerdictsuspicious,
		"unknown":        SecurityBlobEvidenceVerdictunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBlobEvidenceVerdict(input)
	return &out, nil
}
