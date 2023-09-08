package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyDefinitionPolicyType string

const (
	GroupPolicyDefinitionPolicyTypeadmxBacked   GroupPolicyDefinitionPolicyType = "AdmxBacked"
	GroupPolicyDefinitionPolicyTypeadmxIngested GroupPolicyDefinitionPolicyType = "AdmxIngested"
)

func PossibleValuesForGroupPolicyDefinitionPolicyType() []string {
	return []string{
		string(GroupPolicyDefinitionPolicyTypeadmxBacked),
		string(GroupPolicyDefinitionPolicyTypeadmxIngested),
	}
}

func parseGroupPolicyDefinitionPolicyType(input string) (*GroupPolicyDefinitionPolicyType, error) {
	vals := map[string]GroupPolicyDefinitionPolicyType{
		"admxbacked":   GroupPolicyDefinitionPolicyTypeadmxBacked,
		"admxingested": GroupPolicyDefinitionPolicyTypeadmxIngested,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicyDefinitionPolicyType(input)
	return &out, nil
}
