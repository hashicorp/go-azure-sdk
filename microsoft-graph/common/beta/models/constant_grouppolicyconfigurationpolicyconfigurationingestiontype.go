package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyConfigurationPolicyConfigurationIngestionType string

const (
	GroupPolicyConfigurationPolicyConfigurationIngestionTypebuiltIn GroupPolicyConfigurationPolicyConfigurationIngestionType = "BuiltIn"
	GroupPolicyConfigurationPolicyConfigurationIngestionTypecustom  GroupPolicyConfigurationPolicyConfigurationIngestionType = "Custom"
	GroupPolicyConfigurationPolicyConfigurationIngestionTypemixed   GroupPolicyConfigurationPolicyConfigurationIngestionType = "Mixed"
	GroupPolicyConfigurationPolicyConfigurationIngestionTypeunknown GroupPolicyConfigurationPolicyConfigurationIngestionType = "Unknown"
)

func PossibleValuesForGroupPolicyConfigurationPolicyConfigurationIngestionType() []string {
	return []string{
		string(GroupPolicyConfigurationPolicyConfigurationIngestionTypebuiltIn),
		string(GroupPolicyConfigurationPolicyConfigurationIngestionTypecustom),
		string(GroupPolicyConfigurationPolicyConfigurationIngestionTypemixed),
		string(GroupPolicyConfigurationPolicyConfigurationIngestionTypeunknown),
	}
}

func parseGroupPolicyConfigurationPolicyConfigurationIngestionType(input string) (*GroupPolicyConfigurationPolicyConfigurationIngestionType, error) {
	vals := map[string]GroupPolicyConfigurationPolicyConfigurationIngestionType{
		"builtin": GroupPolicyConfigurationPolicyConfigurationIngestionTypebuiltIn,
		"custom":  GroupPolicyConfigurationPolicyConfigurationIngestionTypecustom,
		"mixed":   GroupPolicyConfigurationPolicyConfigurationIngestionTypemixed,
		"unknown": GroupPolicyConfigurationPolicyConfigurationIngestionTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicyConfigurationPolicyConfigurationIngestionType(input)
	return &out, nil
}
