package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDeviceEvidenceRemediationStatus string

const (
	SecurityDeviceEvidenceRemediationStatusblocked    SecurityDeviceEvidenceRemediationStatus = "Blocked"
	SecurityDeviceEvidenceRemediationStatusnone       SecurityDeviceEvidenceRemediationStatus = "None"
	SecurityDeviceEvidenceRemediationStatusnotFound   SecurityDeviceEvidenceRemediationStatus = "NotFound"
	SecurityDeviceEvidenceRemediationStatusprevented  SecurityDeviceEvidenceRemediationStatus = "Prevented"
	SecurityDeviceEvidenceRemediationStatusremediated SecurityDeviceEvidenceRemediationStatus = "Remediated"
)

func PossibleValuesForSecurityDeviceEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityDeviceEvidenceRemediationStatusblocked),
		string(SecurityDeviceEvidenceRemediationStatusnone),
		string(SecurityDeviceEvidenceRemediationStatusnotFound),
		string(SecurityDeviceEvidenceRemediationStatusprevented),
		string(SecurityDeviceEvidenceRemediationStatusremediated),
	}
}

func parseSecurityDeviceEvidenceRemediationStatus(input string) (*SecurityDeviceEvidenceRemediationStatus, error) {
	vals := map[string]SecurityDeviceEvidenceRemediationStatus{
		"blocked":    SecurityDeviceEvidenceRemediationStatusblocked,
		"none":       SecurityDeviceEvidenceRemediationStatusnone,
		"notfound":   SecurityDeviceEvidenceRemediationStatusnotFound,
		"prevented":  SecurityDeviceEvidenceRemediationStatusprevented,
		"remediated": SecurityDeviceEvidenceRemediationStatusremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDeviceEvidenceRemediationStatus(input)
	return &out, nil
}
