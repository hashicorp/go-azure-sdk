package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAlertEvidenceRemediationStatus string

const (
	SecurityAlertEvidenceRemediationStatusblocked    SecurityAlertEvidenceRemediationStatus = "Blocked"
	SecurityAlertEvidenceRemediationStatusnone       SecurityAlertEvidenceRemediationStatus = "None"
	SecurityAlertEvidenceRemediationStatusnotFound   SecurityAlertEvidenceRemediationStatus = "NotFound"
	SecurityAlertEvidenceRemediationStatusprevented  SecurityAlertEvidenceRemediationStatus = "Prevented"
	SecurityAlertEvidenceRemediationStatusremediated SecurityAlertEvidenceRemediationStatus = "Remediated"
)

func PossibleValuesForSecurityAlertEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityAlertEvidenceRemediationStatusblocked),
		string(SecurityAlertEvidenceRemediationStatusnone),
		string(SecurityAlertEvidenceRemediationStatusnotFound),
		string(SecurityAlertEvidenceRemediationStatusprevented),
		string(SecurityAlertEvidenceRemediationStatusremediated),
	}
}

func parseSecurityAlertEvidenceRemediationStatus(input string) (*SecurityAlertEvidenceRemediationStatus, error) {
	vals := map[string]SecurityAlertEvidenceRemediationStatus{
		"blocked":    SecurityAlertEvidenceRemediationStatusblocked,
		"none":       SecurityAlertEvidenceRemediationStatusnone,
		"notfound":   SecurityAlertEvidenceRemediationStatusnotFound,
		"prevented":  SecurityAlertEvidenceRemediationStatusprevented,
		"remediated": SecurityAlertEvidenceRemediationStatusremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAlertEvidenceRemediationStatus(input)
	return &out, nil
}
