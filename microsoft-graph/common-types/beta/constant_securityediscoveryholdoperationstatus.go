package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryHoldOperationStatus string

const (
	SecurityEdiscoveryHoldOperationStatus_Failed             SecurityEdiscoveryHoldOperationStatus = "failed"
	SecurityEdiscoveryHoldOperationStatus_NotStarted         SecurityEdiscoveryHoldOperationStatus = "notStarted"
	SecurityEdiscoveryHoldOperationStatus_PartiallySucceeded SecurityEdiscoveryHoldOperationStatus = "partiallySucceeded"
	SecurityEdiscoveryHoldOperationStatus_Running            SecurityEdiscoveryHoldOperationStatus = "running"
	SecurityEdiscoveryHoldOperationStatus_SubmissionFailed   SecurityEdiscoveryHoldOperationStatus = "submissionFailed"
	SecurityEdiscoveryHoldOperationStatus_Succeeded          SecurityEdiscoveryHoldOperationStatus = "succeeded"
)

func PossibleValuesForSecurityEdiscoveryHoldOperationStatus() []string {
	return []string{
		string(SecurityEdiscoveryHoldOperationStatus_Failed),
		string(SecurityEdiscoveryHoldOperationStatus_NotStarted),
		string(SecurityEdiscoveryHoldOperationStatus_PartiallySucceeded),
		string(SecurityEdiscoveryHoldOperationStatus_Running),
		string(SecurityEdiscoveryHoldOperationStatus_SubmissionFailed),
		string(SecurityEdiscoveryHoldOperationStatus_Succeeded),
	}
}

func (s *SecurityEdiscoveryHoldOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEdiscoveryHoldOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEdiscoveryHoldOperationStatus(input string) (*SecurityEdiscoveryHoldOperationStatus, error) {
	vals := map[string]SecurityEdiscoveryHoldOperationStatus{
		"failed":             SecurityEdiscoveryHoldOperationStatus_Failed,
		"notstarted":         SecurityEdiscoveryHoldOperationStatus_NotStarted,
		"partiallysucceeded": SecurityEdiscoveryHoldOperationStatus_PartiallySucceeded,
		"running":            SecurityEdiscoveryHoldOperationStatus_Running,
		"submissionfailed":   SecurityEdiscoveryHoldOperationStatus_SubmissionFailed,
		"succeeded":          SecurityEdiscoveryHoldOperationStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryHoldOperationStatus(input)
	return &out, nil
}
