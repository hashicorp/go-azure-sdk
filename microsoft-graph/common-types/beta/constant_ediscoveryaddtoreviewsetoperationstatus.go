package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryAddToReviewSetOperationStatus string

const (
	EdiscoveryAddToReviewSetOperationStatus_Failed             EdiscoveryAddToReviewSetOperationStatus = "failed"
	EdiscoveryAddToReviewSetOperationStatus_NotStarted         EdiscoveryAddToReviewSetOperationStatus = "notStarted"
	EdiscoveryAddToReviewSetOperationStatus_PartiallySucceeded EdiscoveryAddToReviewSetOperationStatus = "partiallySucceeded"
	EdiscoveryAddToReviewSetOperationStatus_Running            EdiscoveryAddToReviewSetOperationStatus = "running"
	EdiscoveryAddToReviewSetOperationStatus_SubmissionFailed   EdiscoveryAddToReviewSetOperationStatus = "submissionFailed"
	EdiscoveryAddToReviewSetOperationStatus_Succeeded          EdiscoveryAddToReviewSetOperationStatus = "succeeded"
)

func PossibleValuesForEdiscoveryAddToReviewSetOperationStatus() []string {
	return []string{
		string(EdiscoveryAddToReviewSetOperationStatus_Failed),
		string(EdiscoveryAddToReviewSetOperationStatus_NotStarted),
		string(EdiscoveryAddToReviewSetOperationStatus_PartiallySucceeded),
		string(EdiscoveryAddToReviewSetOperationStatus_Running),
		string(EdiscoveryAddToReviewSetOperationStatus_SubmissionFailed),
		string(EdiscoveryAddToReviewSetOperationStatus_Succeeded),
	}
}

func (s *EdiscoveryAddToReviewSetOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryAddToReviewSetOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryAddToReviewSetOperationStatus(input string) (*EdiscoveryAddToReviewSetOperationStatus, error) {
	vals := map[string]EdiscoveryAddToReviewSetOperationStatus{
		"failed":             EdiscoveryAddToReviewSetOperationStatus_Failed,
		"notstarted":         EdiscoveryAddToReviewSetOperationStatus_NotStarted,
		"partiallysucceeded": EdiscoveryAddToReviewSetOperationStatus_PartiallySucceeded,
		"running":            EdiscoveryAddToReviewSetOperationStatus_Running,
		"submissionfailed":   EdiscoveryAddToReviewSetOperationStatus_SubmissionFailed,
		"succeeded":          EdiscoveryAddToReviewSetOperationStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryAddToReviewSetOperationStatus(input)
	return &out, nil
}
