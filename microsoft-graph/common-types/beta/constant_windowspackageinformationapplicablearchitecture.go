package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPackageInformationApplicableArchitecture string

const (
	WindowsPackageInformationApplicableArchitecture_Arm     WindowsPackageInformationApplicableArchitecture = "arm"
	WindowsPackageInformationApplicableArchitecture_Arm64   WindowsPackageInformationApplicableArchitecture = "arm64"
	WindowsPackageInformationApplicableArchitecture_Neutral WindowsPackageInformationApplicableArchitecture = "neutral"
	WindowsPackageInformationApplicableArchitecture_None    WindowsPackageInformationApplicableArchitecture = "none"
	WindowsPackageInformationApplicableArchitecture_X64     WindowsPackageInformationApplicableArchitecture = "x64"
	WindowsPackageInformationApplicableArchitecture_X86     WindowsPackageInformationApplicableArchitecture = "x86"
)

func PossibleValuesForWindowsPackageInformationApplicableArchitecture() []string {
	return []string{
		string(WindowsPackageInformationApplicableArchitecture_Arm),
		string(WindowsPackageInformationApplicableArchitecture_Arm64),
		string(WindowsPackageInformationApplicableArchitecture_Neutral),
		string(WindowsPackageInformationApplicableArchitecture_None),
		string(WindowsPackageInformationApplicableArchitecture_X64),
		string(WindowsPackageInformationApplicableArchitecture_X86),
	}
}

func (s *WindowsPackageInformationApplicableArchitecture) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPackageInformationApplicableArchitecture(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPackageInformationApplicableArchitecture(input string) (*WindowsPackageInformationApplicableArchitecture, error) {
	vals := map[string]WindowsPackageInformationApplicableArchitecture{
		"arm":     WindowsPackageInformationApplicableArchitecture_Arm,
		"arm64":   WindowsPackageInformationApplicableArchitecture_Arm64,
		"neutral": WindowsPackageInformationApplicableArchitecture_Neutral,
		"none":    WindowsPackageInformationApplicableArchitecture_None,
		"x64":     WindowsPackageInformationApplicableArchitecture_X64,
		"x86":     WindowsPackageInformationApplicableArchitecture_X86,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPackageInformationApplicableArchitecture(input)
	return &out, nil
}
