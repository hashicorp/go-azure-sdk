package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRegistryKeyEvidenceRemediationStatus string

const (
	SecurityRegistryKeyEvidenceRemediationStatusblocked    SecurityRegistryKeyEvidenceRemediationStatus = "Blocked"
	SecurityRegistryKeyEvidenceRemediationStatusnone       SecurityRegistryKeyEvidenceRemediationStatus = "None"
	SecurityRegistryKeyEvidenceRemediationStatusnotFound   SecurityRegistryKeyEvidenceRemediationStatus = "NotFound"
	SecurityRegistryKeyEvidenceRemediationStatusprevented  SecurityRegistryKeyEvidenceRemediationStatus = "Prevented"
	SecurityRegistryKeyEvidenceRemediationStatusremediated SecurityRegistryKeyEvidenceRemediationStatus = "Remediated"
)

func PossibleValuesForSecurityRegistryKeyEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityRegistryKeyEvidenceRemediationStatusblocked),
		string(SecurityRegistryKeyEvidenceRemediationStatusnone),
		string(SecurityRegistryKeyEvidenceRemediationStatusnotFound),
		string(SecurityRegistryKeyEvidenceRemediationStatusprevented),
		string(SecurityRegistryKeyEvidenceRemediationStatusremediated),
	}
}

func parseSecurityRegistryKeyEvidenceRemediationStatus(input string) (*SecurityRegistryKeyEvidenceRemediationStatus, error) {
	vals := map[string]SecurityRegistryKeyEvidenceRemediationStatus{
		"blocked":    SecurityRegistryKeyEvidenceRemediationStatusblocked,
		"none":       SecurityRegistryKeyEvidenceRemediationStatusnone,
		"notfound":   SecurityRegistryKeyEvidenceRemediationStatusnotFound,
		"prevented":  SecurityRegistryKeyEvidenceRemediationStatusprevented,
		"remediated": SecurityRegistryKeyEvidenceRemediationStatusremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRegistryKeyEvidenceRemediationStatus(input)
	return &out, nil
}
