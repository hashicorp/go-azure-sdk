package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityThreatSubmissionStatus string

const (
	SecurityThreatSubmissionStatusfailed     SecurityThreatSubmissionStatus = "Failed"
	SecurityThreatSubmissionStatusnotStarted SecurityThreatSubmissionStatus = "NotStarted"
	SecurityThreatSubmissionStatusrunning    SecurityThreatSubmissionStatus = "Running"
	SecurityThreatSubmissionStatusskipped    SecurityThreatSubmissionStatus = "Skipped"
	SecurityThreatSubmissionStatussucceeded  SecurityThreatSubmissionStatus = "Succeeded"
)

func PossibleValuesForSecurityThreatSubmissionStatus() []string {
	return []string{
		string(SecurityThreatSubmissionStatusfailed),
		string(SecurityThreatSubmissionStatusnotStarted),
		string(SecurityThreatSubmissionStatusrunning),
		string(SecurityThreatSubmissionStatusskipped),
		string(SecurityThreatSubmissionStatussucceeded),
	}
}

func parseSecurityThreatSubmissionStatus(input string) (*SecurityThreatSubmissionStatus, error) {
	vals := map[string]SecurityThreatSubmissionStatus{
		"failed":     SecurityThreatSubmissionStatusfailed,
		"notstarted": SecurityThreatSubmissionStatusnotStarted,
		"running":    SecurityThreatSubmissionStatusrunning,
		"skipped":    SecurityThreatSubmissionStatusskipped,
		"succeeded":  SecurityThreatSubmissionStatussucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityThreatSubmissionStatus(input)
	return &out, nil
}
