package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppInstallExperienceRunAsAccount string

const (
	Win32LobAppInstallExperienceRunAsAccountsystem Win32LobAppInstallExperienceRunAsAccount = "System"
	Win32LobAppInstallExperienceRunAsAccountuser   Win32LobAppInstallExperienceRunAsAccount = "User"
)

func PossibleValuesForWin32LobAppInstallExperienceRunAsAccount() []string {
	return []string{
		string(Win32LobAppInstallExperienceRunAsAccountsystem),
		string(Win32LobAppInstallExperienceRunAsAccountuser),
	}
}

func parseWin32LobAppInstallExperienceRunAsAccount(input string) (*Win32LobAppInstallExperienceRunAsAccount, error) {
	vals := map[string]Win32LobAppInstallExperienceRunAsAccount{
		"system": Win32LobAppInstallExperienceRunAsAccountsystem,
		"user":   Win32LobAppInstallExperienceRunAsAccountuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppInstallExperienceRunAsAccount(input)
	return &out, nil
}
