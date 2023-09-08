package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesControllerEvidenceVerdict string

const (
	SecurityKubernetesControllerEvidenceVerdictmalicious      SecurityKubernetesControllerEvidenceVerdict = "Malicious"
	SecurityKubernetesControllerEvidenceVerdictnoThreatsFound SecurityKubernetesControllerEvidenceVerdict = "NoThreatsFound"
	SecurityKubernetesControllerEvidenceVerdictsuspicious     SecurityKubernetesControllerEvidenceVerdict = "Suspicious"
	SecurityKubernetesControllerEvidenceVerdictunknown        SecurityKubernetesControllerEvidenceVerdict = "Unknown"
)

func PossibleValuesForSecurityKubernetesControllerEvidenceVerdict() []string {
	return []string{
		string(SecurityKubernetesControllerEvidenceVerdictmalicious),
		string(SecurityKubernetesControllerEvidenceVerdictnoThreatsFound),
		string(SecurityKubernetesControllerEvidenceVerdictsuspicious),
		string(SecurityKubernetesControllerEvidenceVerdictunknown),
	}
}

func parseSecurityKubernetesControllerEvidenceVerdict(input string) (*SecurityKubernetesControllerEvidenceVerdict, error) {
	vals := map[string]SecurityKubernetesControllerEvidenceVerdict{
		"malicious":      SecurityKubernetesControllerEvidenceVerdictmalicious,
		"nothreatsfound": SecurityKubernetesControllerEvidenceVerdictnoThreatsFound,
		"suspicious":     SecurityKubernetesControllerEvidenceVerdictsuspicious,
		"unknown":        SecurityKubernetesControllerEvidenceVerdictunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesControllerEvidenceVerdict(input)
	return &out, nil
}
