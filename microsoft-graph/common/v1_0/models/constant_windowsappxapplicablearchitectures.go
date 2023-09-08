package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAppXApplicableArchitectures string

const (
	WindowsAppXApplicableArchitecturesarm     WindowsAppXApplicableArchitectures = "Arm"
	WindowsAppXApplicableArchitecturesneutral WindowsAppXApplicableArchitectures = "Neutral"
	WindowsAppXApplicableArchitecturesnone    WindowsAppXApplicableArchitectures = "None"
	WindowsAppXApplicableArchitecturesx64     WindowsAppXApplicableArchitectures = "X64"
	WindowsAppXApplicableArchitecturesx86     WindowsAppXApplicableArchitectures = "X86"
)

func PossibleValuesForWindowsAppXApplicableArchitectures() []string {
	return []string{
		string(WindowsAppXApplicableArchitecturesarm),
		string(WindowsAppXApplicableArchitecturesneutral),
		string(WindowsAppXApplicableArchitecturesnone),
		string(WindowsAppXApplicableArchitecturesx64),
		string(WindowsAppXApplicableArchitecturesx86),
	}
}

func parseWindowsAppXApplicableArchitectures(input string) (*WindowsAppXApplicableArchitectures, error) {
	vals := map[string]WindowsAppXApplicableArchitectures{
		"arm":     WindowsAppXApplicableArchitecturesarm,
		"neutral": WindowsAppXApplicableArchitecturesneutral,
		"none":    WindowsAppXApplicableArchitecturesnone,
		"x64":     WindowsAppXApplicableArchitecturesx64,
		"x86":     WindowsAppXApplicableArchitecturesx86,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsAppXApplicableArchitectures(input)
	return &out, nil
}
