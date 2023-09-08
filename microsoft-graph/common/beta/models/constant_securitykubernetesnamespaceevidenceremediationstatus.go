package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesNamespaceEvidenceRemediationStatus string

const (
	SecurityKubernetesNamespaceEvidenceRemediationStatusblocked    SecurityKubernetesNamespaceEvidenceRemediationStatus = "Blocked"
	SecurityKubernetesNamespaceEvidenceRemediationStatusnone       SecurityKubernetesNamespaceEvidenceRemediationStatus = "None"
	SecurityKubernetesNamespaceEvidenceRemediationStatusnotFound   SecurityKubernetesNamespaceEvidenceRemediationStatus = "NotFound"
	SecurityKubernetesNamespaceEvidenceRemediationStatusprevented  SecurityKubernetesNamespaceEvidenceRemediationStatus = "Prevented"
	SecurityKubernetesNamespaceEvidenceRemediationStatusremediated SecurityKubernetesNamespaceEvidenceRemediationStatus = "Remediated"
)

func PossibleValuesForSecurityKubernetesNamespaceEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityKubernetesNamespaceEvidenceRemediationStatusblocked),
		string(SecurityKubernetesNamespaceEvidenceRemediationStatusnone),
		string(SecurityKubernetesNamespaceEvidenceRemediationStatusnotFound),
		string(SecurityKubernetesNamespaceEvidenceRemediationStatusprevented),
		string(SecurityKubernetesNamespaceEvidenceRemediationStatusremediated),
	}
}

func parseSecurityKubernetesNamespaceEvidenceRemediationStatus(input string) (*SecurityKubernetesNamespaceEvidenceRemediationStatus, error) {
	vals := map[string]SecurityKubernetesNamespaceEvidenceRemediationStatus{
		"blocked":    SecurityKubernetesNamespaceEvidenceRemediationStatusblocked,
		"none":       SecurityKubernetesNamespaceEvidenceRemediationStatusnone,
		"notfound":   SecurityKubernetesNamespaceEvidenceRemediationStatusnotFound,
		"prevented":  SecurityKubernetesNamespaceEvidenceRemediationStatusprevented,
		"remediated": SecurityKubernetesNamespaceEvidenceRemediationStatusremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesNamespaceEvidenceRemediationStatus(input)
	return &out, nil
}
