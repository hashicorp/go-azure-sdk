package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppPowerShellScriptRuleRuleType string

const (
	Win32LobAppPowerShellScriptRuleRuleTypedetection   Win32LobAppPowerShellScriptRuleRuleType = "Detection"
	Win32LobAppPowerShellScriptRuleRuleTyperequirement Win32LobAppPowerShellScriptRuleRuleType = "Requirement"
)

func PossibleValuesForWin32LobAppPowerShellScriptRuleRuleType() []string {
	return []string{
		string(Win32LobAppPowerShellScriptRuleRuleTypedetection),
		string(Win32LobAppPowerShellScriptRuleRuleTyperequirement),
	}
}

func parseWin32LobAppPowerShellScriptRuleRuleType(input string) (*Win32LobAppPowerShellScriptRuleRuleType, error) {
	vals := map[string]Win32LobAppPowerShellScriptRuleRuleType{
		"detection":   Win32LobAppPowerShellScriptRuleRuleTypedetection,
		"requirement": Win32LobAppPowerShellScriptRuleRuleTyperequirement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppPowerShellScriptRuleRuleType(input)
	return &out, nil
}
