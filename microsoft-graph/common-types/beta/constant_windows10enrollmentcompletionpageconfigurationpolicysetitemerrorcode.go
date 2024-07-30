package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode string

const (
	Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode_Deleted      Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode = "deleted"
	Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode_NoError      Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode = "noError"
	Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode_NotFound     Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode = "notFound"
	Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode_Unauthorized Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode = "unauthorized"
)

func PossibleValuesForWindows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode() []string {
	return []string{
		string(Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode_Deleted),
		string(Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode_NoError),
		string(Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode_NotFound),
		string(Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode_Unauthorized),
	}
}

func (s *Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode(input string) (*Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode, error) {
	vals := map[string]Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode{
		"deleted":      Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode_Deleted,
		"noerror":      Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode_NoError,
		"notfound":     Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode_NotFound,
		"unauthorized": Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode_Unauthorized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode(input)
	return &out, nil
}
