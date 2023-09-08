package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryCaseIndexOperationStatus string

const (
	EdiscoveryCaseIndexOperationStatusfailed             EdiscoveryCaseIndexOperationStatus = "Failed"
	EdiscoveryCaseIndexOperationStatusnotStarted         EdiscoveryCaseIndexOperationStatus = "NotStarted"
	EdiscoveryCaseIndexOperationStatuspartiallySucceeded EdiscoveryCaseIndexOperationStatus = "PartiallySucceeded"
	EdiscoveryCaseIndexOperationStatusrunning            EdiscoveryCaseIndexOperationStatus = "Running"
	EdiscoveryCaseIndexOperationStatussubmissionFailed   EdiscoveryCaseIndexOperationStatus = "SubmissionFailed"
	EdiscoveryCaseIndexOperationStatussucceeded          EdiscoveryCaseIndexOperationStatus = "Succeeded"
)

func PossibleValuesForEdiscoveryCaseIndexOperationStatus() []string {
	return []string{
		string(EdiscoveryCaseIndexOperationStatusfailed),
		string(EdiscoveryCaseIndexOperationStatusnotStarted),
		string(EdiscoveryCaseIndexOperationStatuspartiallySucceeded),
		string(EdiscoveryCaseIndexOperationStatusrunning),
		string(EdiscoveryCaseIndexOperationStatussubmissionFailed),
		string(EdiscoveryCaseIndexOperationStatussucceeded),
	}
}

func parseEdiscoveryCaseIndexOperationStatus(input string) (*EdiscoveryCaseIndexOperationStatus, error) {
	vals := map[string]EdiscoveryCaseIndexOperationStatus{
		"failed":             EdiscoveryCaseIndexOperationStatusfailed,
		"notstarted":         EdiscoveryCaseIndexOperationStatusnotStarted,
		"partiallysucceeded": EdiscoveryCaseIndexOperationStatuspartiallySucceeded,
		"running":            EdiscoveryCaseIndexOperationStatusrunning,
		"submissionfailed":   EdiscoveryCaseIndexOperationStatussubmissionFailed,
		"succeeded":          EdiscoveryCaseIndexOperationStatussucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryCaseIndexOperationStatus(input)
	return &out, nil
}
