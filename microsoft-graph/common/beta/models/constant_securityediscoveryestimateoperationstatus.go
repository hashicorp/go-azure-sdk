package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryEstimateOperationStatus string

const (
	SecurityEdiscoveryEstimateOperationStatusfailed             SecurityEdiscoveryEstimateOperationStatus = "Failed"
	SecurityEdiscoveryEstimateOperationStatusnotStarted         SecurityEdiscoveryEstimateOperationStatus = "NotStarted"
	SecurityEdiscoveryEstimateOperationStatuspartiallySucceeded SecurityEdiscoveryEstimateOperationStatus = "PartiallySucceeded"
	SecurityEdiscoveryEstimateOperationStatusrunning            SecurityEdiscoveryEstimateOperationStatus = "Running"
	SecurityEdiscoveryEstimateOperationStatussubmissionFailed   SecurityEdiscoveryEstimateOperationStatus = "SubmissionFailed"
	SecurityEdiscoveryEstimateOperationStatussucceeded          SecurityEdiscoveryEstimateOperationStatus = "Succeeded"
)

func PossibleValuesForSecurityEdiscoveryEstimateOperationStatus() []string {
	return []string{
		string(SecurityEdiscoveryEstimateOperationStatusfailed),
		string(SecurityEdiscoveryEstimateOperationStatusnotStarted),
		string(SecurityEdiscoveryEstimateOperationStatuspartiallySucceeded),
		string(SecurityEdiscoveryEstimateOperationStatusrunning),
		string(SecurityEdiscoveryEstimateOperationStatussubmissionFailed),
		string(SecurityEdiscoveryEstimateOperationStatussucceeded),
	}
}

func parseSecurityEdiscoveryEstimateOperationStatus(input string) (*SecurityEdiscoveryEstimateOperationStatus, error) {
	vals := map[string]SecurityEdiscoveryEstimateOperationStatus{
		"failed":             SecurityEdiscoveryEstimateOperationStatusfailed,
		"notstarted":         SecurityEdiscoveryEstimateOperationStatusnotStarted,
		"partiallysucceeded": SecurityEdiscoveryEstimateOperationStatuspartiallySucceeded,
		"running":            SecurityEdiscoveryEstimateOperationStatusrunning,
		"submissionfailed":   SecurityEdiscoveryEstimateOperationStatussubmissionFailed,
		"succeeded":          SecurityEdiscoveryEstimateOperationStatussucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryEstimateOperationStatus(input)
	return &out, nil
}
