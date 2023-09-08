package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRegistryKeyEvidenceVerdict string

const (
	SecurityRegistryKeyEvidenceVerdictmalicious      SecurityRegistryKeyEvidenceVerdict = "Malicious"
	SecurityRegistryKeyEvidenceVerdictnoThreatsFound SecurityRegistryKeyEvidenceVerdict = "NoThreatsFound"
	SecurityRegistryKeyEvidenceVerdictsuspicious     SecurityRegistryKeyEvidenceVerdict = "Suspicious"
	SecurityRegistryKeyEvidenceVerdictunknown        SecurityRegistryKeyEvidenceVerdict = "Unknown"
)

func PossibleValuesForSecurityRegistryKeyEvidenceVerdict() []string {
	return []string{
		string(SecurityRegistryKeyEvidenceVerdictmalicious),
		string(SecurityRegistryKeyEvidenceVerdictnoThreatsFound),
		string(SecurityRegistryKeyEvidenceVerdictsuspicious),
		string(SecurityRegistryKeyEvidenceVerdictunknown),
	}
}

func parseSecurityRegistryKeyEvidenceVerdict(input string) (*SecurityRegistryKeyEvidenceVerdict, error) {
	vals := map[string]SecurityRegistryKeyEvidenceVerdict{
		"malicious":      SecurityRegistryKeyEvidenceVerdictmalicious,
		"nothreatsfound": SecurityRegistryKeyEvidenceVerdictnoThreatsFound,
		"suspicious":     SecurityRegistryKeyEvidenceVerdictsuspicious,
		"unknown":        SecurityRegistryKeyEvidenceVerdictunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRegistryKeyEvidenceVerdict(input)
	return &out, nil
}
