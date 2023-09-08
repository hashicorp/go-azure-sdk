package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CrossTenantAccessPolicyTargetConfigurationAccessType string

const (
	CrossTenantAccessPolicyTargetConfigurationAccessTypeallowed CrossTenantAccessPolicyTargetConfigurationAccessType = "Allowed"
	CrossTenantAccessPolicyTargetConfigurationAccessTypeblocked CrossTenantAccessPolicyTargetConfigurationAccessType = "Blocked"
)

func PossibleValuesForCrossTenantAccessPolicyTargetConfigurationAccessType() []string {
	return []string{
		string(CrossTenantAccessPolicyTargetConfigurationAccessTypeallowed),
		string(CrossTenantAccessPolicyTargetConfigurationAccessTypeblocked),
	}
}

func parseCrossTenantAccessPolicyTargetConfigurationAccessType(input string) (*CrossTenantAccessPolicyTargetConfigurationAccessType, error) {
	vals := map[string]CrossTenantAccessPolicyTargetConfigurationAccessType{
		"allowed": CrossTenantAccessPolicyTargetConfigurationAccessTypeallowed,
		"blocked": CrossTenantAccessPolicyTargetConfigurationAccessTypeblocked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CrossTenantAccessPolicyTargetConfigurationAccessType(input)
	return &out, nil
}
