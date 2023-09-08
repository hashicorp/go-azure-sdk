package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyDefinitionFilePolicyType string

const (
	GroupPolicyDefinitionFilePolicyTypeadmxBacked   GroupPolicyDefinitionFilePolicyType = "AdmxBacked"
	GroupPolicyDefinitionFilePolicyTypeadmxIngested GroupPolicyDefinitionFilePolicyType = "AdmxIngested"
)

func PossibleValuesForGroupPolicyDefinitionFilePolicyType() []string {
	return []string{
		string(GroupPolicyDefinitionFilePolicyTypeadmxBacked),
		string(GroupPolicyDefinitionFilePolicyTypeadmxIngested),
	}
}

func parseGroupPolicyDefinitionFilePolicyType(input string) (*GroupPolicyDefinitionFilePolicyType, error) {
	vals := map[string]GroupPolicyDefinitionFilePolicyType{
		"admxbacked":   GroupPolicyDefinitionFilePolicyTypeadmxBacked,
		"admxingested": GroupPolicyDefinitionFilePolicyTypeadmxIngested,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicyDefinitionFilePolicyType(input)
	return &out, nil
}
