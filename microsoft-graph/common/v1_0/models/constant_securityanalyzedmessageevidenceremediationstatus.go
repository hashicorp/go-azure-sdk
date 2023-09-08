package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAnalyzedMessageEvidenceRemediationStatus string

const (
	SecurityAnalyzedMessageEvidenceRemediationStatusblocked    SecurityAnalyzedMessageEvidenceRemediationStatus = "Blocked"
	SecurityAnalyzedMessageEvidenceRemediationStatusnone       SecurityAnalyzedMessageEvidenceRemediationStatus = "None"
	SecurityAnalyzedMessageEvidenceRemediationStatusnotFound   SecurityAnalyzedMessageEvidenceRemediationStatus = "NotFound"
	SecurityAnalyzedMessageEvidenceRemediationStatusprevented  SecurityAnalyzedMessageEvidenceRemediationStatus = "Prevented"
	SecurityAnalyzedMessageEvidenceRemediationStatusremediated SecurityAnalyzedMessageEvidenceRemediationStatus = "Remediated"
)

func PossibleValuesForSecurityAnalyzedMessageEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityAnalyzedMessageEvidenceRemediationStatusblocked),
		string(SecurityAnalyzedMessageEvidenceRemediationStatusnone),
		string(SecurityAnalyzedMessageEvidenceRemediationStatusnotFound),
		string(SecurityAnalyzedMessageEvidenceRemediationStatusprevented),
		string(SecurityAnalyzedMessageEvidenceRemediationStatusremediated),
	}
}

func parseSecurityAnalyzedMessageEvidenceRemediationStatus(input string) (*SecurityAnalyzedMessageEvidenceRemediationStatus, error) {
	vals := map[string]SecurityAnalyzedMessageEvidenceRemediationStatus{
		"blocked":    SecurityAnalyzedMessageEvidenceRemediationStatusblocked,
		"none":       SecurityAnalyzedMessageEvidenceRemediationStatusnone,
		"notfound":   SecurityAnalyzedMessageEvidenceRemediationStatusnotFound,
		"prevented":  SecurityAnalyzedMessageEvidenceRemediationStatusprevented,
		"remediated": SecurityAnalyzedMessageEvidenceRemediationStatusremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAnalyzedMessageEvidenceRemediationStatus(input)
	return &out, nil
}
