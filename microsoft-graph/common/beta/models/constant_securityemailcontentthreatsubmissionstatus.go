package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailContentThreatSubmissionStatus string

const (
	SecurityEmailContentThreatSubmissionStatusfailed     SecurityEmailContentThreatSubmissionStatus = "Failed"
	SecurityEmailContentThreatSubmissionStatusnotStarted SecurityEmailContentThreatSubmissionStatus = "NotStarted"
	SecurityEmailContentThreatSubmissionStatusrunning    SecurityEmailContentThreatSubmissionStatus = "Running"
	SecurityEmailContentThreatSubmissionStatusskipped    SecurityEmailContentThreatSubmissionStatus = "Skipped"
	SecurityEmailContentThreatSubmissionStatussucceeded  SecurityEmailContentThreatSubmissionStatus = "Succeeded"
)

func PossibleValuesForSecurityEmailContentThreatSubmissionStatus() []string {
	return []string{
		string(SecurityEmailContentThreatSubmissionStatusfailed),
		string(SecurityEmailContentThreatSubmissionStatusnotStarted),
		string(SecurityEmailContentThreatSubmissionStatusrunning),
		string(SecurityEmailContentThreatSubmissionStatusskipped),
		string(SecurityEmailContentThreatSubmissionStatussucceeded),
	}
}

func parseSecurityEmailContentThreatSubmissionStatus(input string) (*SecurityEmailContentThreatSubmissionStatus, error) {
	vals := map[string]SecurityEmailContentThreatSubmissionStatus{
		"failed":     SecurityEmailContentThreatSubmissionStatusfailed,
		"notstarted": SecurityEmailContentThreatSubmissionStatusnotStarted,
		"running":    SecurityEmailContentThreatSubmissionStatusrunning,
		"skipped":    SecurityEmailContentThreatSubmissionStatusskipped,
		"succeeded":  SecurityEmailContentThreatSubmissionStatussucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailContentThreatSubmissionStatus(input)
	return &out, nil
}
