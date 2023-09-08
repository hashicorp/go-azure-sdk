package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityProcessEvidenceDetectionStatus string

const (
	SecurityProcessEvidenceDetectionStatusblocked   SecurityProcessEvidenceDetectionStatus = "Blocked"
	SecurityProcessEvidenceDetectionStatusdetected  SecurityProcessEvidenceDetectionStatus = "Detected"
	SecurityProcessEvidenceDetectionStatusprevented SecurityProcessEvidenceDetectionStatus = "Prevented"
)

func PossibleValuesForSecurityProcessEvidenceDetectionStatus() []string {
	return []string{
		string(SecurityProcessEvidenceDetectionStatusblocked),
		string(SecurityProcessEvidenceDetectionStatusdetected),
		string(SecurityProcessEvidenceDetectionStatusprevented),
	}
}

func parseSecurityProcessEvidenceDetectionStatus(input string) (*SecurityProcessEvidenceDetectionStatus, error) {
	vals := map[string]SecurityProcessEvidenceDetectionStatus{
		"blocked":   SecurityProcessEvidenceDetectionStatusblocked,
		"detected":  SecurityProcessEvidenceDetectionStatusdetected,
		"prevented": SecurityProcessEvidenceDetectionStatusprevented,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityProcessEvidenceDetectionStatus(input)
	return &out, nil
}
