package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppRegistryRequirementOperator string

const (
	Win32LobAppRegistryRequirementOperatorequal              Win32LobAppRegistryRequirementOperator = "Equal"
	Win32LobAppRegistryRequirementOperatorgreaterThan        Win32LobAppRegistryRequirementOperator = "GreaterThan"
	Win32LobAppRegistryRequirementOperatorgreaterThanOrEqual Win32LobAppRegistryRequirementOperator = "GreaterThanOrEqual"
	Win32LobAppRegistryRequirementOperatorlessThan           Win32LobAppRegistryRequirementOperator = "LessThan"
	Win32LobAppRegistryRequirementOperatorlessThanOrEqual    Win32LobAppRegistryRequirementOperator = "LessThanOrEqual"
	Win32LobAppRegistryRequirementOperatornotConfigured      Win32LobAppRegistryRequirementOperator = "NotConfigured"
	Win32LobAppRegistryRequirementOperatornotEqual           Win32LobAppRegistryRequirementOperator = "NotEqual"
)

func PossibleValuesForWin32LobAppRegistryRequirementOperator() []string {
	return []string{
		string(Win32LobAppRegistryRequirementOperatorequal),
		string(Win32LobAppRegistryRequirementOperatorgreaterThan),
		string(Win32LobAppRegistryRequirementOperatorgreaterThanOrEqual),
		string(Win32LobAppRegistryRequirementOperatorlessThan),
		string(Win32LobAppRegistryRequirementOperatorlessThanOrEqual),
		string(Win32LobAppRegistryRequirementOperatornotConfigured),
		string(Win32LobAppRegistryRequirementOperatornotEqual),
	}
}

func parseWin32LobAppRegistryRequirementOperator(input string) (*Win32LobAppRegistryRequirementOperator, error) {
	vals := map[string]Win32LobAppRegistryRequirementOperator{
		"equal":              Win32LobAppRegistryRequirementOperatorequal,
		"greaterthan":        Win32LobAppRegistryRequirementOperatorgreaterThan,
		"greaterthanorequal": Win32LobAppRegistryRequirementOperatorgreaterThanOrEqual,
		"lessthan":           Win32LobAppRegistryRequirementOperatorlessThan,
		"lessthanorequal":    Win32LobAppRegistryRequirementOperatorlessThanOrEqual,
		"notconfigured":      Win32LobAppRegistryRequirementOperatornotConfigured,
		"notequal":           Win32LobAppRegistryRequirementOperatornotEqual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppRegistryRequirementOperator(input)
	return &out, nil
}
