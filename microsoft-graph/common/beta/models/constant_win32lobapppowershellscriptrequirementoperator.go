package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppPowerShellScriptRequirementOperator string

const (
	Win32LobAppPowerShellScriptRequirementOperatorequal              Win32LobAppPowerShellScriptRequirementOperator = "Equal"
	Win32LobAppPowerShellScriptRequirementOperatorgreaterThan        Win32LobAppPowerShellScriptRequirementOperator = "GreaterThan"
	Win32LobAppPowerShellScriptRequirementOperatorgreaterThanOrEqual Win32LobAppPowerShellScriptRequirementOperator = "GreaterThanOrEqual"
	Win32LobAppPowerShellScriptRequirementOperatorlessThan           Win32LobAppPowerShellScriptRequirementOperator = "LessThan"
	Win32LobAppPowerShellScriptRequirementOperatorlessThanOrEqual    Win32LobAppPowerShellScriptRequirementOperator = "LessThanOrEqual"
	Win32LobAppPowerShellScriptRequirementOperatornotConfigured      Win32LobAppPowerShellScriptRequirementOperator = "NotConfigured"
	Win32LobAppPowerShellScriptRequirementOperatornotEqual           Win32LobAppPowerShellScriptRequirementOperator = "NotEqual"
)

func PossibleValuesForWin32LobAppPowerShellScriptRequirementOperator() []string {
	return []string{
		string(Win32LobAppPowerShellScriptRequirementOperatorequal),
		string(Win32LobAppPowerShellScriptRequirementOperatorgreaterThan),
		string(Win32LobAppPowerShellScriptRequirementOperatorgreaterThanOrEqual),
		string(Win32LobAppPowerShellScriptRequirementOperatorlessThan),
		string(Win32LobAppPowerShellScriptRequirementOperatorlessThanOrEqual),
		string(Win32LobAppPowerShellScriptRequirementOperatornotConfigured),
		string(Win32LobAppPowerShellScriptRequirementOperatornotEqual),
	}
}

func parseWin32LobAppPowerShellScriptRequirementOperator(input string) (*Win32LobAppPowerShellScriptRequirementOperator, error) {
	vals := map[string]Win32LobAppPowerShellScriptRequirementOperator{
		"equal":              Win32LobAppPowerShellScriptRequirementOperatorequal,
		"greaterthan":        Win32LobAppPowerShellScriptRequirementOperatorgreaterThan,
		"greaterthanorequal": Win32LobAppPowerShellScriptRequirementOperatorgreaterThanOrEqual,
		"lessthan":           Win32LobAppPowerShellScriptRequirementOperatorlessThan,
		"lessthanorequal":    Win32LobAppPowerShellScriptRequirementOperatorlessThanOrEqual,
		"notconfigured":      Win32LobAppPowerShellScriptRequirementOperatornotConfigured,
		"notequal":           Win32LobAppPowerShellScriptRequirementOperatornotEqual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppPowerShellScriptRequirementOperator(input)
	return &out, nil
}
