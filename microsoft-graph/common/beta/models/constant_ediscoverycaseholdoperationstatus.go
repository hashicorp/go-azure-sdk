package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryCaseHoldOperationStatus string

const (
	EdiscoveryCaseHoldOperationStatusfailed             EdiscoveryCaseHoldOperationStatus = "Failed"
	EdiscoveryCaseHoldOperationStatusnotStarted         EdiscoveryCaseHoldOperationStatus = "NotStarted"
	EdiscoveryCaseHoldOperationStatuspartiallySucceeded EdiscoveryCaseHoldOperationStatus = "PartiallySucceeded"
	EdiscoveryCaseHoldOperationStatusrunning            EdiscoveryCaseHoldOperationStatus = "Running"
	EdiscoveryCaseHoldOperationStatussubmissionFailed   EdiscoveryCaseHoldOperationStatus = "SubmissionFailed"
	EdiscoveryCaseHoldOperationStatussucceeded          EdiscoveryCaseHoldOperationStatus = "Succeeded"
)

func PossibleValuesForEdiscoveryCaseHoldOperationStatus() []string {
	return []string{
		string(EdiscoveryCaseHoldOperationStatusfailed),
		string(EdiscoveryCaseHoldOperationStatusnotStarted),
		string(EdiscoveryCaseHoldOperationStatuspartiallySucceeded),
		string(EdiscoveryCaseHoldOperationStatusrunning),
		string(EdiscoveryCaseHoldOperationStatussubmissionFailed),
		string(EdiscoveryCaseHoldOperationStatussucceeded),
	}
}

func parseEdiscoveryCaseHoldOperationStatus(input string) (*EdiscoveryCaseHoldOperationStatus, error) {
	vals := map[string]EdiscoveryCaseHoldOperationStatus{
		"failed":             EdiscoveryCaseHoldOperationStatusfailed,
		"notstarted":         EdiscoveryCaseHoldOperationStatusnotStarted,
		"partiallysucceeded": EdiscoveryCaseHoldOperationStatuspartiallySucceeded,
		"running":            EdiscoveryCaseHoldOperationStatusrunning,
		"submissionfailed":   EdiscoveryCaseHoldOperationStatussubmissionFailed,
		"succeeded":          EdiscoveryCaseHoldOperationStatussucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryCaseHoldOperationStatus(input)
	return &out, nil
}
