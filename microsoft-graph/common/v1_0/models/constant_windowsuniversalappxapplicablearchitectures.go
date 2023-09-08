package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUniversalAppXApplicableArchitectures string

const (
	WindowsUniversalAppXApplicableArchitecturesarm     WindowsUniversalAppXApplicableArchitectures = "Arm"
	WindowsUniversalAppXApplicableArchitecturesneutral WindowsUniversalAppXApplicableArchitectures = "Neutral"
	WindowsUniversalAppXApplicableArchitecturesnone    WindowsUniversalAppXApplicableArchitectures = "None"
	WindowsUniversalAppXApplicableArchitecturesx64     WindowsUniversalAppXApplicableArchitectures = "X64"
	WindowsUniversalAppXApplicableArchitecturesx86     WindowsUniversalAppXApplicableArchitectures = "X86"
)

func PossibleValuesForWindowsUniversalAppXApplicableArchitectures() []string {
	return []string{
		string(WindowsUniversalAppXApplicableArchitecturesarm),
		string(WindowsUniversalAppXApplicableArchitecturesneutral),
		string(WindowsUniversalAppXApplicableArchitecturesnone),
		string(WindowsUniversalAppXApplicableArchitecturesx64),
		string(WindowsUniversalAppXApplicableArchitecturesx86),
	}
}

func parseWindowsUniversalAppXApplicableArchitectures(input string) (*WindowsUniversalAppXApplicableArchitectures, error) {
	vals := map[string]WindowsUniversalAppXApplicableArchitectures{
		"arm":     WindowsUniversalAppXApplicableArchitecturesarm,
		"neutral": WindowsUniversalAppXApplicableArchitecturesneutral,
		"none":    WindowsUniversalAppXApplicableArchitecturesnone,
		"x64":     WindowsUniversalAppXApplicableArchitecturesx64,
		"x86":     WindowsUniversalAppXApplicableArchitecturesx86,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUniversalAppXApplicableArchitectures(input)
	return &out, nil
}
