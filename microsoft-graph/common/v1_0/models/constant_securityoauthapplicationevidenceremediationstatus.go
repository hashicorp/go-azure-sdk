package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityOauthApplicationEvidenceRemediationStatus string

const (
	SecurityOauthApplicationEvidenceRemediationStatusblocked    SecurityOauthApplicationEvidenceRemediationStatus = "Blocked"
	SecurityOauthApplicationEvidenceRemediationStatusnone       SecurityOauthApplicationEvidenceRemediationStatus = "None"
	SecurityOauthApplicationEvidenceRemediationStatusnotFound   SecurityOauthApplicationEvidenceRemediationStatus = "NotFound"
	SecurityOauthApplicationEvidenceRemediationStatusprevented  SecurityOauthApplicationEvidenceRemediationStatus = "Prevented"
	SecurityOauthApplicationEvidenceRemediationStatusremediated SecurityOauthApplicationEvidenceRemediationStatus = "Remediated"
)

func PossibleValuesForSecurityOauthApplicationEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityOauthApplicationEvidenceRemediationStatusblocked),
		string(SecurityOauthApplicationEvidenceRemediationStatusnone),
		string(SecurityOauthApplicationEvidenceRemediationStatusnotFound),
		string(SecurityOauthApplicationEvidenceRemediationStatusprevented),
		string(SecurityOauthApplicationEvidenceRemediationStatusremediated),
	}
}

func parseSecurityOauthApplicationEvidenceRemediationStatus(input string) (*SecurityOauthApplicationEvidenceRemediationStatus, error) {
	vals := map[string]SecurityOauthApplicationEvidenceRemediationStatus{
		"blocked":    SecurityOauthApplicationEvidenceRemediationStatusblocked,
		"none":       SecurityOauthApplicationEvidenceRemediationStatusnone,
		"notfound":   SecurityOauthApplicationEvidenceRemediationStatusnotFound,
		"prevented":  SecurityOauthApplicationEvidenceRemediationStatusprevented,
		"remediated": SecurityOauthApplicationEvidenceRemediationStatusremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityOauthApplicationEvidenceRemediationStatus(input)
	return &out, nil
}
