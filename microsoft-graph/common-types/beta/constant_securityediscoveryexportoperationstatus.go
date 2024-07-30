package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryExportOperationStatus string

const (
	SecurityEdiscoveryExportOperationStatus_Failed             SecurityEdiscoveryExportOperationStatus = "failed"
	SecurityEdiscoveryExportOperationStatus_NotStarted         SecurityEdiscoveryExportOperationStatus = "notStarted"
	SecurityEdiscoveryExportOperationStatus_PartiallySucceeded SecurityEdiscoveryExportOperationStatus = "partiallySucceeded"
	SecurityEdiscoveryExportOperationStatus_Running            SecurityEdiscoveryExportOperationStatus = "running"
	SecurityEdiscoveryExportOperationStatus_SubmissionFailed   SecurityEdiscoveryExportOperationStatus = "submissionFailed"
	SecurityEdiscoveryExportOperationStatus_Succeeded          SecurityEdiscoveryExportOperationStatus = "succeeded"
)

func PossibleValuesForSecurityEdiscoveryExportOperationStatus() []string {
	return []string{
		string(SecurityEdiscoveryExportOperationStatus_Failed),
		string(SecurityEdiscoveryExportOperationStatus_NotStarted),
		string(SecurityEdiscoveryExportOperationStatus_PartiallySucceeded),
		string(SecurityEdiscoveryExportOperationStatus_Running),
		string(SecurityEdiscoveryExportOperationStatus_SubmissionFailed),
		string(SecurityEdiscoveryExportOperationStatus_Succeeded),
	}
}

func (s *SecurityEdiscoveryExportOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEdiscoveryExportOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEdiscoveryExportOperationStatus(input string) (*SecurityEdiscoveryExportOperationStatus, error) {
	vals := map[string]SecurityEdiscoveryExportOperationStatus{
		"failed":             SecurityEdiscoveryExportOperationStatus_Failed,
		"notstarted":         SecurityEdiscoveryExportOperationStatus_NotStarted,
		"partiallysucceeded": SecurityEdiscoveryExportOperationStatus_PartiallySucceeded,
		"running":            SecurityEdiscoveryExportOperationStatus_Running,
		"submissionfailed":   SecurityEdiscoveryExportOperationStatus_SubmissionFailed,
		"succeeded":          SecurityEdiscoveryExportOperationStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryExportOperationStatus(input)
	return &out, nil
}
