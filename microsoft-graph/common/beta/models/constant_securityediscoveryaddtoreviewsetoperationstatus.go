package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryAddToReviewSetOperationStatus string

const (
	SecurityEdiscoveryAddToReviewSetOperationStatusfailed             SecurityEdiscoveryAddToReviewSetOperationStatus = "Failed"
	SecurityEdiscoveryAddToReviewSetOperationStatusnotStarted         SecurityEdiscoveryAddToReviewSetOperationStatus = "NotStarted"
	SecurityEdiscoveryAddToReviewSetOperationStatuspartiallySucceeded SecurityEdiscoveryAddToReviewSetOperationStatus = "PartiallySucceeded"
	SecurityEdiscoveryAddToReviewSetOperationStatusrunning            SecurityEdiscoveryAddToReviewSetOperationStatus = "Running"
	SecurityEdiscoveryAddToReviewSetOperationStatussubmissionFailed   SecurityEdiscoveryAddToReviewSetOperationStatus = "SubmissionFailed"
	SecurityEdiscoveryAddToReviewSetOperationStatussucceeded          SecurityEdiscoveryAddToReviewSetOperationStatus = "Succeeded"
)

func PossibleValuesForSecurityEdiscoveryAddToReviewSetOperationStatus() []string {
	return []string{
		string(SecurityEdiscoveryAddToReviewSetOperationStatusfailed),
		string(SecurityEdiscoveryAddToReviewSetOperationStatusnotStarted),
		string(SecurityEdiscoveryAddToReviewSetOperationStatuspartiallySucceeded),
		string(SecurityEdiscoveryAddToReviewSetOperationStatusrunning),
		string(SecurityEdiscoveryAddToReviewSetOperationStatussubmissionFailed),
		string(SecurityEdiscoveryAddToReviewSetOperationStatussucceeded),
	}
}

func parseSecurityEdiscoveryAddToReviewSetOperationStatus(input string) (*SecurityEdiscoveryAddToReviewSetOperationStatus, error) {
	vals := map[string]SecurityEdiscoveryAddToReviewSetOperationStatus{
		"failed":             SecurityEdiscoveryAddToReviewSetOperationStatusfailed,
		"notstarted":         SecurityEdiscoveryAddToReviewSetOperationStatusnotStarted,
		"partiallysucceeded": SecurityEdiscoveryAddToReviewSetOperationStatuspartiallySucceeded,
		"running":            SecurityEdiscoveryAddToReviewSetOperationStatusrunning,
		"submissionfailed":   SecurityEdiscoveryAddToReviewSetOperationStatussubmissionFailed,
		"succeeded":          SecurityEdiscoveryAddToReviewSetOperationStatussucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryAddToReviewSetOperationStatus(input)
	return &out, nil
}
