package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUrlThreatSubmissionStatus string

const (
	SecurityUrlThreatSubmissionStatusfailed     SecurityUrlThreatSubmissionStatus = "Failed"
	SecurityUrlThreatSubmissionStatusnotStarted SecurityUrlThreatSubmissionStatus = "NotStarted"
	SecurityUrlThreatSubmissionStatusrunning    SecurityUrlThreatSubmissionStatus = "Running"
	SecurityUrlThreatSubmissionStatusskipped    SecurityUrlThreatSubmissionStatus = "Skipped"
	SecurityUrlThreatSubmissionStatussucceeded  SecurityUrlThreatSubmissionStatus = "Succeeded"
)

func PossibleValuesForSecurityUrlThreatSubmissionStatus() []string {
	return []string{
		string(SecurityUrlThreatSubmissionStatusfailed),
		string(SecurityUrlThreatSubmissionStatusnotStarted),
		string(SecurityUrlThreatSubmissionStatusrunning),
		string(SecurityUrlThreatSubmissionStatusskipped),
		string(SecurityUrlThreatSubmissionStatussucceeded),
	}
}

func parseSecurityUrlThreatSubmissionStatus(input string) (*SecurityUrlThreatSubmissionStatus, error) {
	vals := map[string]SecurityUrlThreatSubmissionStatus{
		"failed":     SecurityUrlThreatSubmissionStatusfailed,
		"notstarted": SecurityUrlThreatSubmissionStatusnotStarted,
		"running":    SecurityUrlThreatSubmissionStatusrunning,
		"skipped":    SecurityUrlThreatSubmissionStatusskipped,
		"succeeded":  SecurityUrlThreatSubmissionStatussucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityUrlThreatSubmissionStatus(input)
	return &out, nil
}
