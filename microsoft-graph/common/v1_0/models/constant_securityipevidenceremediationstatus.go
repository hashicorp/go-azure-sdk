package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIpEvidenceRemediationStatus string

const (
	SecurityIpEvidenceRemediationStatusblocked    SecurityIpEvidenceRemediationStatus = "Blocked"
	SecurityIpEvidenceRemediationStatusnone       SecurityIpEvidenceRemediationStatus = "None"
	SecurityIpEvidenceRemediationStatusnotFound   SecurityIpEvidenceRemediationStatus = "NotFound"
	SecurityIpEvidenceRemediationStatusprevented  SecurityIpEvidenceRemediationStatus = "Prevented"
	SecurityIpEvidenceRemediationStatusremediated SecurityIpEvidenceRemediationStatus = "Remediated"
)

func PossibleValuesForSecurityIpEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityIpEvidenceRemediationStatusblocked),
		string(SecurityIpEvidenceRemediationStatusnone),
		string(SecurityIpEvidenceRemediationStatusnotFound),
		string(SecurityIpEvidenceRemediationStatusprevented),
		string(SecurityIpEvidenceRemediationStatusremediated),
	}
}

func parseSecurityIpEvidenceRemediationStatus(input string) (*SecurityIpEvidenceRemediationStatus, error) {
	vals := map[string]SecurityIpEvidenceRemediationStatus{
		"blocked":    SecurityIpEvidenceRemediationStatusblocked,
		"none":       SecurityIpEvidenceRemediationStatusnone,
		"notfound":   SecurityIpEvidenceRemediationStatusnotFound,
		"prevented":  SecurityIpEvidenceRemediationStatusprevented,
		"remediated": SecurityIpEvidenceRemediationStatusremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityIpEvidenceRemediationStatus(input)
	return &out, nil
}
