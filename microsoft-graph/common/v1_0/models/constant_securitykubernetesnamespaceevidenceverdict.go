package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesNamespaceEvidenceVerdict string

const (
	SecurityKubernetesNamespaceEvidenceVerdictmalicious      SecurityKubernetesNamespaceEvidenceVerdict = "Malicious"
	SecurityKubernetesNamespaceEvidenceVerdictnoThreatsFound SecurityKubernetesNamespaceEvidenceVerdict = "NoThreatsFound"
	SecurityKubernetesNamespaceEvidenceVerdictsuspicious     SecurityKubernetesNamespaceEvidenceVerdict = "Suspicious"
	SecurityKubernetesNamespaceEvidenceVerdictunknown        SecurityKubernetesNamespaceEvidenceVerdict = "Unknown"
)

func PossibleValuesForSecurityKubernetesNamespaceEvidenceVerdict() []string {
	return []string{
		string(SecurityKubernetesNamespaceEvidenceVerdictmalicious),
		string(SecurityKubernetesNamespaceEvidenceVerdictnoThreatsFound),
		string(SecurityKubernetesNamespaceEvidenceVerdictsuspicious),
		string(SecurityKubernetesNamespaceEvidenceVerdictunknown),
	}
}

func parseSecurityKubernetesNamespaceEvidenceVerdict(input string) (*SecurityKubernetesNamespaceEvidenceVerdict, error) {
	vals := map[string]SecurityKubernetesNamespaceEvidenceVerdict{
		"malicious":      SecurityKubernetesNamespaceEvidenceVerdictmalicious,
		"nothreatsfound": SecurityKubernetesNamespaceEvidenceVerdictnoThreatsFound,
		"suspicious":     SecurityKubernetesNamespaceEvidenceVerdictsuspicious,
		"unknown":        SecurityKubernetesNamespaceEvidenceVerdictunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesNamespaceEvidenceVerdict(input)
	return &out, nil
}
