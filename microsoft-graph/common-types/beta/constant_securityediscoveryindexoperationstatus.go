package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryIndexOperationStatus string

const (
	SecurityEdiscoveryIndexOperationStatus_Failed             SecurityEdiscoveryIndexOperationStatus = "failed"
	SecurityEdiscoveryIndexOperationStatus_NotStarted         SecurityEdiscoveryIndexOperationStatus = "notStarted"
	SecurityEdiscoveryIndexOperationStatus_PartiallySucceeded SecurityEdiscoveryIndexOperationStatus = "partiallySucceeded"
	SecurityEdiscoveryIndexOperationStatus_Running            SecurityEdiscoveryIndexOperationStatus = "running"
	SecurityEdiscoveryIndexOperationStatus_SubmissionFailed   SecurityEdiscoveryIndexOperationStatus = "submissionFailed"
	SecurityEdiscoveryIndexOperationStatus_Succeeded          SecurityEdiscoveryIndexOperationStatus = "succeeded"
)

func PossibleValuesForSecurityEdiscoveryIndexOperationStatus() []string {
	return []string{
		string(SecurityEdiscoveryIndexOperationStatus_Failed),
		string(SecurityEdiscoveryIndexOperationStatus_NotStarted),
		string(SecurityEdiscoveryIndexOperationStatus_PartiallySucceeded),
		string(SecurityEdiscoveryIndexOperationStatus_Running),
		string(SecurityEdiscoveryIndexOperationStatus_SubmissionFailed),
		string(SecurityEdiscoveryIndexOperationStatus_Succeeded),
	}
}

func (s *SecurityEdiscoveryIndexOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEdiscoveryIndexOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEdiscoveryIndexOperationStatus(input string) (*SecurityEdiscoveryIndexOperationStatus, error) {
	vals := map[string]SecurityEdiscoveryIndexOperationStatus{
		"failed":             SecurityEdiscoveryIndexOperationStatus_Failed,
		"notstarted":         SecurityEdiscoveryIndexOperationStatus_NotStarted,
		"partiallysucceeded": SecurityEdiscoveryIndexOperationStatus_PartiallySucceeded,
		"running":            SecurityEdiscoveryIndexOperationStatus_Running,
		"submissionfailed":   SecurityEdiscoveryIndexOperationStatus_SubmissionFailed,
		"succeeded":          SecurityEdiscoveryIndexOperationStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryIndexOperationStatus(input)
	return &out, nil
}
