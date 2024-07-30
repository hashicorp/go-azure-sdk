package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus string

const (
	Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus_Error          Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus = "error"
	Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus_NotAssigned    Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus = "notAssigned"
	Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus_PartialSuccess Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus = "partialSuccess"
	Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus_Success        Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus = "success"
	Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus_Unknown        Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus = "unknown"
	Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus_Validating     Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus = "validating"
)

func PossibleValuesForWindows10EnrollmentCompletionPageConfigurationPolicySetItemStatus() []string {
	return []string{
		string(Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus_Error),
		string(Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus_NotAssigned),
		string(Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus_PartialSuccess),
		string(Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus_Success),
		string(Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus_Unknown),
		string(Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus_Validating),
	}
}

func (s *Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EnrollmentCompletionPageConfigurationPolicySetItemStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EnrollmentCompletionPageConfigurationPolicySetItemStatus(input string) (*Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus, error) {
	vals := map[string]Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus{
		"error":          Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus_Error,
		"notassigned":    Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus_NotAssigned,
		"partialsuccess": Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus_PartialSuccess,
		"success":        Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus_Success,
		"unknown":        Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus_Unknown,
		"validating":     Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus_Validating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus(input)
	return &out, nil
}
