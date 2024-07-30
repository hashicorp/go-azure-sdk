package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryCaseExportOperationStatus string

const (
	EdiscoveryCaseExportOperationStatus_Failed             EdiscoveryCaseExportOperationStatus = "failed"
	EdiscoveryCaseExportOperationStatus_NotStarted         EdiscoveryCaseExportOperationStatus = "notStarted"
	EdiscoveryCaseExportOperationStatus_PartiallySucceeded EdiscoveryCaseExportOperationStatus = "partiallySucceeded"
	EdiscoveryCaseExportOperationStatus_Running            EdiscoveryCaseExportOperationStatus = "running"
	EdiscoveryCaseExportOperationStatus_SubmissionFailed   EdiscoveryCaseExportOperationStatus = "submissionFailed"
	EdiscoveryCaseExportOperationStatus_Succeeded          EdiscoveryCaseExportOperationStatus = "succeeded"
)

func PossibleValuesForEdiscoveryCaseExportOperationStatus() []string {
	return []string{
		string(EdiscoveryCaseExportOperationStatus_Failed),
		string(EdiscoveryCaseExportOperationStatus_NotStarted),
		string(EdiscoveryCaseExportOperationStatus_PartiallySucceeded),
		string(EdiscoveryCaseExportOperationStatus_Running),
		string(EdiscoveryCaseExportOperationStatus_SubmissionFailed),
		string(EdiscoveryCaseExportOperationStatus_Succeeded),
	}
}

func (s *EdiscoveryCaseExportOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryCaseExportOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryCaseExportOperationStatus(input string) (*EdiscoveryCaseExportOperationStatus, error) {
	vals := map[string]EdiscoveryCaseExportOperationStatus{
		"failed":             EdiscoveryCaseExportOperationStatus_Failed,
		"notstarted":         EdiscoveryCaseExportOperationStatus_NotStarted,
		"partiallysucceeded": EdiscoveryCaseExportOperationStatus_PartiallySucceeded,
		"running":            EdiscoveryCaseExportOperationStatus_Running,
		"submissionfailed":   EdiscoveryCaseExportOperationStatus_SubmissionFailed,
		"succeeded":          EdiscoveryCaseExportOperationStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryCaseExportOperationStatus(input)
	return &out, nil
}
