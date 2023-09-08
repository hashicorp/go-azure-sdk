package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBlobContainerEvidenceRemediationStatus string

const (
	SecurityBlobContainerEvidenceRemediationStatusblocked    SecurityBlobContainerEvidenceRemediationStatus = "Blocked"
	SecurityBlobContainerEvidenceRemediationStatusnone       SecurityBlobContainerEvidenceRemediationStatus = "None"
	SecurityBlobContainerEvidenceRemediationStatusnotFound   SecurityBlobContainerEvidenceRemediationStatus = "NotFound"
	SecurityBlobContainerEvidenceRemediationStatusprevented  SecurityBlobContainerEvidenceRemediationStatus = "Prevented"
	SecurityBlobContainerEvidenceRemediationStatusremediated SecurityBlobContainerEvidenceRemediationStatus = "Remediated"
)

func PossibleValuesForSecurityBlobContainerEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityBlobContainerEvidenceRemediationStatusblocked),
		string(SecurityBlobContainerEvidenceRemediationStatusnone),
		string(SecurityBlobContainerEvidenceRemediationStatusnotFound),
		string(SecurityBlobContainerEvidenceRemediationStatusprevented),
		string(SecurityBlobContainerEvidenceRemediationStatusremediated),
	}
}

func parseSecurityBlobContainerEvidenceRemediationStatus(input string) (*SecurityBlobContainerEvidenceRemediationStatus, error) {
	vals := map[string]SecurityBlobContainerEvidenceRemediationStatus{
		"blocked":    SecurityBlobContainerEvidenceRemediationStatusblocked,
		"none":       SecurityBlobContainerEvidenceRemediationStatusnone,
		"notfound":   SecurityBlobContainerEvidenceRemediationStatusnotFound,
		"prevented":  SecurityBlobContainerEvidenceRemediationStatusprevented,
		"remediated": SecurityBlobContainerEvidenceRemediationStatusremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBlobContainerEvidenceRemediationStatus(input)
	return &out, nil
}
