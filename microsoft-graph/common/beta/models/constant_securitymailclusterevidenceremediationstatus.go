package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityMailClusterEvidenceRemediationStatus string

const (
	SecurityMailClusterEvidenceRemediationStatusblocked    SecurityMailClusterEvidenceRemediationStatus = "Blocked"
	SecurityMailClusterEvidenceRemediationStatusnone       SecurityMailClusterEvidenceRemediationStatus = "None"
	SecurityMailClusterEvidenceRemediationStatusnotFound   SecurityMailClusterEvidenceRemediationStatus = "NotFound"
	SecurityMailClusterEvidenceRemediationStatusprevented  SecurityMailClusterEvidenceRemediationStatus = "Prevented"
	SecurityMailClusterEvidenceRemediationStatusremediated SecurityMailClusterEvidenceRemediationStatus = "Remediated"
)

func PossibleValuesForSecurityMailClusterEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityMailClusterEvidenceRemediationStatusblocked),
		string(SecurityMailClusterEvidenceRemediationStatusnone),
		string(SecurityMailClusterEvidenceRemediationStatusnotFound),
		string(SecurityMailClusterEvidenceRemediationStatusprevented),
		string(SecurityMailClusterEvidenceRemediationStatusremediated),
	}
}

func parseSecurityMailClusterEvidenceRemediationStatus(input string) (*SecurityMailClusterEvidenceRemediationStatus, error) {
	vals := map[string]SecurityMailClusterEvidenceRemediationStatus{
		"blocked":    SecurityMailClusterEvidenceRemediationStatusblocked,
		"none":       SecurityMailClusterEvidenceRemediationStatusnone,
		"notfound":   SecurityMailClusterEvidenceRemediationStatusnotFound,
		"prevented":  SecurityMailClusterEvidenceRemediationStatusprevented,
		"remediated": SecurityMailClusterEvidenceRemediationStatusremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityMailClusterEvidenceRemediationStatus(input)
	return &out, nil
}
