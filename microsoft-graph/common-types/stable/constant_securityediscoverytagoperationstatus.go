package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryTagOperationStatus string

const (
	SecurityEdiscoveryTagOperationStatus_Failed             SecurityEdiscoveryTagOperationStatus = "failed"
	SecurityEdiscoveryTagOperationStatus_NotStarted         SecurityEdiscoveryTagOperationStatus = "notStarted"
	SecurityEdiscoveryTagOperationStatus_PartiallySucceeded SecurityEdiscoveryTagOperationStatus = "partiallySucceeded"
	SecurityEdiscoveryTagOperationStatus_Running            SecurityEdiscoveryTagOperationStatus = "running"
	SecurityEdiscoveryTagOperationStatus_SubmissionFailed   SecurityEdiscoveryTagOperationStatus = "submissionFailed"
	SecurityEdiscoveryTagOperationStatus_Succeeded          SecurityEdiscoveryTagOperationStatus = "succeeded"
)

func PossibleValuesForSecurityEdiscoveryTagOperationStatus() []string {
	return []string{
		string(SecurityEdiscoveryTagOperationStatus_Failed),
		string(SecurityEdiscoveryTagOperationStatus_NotStarted),
		string(SecurityEdiscoveryTagOperationStatus_PartiallySucceeded),
		string(SecurityEdiscoveryTagOperationStatus_Running),
		string(SecurityEdiscoveryTagOperationStatus_SubmissionFailed),
		string(SecurityEdiscoveryTagOperationStatus_Succeeded),
	}
}

func (s *SecurityEdiscoveryTagOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEdiscoveryTagOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEdiscoveryTagOperationStatus(input string) (*SecurityEdiscoveryTagOperationStatus, error) {
	vals := map[string]SecurityEdiscoveryTagOperationStatus{
		"failed":             SecurityEdiscoveryTagOperationStatus_Failed,
		"notstarted":         SecurityEdiscoveryTagOperationStatus_NotStarted,
		"partiallysucceeded": SecurityEdiscoveryTagOperationStatus_PartiallySucceeded,
		"running":            SecurityEdiscoveryTagOperationStatus_Running,
		"submissionfailed":   SecurityEdiscoveryTagOperationStatus_SubmissionFailed,
		"succeeded":          SecurityEdiscoveryTagOperationStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryTagOperationStatus(input)
	return &out, nil
}
