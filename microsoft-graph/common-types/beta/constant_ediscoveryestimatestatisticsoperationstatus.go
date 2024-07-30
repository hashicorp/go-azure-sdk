package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryEstimateStatisticsOperationStatus string

const (
	EdiscoveryEstimateStatisticsOperationStatus_Failed             EdiscoveryEstimateStatisticsOperationStatus = "failed"
	EdiscoveryEstimateStatisticsOperationStatus_NotStarted         EdiscoveryEstimateStatisticsOperationStatus = "notStarted"
	EdiscoveryEstimateStatisticsOperationStatus_PartiallySucceeded EdiscoveryEstimateStatisticsOperationStatus = "partiallySucceeded"
	EdiscoveryEstimateStatisticsOperationStatus_Running            EdiscoveryEstimateStatisticsOperationStatus = "running"
	EdiscoveryEstimateStatisticsOperationStatus_SubmissionFailed   EdiscoveryEstimateStatisticsOperationStatus = "submissionFailed"
	EdiscoveryEstimateStatisticsOperationStatus_Succeeded          EdiscoveryEstimateStatisticsOperationStatus = "succeeded"
)

func PossibleValuesForEdiscoveryEstimateStatisticsOperationStatus() []string {
	return []string{
		string(EdiscoveryEstimateStatisticsOperationStatus_Failed),
		string(EdiscoveryEstimateStatisticsOperationStatus_NotStarted),
		string(EdiscoveryEstimateStatisticsOperationStatus_PartiallySucceeded),
		string(EdiscoveryEstimateStatisticsOperationStatus_Running),
		string(EdiscoveryEstimateStatisticsOperationStatus_SubmissionFailed),
		string(EdiscoveryEstimateStatisticsOperationStatus_Succeeded),
	}
}

func (s *EdiscoveryEstimateStatisticsOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryEstimateStatisticsOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryEstimateStatisticsOperationStatus(input string) (*EdiscoveryEstimateStatisticsOperationStatus, error) {
	vals := map[string]EdiscoveryEstimateStatisticsOperationStatus{
		"failed":             EdiscoveryEstimateStatisticsOperationStatus_Failed,
		"notstarted":         EdiscoveryEstimateStatisticsOperationStatus_NotStarted,
		"partiallysucceeded": EdiscoveryEstimateStatisticsOperationStatus_PartiallySucceeded,
		"running":            EdiscoveryEstimateStatisticsOperationStatus_Running,
		"submissionfailed":   EdiscoveryEstimateStatisticsOperationStatus_SubmissionFailed,
		"succeeded":          EdiscoveryEstimateStatisticsOperationStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryEstimateStatisticsOperationStatus(input)
	return &out, nil
}
