package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppApplicableArchitectures string

const (
	Win32LobAppApplicableArchitecturesarm     Win32LobAppApplicableArchitectures = "Arm"
	Win32LobAppApplicableArchitecturesneutral Win32LobAppApplicableArchitectures = "Neutral"
	Win32LobAppApplicableArchitecturesnone    Win32LobAppApplicableArchitectures = "None"
	Win32LobAppApplicableArchitecturesx64     Win32LobAppApplicableArchitectures = "X64"
	Win32LobAppApplicableArchitecturesx86     Win32LobAppApplicableArchitectures = "X86"
)

func PossibleValuesForWin32LobAppApplicableArchitectures() []string {
	return []string{
		string(Win32LobAppApplicableArchitecturesarm),
		string(Win32LobAppApplicableArchitecturesneutral),
		string(Win32LobAppApplicableArchitecturesnone),
		string(Win32LobAppApplicableArchitecturesx64),
		string(Win32LobAppApplicableArchitecturesx86),
	}
}

func parseWin32LobAppApplicableArchitectures(input string) (*Win32LobAppApplicableArchitectures, error) {
	vals := map[string]Win32LobAppApplicableArchitectures{
		"arm":     Win32LobAppApplicableArchitecturesarm,
		"neutral": Win32LobAppApplicableArchitecturesneutral,
		"none":    Win32LobAppApplicableArchitecturesnone,
		"x64":     Win32LobAppApplicableArchitecturesx64,
		"x86":     Win32LobAppApplicableArchitecturesx86,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppApplicableArchitectures(input)
	return &out, nil
}
