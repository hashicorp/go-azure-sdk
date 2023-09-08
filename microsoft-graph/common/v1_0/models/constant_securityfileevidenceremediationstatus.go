package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileEvidenceRemediationStatus string

const (
	SecurityFileEvidenceRemediationStatusblocked    SecurityFileEvidenceRemediationStatus = "Blocked"
	SecurityFileEvidenceRemediationStatusnone       SecurityFileEvidenceRemediationStatus = "None"
	SecurityFileEvidenceRemediationStatusnotFound   SecurityFileEvidenceRemediationStatus = "NotFound"
	SecurityFileEvidenceRemediationStatusprevented  SecurityFileEvidenceRemediationStatus = "Prevented"
	SecurityFileEvidenceRemediationStatusremediated SecurityFileEvidenceRemediationStatus = "Remediated"
)

func PossibleValuesForSecurityFileEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityFileEvidenceRemediationStatusblocked),
		string(SecurityFileEvidenceRemediationStatusnone),
		string(SecurityFileEvidenceRemediationStatusnotFound),
		string(SecurityFileEvidenceRemediationStatusprevented),
		string(SecurityFileEvidenceRemediationStatusremediated),
	}
}

func parseSecurityFileEvidenceRemediationStatus(input string) (*SecurityFileEvidenceRemediationStatus, error) {
	vals := map[string]SecurityFileEvidenceRemediationStatus{
		"blocked":    SecurityFileEvidenceRemediationStatusblocked,
		"none":       SecurityFileEvidenceRemediationStatusnone,
		"notfound":   SecurityFileEvidenceRemediationStatusnotFound,
		"prevented":  SecurityFileEvidenceRemediationStatusprevented,
		"remediated": SecurityFileEvidenceRemediationStatusremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileEvidenceRemediationStatus(input)
	return &out, nil
}
