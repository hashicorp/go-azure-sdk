package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81AppXApplicableArchitectures string

const (
	WindowsPhone81AppXApplicableArchitecturesarm     WindowsPhone81AppXApplicableArchitectures = "Arm"
	WindowsPhone81AppXApplicableArchitecturesarm64   WindowsPhone81AppXApplicableArchitectures = "Arm64"
	WindowsPhone81AppXApplicableArchitecturesneutral WindowsPhone81AppXApplicableArchitectures = "Neutral"
	WindowsPhone81AppXApplicableArchitecturesnone    WindowsPhone81AppXApplicableArchitectures = "None"
	WindowsPhone81AppXApplicableArchitecturesx64     WindowsPhone81AppXApplicableArchitectures = "X64"
	WindowsPhone81AppXApplicableArchitecturesx86     WindowsPhone81AppXApplicableArchitectures = "X86"
)

func PossibleValuesForWindowsPhone81AppXApplicableArchitectures() []string {
	return []string{
		string(WindowsPhone81AppXApplicableArchitecturesarm),
		string(WindowsPhone81AppXApplicableArchitecturesarm64),
		string(WindowsPhone81AppXApplicableArchitecturesneutral),
		string(WindowsPhone81AppXApplicableArchitecturesnone),
		string(WindowsPhone81AppXApplicableArchitecturesx64),
		string(WindowsPhone81AppXApplicableArchitecturesx86),
	}
}

func parseWindowsPhone81AppXApplicableArchitectures(input string) (*WindowsPhone81AppXApplicableArchitectures, error) {
	vals := map[string]WindowsPhone81AppXApplicableArchitectures{
		"arm":     WindowsPhone81AppXApplicableArchitecturesarm,
		"arm64":   WindowsPhone81AppXApplicableArchitecturesarm64,
		"neutral": WindowsPhone81AppXApplicableArchitecturesneutral,
		"none":    WindowsPhone81AppXApplicableArchitecturesnone,
		"x64":     WindowsPhone81AppXApplicableArchitecturesx64,
		"x86":     WindowsPhone81AppXApplicableArchitecturesx86,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81AppXApplicableArchitectures(input)
	return &out, nil
}
