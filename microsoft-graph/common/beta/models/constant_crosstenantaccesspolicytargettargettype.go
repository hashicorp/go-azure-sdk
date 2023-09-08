package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CrossTenantAccessPolicyTargetTargetType string

const (
	CrossTenantAccessPolicyTargetTargetTypeapplication CrossTenantAccessPolicyTargetTargetType = "Application"
	CrossTenantAccessPolicyTargetTargetTypegroup       CrossTenantAccessPolicyTargetTargetType = "Group"
	CrossTenantAccessPolicyTargetTargetTypeuser        CrossTenantAccessPolicyTargetTargetType = "User"
)

func PossibleValuesForCrossTenantAccessPolicyTargetTargetType() []string {
	return []string{
		string(CrossTenantAccessPolicyTargetTargetTypeapplication),
		string(CrossTenantAccessPolicyTargetTargetTypegroup),
		string(CrossTenantAccessPolicyTargetTargetTypeuser),
	}
}

func parseCrossTenantAccessPolicyTargetTargetType(input string) (*CrossTenantAccessPolicyTargetTargetType, error) {
	vals := map[string]CrossTenantAccessPolicyTargetTargetType{
		"application": CrossTenantAccessPolicyTargetTargetTypeapplication,
		"group":       CrossTenantAccessPolicyTargetTargetTypegroup,
		"user":        CrossTenantAccessPolicyTargetTargetTypeuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CrossTenantAccessPolicyTargetTargetType(input)
	return &out, nil
}
