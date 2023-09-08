package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUserEvidenceVerdict string

const (
	SecurityUserEvidenceVerdictmalicious      SecurityUserEvidenceVerdict = "Malicious"
	SecurityUserEvidenceVerdictnoThreatsFound SecurityUserEvidenceVerdict = "NoThreatsFound"
	SecurityUserEvidenceVerdictsuspicious     SecurityUserEvidenceVerdict = "Suspicious"
	SecurityUserEvidenceVerdictunknown        SecurityUserEvidenceVerdict = "Unknown"
)

func PossibleValuesForSecurityUserEvidenceVerdict() []string {
	return []string{
		string(SecurityUserEvidenceVerdictmalicious),
		string(SecurityUserEvidenceVerdictnoThreatsFound),
		string(SecurityUserEvidenceVerdictsuspicious),
		string(SecurityUserEvidenceVerdictunknown),
	}
}

func parseSecurityUserEvidenceVerdict(input string) (*SecurityUserEvidenceVerdict, error) {
	vals := map[string]SecurityUserEvidenceVerdict{
		"malicious":      SecurityUserEvidenceVerdictmalicious,
		"nothreatsfound": SecurityUserEvidenceVerdictnoThreatsFound,
		"suspicious":     SecurityUserEvidenceVerdictsuspicious,
		"unknown":        SecurityUserEvidenceVerdictunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityUserEvidenceVerdict(input)
	return &out, nil
}
