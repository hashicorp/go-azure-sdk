package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesServiceEvidenceVerdict string

const (
	SecurityKubernetesServiceEvidenceVerdictmalicious      SecurityKubernetesServiceEvidenceVerdict = "Malicious"
	SecurityKubernetesServiceEvidenceVerdictnoThreatsFound SecurityKubernetesServiceEvidenceVerdict = "NoThreatsFound"
	SecurityKubernetesServiceEvidenceVerdictsuspicious     SecurityKubernetesServiceEvidenceVerdict = "Suspicious"
	SecurityKubernetesServiceEvidenceVerdictunknown        SecurityKubernetesServiceEvidenceVerdict = "Unknown"
)

func PossibleValuesForSecurityKubernetesServiceEvidenceVerdict() []string {
	return []string{
		string(SecurityKubernetesServiceEvidenceVerdictmalicious),
		string(SecurityKubernetesServiceEvidenceVerdictnoThreatsFound),
		string(SecurityKubernetesServiceEvidenceVerdictsuspicious),
		string(SecurityKubernetesServiceEvidenceVerdictunknown),
	}
}

func parseSecurityKubernetesServiceEvidenceVerdict(input string) (*SecurityKubernetesServiceEvidenceVerdict, error) {
	vals := map[string]SecurityKubernetesServiceEvidenceVerdict{
		"malicious":      SecurityKubernetesServiceEvidenceVerdictmalicious,
		"nothreatsfound": SecurityKubernetesServiceEvidenceVerdictnoThreatsFound,
		"suspicious":     SecurityKubernetesServiceEvidenceVerdictsuspicious,
		"unknown":        SecurityKubernetesServiceEvidenceVerdictunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesServiceEvidenceVerdict(input)
	return &out, nil
}
