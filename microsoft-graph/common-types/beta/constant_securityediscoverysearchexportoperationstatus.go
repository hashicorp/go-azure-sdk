package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoverySearchExportOperationStatus string

const (
	SecurityEdiscoverySearchExportOperationStatus_Failed             SecurityEdiscoverySearchExportOperationStatus = "failed"
	SecurityEdiscoverySearchExportOperationStatus_NotStarted         SecurityEdiscoverySearchExportOperationStatus = "notStarted"
	SecurityEdiscoverySearchExportOperationStatus_PartiallySucceeded SecurityEdiscoverySearchExportOperationStatus = "partiallySucceeded"
	SecurityEdiscoverySearchExportOperationStatus_Running            SecurityEdiscoverySearchExportOperationStatus = "running"
	SecurityEdiscoverySearchExportOperationStatus_SubmissionFailed   SecurityEdiscoverySearchExportOperationStatus = "submissionFailed"
	SecurityEdiscoverySearchExportOperationStatus_Succeeded          SecurityEdiscoverySearchExportOperationStatus = "succeeded"
)

func PossibleValuesForSecurityEdiscoverySearchExportOperationStatus() []string {
	return []string{
		string(SecurityEdiscoverySearchExportOperationStatus_Failed),
		string(SecurityEdiscoverySearchExportOperationStatus_NotStarted),
		string(SecurityEdiscoverySearchExportOperationStatus_PartiallySucceeded),
		string(SecurityEdiscoverySearchExportOperationStatus_Running),
		string(SecurityEdiscoverySearchExportOperationStatus_SubmissionFailed),
		string(SecurityEdiscoverySearchExportOperationStatus_Succeeded),
	}
}

func (s *SecurityEdiscoverySearchExportOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEdiscoverySearchExportOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEdiscoverySearchExportOperationStatus(input string) (*SecurityEdiscoverySearchExportOperationStatus, error) {
	vals := map[string]SecurityEdiscoverySearchExportOperationStatus{
		"failed":             SecurityEdiscoverySearchExportOperationStatus_Failed,
		"notstarted":         SecurityEdiscoverySearchExportOperationStatus_NotStarted,
		"partiallysucceeded": SecurityEdiscoverySearchExportOperationStatus_PartiallySucceeded,
		"running":            SecurityEdiscoverySearchExportOperationStatus_Running,
		"submissionfailed":   SecurityEdiscoverySearchExportOperationStatus_SubmissionFailed,
		"succeeded":          SecurityEdiscoverySearchExportOperationStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoverySearchExportOperationStatus(input)
	return &out, nil
}
