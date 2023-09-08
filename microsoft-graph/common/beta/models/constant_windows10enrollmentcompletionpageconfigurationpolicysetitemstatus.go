package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus string

const (
	Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatuserror          Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus = "Error"
	Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatusnotAssigned    Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus = "NotAssigned"
	Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatuspartialSuccess Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus = "PartialSuccess"
	Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatussuccess        Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus = "Success"
	Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatusunknown        Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus = "Unknown"
	Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatusvalidating     Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus = "Validating"
)

func PossibleValuesForWindows10EnrollmentCompletionPageConfigurationPolicySetItemStatus() []string {
	return []string{
		string(Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatuserror),
		string(Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatusnotAssigned),
		string(Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatuspartialSuccess),
		string(Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatussuccess),
		string(Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatusunknown),
		string(Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatusvalidating),
	}
}

func parseWindows10EnrollmentCompletionPageConfigurationPolicySetItemStatus(input string) (*Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus, error) {
	vals := map[string]Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus{
		"error":          Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatuserror,
		"notassigned":    Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatusnotAssigned,
		"partialsuccess": Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatuspartialSuccess,
		"success":        Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatussuccess,
		"unknown":        Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatusunknown,
		"validating":     Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatusvalidating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EnrollmentCompletionPageConfigurationPolicySetItemStatus(input)
	return &out, nil
}
