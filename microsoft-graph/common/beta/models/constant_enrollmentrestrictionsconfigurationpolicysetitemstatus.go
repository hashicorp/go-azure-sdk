package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrollmentRestrictionsConfigurationPolicySetItemStatus string

const (
	EnrollmentRestrictionsConfigurationPolicySetItemStatuserror          EnrollmentRestrictionsConfigurationPolicySetItemStatus = "Error"
	EnrollmentRestrictionsConfigurationPolicySetItemStatusnotAssigned    EnrollmentRestrictionsConfigurationPolicySetItemStatus = "NotAssigned"
	EnrollmentRestrictionsConfigurationPolicySetItemStatuspartialSuccess EnrollmentRestrictionsConfigurationPolicySetItemStatus = "PartialSuccess"
	EnrollmentRestrictionsConfigurationPolicySetItemStatussuccess        EnrollmentRestrictionsConfigurationPolicySetItemStatus = "Success"
	EnrollmentRestrictionsConfigurationPolicySetItemStatusunknown        EnrollmentRestrictionsConfigurationPolicySetItemStatus = "Unknown"
	EnrollmentRestrictionsConfigurationPolicySetItemStatusvalidating     EnrollmentRestrictionsConfigurationPolicySetItemStatus = "Validating"
)

func PossibleValuesForEnrollmentRestrictionsConfigurationPolicySetItemStatus() []string {
	return []string{
		string(EnrollmentRestrictionsConfigurationPolicySetItemStatuserror),
		string(EnrollmentRestrictionsConfigurationPolicySetItemStatusnotAssigned),
		string(EnrollmentRestrictionsConfigurationPolicySetItemStatuspartialSuccess),
		string(EnrollmentRestrictionsConfigurationPolicySetItemStatussuccess),
		string(EnrollmentRestrictionsConfigurationPolicySetItemStatusunknown),
		string(EnrollmentRestrictionsConfigurationPolicySetItemStatusvalidating),
	}
}

func parseEnrollmentRestrictionsConfigurationPolicySetItemStatus(input string) (*EnrollmentRestrictionsConfigurationPolicySetItemStatus, error) {
	vals := map[string]EnrollmentRestrictionsConfigurationPolicySetItemStatus{
		"error":          EnrollmentRestrictionsConfigurationPolicySetItemStatuserror,
		"notassigned":    EnrollmentRestrictionsConfigurationPolicySetItemStatusnotAssigned,
		"partialsuccess": EnrollmentRestrictionsConfigurationPolicySetItemStatuspartialSuccess,
		"success":        EnrollmentRestrictionsConfigurationPolicySetItemStatussuccess,
		"unknown":        EnrollmentRestrictionsConfigurationPolicySetItemStatusunknown,
		"validating":     EnrollmentRestrictionsConfigurationPolicySetItemStatusvalidating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnrollmentRestrictionsConfigurationPolicySetItemStatus(input)
	return &out, nil
}
