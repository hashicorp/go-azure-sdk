package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAzureResourceEvidenceVerdict string

const (
	SecurityAzureResourceEvidenceVerdictmalicious      SecurityAzureResourceEvidenceVerdict = "Malicious"
	SecurityAzureResourceEvidenceVerdictnoThreatsFound SecurityAzureResourceEvidenceVerdict = "NoThreatsFound"
	SecurityAzureResourceEvidenceVerdictsuspicious     SecurityAzureResourceEvidenceVerdict = "Suspicious"
	SecurityAzureResourceEvidenceVerdictunknown        SecurityAzureResourceEvidenceVerdict = "Unknown"
)

func PossibleValuesForSecurityAzureResourceEvidenceVerdict() []string {
	return []string{
		string(SecurityAzureResourceEvidenceVerdictmalicious),
		string(SecurityAzureResourceEvidenceVerdictnoThreatsFound),
		string(SecurityAzureResourceEvidenceVerdictsuspicious),
		string(SecurityAzureResourceEvidenceVerdictunknown),
	}
}

func parseSecurityAzureResourceEvidenceVerdict(input string) (*SecurityAzureResourceEvidenceVerdict, error) {
	vals := map[string]SecurityAzureResourceEvidenceVerdict{
		"malicious":      SecurityAzureResourceEvidenceVerdictmalicious,
		"nothreatsfound": SecurityAzureResourceEvidenceVerdictnoThreatsFound,
		"suspicious":     SecurityAzureResourceEvidenceVerdictsuspicious,
		"unknown":        SecurityAzureResourceEvidenceVerdictunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAzureResourceEvidenceVerdict(input)
	return &out, nil
}
