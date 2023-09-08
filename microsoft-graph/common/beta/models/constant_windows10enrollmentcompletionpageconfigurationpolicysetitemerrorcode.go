package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode string

const (
	Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCodedeleted      Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode = "Deleted"
	Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCodenoError      Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode = "NoError"
	Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCodenotFound     Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode = "NotFound"
	Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCodeunauthorized Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode = "Unauthorized"
)

func PossibleValuesForWindows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode() []string {
	return []string{
		string(Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCodedeleted),
		string(Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCodenoError),
		string(Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCodenotFound),
		string(Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCodeunauthorized),
	}
}

func parseWindows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode(input string) (*Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode, error) {
	vals := map[string]Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode{
		"deleted":      Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCodedeleted,
		"noerror":      Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCodenoError,
		"notfound":     Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCodenotFound,
		"unauthorized": Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCodeunauthorized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EnrollmentCompletionPageConfigurationPolicySetItemErrorCode(input)
	return &out, nil
}
