package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityProcessEvidenceRemediationStatus string

const (
	SecurityProcessEvidenceRemediationStatusblocked    SecurityProcessEvidenceRemediationStatus = "Blocked"
	SecurityProcessEvidenceRemediationStatusnone       SecurityProcessEvidenceRemediationStatus = "None"
	SecurityProcessEvidenceRemediationStatusnotFound   SecurityProcessEvidenceRemediationStatus = "NotFound"
	SecurityProcessEvidenceRemediationStatusprevented  SecurityProcessEvidenceRemediationStatus = "Prevented"
	SecurityProcessEvidenceRemediationStatusremediated SecurityProcessEvidenceRemediationStatus = "Remediated"
)

func PossibleValuesForSecurityProcessEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityProcessEvidenceRemediationStatusblocked),
		string(SecurityProcessEvidenceRemediationStatusnone),
		string(SecurityProcessEvidenceRemediationStatusnotFound),
		string(SecurityProcessEvidenceRemediationStatusprevented),
		string(SecurityProcessEvidenceRemediationStatusremediated),
	}
}

func parseSecurityProcessEvidenceRemediationStatus(input string) (*SecurityProcessEvidenceRemediationStatus, error) {
	vals := map[string]SecurityProcessEvidenceRemediationStatus{
		"blocked":    SecurityProcessEvidenceRemediationStatusblocked,
		"none":       SecurityProcessEvidenceRemediationStatusnone,
		"notfound":   SecurityProcessEvidenceRemediationStatusnotFound,
		"prevented":  SecurityProcessEvidenceRemediationStatusprevented,
		"remediated": SecurityProcessEvidenceRemediationStatusremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityProcessEvidenceRemediationStatus(input)
	return &out, nil
}
