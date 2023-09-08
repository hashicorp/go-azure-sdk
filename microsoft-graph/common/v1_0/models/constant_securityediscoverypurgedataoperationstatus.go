package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryPurgeDataOperationStatus string

const (
	SecurityEdiscoveryPurgeDataOperationStatusfailed             SecurityEdiscoveryPurgeDataOperationStatus = "Failed"
	SecurityEdiscoveryPurgeDataOperationStatusnotStarted         SecurityEdiscoveryPurgeDataOperationStatus = "NotStarted"
	SecurityEdiscoveryPurgeDataOperationStatuspartiallySucceeded SecurityEdiscoveryPurgeDataOperationStatus = "PartiallySucceeded"
	SecurityEdiscoveryPurgeDataOperationStatusrunning            SecurityEdiscoveryPurgeDataOperationStatus = "Running"
	SecurityEdiscoveryPurgeDataOperationStatussubmissionFailed   SecurityEdiscoveryPurgeDataOperationStatus = "SubmissionFailed"
	SecurityEdiscoveryPurgeDataOperationStatussucceeded          SecurityEdiscoveryPurgeDataOperationStatus = "Succeeded"
)

func PossibleValuesForSecurityEdiscoveryPurgeDataOperationStatus() []string {
	return []string{
		string(SecurityEdiscoveryPurgeDataOperationStatusfailed),
		string(SecurityEdiscoveryPurgeDataOperationStatusnotStarted),
		string(SecurityEdiscoveryPurgeDataOperationStatuspartiallySucceeded),
		string(SecurityEdiscoveryPurgeDataOperationStatusrunning),
		string(SecurityEdiscoveryPurgeDataOperationStatussubmissionFailed),
		string(SecurityEdiscoveryPurgeDataOperationStatussucceeded),
	}
}

func parseSecurityEdiscoveryPurgeDataOperationStatus(input string) (*SecurityEdiscoveryPurgeDataOperationStatus, error) {
	vals := map[string]SecurityEdiscoveryPurgeDataOperationStatus{
		"failed":             SecurityEdiscoveryPurgeDataOperationStatusfailed,
		"notstarted":         SecurityEdiscoveryPurgeDataOperationStatusnotStarted,
		"partiallysucceeded": SecurityEdiscoveryPurgeDataOperationStatuspartiallySucceeded,
		"running":            SecurityEdiscoveryPurgeDataOperationStatusrunning,
		"submissionfailed":   SecurityEdiscoveryPurgeDataOperationStatussubmissionFailed,
		"succeeded":          SecurityEdiscoveryPurgeDataOperationStatussucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryPurgeDataOperationStatus(input)
	return &out, nil
}
