package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesControllerEvidenceRemediationStatus string

const (
	SecurityKubernetesControllerEvidenceRemediationStatusblocked    SecurityKubernetesControllerEvidenceRemediationStatus = "Blocked"
	SecurityKubernetesControllerEvidenceRemediationStatusnone       SecurityKubernetesControllerEvidenceRemediationStatus = "None"
	SecurityKubernetesControllerEvidenceRemediationStatusnotFound   SecurityKubernetesControllerEvidenceRemediationStatus = "NotFound"
	SecurityKubernetesControllerEvidenceRemediationStatusprevented  SecurityKubernetesControllerEvidenceRemediationStatus = "Prevented"
	SecurityKubernetesControllerEvidenceRemediationStatusremediated SecurityKubernetesControllerEvidenceRemediationStatus = "Remediated"
)

func PossibleValuesForSecurityKubernetesControllerEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityKubernetesControllerEvidenceRemediationStatusblocked),
		string(SecurityKubernetesControllerEvidenceRemediationStatusnone),
		string(SecurityKubernetesControllerEvidenceRemediationStatusnotFound),
		string(SecurityKubernetesControllerEvidenceRemediationStatusprevented),
		string(SecurityKubernetesControllerEvidenceRemediationStatusremediated),
	}
}

func parseSecurityKubernetesControllerEvidenceRemediationStatus(input string) (*SecurityKubernetesControllerEvidenceRemediationStatus, error) {
	vals := map[string]SecurityKubernetesControllerEvidenceRemediationStatus{
		"blocked":    SecurityKubernetesControllerEvidenceRemediationStatusblocked,
		"none":       SecurityKubernetesControllerEvidenceRemediationStatusnone,
		"notfound":   SecurityKubernetesControllerEvidenceRemediationStatusnotFound,
		"prevented":  SecurityKubernetesControllerEvidenceRemediationStatusprevented,
		"remediated": SecurityKubernetesControllerEvidenceRemediationStatusremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesControllerEvidenceRemediationStatus(input)
	return &out, nil
}
