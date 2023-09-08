package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileUrlThreatSubmissionStatus string

const (
	SecurityFileUrlThreatSubmissionStatusfailed     SecurityFileUrlThreatSubmissionStatus = "Failed"
	SecurityFileUrlThreatSubmissionStatusnotStarted SecurityFileUrlThreatSubmissionStatus = "NotStarted"
	SecurityFileUrlThreatSubmissionStatusrunning    SecurityFileUrlThreatSubmissionStatus = "Running"
	SecurityFileUrlThreatSubmissionStatusskipped    SecurityFileUrlThreatSubmissionStatus = "Skipped"
	SecurityFileUrlThreatSubmissionStatussucceeded  SecurityFileUrlThreatSubmissionStatus = "Succeeded"
)

func PossibleValuesForSecurityFileUrlThreatSubmissionStatus() []string {
	return []string{
		string(SecurityFileUrlThreatSubmissionStatusfailed),
		string(SecurityFileUrlThreatSubmissionStatusnotStarted),
		string(SecurityFileUrlThreatSubmissionStatusrunning),
		string(SecurityFileUrlThreatSubmissionStatusskipped),
		string(SecurityFileUrlThreatSubmissionStatussucceeded),
	}
}

func parseSecurityFileUrlThreatSubmissionStatus(input string) (*SecurityFileUrlThreatSubmissionStatus, error) {
	vals := map[string]SecurityFileUrlThreatSubmissionStatus{
		"failed":     SecurityFileUrlThreatSubmissionStatusfailed,
		"notstarted": SecurityFileUrlThreatSubmissionStatusnotStarted,
		"running":    SecurityFileUrlThreatSubmissionStatusrunning,
		"skipped":    SecurityFileUrlThreatSubmissionStatusskipped,
		"succeeded":  SecurityFileUrlThreatSubmissionStatussucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileUrlThreatSubmissionStatus(input)
	return &out, nil
}
