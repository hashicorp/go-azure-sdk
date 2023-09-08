package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppRequirementOperator string

const (
	Win32LobAppRequirementOperatorequal              Win32LobAppRequirementOperator = "Equal"
	Win32LobAppRequirementOperatorgreaterThan        Win32LobAppRequirementOperator = "GreaterThan"
	Win32LobAppRequirementOperatorgreaterThanOrEqual Win32LobAppRequirementOperator = "GreaterThanOrEqual"
	Win32LobAppRequirementOperatorlessThan           Win32LobAppRequirementOperator = "LessThan"
	Win32LobAppRequirementOperatorlessThanOrEqual    Win32LobAppRequirementOperator = "LessThanOrEqual"
	Win32LobAppRequirementOperatornotConfigured      Win32LobAppRequirementOperator = "NotConfigured"
	Win32LobAppRequirementOperatornotEqual           Win32LobAppRequirementOperator = "NotEqual"
)

func PossibleValuesForWin32LobAppRequirementOperator() []string {
	return []string{
		string(Win32LobAppRequirementOperatorequal),
		string(Win32LobAppRequirementOperatorgreaterThan),
		string(Win32LobAppRequirementOperatorgreaterThanOrEqual),
		string(Win32LobAppRequirementOperatorlessThan),
		string(Win32LobAppRequirementOperatorlessThanOrEqual),
		string(Win32LobAppRequirementOperatornotConfigured),
		string(Win32LobAppRequirementOperatornotEqual),
	}
}

func parseWin32LobAppRequirementOperator(input string) (*Win32LobAppRequirementOperator, error) {
	vals := map[string]Win32LobAppRequirementOperator{
		"equal":              Win32LobAppRequirementOperatorequal,
		"greaterthan":        Win32LobAppRequirementOperatorgreaterThan,
		"greaterthanorequal": Win32LobAppRequirementOperatorgreaterThanOrEqual,
		"lessthan":           Win32LobAppRequirementOperatorlessThan,
		"lessthanorequal":    Win32LobAppRequirementOperatorlessThanOrEqual,
		"notconfigured":      Win32LobAppRequirementOperatornotConfigured,
		"notequal":           Win32LobAppRequirementOperatornotEqual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppRequirementOperator(input)
	return &out, nil
}
