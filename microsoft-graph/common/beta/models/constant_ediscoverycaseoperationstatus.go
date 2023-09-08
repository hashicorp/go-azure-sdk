package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryCaseOperationStatus string

const (
	EdiscoveryCaseOperationStatusfailed             EdiscoveryCaseOperationStatus = "Failed"
	EdiscoveryCaseOperationStatusnotStarted         EdiscoveryCaseOperationStatus = "NotStarted"
	EdiscoveryCaseOperationStatuspartiallySucceeded EdiscoveryCaseOperationStatus = "PartiallySucceeded"
	EdiscoveryCaseOperationStatusrunning            EdiscoveryCaseOperationStatus = "Running"
	EdiscoveryCaseOperationStatussubmissionFailed   EdiscoveryCaseOperationStatus = "SubmissionFailed"
	EdiscoveryCaseOperationStatussucceeded          EdiscoveryCaseOperationStatus = "Succeeded"
)

func PossibleValuesForEdiscoveryCaseOperationStatus() []string {
	return []string{
		string(EdiscoveryCaseOperationStatusfailed),
		string(EdiscoveryCaseOperationStatusnotStarted),
		string(EdiscoveryCaseOperationStatuspartiallySucceeded),
		string(EdiscoveryCaseOperationStatusrunning),
		string(EdiscoveryCaseOperationStatussubmissionFailed),
		string(EdiscoveryCaseOperationStatussucceeded),
	}
}

func parseEdiscoveryCaseOperationStatus(input string) (*EdiscoveryCaseOperationStatus, error) {
	vals := map[string]EdiscoveryCaseOperationStatus{
		"failed":             EdiscoveryCaseOperationStatusfailed,
		"notstarted":         EdiscoveryCaseOperationStatusnotStarted,
		"partiallysucceeded": EdiscoveryCaseOperationStatuspartiallySucceeded,
		"running":            EdiscoveryCaseOperationStatusrunning,
		"submissionfailed":   EdiscoveryCaseOperationStatussubmissionFailed,
		"succeeded":          EdiscoveryCaseOperationStatussucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryCaseOperationStatus(input)
	return &out, nil
}
