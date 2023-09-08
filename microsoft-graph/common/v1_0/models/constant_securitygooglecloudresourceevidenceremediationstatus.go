package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityGoogleCloudResourceEvidenceRemediationStatus string

const (
	SecurityGoogleCloudResourceEvidenceRemediationStatusblocked    SecurityGoogleCloudResourceEvidenceRemediationStatus = "Blocked"
	SecurityGoogleCloudResourceEvidenceRemediationStatusnone       SecurityGoogleCloudResourceEvidenceRemediationStatus = "None"
	SecurityGoogleCloudResourceEvidenceRemediationStatusnotFound   SecurityGoogleCloudResourceEvidenceRemediationStatus = "NotFound"
	SecurityGoogleCloudResourceEvidenceRemediationStatusprevented  SecurityGoogleCloudResourceEvidenceRemediationStatus = "Prevented"
	SecurityGoogleCloudResourceEvidenceRemediationStatusremediated SecurityGoogleCloudResourceEvidenceRemediationStatus = "Remediated"
)

func PossibleValuesForSecurityGoogleCloudResourceEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityGoogleCloudResourceEvidenceRemediationStatusblocked),
		string(SecurityGoogleCloudResourceEvidenceRemediationStatusnone),
		string(SecurityGoogleCloudResourceEvidenceRemediationStatusnotFound),
		string(SecurityGoogleCloudResourceEvidenceRemediationStatusprevented),
		string(SecurityGoogleCloudResourceEvidenceRemediationStatusremediated),
	}
}

func parseSecurityGoogleCloudResourceEvidenceRemediationStatus(input string) (*SecurityGoogleCloudResourceEvidenceRemediationStatus, error) {
	vals := map[string]SecurityGoogleCloudResourceEvidenceRemediationStatus{
		"blocked":    SecurityGoogleCloudResourceEvidenceRemediationStatusblocked,
		"none":       SecurityGoogleCloudResourceEvidenceRemediationStatusnone,
		"notfound":   SecurityGoogleCloudResourceEvidenceRemediationStatusnotFound,
		"prevented":  SecurityGoogleCloudResourceEvidenceRemediationStatusprevented,
		"remediated": SecurityGoogleCloudResourceEvidenceRemediationStatusremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityGoogleCloudResourceEvidenceRemediationStatus(input)
	return &out, nil
}
