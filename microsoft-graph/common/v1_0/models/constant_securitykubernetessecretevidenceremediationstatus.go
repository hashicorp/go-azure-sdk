package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesSecretEvidenceRemediationStatus string

const (
	SecurityKubernetesSecretEvidenceRemediationStatusblocked    SecurityKubernetesSecretEvidenceRemediationStatus = "Blocked"
	SecurityKubernetesSecretEvidenceRemediationStatusnone       SecurityKubernetesSecretEvidenceRemediationStatus = "None"
	SecurityKubernetesSecretEvidenceRemediationStatusnotFound   SecurityKubernetesSecretEvidenceRemediationStatus = "NotFound"
	SecurityKubernetesSecretEvidenceRemediationStatusprevented  SecurityKubernetesSecretEvidenceRemediationStatus = "Prevented"
	SecurityKubernetesSecretEvidenceRemediationStatusremediated SecurityKubernetesSecretEvidenceRemediationStatus = "Remediated"
)

func PossibleValuesForSecurityKubernetesSecretEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityKubernetesSecretEvidenceRemediationStatusblocked),
		string(SecurityKubernetesSecretEvidenceRemediationStatusnone),
		string(SecurityKubernetesSecretEvidenceRemediationStatusnotFound),
		string(SecurityKubernetesSecretEvidenceRemediationStatusprevented),
		string(SecurityKubernetesSecretEvidenceRemediationStatusremediated),
	}
}

func parseSecurityKubernetesSecretEvidenceRemediationStatus(input string) (*SecurityKubernetesSecretEvidenceRemediationStatus, error) {
	vals := map[string]SecurityKubernetesSecretEvidenceRemediationStatus{
		"blocked":    SecurityKubernetesSecretEvidenceRemediationStatusblocked,
		"none":       SecurityKubernetesSecretEvidenceRemediationStatusnone,
		"notfound":   SecurityKubernetesSecretEvidenceRemediationStatusnotFound,
		"prevented":  SecurityKubernetesSecretEvidenceRemediationStatusprevented,
		"remediated": SecurityKubernetesSecretEvidenceRemediationStatusremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesSecretEvidenceRemediationStatus(input)
	return &out, nil
}
