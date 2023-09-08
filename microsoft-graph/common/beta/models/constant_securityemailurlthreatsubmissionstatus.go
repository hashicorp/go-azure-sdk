package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailUrlThreatSubmissionStatus string

const (
	SecurityEmailUrlThreatSubmissionStatusfailed     SecurityEmailUrlThreatSubmissionStatus = "Failed"
	SecurityEmailUrlThreatSubmissionStatusnotStarted SecurityEmailUrlThreatSubmissionStatus = "NotStarted"
	SecurityEmailUrlThreatSubmissionStatusrunning    SecurityEmailUrlThreatSubmissionStatus = "Running"
	SecurityEmailUrlThreatSubmissionStatusskipped    SecurityEmailUrlThreatSubmissionStatus = "Skipped"
	SecurityEmailUrlThreatSubmissionStatussucceeded  SecurityEmailUrlThreatSubmissionStatus = "Succeeded"
)

func PossibleValuesForSecurityEmailUrlThreatSubmissionStatus() []string {
	return []string{
		string(SecurityEmailUrlThreatSubmissionStatusfailed),
		string(SecurityEmailUrlThreatSubmissionStatusnotStarted),
		string(SecurityEmailUrlThreatSubmissionStatusrunning),
		string(SecurityEmailUrlThreatSubmissionStatusskipped),
		string(SecurityEmailUrlThreatSubmissionStatussucceeded),
	}
}

func parseSecurityEmailUrlThreatSubmissionStatus(input string) (*SecurityEmailUrlThreatSubmissionStatus, error) {
	vals := map[string]SecurityEmailUrlThreatSubmissionStatus{
		"failed":     SecurityEmailUrlThreatSubmissionStatusfailed,
		"notstarted": SecurityEmailUrlThreatSubmissionStatusnotStarted,
		"running":    SecurityEmailUrlThreatSubmissionStatusrunning,
		"skipped":    SecurityEmailUrlThreatSubmissionStatusskipped,
		"succeeded":  SecurityEmailUrlThreatSubmissionStatussucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailUrlThreatSubmissionStatus(input)
	return &out, nil
}
