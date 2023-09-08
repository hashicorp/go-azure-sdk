package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetResourceGroupType string

const (
	TargetResourceGroupTypeazureAD       TargetResourceGroupType = "AzureAD"
	TargetResourceGroupTypeunifiedGroups TargetResourceGroupType = "UnifiedGroups"
)

func PossibleValuesForTargetResourceGroupType() []string {
	return []string{
		string(TargetResourceGroupTypeazureAD),
		string(TargetResourceGroupTypeunifiedGroups),
	}
}

func parseTargetResourceGroupType(input string) (*TargetResourceGroupType, error) {
	vals := map[string]TargetResourceGroupType{
		"azuread":       TargetResourceGroupTypeazureAD,
		"unifiedgroups": TargetResourceGroupTypeunifiedGroups,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetResourceGroupType(input)
	return &out, nil
}
