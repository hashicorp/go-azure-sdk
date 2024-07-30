package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryEstimateOperationStatus string

const (
	SecurityEdiscoveryEstimateOperationStatus_Failed             SecurityEdiscoveryEstimateOperationStatus = "failed"
	SecurityEdiscoveryEstimateOperationStatus_NotStarted         SecurityEdiscoveryEstimateOperationStatus = "notStarted"
	SecurityEdiscoveryEstimateOperationStatus_PartiallySucceeded SecurityEdiscoveryEstimateOperationStatus = "partiallySucceeded"
	SecurityEdiscoveryEstimateOperationStatus_Running            SecurityEdiscoveryEstimateOperationStatus = "running"
	SecurityEdiscoveryEstimateOperationStatus_SubmissionFailed   SecurityEdiscoveryEstimateOperationStatus = "submissionFailed"
	SecurityEdiscoveryEstimateOperationStatus_Succeeded          SecurityEdiscoveryEstimateOperationStatus = "succeeded"
)

func PossibleValuesForSecurityEdiscoveryEstimateOperationStatus() []string {
	return []string{
		string(SecurityEdiscoveryEstimateOperationStatus_Failed),
		string(SecurityEdiscoveryEstimateOperationStatus_NotStarted),
		string(SecurityEdiscoveryEstimateOperationStatus_PartiallySucceeded),
		string(SecurityEdiscoveryEstimateOperationStatus_Running),
		string(SecurityEdiscoveryEstimateOperationStatus_SubmissionFailed),
		string(SecurityEdiscoveryEstimateOperationStatus_Succeeded),
	}
}

func (s *SecurityEdiscoveryEstimateOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEdiscoveryEstimateOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEdiscoveryEstimateOperationStatus(input string) (*SecurityEdiscoveryEstimateOperationStatus, error) {
	vals := map[string]SecurityEdiscoveryEstimateOperationStatus{
		"failed":             SecurityEdiscoveryEstimateOperationStatus_Failed,
		"notstarted":         SecurityEdiscoveryEstimateOperationStatus_NotStarted,
		"partiallysucceeded": SecurityEdiscoveryEstimateOperationStatus_PartiallySucceeded,
		"running":            SecurityEdiscoveryEstimateOperationStatus_Running,
		"submissionfailed":   SecurityEdiscoveryEstimateOperationStatus_SubmissionFailed,
		"succeeded":          SecurityEdiscoveryEstimateOperationStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryEstimateOperationStatus(input)
	return &out, nil
}
