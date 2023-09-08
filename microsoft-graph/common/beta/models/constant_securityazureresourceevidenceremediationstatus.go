package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAzureResourceEvidenceRemediationStatus string

const (
	SecurityAzureResourceEvidenceRemediationStatusblocked    SecurityAzureResourceEvidenceRemediationStatus = "Blocked"
	SecurityAzureResourceEvidenceRemediationStatusnone       SecurityAzureResourceEvidenceRemediationStatus = "None"
	SecurityAzureResourceEvidenceRemediationStatusnotFound   SecurityAzureResourceEvidenceRemediationStatus = "NotFound"
	SecurityAzureResourceEvidenceRemediationStatusprevented  SecurityAzureResourceEvidenceRemediationStatus = "Prevented"
	SecurityAzureResourceEvidenceRemediationStatusremediated SecurityAzureResourceEvidenceRemediationStatus = "Remediated"
)

func PossibleValuesForSecurityAzureResourceEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityAzureResourceEvidenceRemediationStatusblocked),
		string(SecurityAzureResourceEvidenceRemediationStatusnone),
		string(SecurityAzureResourceEvidenceRemediationStatusnotFound),
		string(SecurityAzureResourceEvidenceRemediationStatusprevented),
		string(SecurityAzureResourceEvidenceRemediationStatusremediated),
	}
}

func parseSecurityAzureResourceEvidenceRemediationStatus(input string) (*SecurityAzureResourceEvidenceRemediationStatus, error) {
	vals := map[string]SecurityAzureResourceEvidenceRemediationStatus{
		"blocked":    SecurityAzureResourceEvidenceRemediationStatusblocked,
		"none":       SecurityAzureResourceEvidenceRemediationStatusnone,
		"notfound":   SecurityAzureResourceEvidenceRemediationStatusnotFound,
		"prevented":  SecurityAzureResourceEvidenceRemediationStatusprevented,
		"remediated": SecurityAzureResourceEvidenceRemediationStatusremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAzureResourceEvidenceRemediationStatus(input)
	return &out, nil
}
