package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncState string

const (
	WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncStatefailure WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncState = "Failure"
	WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncStatepending WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncState = "Pending"
	WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncStatesuccess WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncState = "Success"
)

func PossibleValuesForWindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncState() []string {
	return []string{
		string(WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncStatefailure),
		string(WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncStatepending),
		string(WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncStatesuccess),
	}
}

func parseWindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncState(input string) (*WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncState, error) {
	vals := map[string]WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncState{
		"failure": WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncStatefailure,
		"pending": WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncStatepending,
		"success": WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncStatesuccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncState(input)
	return &out, nil
}
