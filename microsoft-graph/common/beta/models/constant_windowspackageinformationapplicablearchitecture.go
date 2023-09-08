package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPackageInformationApplicableArchitecture string

const (
	WindowsPackageInformationApplicableArchitecturearm     WindowsPackageInformationApplicableArchitecture = "Arm"
	WindowsPackageInformationApplicableArchitecturearm64   WindowsPackageInformationApplicableArchitecture = "Arm64"
	WindowsPackageInformationApplicableArchitectureneutral WindowsPackageInformationApplicableArchitecture = "Neutral"
	WindowsPackageInformationApplicableArchitecturenone    WindowsPackageInformationApplicableArchitecture = "None"
	WindowsPackageInformationApplicableArchitecturex64     WindowsPackageInformationApplicableArchitecture = "X64"
	WindowsPackageInformationApplicableArchitecturex86     WindowsPackageInformationApplicableArchitecture = "X86"
)

func PossibleValuesForWindowsPackageInformationApplicableArchitecture() []string {
	return []string{
		string(WindowsPackageInformationApplicableArchitecturearm),
		string(WindowsPackageInformationApplicableArchitecturearm64),
		string(WindowsPackageInformationApplicableArchitectureneutral),
		string(WindowsPackageInformationApplicableArchitecturenone),
		string(WindowsPackageInformationApplicableArchitecturex64),
		string(WindowsPackageInformationApplicableArchitecturex86),
	}
}

func parseWindowsPackageInformationApplicableArchitecture(input string) (*WindowsPackageInformationApplicableArchitecture, error) {
	vals := map[string]WindowsPackageInformationApplicableArchitecture{
		"arm":     WindowsPackageInformationApplicableArchitecturearm,
		"arm64":   WindowsPackageInformationApplicableArchitecturearm64,
		"neutral": WindowsPackageInformationApplicableArchitectureneutral,
		"none":    WindowsPackageInformationApplicableArchitecturenone,
		"x64":     WindowsPackageInformationApplicableArchitecturex64,
		"x86":     WindowsPackageInformationApplicableArchitecturex86,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPackageInformationApplicableArchitecture(input)
	return &out, nil
}
