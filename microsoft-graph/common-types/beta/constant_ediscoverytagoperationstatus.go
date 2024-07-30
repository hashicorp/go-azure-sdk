package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryTagOperationStatus string

const (
	EdiscoveryTagOperationStatus_Failed             EdiscoveryTagOperationStatus = "failed"
	EdiscoveryTagOperationStatus_NotStarted         EdiscoveryTagOperationStatus = "notStarted"
	EdiscoveryTagOperationStatus_PartiallySucceeded EdiscoveryTagOperationStatus = "partiallySucceeded"
	EdiscoveryTagOperationStatus_Running            EdiscoveryTagOperationStatus = "running"
	EdiscoveryTagOperationStatus_SubmissionFailed   EdiscoveryTagOperationStatus = "submissionFailed"
	EdiscoveryTagOperationStatus_Succeeded          EdiscoveryTagOperationStatus = "succeeded"
)

func PossibleValuesForEdiscoveryTagOperationStatus() []string {
	return []string{
		string(EdiscoveryTagOperationStatus_Failed),
		string(EdiscoveryTagOperationStatus_NotStarted),
		string(EdiscoveryTagOperationStatus_PartiallySucceeded),
		string(EdiscoveryTagOperationStatus_Running),
		string(EdiscoveryTagOperationStatus_SubmissionFailed),
		string(EdiscoveryTagOperationStatus_Succeeded),
	}
}

func (s *EdiscoveryTagOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryTagOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryTagOperationStatus(input string) (*EdiscoveryTagOperationStatus, error) {
	vals := map[string]EdiscoveryTagOperationStatus{
		"failed":             EdiscoveryTagOperationStatus_Failed,
		"notstarted":         EdiscoveryTagOperationStatus_NotStarted,
		"partiallysucceeded": EdiscoveryTagOperationStatus_PartiallySucceeded,
		"running":            EdiscoveryTagOperationStatus_Running,
		"submissionfailed":   EdiscoveryTagOperationStatus_SubmissionFailed,
		"succeeded":          EdiscoveryTagOperationStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryTagOperationStatus(input)
	return &out, nil
}
