package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryPurgeDataOperationStatus string

const (
	EdiscoveryPurgeDataOperationStatusfailed             EdiscoveryPurgeDataOperationStatus = "Failed"
	EdiscoveryPurgeDataOperationStatusnotStarted         EdiscoveryPurgeDataOperationStatus = "NotStarted"
	EdiscoveryPurgeDataOperationStatuspartiallySucceeded EdiscoveryPurgeDataOperationStatus = "PartiallySucceeded"
	EdiscoveryPurgeDataOperationStatusrunning            EdiscoveryPurgeDataOperationStatus = "Running"
	EdiscoveryPurgeDataOperationStatussubmissionFailed   EdiscoveryPurgeDataOperationStatus = "SubmissionFailed"
	EdiscoveryPurgeDataOperationStatussucceeded          EdiscoveryPurgeDataOperationStatus = "Succeeded"
)

func PossibleValuesForEdiscoveryPurgeDataOperationStatus() []string {
	return []string{
		string(EdiscoveryPurgeDataOperationStatusfailed),
		string(EdiscoveryPurgeDataOperationStatusnotStarted),
		string(EdiscoveryPurgeDataOperationStatuspartiallySucceeded),
		string(EdiscoveryPurgeDataOperationStatusrunning),
		string(EdiscoveryPurgeDataOperationStatussubmissionFailed),
		string(EdiscoveryPurgeDataOperationStatussucceeded),
	}
}

func parseEdiscoveryPurgeDataOperationStatus(input string) (*EdiscoveryPurgeDataOperationStatus, error) {
	vals := map[string]EdiscoveryPurgeDataOperationStatus{
		"failed":             EdiscoveryPurgeDataOperationStatusfailed,
		"notstarted":         EdiscoveryPurgeDataOperationStatusnotStarted,
		"partiallysucceeded": EdiscoveryPurgeDataOperationStatuspartiallySucceeded,
		"running":            EdiscoveryPurgeDataOperationStatusrunning,
		"submissionfailed":   EdiscoveryPurgeDataOperationStatussubmissionFailed,
		"succeeded":          EdiscoveryPurgeDataOperationStatussucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryPurgeDataOperationStatus(input)
	return &out, nil
}
