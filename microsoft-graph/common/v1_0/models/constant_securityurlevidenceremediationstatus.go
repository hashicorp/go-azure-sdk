package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUrlEvidenceRemediationStatus string

const (
	SecurityUrlEvidenceRemediationStatusblocked    SecurityUrlEvidenceRemediationStatus = "Blocked"
	SecurityUrlEvidenceRemediationStatusnone       SecurityUrlEvidenceRemediationStatus = "None"
	SecurityUrlEvidenceRemediationStatusnotFound   SecurityUrlEvidenceRemediationStatus = "NotFound"
	SecurityUrlEvidenceRemediationStatusprevented  SecurityUrlEvidenceRemediationStatus = "Prevented"
	SecurityUrlEvidenceRemediationStatusremediated SecurityUrlEvidenceRemediationStatus = "Remediated"
)

func PossibleValuesForSecurityUrlEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityUrlEvidenceRemediationStatusblocked),
		string(SecurityUrlEvidenceRemediationStatusnone),
		string(SecurityUrlEvidenceRemediationStatusnotFound),
		string(SecurityUrlEvidenceRemediationStatusprevented),
		string(SecurityUrlEvidenceRemediationStatusremediated),
	}
}

func parseSecurityUrlEvidenceRemediationStatus(input string) (*SecurityUrlEvidenceRemediationStatus, error) {
	vals := map[string]SecurityUrlEvidenceRemediationStatus{
		"blocked":    SecurityUrlEvidenceRemediationStatusblocked,
		"none":       SecurityUrlEvidenceRemediationStatusnone,
		"notfound":   SecurityUrlEvidenceRemediationStatusnotFound,
		"prevented":  SecurityUrlEvidenceRemediationStatusprevented,
		"remediated": SecurityUrlEvidenceRemediationStatusremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityUrlEvidenceRemediationStatus(input)
	return &out, nil
}
