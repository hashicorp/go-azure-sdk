package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRegistryValueEvidenceVerdict string

const (
	SecurityRegistryValueEvidenceVerdictmalicious      SecurityRegistryValueEvidenceVerdict = "Malicious"
	SecurityRegistryValueEvidenceVerdictnoThreatsFound SecurityRegistryValueEvidenceVerdict = "NoThreatsFound"
	SecurityRegistryValueEvidenceVerdictsuspicious     SecurityRegistryValueEvidenceVerdict = "Suspicious"
	SecurityRegistryValueEvidenceVerdictunknown        SecurityRegistryValueEvidenceVerdict = "Unknown"
)

func PossibleValuesForSecurityRegistryValueEvidenceVerdict() []string {
	return []string{
		string(SecurityRegistryValueEvidenceVerdictmalicious),
		string(SecurityRegistryValueEvidenceVerdictnoThreatsFound),
		string(SecurityRegistryValueEvidenceVerdictsuspicious),
		string(SecurityRegistryValueEvidenceVerdictunknown),
	}
}

func parseSecurityRegistryValueEvidenceVerdict(input string) (*SecurityRegistryValueEvidenceVerdict, error) {
	vals := map[string]SecurityRegistryValueEvidenceVerdict{
		"malicious":      SecurityRegistryValueEvidenceVerdictmalicious,
		"nothreatsfound": SecurityRegistryValueEvidenceVerdictnoThreatsFound,
		"suspicious":     SecurityRegistryValueEvidenceVerdictsuspicious,
		"unknown":        SecurityRegistryValueEvidenceVerdictunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRegistryValueEvidenceVerdict(input)
	return &out, nil
}
