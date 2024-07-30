package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryAddToReviewSetOperationStatus string

const (
	SecurityEdiscoveryAddToReviewSetOperationStatus_Failed             SecurityEdiscoveryAddToReviewSetOperationStatus = "failed"
	SecurityEdiscoveryAddToReviewSetOperationStatus_NotStarted         SecurityEdiscoveryAddToReviewSetOperationStatus = "notStarted"
	SecurityEdiscoveryAddToReviewSetOperationStatus_PartiallySucceeded SecurityEdiscoveryAddToReviewSetOperationStatus = "partiallySucceeded"
	SecurityEdiscoveryAddToReviewSetOperationStatus_Running            SecurityEdiscoveryAddToReviewSetOperationStatus = "running"
	SecurityEdiscoveryAddToReviewSetOperationStatus_SubmissionFailed   SecurityEdiscoveryAddToReviewSetOperationStatus = "submissionFailed"
	SecurityEdiscoveryAddToReviewSetOperationStatus_Succeeded          SecurityEdiscoveryAddToReviewSetOperationStatus = "succeeded"
)

func PossibleValuesForSecurityEdiscoveryAddToReviewSetOperationStatus() []string {
	return []string{
		string(SecurityEdiscoveryAddToReviewSetOperationStatus_Failed),
		string(SecurityEdiscoveryAddToReviewSetOperationStatus_NotStarted),
		string(SecurityEdiscoveryAddToReviewSetOperationStatus_PartiallySucceeded),
		string(SecurityEdiscoveryAddToReviewSetOperationStatus_Running),
		string(SecurityEdiscoveryAddToReviewSetOperationStatus_SubmissionFailed),
		string(SecurityEdiscoveryAddToReviewSetOperationStatus_Succeeded),
	}
}

func (s *SecurityEdiscoveryAddToReviewSetOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEdiscoveryAddToReviewSetOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEdiscoveryAddToReviewSetOperationStatus(input string) (*SecurityEdiscoveryAddToReviewSetOperationStatus, error) {
	vals := map[string]SecurityEdiscoveryAddToReviewSetOperationStatus{
		"failed":             SecurityEdiscoveryAddToReviewSetOperationStatus_Failed,
		"notstarted":         SecurityEdiscoveryAddToReviewSetOperationStatus_NotStarted,
		"partiallysucceeded": SecurityEdiscoveryAddToReviewSetOperationStatus_PartiallySucceeded,
		"running":            SecurityEdiscoveryAddToReviewSetOperationStatus_Running,
		"submissionfailed":   SecurityEdiscoveryAddToReviewSetOperationStatus_SubmissionFailed,
		"succeeded":          SecurityEdiscoveryAddToReviewSetOperationStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryAddToReviewSetOperationStatus(input)
	return &out, nil
}
