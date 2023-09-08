package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBlobContainerEvidenceVerdict string

const (
	SecurityBlobContainerEvidenceVerdictmalicious      SecurityBlobContainerEvidenceVerdict = "Malicious"
	SecurityBlobContainerEvidenceVerdictnoThreatsFound SecurityBlobContainerEvidenceVerdict = "NoThreatsFound"
	SecurityBlobContainerEvidenceVerdictsuspicious     SecurityBlobContainerEvidenceVerdict = "Suspicious"
	SecurityBlobContainerEvidenceVerdictunknown        SecurityBlobContainerEvidenceVerdict = "Unknown"
)

func PossibleValuesForSecurityBlobContainerEvidenceVerdict() []string {
	return []string{
		string(SecurityBlobContainerEvidenceVerdictmalicious),
		string(SecurityBlobContainerEvidenceVerdictnoThreatsFound),
		string(SecurityBlobContainerEvidenceVerdictsuspicious),
		string(SecurityBlobContainerEvidenceVerdictunknown),
	}
}

func parseSecurityBlobContainerEvidenceVerdict(input string) (*SecurityBlobContainerEvidenceVerdict, error) {
	vals := map[string]SecurityBlobContainerEvidenceVerdict{
		"malicious":      SecurityBlobContainerEvidenceVerdictmalicious,
		"nothreatsfound": SecurityBlobContainerEvidenceVerdictnoThreatsFound,
		"suspicious":     SecurityBlobContainerEvidenceVerdictsuspicious,
		"unknown":        SecurityBlobContainerEvidenceVerdictunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBlobContainerEvidenceVerdict(input)
	return &out, nil
}
