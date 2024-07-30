package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryCaseHoldOperationStatus string

const (
	EdiscoveryCaseHoldOperationStatus_Failed             EdiscoveryCaseHoldOperationStatus = "failed"
	EdiscoveryCaseHoldOperationStatus_NotStarted         EdiscoveryCaseHoldOperationStatus = "notStarted"
	EdiscoveryCaseHoldOperationStatus_PartiallySucceeded EdiscoveryCaseHoldOperationStatus = "partiallySucceeded"
	EdiscoveryCaseHoldOperationStatus_Running            EdiscoveryCaseHoldOperationStatus = "running"
	EdiscoveryCaseHoldOperationStatus_SubmissionFailed   EdiscoveryCaseHoldOperationStatus = "submissionFailed"
	EdiscoveryCaseHoldOperationStatus_Succeeded          EdiscoveryCaseHoldOperationStatus = "succeeded"
)

func PossibleValuesForEdiscoveryCaseHoldOperationStatus() []string {
	return []string{
		string(EdiscoveryCaseHoldOperationStatus_Failed),
		string(EdiscoveryCaseHoldOperationStatus_NotStarted),
		string(EdiscoveryCaseHoldOperationStatus_PartiallySucceeded),
		string(EdiscoveryCaseHoldOperationStatus_Running),
		string(EdiscoveryCaseHoldOperationStatus_SubmissionFailed),
		string(EdiscoveryCaseHoldOperationStatus_Succeeded),
	}
}

func (s *EdiscoveryCaseHoldOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryCaseHoldOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryCaseHoldOperationStatus(input string) (*EdiscoveryCaseHoldOperationStatus, error) {
	vals := map[string]EdiscoveryCaseHoldOperationStatus{
		"failed":             EdiscoveryCaseHoldOperationStatus_Failed,
		"notstarted":         EdiscoveryCaseHoldOperationStatus_NotStarted,
		"partiallysucceeded": EdiscoveryCaseHoldOperationStatus_PartiallySucceeded,
		"running":            EdiscoveryCaseHoldOperationStatus_Running,
		"submissionfailed":   EdiscoveryCaseHoldOperationStatus_SubmissionFailed,
		"succeeded":          EdiscoveryCaseHoldOperationStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryCaseHoldOperationStatus(input)
	return &out, nil
}
