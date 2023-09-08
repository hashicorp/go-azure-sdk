package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryTagOperationStatus string

const (
	SecurityEdiscoveryTagOperationStatusfailed             SecurityEdiscoveryTagOperationStatus = "Failed"
	SecurityEdiscoveryTagOperationStatusnotStarted         SecurityEdiscoveryTagOperationStatus = "NotStarted"
	SecurityEdiscoveryTagOperationStatuspartiallySucceeded SecurityEdiscoveryTagOperationStatus = "PartiallySucceeded"
	SecurityEdiscoveryTagOperationStatusrunning            SecurityEdiscoveryTagOperationStatus = "Running"
	SecurityEdiscoveryTagOperationStatussubmissionFailed   SecurityEdiscoveryTagOperationStatus = "SubmissionFailed"
	SecurityEdiscoveryTagOperationStatussucceeded          SecurityEdiscoveryTagOperationStatus = "Succeeded"
)

func PossibleValuesForSecurityEdiscoveryTagOperationStatus() []string {
	return []string{
		string(SecurityEdiscoveryTagOperationStatusfailed),
		string(SecurityEdiscoveryTagOperationStatusnotStarted),
		string(SecurityEdiscoveryTagOperationStatuspartiallySucceeded),
		string(SecurityEdiscoveryTagOperationStatusrunning),
		string(SecurityEdiscoveryTagOperationStatussubmissionFailed),
		string(SecurityEdiscoveryTagOperationStatussucceeded),
	}
}

func parseSecurityEdiscoveryTagOperationStatus(input string) (*SecurityEdiscoveryTagOperationStatus, error) {
	vals := map[string]SecurityEdiscoveryTagOperationStatus{
		"failed":             SecurityEdiscoveryTagOperationStatusfailed,
		"notstarted":         SecurityEdiscoveryTagOperationStatusnotStarted,
		"partiallysucceeded": SecurityEdiscoveryTagOperationStatuspartiallySucceeded,
		"running":            SecurityEdiscoveryTagOperationStatusrunning,
		"submissionfailed":   SecurityEdiscoveryTagOperationStatussubmissionFailed,
		"succeeded":          SecurityEdiscoveryTagOperationStatussucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryTagOperationStatus(input)
	return &out, nil
}
