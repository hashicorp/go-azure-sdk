package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryExportOperationStatus string

const (
	SecurityEdiscoveryExportOperationStatusfailed             SecurityEdiscoveryExportOperationStatus = "Failed"
	SecurityEdiscoveryExportOperationStatusnotStarted         SecurityEdiscoveryExportOperationStatus = "NotStarted"
	SecurityEdiscoveryExportOperationStatuspartiallySucceeded SecurityEdiscoveryExportOperationStatus = "PartiallySucceeded"
	SecurityEdiscoveryExportOperationStatusrunning            SecurityEdiscoveryExportOperationStatus = "Running"
	SecurityEdiscoveryExportOperationStatussubmissionFailed   SecurityEdiscoveryExportOperationStatus = "SubmissionFailed"
	SecurityEdiscoveryExportOperationStatussucceeded          SecurityEdiscoveryExportOperationStatus = "Succeeded"
)

func PossibleValuesForSecurityEdiscoveryExportOperationStatus() []string {
	return []string{
		string(SecurityEdiscoveryExportOperationStatusfailed),
		string(SecurityEdiscoveryExportOperationStatusnotStarted),
		string(SecurityEdiscoveryExportOperationStatuspartiallySucceeded),
		string(SecurityEdiscoveryExportOperationStatusrunning),
		string(SecurityEdiscoveryExportOperationStatussubmissionFailed),
		string(SecurityEdiscoveryExportOperationStatussucceeded),
	}
}

func parseSecurityEdiscoveryExportOperationStatus(input string) (*SecurityEdiscoveryExportOperationStatus, error) {
	vals := map[string]SecurityEdiscoveryExportOperationStatus{
		"failed":             SecurityEdiscoveryExportOperationStatusfailed,
		"notstarted":         SecurityEdiscoveryExportOperationStatusnotStarted,
		"partiallysucceeded": SecurityEdiscoveryExportOperationStatuspartiallySucceeded,
		"running":            SecurityEdiscoveryExportOperationStatusrunning,
		"submissionfailed":   SecurityEdiscoveryExportOperationStatussubmissionFailed,
		"succeeded":          SecurityEdiscoveryExportOperationStatussucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryExportOperationStatus(input)
	return &out, nil
}
