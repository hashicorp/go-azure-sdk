package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesServiceEvidenceRemediationStatus string

const (
	SecurityKubernetesServiceEvidenceRemediationStatusblocked    SecurityKubernetesServiceEvidenceRemediationStatus = "Blocked"
	SecurityKubernetesServiceEvidenceRemediationStatusnone       SecurityKubernetesServiceEvidenceRemediationStatus = "None"
	SecurityKubernetesServiceEvidenceRemediationStatusnotFound   SecurityKubernetesServiceEvidenceRemediationStatus = "NotFound"
	SecurityKubernetesServiceEvidenceRemediationStatusprevented  SecurityKubernetesServiceEvidenceRemediationStatus = "Prevented"
	SecurityKubernetesServiceEvidenceRemediationStatusremediated SecurityKubernetesServiceEvidenceRemediationStatus = "Remediated"
)

func PossibleValuesForSecurityKubernetesServiceEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityKubernetesServiceEvidenceRemediationStatusblocked),
		string(SecurityKubernetesServiceEvidenceRemediationStatusnone),
		string(SecurityKubernetesServiceEvidenceRemediationStatusnotFound),
		string(SecurityKubernetesServiceEvidenceRemediationStatusprevented),
		string(SecurityKubernetesServiceEvidenceRemediationStatusremediated),
	}
}

func parseSecurityKubernetesServiceEvidenceRemediationStatus(input string) (*SecurityKubernetesServiceEvidenceRemediationStatus, error) {
	vals := map[string]SecurityKubernetesServiceEvidenceRemediationStatus{
		"blocked":    SecurityKubernetesServiceEvidenceRemediationStatusblocked,
		"none":       SecurityKubernetesServiceEvidenceRemediationStatusnone,
		"notfound":   SecurityKubernetesServiceEvidenceRemediationStatusnotFound,
		"prevented":  SecurityKubernetesServiceEvidenceRemediationStatusprevented,
		"remediated": SecurityKubernetesServiceEvidenceRemediationStatusremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesServiceEvidenceRemediationStatus(input)
	return &out, nil
}
