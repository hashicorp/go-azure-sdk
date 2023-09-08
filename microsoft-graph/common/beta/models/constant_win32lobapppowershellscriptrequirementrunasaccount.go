package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppPowerShellScriptRequirementRunAsAccount string

const (
	Win32LobAppPowerShellScriptRequirementRunAsAccountsystem Win32LobAppPowerShellScriptRequirementRunAsAccount = "System"
	Win32LobAppPowerShellScriptRequirementRunAsAccountuser   Win32LobAppPowerShellScriptRequirementRunAsAccount = "User"
)

func PossibleValuesForWin32LobAppPowerShellScriptRequirementRunAsAccount() []string {
	return []string{
		string(Win32LobAppPowerShellScriptRequirementRunAsAccountsystem),
		string(Win32LobAppPowerShellScriptRequirementRunAsAccountuser),
	}
}

func parseWin32LobAppPowerShellScriptRequirementRunAsAccount(input string) (*Win32LobAppPowerShellScriptRequirementRunAsAccount, error) {
	vals := map[string]Win32LobAppPowerShellScriptRequirementRunAsAccount{
		"system": Win32LobAppPowerShellScriptRequirementRunAsAccountsystem,
		"user":   Win32LobAppPowerShellScriptRequirementRunAsAccountuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppPowerShellScriptRequirementRunAsAccount(input)
	return &out, nil
}
