package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyDefinitionClassType string

const (
	GroupPolicyDefinitionClassTypemachine GroupPolicyDefinitionClassType = "Machine"
	GroupPolicyDefinitionClassTypeuser    GroupPolicyDefinitionClassType = "User"
)

func PossibleValuesForGroupPolicyDefinitionClassType() []string {
	return []string{
		string(GroupPolicyDefinitionClassTypemachine),
		string(GroupPolicyDefinitionClassTypeuser),
	}
}

func parseGroupPolicyDefinitionClassType(input string) (*GroupPolicyDefinitionClassType, error) {
	vals := map[string]GroupPolicyDefinitionClassType{
		"machine": GroupPolicyDefinitionClassTypemachine,
		"user":    GroupPolicyDefinitionClassTypeuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicyDefinitionClassType(input)
	return &out, nil
}
