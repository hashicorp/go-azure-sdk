package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySecurityGroupEvidenceRemediationStatus string

const (
	SecuritySecurityGroupEvidenceRemediationStatusblocked    SecuritySecurityGroupEvidenceRemediationStatus = "Blocked"
	SecuritySecurityGroupEvidenceRemediationStatusnone       SecuritySecurityGroupEvidenceRemediationStatus = "None"
	SecuritySecurityGroupEvidenceRemediationStatusnotFound   SecuritySecurityGroupEvidenceRemediationStatus = "NotFound"
	SecuritySecurityGroupEvidenceRemediationStatusprevented  SecuritySecurityGroupEvidenceRemediationStatus = "Prevented"
	SecuritySecurityGroupEvidenceRemediationStatusremediated SecuritySecurityGroupEvidenceRemediationStatus = "Remediated"
)

func PossibleValuesForSecuritySecurityGroupEvidenceRemediationStatus() []string {
	return []string{
		string(SecuritySecurityGroupEvidenceRemediationStatusblocked),
		string(SecuritySecurityGroupEvidenceRemediationStatusnone),
		string(SecuritySecurityGroupEvidenceRemediationStatusnotFound),
		string(SecuritySecurityGroupEvidenceRemediationStatusprevented),
		string(SecuritySecurityGroupEvidenceRemediationStatusremediated),
	}
}

func parseSecuritySecurityGroupEvidenceRemediationStatus(input string) (*SecuritySecurityGroupEvidenceRemediationStatus, error) {
	vals := map[string]SecuritySecurityGroupEvidenceRemediationStatus{
		"blocked":    SecuritySecurityGroupEvidenceRemediationStatusblocked,
		"none":       SecuritySecurityGroupEvidenceRemediationStatusnone,
		"notfound":   SecuritySecurityGroupEvidenceRemediationStatusnotFound,
		"prevented":  SecuritySecurityGroupEvidenceRemediationStatusprevented,
		"remediated": SecuritySecurityGroupEvidenceRemediationStatusremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecuritySecurityGroupEvidenceRemediationStatus(input)
	return &out, nil
}
