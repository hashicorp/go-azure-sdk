package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityCaseOperationStatus string

const (
	SecurityCaseOperationStatusfailed             SecurityCaseOperationStatus = "Failed"
	SecurityCaseOperationStatusnotStarted         SecurityCaseOperationStatus = "NotStarted"
	SecurityCaseOperationStatuspartiallySucceeded SecurityCaseOperationStatus = "PartiallySucceeded"
	SecurityCaseOperationStatusrunning            SecurityCaseOperationStatus = "Running"
	SecurityCaseOperationStatussubmissionFailed   SecurityCaseOperationStatus = "SubmissionFailed"
	SecurityCaseOperationStatussucceeded          SecurityCaseOperationStatus = "Succeeded"
)

func PossibleValuesForSecurityCaseOperationStatus() []string {
	return []string{
		string(SecurityCaseOperationStatusfailed),
		string(SecurityCaseOperationStatusnotStarted),
		string(SecurityCaseOperationStatuspartiallySucceeded),
		string(SecurityCaseOperationStatusrunning),
		string(SecurityCaseOperationStatussubmissionFailed),
		string(SecurityCaseOperationStatussucceeded),
	}
}

func parseSecurityCaseOperationStatus(input string) (*SecurityCaseOperationStatus, error) {
	vals := map[string]SecurityCaseOperationStatus{
		"failed":             SecurityCaseOperationStatusfailed,
		"notstarted":         SecurityCaseOperationStatusnotStarted,
		"partiallysucceeded": SecurityCaseOperationStatuspartiallySucceeded,
		"running":            SecurityCaseOperationStatusrunning,
		"submissionfailed":   SecurityCaseOperationStatussubmissionFailed,
		"succeeded":          SecurityCaseOperationStatussucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityCaseOperationStatus(input)
	return &out, nil
}
