package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OfficeSuiteAppOfficePlatformArchitecture string

const (
	OfficeSuiteAppOfficePlatformArchitecturearm     OfficeSuiteAppOfficePlatformArchitecture = "Arm"
	OfficeSuiteAppOfficePlatformArchitecturearm64   OfficeSuiteAppOfficePlatformArchitecture = "Arm64"
	OfficeSuiteAppOfficePlatformArchitectureneutral OfficeSuiteAppOfficePlatformArchitecture = "Neutral"
	OfficeSuiteAppOfficePlatformArchitecturenone    OfficeSuiteAppOfficePlatformArchitecture = "None"
	OfficeSuiteAppOfficePlatformArchitecturex64     OfficeSuiteAppOfficePlatformArchitecture = "X64"
	OfficeSuiteAppOfficePlatformArchitecturex86     OfficeSuiteAppOfficePlatformArchitecture = "X86"
)

func PossibleValuesForOfficeSuiteAppOfficePlatformArchitecture() []string {
	return []string{
		string(OfficeSuiteAppOfficePlatformArchitecturearm),
		string(OfficeSuiteAppOfficePlatformArchitecturearm64),
		string(OfficeSuiteAppOfficePlatformArchitectureneutral),
		string(OfficeSuiteAppOfficePlatformArchitecturenone),
		string(OfficeSuiteAppOfficePlatformArchitecturex64),
		string(OfficeSuiteAppOfficePlatformArchitecturex86),
	}
}

func parseOfficeSuiteAppOfficePlatformArchitecture(input string) (*OfficeSuiteAppOfficePlatformArchitecture, error) {
	vals := map[string]OfficeSuiteAppOfficePlatformArchitecture{
		"arm":     OfficeSuiteAppOfficePlatformArchitecturearm,
		"arm64":   OfficeSuiteAppOfficePlatformArchitecturearm64,
		"neutral": OfficeSuiteAppOfficePlatformArchitectureneutral,
		"none":    OfficeSuiteAppOfficePlatformArchitecturenone,
		"x64":     OfficeSuiteAppOfficePlatformArchitecturex64,
		"x86":     OfficeSuiteAppOfficePlatformArchitecturex86,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OfficeSuiteAppOfficePlatformArchitecture(input)
	return &out, nil
}
