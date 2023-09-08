package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryEstimateStatisticsOperationStatus string

const (
	EdiscoveryEstimateStatisticsOperationStatusfailed             EdiscoveryEstimateStatisticsOperationStatus = "Failed"
	EdiscoveryEstimateStatisticsOperationStatusnotStarted         EdiscoveryEstimateStatisticsOperationStatus = "NotStarted"
	EdiscoveryEstimateStatisticsOperationStatuspartiallySucceeded EdiscoveryEstimateStatisticsOperationStatus = "PartiallySucceeded"
	EdiscoveryEstimateStatisticsOperationStatusrunning            EdiscoveryEstimateStatisticsOperationStatus = "Running"
	EdiscoveryEstimateStatisticsOperationStatussubmissionFailed   EdiscoveryEstimateStatisticsOperationStatus = "SubmissionFailed"
	EdiscoveryEstimateStatisticsOperationStatussucceeded          EdiscoveryEstimateStatisticsOperationStatus = "Succeeded"
)

func PossibleValuesForEdiscoveryEstimateStatisticsOperationStatus() []string {
	return []string{
		string(EdiscoveryEstimateStatisticsOperationStatusfailed),
		string(EdiscoveryEstimateStatisticsOperationStatusnotStarted),
		string(EdiscoveryEstimateStatisticsOperationStatuspartiallySucceeded),
		string(EdiscoveryEstimateStatisticsOperationStatusrunning),
		string(EdiscoveryEstimateStatisticsOperationStatussubmissionFailed),
		string(EdiscoveryEstimateStatisticsOperationStatussucceeded),
	}
}

func parseEdiscoveryEstimateStatisticsOperationStatus(input string) (*EdiscoveryEstimateStatisticsOperationStatus, error) {
	vals := map[string]EdiscoveryEstimateStatisticsOperationStatus{
		"failed":             EdiscoveryEstimateStatisticsOperationStatusfailed,
		"notstarted":         EdiscoveryEstimateStatisticsOperationStatusnotStarted,
		"partiallysucceeded": EdiscoveryEstimateStatisticsOperationStatuspartiallySucceeded,
		"running":            EdiscoveryEstimateStatisticsOperationStatusrunning,
		"submissionfailed":   EdiscoveryEstimateStatisticsOperationStatussubmissionFailed,
		"succeeded":          EdiscoveryEstimateStatisticsOperationStatussucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryEstimateStatisticsOperationStatus(input)
	return &out, nil
}
