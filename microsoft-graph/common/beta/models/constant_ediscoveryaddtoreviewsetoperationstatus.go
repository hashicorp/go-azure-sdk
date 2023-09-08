package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryAddToReviewSetOperationStatus string

const (
	EdiscoveryAddToReviewSetOperationStatusfailed             EdiscoveryAddToReviewSetOperationStatus = "Failed"
	EdiscoveryAddToReviewSetOperationStatusnotStarted         EdiscoveryAddToReviewSetOperationStatus = "NotStarted"
	EdiscoveryAddToReviewSetOperationStatuspartiallySucceeded EdiscoveryAddToReviewSetOperationStatus = "PartiallySucceeded"
	EdiscoveryAddToReviewSetOperationStatusrunning            EdiscoveryAddToReviewSetOperationStatus = "Running"
	EdiscoveryAddToReviewSetOperationStatussubmissionFailed   EdiscoveryAddToReviewSetOperationStatus = "SubmissionFailed"
	EdiscoveryAddToReviewSetOperationStatussucceeded          EdiscoveryAddToReviewSetOperationStatus = "Succeeded"
)

func PossibleValuesForEdiscoveryAddToReviewSetOperationStatus() []string {
	return []string{
		string(EdiscoveryAddToReviewSetOperationStatusfailed),
		string(EdiscoveryAddToReviewSetOperationStatusnotStarted),
		string(EdiscoveryAddToReviewSetOperationStatuspartiallySucceeded),
		string(EdiscoveryAddToReviewSetOperationStatusrunning),
		string(EdiscoveryAddToReviewSetOperationStatussubmissionFailed),
		string(EdiscoveryAddToReviewSetOperationStatussucceeded),
	}
}

func parseEdiscoveryAddToReviewSetOperationStatus(input string) (*EdiscoveryAddToReviewSetOperationStatus, error) {
	vals := map[string]EdiscoveryAddToReviewSetOperationStatus{
		"failed":             EdiscoveryAddToReviewSetOperationStatusfailed,
		"notstarted":         EdiscoveryAddToReviewSetOperationStatusnotStarted,
		"partiallysucceeded": EdiscoveryAddToReviewSetOperationStatuspartiallySucceeded,
		"running":            EdiscoveryAddToReviewSetOperationStatusrunning,
		"submissionfailed":   EdiscoveryAddToReviewSetOperationStatussubmissionFailed,
		"succeeded":          EdiscoveryAddToReviewSetOperationStatussucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryAddToReviewSetOperationStatus(input)
	return &out, nil
}
