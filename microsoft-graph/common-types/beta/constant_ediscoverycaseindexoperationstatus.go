package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryCaseIndexOperationStatus string

const (
	EdiscoveryCaseIndexOperationStatus_Failed             EdiscoveryCaseIndexOperationStatus = "failed"
	EdiscoveryCaseIndexOperationStatus_NotStarted         EdiscoveryCaseIndexOperationStatus = "notStarted"
	EdiscoveryCaseIndexOperationStatus_PartiallySucceeded EdiscoveryCaseIndexOperationStatus = "partiallySucceeded"
	EdiscoveryCaseIndexOperationStatus_Running            EdiscoveryCaseIndexOperationStatus = "running"
	EdiscoveryCaseIndexOperationStatus_SubmissionFailed   EdiscoveryCaseIndexOperationStatus = "submissionFailed"
	EdiscoveryCaseIndexOperationStatus_Succeeded          EdiscoveryCaseIndexOperationStatus = "succeeded"
)

func PossibleValuesForEdiscoveryCaseIndexOperationStatus() []string {
	return []string{
		string(EdiscoveryCaseIndexOperationStatus_Failed),
		string(EdiscoveryCaseIndexOperationStatus_NotStarted),
		string(EdiscoveryCaseIndexOperationStatus_PartiallySucceeded),
		string(EdiscoveryCaseIndexOperationStatus_Running),
		string(EdiscoveryCaseIndexOperationStatus_SubmissionFailed),
		string(EdiscoveryCaseIndexOperationStatus_Succeeded),
	}
}

func (s *EdiscoveryCaseIndexOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryCaseIndexOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryCaseIndexOperationStatus(input string) (*EdiscoveryCaseIndexOperationStatus, error) {
	vals := map[string]EdiscoveryCaseIndexOperationStatus{
		"failed":             EdiscoveryCaseIndexOperationStatus_Failed,
		"notstarted":         EdiscoveryCaseIndexOperationStatus_NotStarted,
		"partiallysucceeded": EdiscoveryCaseIndexOperationStatus_PartiallySucceeded,
		"running":            EdiscoveryCaseIndexOperationStatus_Running,
		"submissionfailed":   EdiscoveryCaseIndexOperationStatus_SubmissionFailed,
		"succeeded":          EdiscoveryCaseIndexOperationStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryCaseIndexOperationStatus(input)
	return &out, nil
}
