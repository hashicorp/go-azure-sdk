package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrollmentRestrictionsConfigurationPolicySetItemErrorCode string

const (
	EnrollmentRestrictionsConfigurationPolicySetItemErrorCodedeleted      EnrollmentRestrictionsConfigurationPolicySetItemErrorCode = "Deleted"
	EnrollmentRestrictionsConfigurationPolicySetItemErrorCodenoError      EnrollmentRestrictionsConfigurationPolicySetItemErrorCode = "NoError"
	EnrollmentRestrictionsConfigurationPolicySetItemErrorCodenotFound     EnrollmentRestrictionsConfigurationPolicySetItemErrorCode = "NotFound"
	EnrollmentRestrictionsConfigurationPolicySetItemErrorCodeunauthorized EnrollmentRestrictionsConfigurationPolicySetItemErrorCode = "Unauthorized"
)

func PossibleValuesForEnrollmentRestrictionsConfigurationPolicySetItemErrorCode() []string {
	return []string{
		string(EnrollmentRestrictionsConfigurationPolicySetItemErrorCodedeleted),
		string(EnrollmentRestrictionsConfigurationPolicySetItemErrorCodenoError),
		string(EnrollmentRestrictionsConfigurationPolicySetItemErrorCodenotFound),
		string(EnrollmentRestrictionsConfigurationPolicySetItemErrorCodeunauthorized),
	}
}

func parseEnrollmentRestrictionsConfigurationPolicySetItemErrorCode(input string) (*EnrollmentRestrictionsConfigurationPolicySetItemErrorCode, error) {
	vals := map[string]EnrollmentRestrictionsConfigurationPolicySetItemErrorCode{
		"deleted":      EnrollmentRestrictionsConfigurationPolicySetItemErrorCodedeleted,
		"noerror":      EnrollmentRestrictionsConfigurationPolicySetItemErrorCodenoError,
		"notfound":     EnrollmentRestrictionsConfigurationPolicySetItemErrorCodenotFound,
		"unauthorized": EnrollmentRestrictionsConfigurationPolicySetItemErrorCodeunauthorized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnrollmentRestrictionsConfigurationPolicySetItemErrorCode(input)
	return &out, nil
}
