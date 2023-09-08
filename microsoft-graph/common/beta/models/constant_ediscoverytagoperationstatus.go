package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryTagOperationStatus string

const (
	EdiscoveryTagOperationStatusfailed             EdiscoveryTagOperationStatus = "Failed"
	EdiscoveryTagOperationStatusnotStarted         EdiscoveryTagOperationStatus = "NotStarted"
	EdiscoveryTagOperationStatuspartiallySucceeded EdiscoveryTagOperationStatus = "PartiallySucceeded"
	EdiscoveryTagOperationStatusrunning            EdiscoveryTagOperationStatus = "Running"
	EdiscoveryTagOperationStatussubmissionFailed   EdiscoveryTagOperationStatus = "SubmissionFailed"
	EdiscoveryTagOperationStatussucceeded          EdiscoveryTagOperationStatus = "Succeeded"
)

func PossibleValuesForEdiscoveryTagOperationStatus() []string {
	return []string{
		string(EdiscoveryTagOperationStatusfailed),
		string(EdiscoveryTagOperationStatusnotStarted),
		string(EdiscoveryTagOperationStatuspartiallySucceeded),
		string(EdiscoveryTagOperationStatusrunning),
		string(EdiscoveryTagOperationStatussubmissionFailed),
		string(EdiscoveryTagOperationStatussucceeded),
	}
}

func parseEdiscoveryTagOperationStatus(input string) (*EdiscoveryTagOperationStatus, error) {
	vals := map[string]EdiscoveryTagOperationStatus{
		"failed":             EdiscoveryTagOperationStatusfailed,
		"notstarted":         EdiscoveryTagOperationStatusnotStarted,
		"partiallysucceeded": EdiscoveryTagOperationStatuspartiallySucceeded,
		"running":            EdiscoveryTagOperationStatusrunning,
		"submissionfailed":   EdiscoveryTagOperationStatussubmissionFailed,
		"succeeded":          EdiscoveryTagOperationStatussucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryTagOperationStatus(input)
	return &out, nil
}
