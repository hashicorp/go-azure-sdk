package managementpolicies

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagementPolicyName string

const (
	ManagementPolicyNameDefault ManagementPolicyName = "default"
)

func PossibleValuesForManagementPolicyName() []string {
	return []string{
		string(ManagementPolicyNameDefault),
	}
}

func parseManagementPolicyName(input string) (*ManagementPolicyName, error) {
	vals := map[string]ManagementPolicyName{
		"default": ManagementPolicyNameDefault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagementPolicyName(input)
	return &out, nil
}

type RuleType string

const (
	RuleTypeLifecycle RuleType = "Lifecycle"
)

func PossibleValuesForRuleType() []string {
	return []string{
		string(RuleTypeLifecycle),
	}
}

func parseRuleType(input string) (*RuleType, error) {
	vals := map[string]RuleType{
		"lifecycle": RuleTypeLifecycle,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RuleType(input)
	return &out, nil
}
