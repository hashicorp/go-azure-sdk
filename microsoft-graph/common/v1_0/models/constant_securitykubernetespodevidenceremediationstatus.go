package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesPodEvidenceRemediationStatus string

const (
	SecurityKubernetesPodEvidenceRemediationStatusblocked    SecurityKubernetesPodEvidenceRemediationStatus = "Blocked"
	SecurityKubernetesPodEvidenceRemediationStatusnone       SecurityKubernetesPodEvidenceRemediationStatus = "None"
	SecurityKubernetesPodEvidenceRemediationStatusnotFound   SecurityKubernetesPodEvidenceRemediationStatus = "NotFound"
	SecurityKubernetesPodEvidenceRemediationStatusprevented  SecurityKubernetesPodEvidenceRemediationStatus = "Prevented"
	SecurityKubernetesPodEvidenceRemediationStatusremediated SecurityKubernetesPodEvidenceRemediationStatus = "Remediated"
)

func PossibleValuesForSecurityKubernetesPodEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityKubernetesPodEvidenceRemediationStatusblocked),
		string(SecurityKubernetesPodEvidenceRemediationStatusnone),
		string(SecurityKubernetesPodEvidenceRemediationStatusnotFound),
		string(SecurityKubernetesPodEvidenceRemediationStatusprevented),
		string(SecurityKubernetesPodEvidenceRemediationStatusremediated),
	}
}

func parseSecurityKubernetesPodEvidenceRemediationStatus(input string) (*SecurityKubernetesPodEvidenceRemediationStatus, error) {
	vals := map[string]SecurityKubernetesPodEvidenceRemediationStatus{
		"blocked":    SecurityKubernetesPodEvidenceRemediationStatusblocked,
		"none":       SecurityKubernetesPodEvidenceRemediationStatusnone,
		"notfound":   SecurityKubernetesPodEvidenceRemediationStatusnotFound,
		"prevented":  SecurityKubernetesPodEvidenceRemediationStatusprevented,
		"remediated": SecurityKubernetesPodEvidenceRemediationStatusremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesPodEvidenceRemediationStatus(input)
	return &out, nil
}
