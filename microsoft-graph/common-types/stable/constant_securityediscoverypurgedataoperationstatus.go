package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryPurgeDataOperationStatus string

const (
	SecurityEdiscoveryPurgeDataOperationStatus_Failed             SecurityEdiscoveryPurgeDataOperationStatus = "failed"
	SecurityEdiscoveryPurgeDataOperationStatus_NotStarted         SecurityEdiscoveryPurgeDataOperationStatus = "notStarted"
	SecurityEdiscoveryPurgeDataOperationStatus_PartiallySucceeded SecurityEdiscoveryPurgeDataOperationStatus = "partiallySucceeded"
	SecurityEdiscoveryPurgeDataOperationStatus_Running            SecurityEdiscoveryPurgeDataOperationStatus = "running"
	SecurityEdiscoveryPurgeDataOperationStatus_SubmissionFailed   SecurityEdiscoveryPurgeDataOperationStatus = "submissionFailed"
	SecurityEdiscoveryPurgeDataOperationStatus_Succeeded          SecurityEdiscoveryPurgeDataOperationStatus = "succeeded"
)

func PossibleValuesForSecurityEdiscoveryPurgeDataOperationStatus() []string {
	return []string{
		string(SecurityEdiscoveryPurgeDataOperationStatus_Failed),
		string(SecurityEdiscoveryPurgeDataOperationStatus_NotStarted),
		string(SecurityEdiscoveryPurgeDataOperationStatus_PartiallySucceeded),
		string(SecurityEdiscoveryPurgeDataOperationStatus_Running),
		string(SecurityEdiscoveryPurgeDataOperationStatus_SubmissionFailed),
		string(SecurityEdiscoveryPurgeDataOperationStatus_Succeeded),
	}
}

func (s *SecurityEdiscoveryPurgeDataOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEdiscoveryPurgeDataOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEdiscoveryPurgeDataOperationStatus(input string) (*SecurityEdiscoveryPurgeDataOperationStatus, error) {
	vals := map[string]SecurityEdiscoveryPurgeDataOperationStatus{
		"failed":             SecurityEdiscoveryPurgeDataOperationStatus_Failed,
		"notstarted":         SecurityEdiscoveryPurgeDataOperationStatus_NotStarted,
		"partiallysucceeded": SecurityEdiscoveryPurgeDataOperationStatus_PartiallySucceeded,
		"running":            SecurityEdiscoveryPurgeDataOperationStatus_Running,
		"submissionfailed":   SecurityEdiscoveryPurgeDataOperationStatus_SubmissionFailed,
		"succeeded":          SecurityEdiscoveryPurgeDataOperationStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryPurgeDataOperationStatus(input)
	return &out, nil
}
