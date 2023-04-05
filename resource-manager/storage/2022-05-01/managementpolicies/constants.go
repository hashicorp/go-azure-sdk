package managementpolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RuleType string

const (
	RuleTypeLifecycle RuleType = "Lifecycle"
)

func PossibleValuesForRuleType() []string {
	return []string{
		string(RuleTypeLifecycle),
	}
}
