package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContainerImageEvidenceRemediationStatus string

const (
	SecurityContainerImageEvidenceRemediationStatusblocked    SecurityContainerImageEvidenceRemediationStatus = "Blocked"
	SecurityContainerImageEvidenceRemediationStatusnone       SecurityContainerImageEvidenceRemediationStatus = "None"
	SecurityContainerImageEvidenceRemediationStatusnotFound   SecurityContainerImageEvidenceRemediationStatus = "NotFound"
	SecurityContainerImageEvidenceRemediationStatusprevented  SecurityContainerImageEvidenceRemediationStatus = "Prevented"
	SecurityContainerImageEvidenceRemediationStatusremediated SecurityContainerImageEvidenceRemediationStatus = "Remediated"
)

func PossibleValuesForSecurityContainerImageEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityContainerImageEvidenceRemediationStatusblocked),
		string(SecurityContainerImageEvidenceRemediationStatusnone),
		string(SecurityContainerImageEvidenceRemediationStatusnotFound),
		string(SecurityContainerImageEvidenceRemediationStatusprevented),
		string(SecurityContainerImageEvidenceRemediationStatusremediated),
	}
}

func parseSecurityContainerImageEvidenceRemediationStatus(input string) (*SecurityContainerImageEvidenceRemediationStatus, error) {
	vals := map[string]SecurityContainerImageEvidenceRemediationStatus{
		"blocked":    SecurityContainerImageEvidenceRemediationStatusblocked,
		"none":       SecurityContainerImageEvidenceRemediationStatusnone,
		"notfound":   SecurityContainerImageEvidenceRemediationStatusnotFound,
		"prevented":  SecurityContainerImageEvidenceRemediationStatusprevented,
		"remediated": SecurityContainerImageEvidenceRemediationStatusremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityContainerImageEvidenceRemediationStatus(input)
	return &out, nil
}
