package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppFileSystemRuleOperator string

const (
	Win32LobAppFileSystemRuleOperatorequal              Win32LobAppFileSystemRuleOperator = "Equal"
	Win32LobAppFileSystemRuleOperatorgreaterThan        Win32LobAppFileSystemRuleOperator = "GreaterThan"
	Win32LobAppFileSystemRuleOperatorgreaterThanOrEqual Win32LobAppFileSystemRuleOperator = "GreaterThanOrEqual"
	Win32LobAppFileSystemRuleOperatorlessThan           Win32LobAppFileSystemRuleOperator = "LessThan"
	Win32LobAppFileSystemRuleOperatorlessThanOrEqual    Win32LobAppFileSystemRuleOperator = "LessThanOrEqual"
	Win32LobAppFileSystemRuleOperatornotConfigured      Win32LobAppFileSystemRuleOperator = "NotConfigured"
	Win32LobAppFileSystemRuleOperatornotEqual           Win32LobAppFileSystemRuleOperator = "NotEqual"
)

func PossibleValuesForWin32LobAppFileSystemRuleOperator() []string {
	return []string{
		string(Win32LobAppFileSystemRuleOperatorequal),
		string(Win32LobAppFileSystemRuleOperatorgreaterThan),
		string(Win32LobAppFileSystemRuleOperatorgreaterThanOrEqual),
		string(Win32LobAppFileSystemRuleOperatorlessThan),
		string(Win32LobAppFileSystemRuleOperatorlessThanOrEqual),
		string(Win32LobAppFileSystemRuleOperatornotConfigured),
		string(Win32LobAppFileSystemRuleOperatornotEqual),
	}
}

func parseWin32LobAppFileSystemRuleOperator(input string) (*Win32LobAppFileSystemRuleOperator, error) {
	vals := map[string]Win32LobAppFileSystemRuleOperator{
		"equal":              Win32LobAppFileSystemRuleOperatorequal,
		"greaterthan":        Win32LobAppFileSystemRuleOperatorgreaterThan,
		"greaterthanorequal": Win32LobAppFileSystemRuleOperatorgreaterThanOrEqual,
		"lessthan":           Win32LobAppFileSystemRuleOperatorlessThan,
		"lessthanorequal":    Win32LobAppFileSystemRuleOperatorlessThanOrEqual,
		"notconfigured":      Win32LobAppFileSystemRuleOperatornotConfigured,
		"notequal":           Win32LobAppFileSystemRuleOperatornotEqual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppFileSystemRuleOperator(input)
	return &out, nil
}
