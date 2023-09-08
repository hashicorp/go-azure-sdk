package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileContentThreatSubmissionStatus string

const (
	SecurityFileContentThreatSubmissionStatusfailed     SecurityFileContentThreatSubmissionStatus = "Failed"
	SecurityFileContentThreatSubmissionStatusnotStarted SecurityFileContentThreatSubmissionStatus = "NotStarted"
	SecurityFileContentThreatSubmissionStatusrunning    SecurityFileContentThreatSubmissionStatus = "Running"
	SecurityFileContentThreatSubmissionStatusskipped    SecurityFileContentThreatSubmissionStatus = "Skipped"
	SecurityFileContentThreatSubmissionStatussucceeded  SecurityFileContentThreatSubmissionStatus = "Succeeded"
)

func PossibleValuesForSecurityFileContentThreatSubmissionStatus() []string {
	return []string{
		string(SecurityFileContentThreatSubmissionStatusfailed),
		string(SecurityFileContentThreatSubmissionStatusnotStarted),
		string(SecurityFileContentThreatSubmissionStatusrunning),
		string(SecurityFileContentThreatSubmissionStatusskipped),
		string(SecurityFileContentThreatSubmissionStatussucceeded),
	}
}

func parseSecurityFileContentThreatSubmissionStatus(input string) (*SecurityFileContentThreatSubmissionStatus, error) {
	vals := map[string]SecurityFileContentThreatSubmissionStatus{
		"failed":     SecurityFileContentThreatSubmissionStatusfailed,
		"notstarted": SecurityFileContentThreatSubmissionStatusnotStarted,
		"running":    SecurityFileContentThreatSubmissionStatusrunning,
		"skipped":    SecurityFileContentThreatSubmissionStatusskipped,
		"succeeded":  SecurityFileContentThreatSubmissionStatussucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileContentThreatSubmissionStatus(input)
	return &out, nil
}
