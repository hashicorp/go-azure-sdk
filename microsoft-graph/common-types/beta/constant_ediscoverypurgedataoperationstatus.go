package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryPurgeDataOperationStatus string

const (
	EdiscoveryPurgeDataOperationStatus_Failed             EdiscoveryPurgeDataOperationStatus = "failed"
	EdiscoveryPurgeDataOperationStatus_NotStarted         EdiscoveryPurgeDataOperationStatus = "notStarted"
	EdiscoveryPurgeDataOperationStatus_PartiallySucceeded EdiscoveryPurgeDataOperationStatus = "partiallySucceeded"
	EdiscoveryPurgeDataOperationStatus_Running            EdiscoveryPurgeDataOperationStatus = "running"
	EdiscoveryPurgeDataOperationStatus_SubmissionFailed   EdiscoveryPurgeDataOperationStatus = "submissionFailed"
	EdiscoveryPurgeDataOperationStatus_Succeeded          EdiscoveryPurgeDataOperationStatus = "succeeded"
)

func PossibleValuesForEdiscoveryPurgeDataOperationStatus() []string {
	return []string{
		string(EdiscoveryPurgeDataOperationStatus_Failed),
		string(EdiscoveryPurgeDataOperationStatus_NotStarted),
		string(EdiscoveryPurgeDataOperationStatus_PartiallySucceeded),
		string(EdiscoveryPurgeDataOperationStatus_Running),
		string(EdiscoveryPurgeDataOperationStatus_SubmissionFailed),
		string(EdiscoveryPurgeDataOperationStatus_Succeeded),
	}
}

func (s *EdiscoveryPurgeDataOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryPurgeDataOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryPurgeDataOperationStatus(input string) (*EdiscoveryPurgeDataOperationStatus, error) {
	vals := map[string]EdiscoveryPurgeDataOperationStatus{
		"failed":             EdiscoveryPurgeDataOperationStatus_Failed,
		"notstarted":         EdiscoveryPurgeDataOperationStatus_NotStarted,
		"partiallysucceeded": EdiscoveryPurgeDataOperationStatus_PartiallySucceeded,
		"running":            EdiscoveryPurgeDataOperationStatus_Running,
		"submissionfailed":   EdiscoveryPurgeDataOperationStatus_SubmissionFailed,
		"succeeded":          EdiscoveryPurgeDataOperationStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryPurgeDataOperationStatus(input)
	return &out, nil
}
