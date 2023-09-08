package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAmazonResourceEvidenceRemediationStatus string

const (
	SecurityAmazonResourceEvidenceRemediationStatusblocked    SecurityAmazonResourceEvidenceRemediationStatus = "Blocked"
	SecurityAmazonResourceEvidenceRemediationStatusnone       SecurityAmazonResourceEvidenceRemediationStatus = "None"
	SecurityAmazonResourceEvidenceRemediationStatusnotFound   SecurityAmazonResourceEvidenceRemediationStatus = "NotFound"
	SecurityAmazonResourceEvidenceRemediationStatusprevented  SecurityAmazonResourceEvidenceRemediationStatus = "Prevented"
	SecurityAmazonResourceEvidenceRemediationStatusremediated SecurityAmazonResourceEvidenceRemediationStatus = "Remediated"
)

func PossibleValuesForSecurityAmazonResourceEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityAmazonResourceEvidenceRemediationStatusblocked),
		string(SecurityAmazonResourceEvidenceRemediationStatusnone),
		string(SecurityAmazonResourceEvidenceRemediationStatusnotFound),
		string(SecurityAmazonResourceEvidenceRemediationStatusprevented),
		string(SecurityAmazonResourceEvidenceRemediationStatusremediated),
	}
}

func parseSecurityAmazonResourceEvidenceRemediationStatus(input string) (*SecurityAmazonResourceEvidenceRemediationStatus, error) {
	vals := map[string]SecurityAmazonResourceEvidenceRemediationStatus{
		"blocked":    SecurityAmazonResourceEvidenceRemediationStatusblocked,
		"none":       SecurityAmazonResourceEvidenceRemediationStatusnone,
		"notfound":   SecurityAmazonResourceEvidenceRemediationStatusnotFound,
		"prevented":  SecurityAmazonResourceEvidenceRemediationStatusprevented,
		"remediated": SecurityAmazonResourceEvidenceRemediationStatusremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAmazonResourceEvidenceRemediationStatus(input)
	return &out, nil
}
