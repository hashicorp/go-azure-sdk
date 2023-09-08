package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedPCConfigurationSetEduPolicies string

const (
	SharedPCConfigurationSetEduPoliciesdisabled      SharedPCConfigurationSetEduPolicies = "Disabled"
	SharedPCConfigurationSetEduPoliciesenabled       SharedPCConfigurationSetEduPolicies = "Enabled"
	SharedPCConfigurationSetEduPoliciesnotConfigured SharedPCConfigurationSetEduPolicies = "NotConfigured"
)

func PossibleValuesForSharedPCConfigurationSetEduPolicies() []string {
	return []string{
		string(SharedPCConfigurationSetEduPoliciesdisabled),
		string(SharedPCConfigurationSetEduPoliciesenabled),
		string(SharedPCConfigurationSetEduPoliciesnotConfigured),
	}
}

func parseSharedPCConfigurationSetEduPolicies(input string) (*SharedPCConfigurationSetEduPolicies, error) {
	vals := map[string]SharedPCConfigurationSetEduPolicies{
		"disabled":      SharedPCConfigurationSetEduPoliciesdisabled,
		"enabled":       SharedPCConfigurationSetEduPoliciesenabled,
		"notconfigured": SharedPCConfigurationSetEduPoliciesnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharedPCConfigurationSetEduPolicies(input)
	return &out, nil
}
