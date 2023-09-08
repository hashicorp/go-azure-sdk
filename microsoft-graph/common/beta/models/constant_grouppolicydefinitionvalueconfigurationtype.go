package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyDefinitionValueConfigurationType string

const (
	GroupPolicyDefinitionValueConfigurationTypepolicy     GroupPolicyDefinitionValueConfigurationType = "Policy"
	GroupPolicyDefinitionValueConfigurationTypepreference GroupPolicyDefinitionValueConfigurationType = "Preference"
)

func PossibleValuesForGroupPolicyDefinitionValueConfigurationType() []string {
	return []string{
		string(GroupPolicyDefinitionValueConfigurationTypepolicy),
		string(GroupPolicyDefinitionValueConfigurationTypepreference),
	}
}

func parseGroupPolicyDefinitionValueConfigurationType(input string) (*GroupPolicyDefinitionValueConfigurationType, error) {
	vals := map[string]GroupPolicyDefinitionValueConfigurationType{
		"policy":     GroupPolicyDefinitionValueConfigurationTypepolicy,
		"preference": GroupPolicyDefinitionValueConfigurationTypepreference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicyDefinitionValueConfigurationType(input)
	return &out, nil
}
