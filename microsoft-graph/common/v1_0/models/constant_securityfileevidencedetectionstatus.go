package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileEvidenceDetectionStatus string

const (
	SecurityFileEvidenceDetectionStatusblocked   SecurityFileEvidenceDetectionStatus = "Blocked"
	SecurityFileEvidenceDetectionStatusdetected  SecurityFileEvidenceDetectionStatus = "Detected"
	SecurityFileEvidenceDetectionStatusprevented SecurityFileEvidenceDetectionStatus = "Prevented"
)

func PossibleValuesForSecurityFileEvidenceDetectionStatus() []string {
	return []string{
		string(SecurityFileEvidenceDetectionStatusblocked),
		string(SecurityFileEvidenceDetectionStatusdetected),
		string(SecurityFileEvidenceDetectionStatusprevented),
	}
}

func parseSecurityFileEvidenceDetectionStatus(input string) (*SecurityFileEvidenceDetectionStatus, error) {
	vals := map[string]SecurityFileEvidenceDetectionStatus{
		"blocked":   SecurityFileEvidenceDetectionStatusblocked,
		"detected":  SecurityFileEvidenceDetectionStatusdetected,
		"prevented": SecurityFileEvidenceDetectionStatusprevented,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileEvidenceDetectionStatus(input)
	return &out, nil
}
