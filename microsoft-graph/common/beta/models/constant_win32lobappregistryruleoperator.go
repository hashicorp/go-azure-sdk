package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppRegistryRuleOperator string

const (
	Win32LobAppRegistryRuleOperatorequal              Win32LobAppRegistryRuleOperator = "Equal"
	Win32LobAppRegistryRuleOperatorgreaterThan        Win32LobAppRegistryRuleOperator = "GreaterThan"
	Win32LobAppRegistryRuleOperatorgreaterThanOrEqual Win32LobAppRegistryRuleOperator = "GreaterThanOrEqual"
	Win32LobAppRegistryRuleOperatorlessThan           Win32LobAppRegistryRuleOperator = "LessThan"
	Win32LobAppRegistryRuleOperatorlessThanOrEqual    Win32LobAppRegistryRuleOperator = "LessThanOrEqual"
	Win32LobAppRegistryRuleOperatornotConfigured      Win32LobAppRegistryRuleOperator = "NotConfigured"
	Win32LobAppRegistryRuleOperatornotEqual           Win32LobAppRegistryRuleOperator = "NotEqual"
)

func PossibleValuesForWin32LobAppRegistryRuleOperator() []string {
	return []string{
		string(Win32LobAppRegistryRuleOperatorequal),
		string(Win32LobAppRegistryRuleOperatorgreaterThan),
		string(Win32LobAppRegistryRuleOperatorgreaterThanOrEqual),
		string(Win32LobAppRegistryRuleOperatorlessThan),
		string(Win32LobAppRegistryRuleOperatorlessThanOrEqual),
		string(Win32LobAppRegistryRuleOperatornotConfigured),
		string(Win32LobAppRegistryRuleOperatornotEqual),
	}
}

func parseWin32LobAppRegistryRuleOperator(input string) (*Win32LobAppRegistryRuleOperator, error) {
	vals := map[string]Win32LobAppRegistryRuleOperator{
		"equal":              Win32LobAppRegistryRuleOperatorequal,
		"greaterthan":        Win32LobAppRegistryRuleOperatorgreaterThan,
		"greaterthanorequal": Win32LobAppRegistryRuleOperatorgreaterThanOrEqual,
		"lessthan":           Win32LobAppRegistryRuleOperatorlessThan,
		"lessthanorequal":    Win32LobAppRegistryRuleOperatorlessThanOrEqual,
		"notconfigured":      Win32LobAppRegistryRuleOperatornotConfigured,
		"notequal":           Win32LobAppRegistryRuleOperatornotEqual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppRegistryRuleOperator(input)
	return &out, nil
}
