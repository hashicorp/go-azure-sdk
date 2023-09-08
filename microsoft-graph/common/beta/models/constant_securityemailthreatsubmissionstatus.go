package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailThreatSubmissionStatus string

const (
	SecurityEmailThreatSubmissionStatusfailed     SecurityEmailThreatSubmissionStatus = "Failed"
	SecurityEmailThreatSubmissionStatusnotStarted SecurityEmailThreatSubmissionStatus = "NotStarted"
	SecurityEmailThreatSubmissionStatusrunning    SecurityEmailThreatSubmissionStatus = "Running"
	SecurityEmailThreatSubmissionStatusskipped    SecurityEmailThreatSubmissionStatus = "Skipped"
	SecurityEmailThreatSubmissionStatussucceeded  SecurityEmailThreatSubmissionStatus = "Succeeded"
)

func PossibleValuesForSecurityEmailThreatSubmissionStatus() []string {
	return []string{
		string(SecurityEmailThreatSubmissionStatusfailed),
		string(SecurityEmailThreatSubmissionStatusnotStarted),
		string(SecurityEmailThreatSubmissionStatusrunning),
		string(SecurityEmailThreatSubmissionStatusskipped),
		string(SecurityEmailThreatSubmissionStatussucceeded),
	}
}

func parseSecurityEmailThreatSubmissionStatus(input string) (*SecurityEmailThreatSubmissionStatus, error) {
	vals := map[string]SecurityEmailThreatSubmissionStatus{
		"failed":     SecurityEmailThreatSubmissionStatusfailed,
		"notstarted": SecurityEmailThreatSubmissionStatusnotStarted,
		"running":    SecurityEmailThreatSubmissionStatusrunning,
		"skipped":    SecurityEmailThreatSubmissionStatusskipped,
		"succeeded":  SecurityEmailThreatSubmissionStatussucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailThreatSubmissionStatus(input)
	return &out, nil
}
