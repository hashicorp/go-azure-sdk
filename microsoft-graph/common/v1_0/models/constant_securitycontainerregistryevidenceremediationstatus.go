package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContainerRegistryEvidenceRemediationStatus string

const (
	SecurityContainerRegistryEvidenceRemediationStatusblocked    SecurityContainerRegistryEvidenceRemediationStatus = "Blocked"
	SecurityContainerRegistryEvidenceRemediationStatusnone       SecurityContainerRegistryEvidenceRemediationStatus = "None"
	SecurityContainerRegistryEvidenceRemediationStatusnotFound   SecurityContainerRegistryEvidenceRemediationStatus = "NotFound"
	SecurityContainerRegistryEvidenceRemediationStatusprevented  SecurityContainerRegistryEvidenceRemediationStatus = "Prevented"
	SecurityContainerRegistryEvidenceRemediationStatusremediated SecurityContainerRegistryEvidenceRemediationStatus = "Remediated"
)

func PossibleValuesForSecurityContainerRegistryEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityContainerRegistryEvidenceRemediationStatusblocked),
		string(SecurityContainerRegistryEvidenceRemediationStatusnone),
		string(SecurityContainerRegistryEvidenceRemediationStatusnotFound),
		string(SecurityContainerRegistryEvidenceRemediationStatusprevented),
		string(SecurityContainerRegistryEvidenceRemediationStatusremediated),
	}
}

func parseSecurityContainerRegistryEvidenceRemediationStatus(input string) (*SecurityContainerRegistryEvidenceRemediationStatus, error) {
	vals := map[string]SecurityContainerRegistryEvidenceRemediationStatus{
		"blocked":    SecurityContainerRegistryEvidenceRemediationStatusblocked,
		"none":       SecurityContainerRegistryEvidenceRemediationStatusnone,
		"notfound":   SecurityContainerRegistryEvidenceRemediationStatusnotFound,
		"prevented":  SecurityContainerRegistryEvidenceRemediationStatusprevented,
		"remediated": SecurityContainerRegistryEvidenceRemediationStatusremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityContainerRegistryEvidenceRemediationStatus(input)
	return &out, nil
}
