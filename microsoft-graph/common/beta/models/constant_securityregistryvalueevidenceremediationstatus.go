package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRegistryValueEvidenceRemediationStatus string

const (
	SecurityRegistryValueEvidenceRemediationStatusblocked    SecurityRegistryValueEvidenceRemediationStatus = "Blocked"
	SecurityRegistryValueEvidenceRemediationStatusnone       SecurityRegistryValueEvidenceRemediationStatus = "None"
	SecurityRegistryValueEvidenceRemediationStatusnotFound   SecurityRegistryValueEvidenceRemediationStatus = "NotFound"
	SecurityRegistryValueEvidenceRemediationStatusprevented  SecurityRegistryValueEvidenceRemediationStatus = "Prevented"
	SecurityRegistryValueEvidenceRemediationStatusremediated SecurityRegistryValueEvidenceRemediationStatus = "Remediated"
)

func PossibleValuesForSecurityRegistryValueEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityRegistryValueEvidenceRemediationStatusblocked),
		string(SecurityRegistryValueEvidenceRemediationStatusnone),
		string(SecurityRegistryValueEvidenceRemediationStatusnotFound),
		string(SecurityRegistryValueEvidenceRemediationStatusprevented),
		string(SecurityRegistryValueEvidenceRemediationStatusremediated),
	}
}

func parseSecurityRegistryValueEvidenceRemediationStatus(input string) (*SecurityRegistryValueEvidenceRemediationStatus, error) {
	vals := map[string]SecurityRegistryValueEvidenceRemediationStatus{
		"blocked":    SecurityRegistryValueEvidenceRemediationStatusblocked,
		"none":       SecurityRegistryValueEvidenceRemediationStatusnone,
		"notfound":   SecurityRegistryValueEvidenceRemediationStatusnotFound,
		"prevented":  SecurityRegistryValueEvidenceRemediationStatusprevented,
		"remediated": SecurityRegistryValueEvidenceRemediationStatusremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRegistryValueEvidenceRemediationStatus(input)
	return &out, nil
}
