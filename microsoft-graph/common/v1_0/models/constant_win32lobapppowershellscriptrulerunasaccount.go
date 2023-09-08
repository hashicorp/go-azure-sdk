package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppPowerShellScriptRuleRunAsAccount string

const (
	Win32LobAppPowerShellScriptRuleRunAsAccountsystem Win32LobAppPowerShellScriptRuleRunAsAccount = "System"
	Win32LobAppPowerShellScriptRuleRunAsAccountuser   Win32LobAppPowerShellScriptRuleRunAsAccount = "User"
)

func PossibleValuesForWin32LobAppPowerShellScriptRuleRunAsAccount() []string {
	return []string{
		string(Win32LobAppPowerShellScriptRuleRunAsAccountsystem),
		string(Win32LobAppPowerShellScriptRuleRunAsAccountuser),
	}
}

func parseWin32LobAppPowerShellScriptRuleRunAsAccount(input string) (*Win32LobAppPowerShellScriptRuleRunAsAccount, error) {
	vals := map[string]Win32LobAppPowerShellScriptRuleRunAsAccount{
		"system": Win32LobAppPowerShellScriptRuleRunAsAccountsystem,
		"user":   Win32LobAppPowerShellScriptRuleRunAsAccountuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppPowerShellScriptRuleRunAsAccount(input)
	return &out, nil
}
