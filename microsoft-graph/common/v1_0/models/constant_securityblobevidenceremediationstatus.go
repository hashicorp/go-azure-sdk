package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBlobEvidenceRemediationStatus string

const (
	SecurityBlobEvidenceRemediationStatusblocked    SecurityBlobEvidenceRemediationStatus = "Blocked"
	SecurityBlobEvidenceRemediationStatusnone       SecurityBlobEvidenceRemediationStatus = "None"
	SecurityBlobEvidenceRemediationStatusnotFound   SecurityBlobEvidenceRemediationStatus = "NotFound"
	SecurityBlobEvidenceRemediationStatusprevented  SecurityBlobEvidenceRemediationStatus = "Prevented"
	SecurityBlobEvidenceRemediationStatusremediated SecurityBlobEvidenceRemediationStatus = "Remediated"
)

func PossibleValuesForSecurityBlobEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityBlobEvidenceRemediationStatusblocked),
		string(SecurityBlobEvidenceRemediationStatusnone),
		string(SecurityBlobEvidenceRemediationStatusnotFound),
		string(SecurityBlobEvidenceRemediationStatusprevented),
		string(SecurityBlobEvidenceRemediationStatusremediated),
	}
}

func parseSecurityBlobEvidenceRemediationStatus(input string) (*SecurityBlobEvidenceRemediationStatus, error) {
	vals := map[string]SecurityBlobEvidenceRemediationStatus{
		"blocked":    SecurityBlobEvidenceRemediationStatusblocked,
		"none":       SecurityBlobEvidenceRemediationStatusnone,
		"notfound":   SecurityBlobEvidenceRemediationStatusnotFound,
		"prevented":  SecurityBlobEvidenceRemediationStatusprevented,
		"remediated": SecurityBlobEvidenceRemediationStatusremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBlobEvidenceRemediationStatus(input)
	return &out, nil
}
