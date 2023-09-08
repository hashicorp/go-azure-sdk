package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityCloudApplicationEvidenceRemediationStatus string

const (
	SecurityCloudApplicationEvidenceRemediationStatusblocked    SecurityCloudApplicationEvidenceRemediationStatus = "Blocked"
	SecurityCloudApplicationEvidenceRemediationStatusnone       SecurityCloudApplicationEvidenceRemediationStatus = "None"
	SecurityCloudApplicationEvidenceRemediationStatusnotFound   SecurityCloudApplicationEvidenceRemediationStatus = "NotFound"
	SecurityCloudApplicationEvidenceRemediationStatusprevented  SecurityCloudApplicationEvidenceRemediationStatus = "Prevented"
	SecurityCloudApplicationEvidenceRemediationStatusremediated SecurityCloudApplicationEvidenceRemediationStatus = "Remediated"
)

func PossibleValuesForSecurityCloudApplicationEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityCloudApplicationEvidenceRemediationStatusblocked),
		string(SecurityCloudApplicationEvidenceRemediationStatusnone),
		string(SecurityCloudApplicationEvidenceRemediationStatusnotFound),
		string(SecurityCloudApplicationEvidenceRemediationStatusprevented),
		string(SecurityCloudApplicationEvidenceRemediationStatusremediated),
	}
}

func parseSecurityCloudApplicationEvidenceRemediationStatus(input string) (*SecurityCloudApplicationEvidenceRemediationStatus, error) {
	vals := map[string]SecurityCloudApplicationEvidenceRemediationStatus{
		"blocked":    SecurityCloudApplicationEvidenceRemediationStatusblocked,
		"none":       SecurityCloudApplicationEvidenceRemediationStatusnone,
		"notfound":   SecurityCloudApplicationEvidenceRemediationStatusnotFound,
		"prevented":  SecurityCloudApplicationEvidenceRemediationStatusprevented,
		"remediated": SecurityCloudApplicationEvidenceRemediationStatusremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityCloudApplicationEvidenceRemediationStatus(input)
	return &out, nil
}
