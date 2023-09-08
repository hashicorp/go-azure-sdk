package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateStateStatus string

const (
	WindowsUpdateStateStatusfailed              WindowsUpdateStateStatus = "Failed"
	WindowsUpdateStateStatuspendingInstallation WindowsUpdateStateStatus = "PendingInstallation"
	WindowsUpdateStateStatuspendingReboot       WindowsUpdateStateStatus = "PendingReboot"
	WindowsUpdateStateStatusupToDate            WindowsUpdateStateStatus = "UpToDate"
)

func PossibleValuesForWindowsUpdateStateStatus() []string {
	return []string{
		string(WindowsUpdateStateStatusfailed),
		string(WindowsUpdateStateStatuspendingInstallation),
		string(WindowsUpdateStateStatuspendingReboot),
		string(WindowsUpdateStateStatusupToDate),
	}
}

func parseWindowsUpdateStateStatus(input string) (*WindowsUpdateStateStatus, error) {
	vals := map[string]WindowsUpdateStateStatus{
		"failed":              WindowsUpdateStateStatusfailed,
		"pendinginstallation": WindowsUpdateStateStatuspendingInstallation,
		"pendingreboot":       WindowsUpdateStateStatuspendingReboot,
		"uptodate":            WindowsUpdateStateStatusupToDate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdateStateStatus(input)
	return &out, nil
}
