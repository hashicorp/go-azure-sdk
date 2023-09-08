package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesServiceAccountEvidenceVerdict string

const (
	SecurityKubernetesServiceAccountEvidenceVerdictmalicious      SecurityKubernetesServiceAccountEvidenceVerdict = "Malicious"
	SecurityKubernetesServiceAccountEvidenceVerdictnoThreatsFound SecurityKubernetesServiceAccountEvidenceVerdict = "NoThreatsFound"
	SecurityKubernetesServiceAccountEvidenceVerdictsuspicious     SecurityKubernetesServiceAccountEvidenceVerdict = "Suspicious"
	SecurityKubernetesServiceAccountEvidenceVerdictunknown        SecurityKubernetesServiceAccountEvidenceVerdict = "Unknown"
)

func PossibleValuesForSecurityKubernetesServiceAccountEvidenceVerdict() []string {
	return []string{
		string(SecurityKubernetesServiceAccountEvidenceVerdictmalicious),
		string(SecurityKubernetesServiceAccountEvidenceVerdictnoThreatsFound),
		string(SecurityKubernetesServiceAccountEvidenceVerdictsuspicious),
		string(SecurityKubernetesServiceAccountEvidenceVerdictunknown),
	}
}

func parseSecurityKubernetesServiceAccountEvidenceVerdict(input string) (*SecurityKubernetesServiceAccountEvidenceVerdict, error) {
	vals := map[string]SecurityKubernetesServiceAccountEvidenceVerdict{
		"malicious":      SecurityKubernetesServiceAccountEvidenceVerdictmalicious,
		"nothreatsfound": SecurityKubernetesServiceAccountEvidenceVerdictnoThreatsFound,
		"suspicious":     SecurityKubernetesServiceAccountEvidenceVerdictsuspicious,
		"unknown":        SecurityKubernetesServiceAccountEvidenceVerdictunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesServiceAccountEvidenceVerdict(input)
	return &out, nil
}
