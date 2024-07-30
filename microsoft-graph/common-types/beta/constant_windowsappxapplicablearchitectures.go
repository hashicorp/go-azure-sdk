package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAppXApplicableArchitectures string

const (
	WindowsAppXApplicableArchitectures_Arm     WindowsAppXApplicableArchitectures = "arm"
	WindowsAppXApplicableArchitectures_Arm64   WindowsAppXApplicableArchitectures = "arm64"
	WindowsAppXApplicableArchitectures_Neutral WindowsAppXApplicableArchitectures = "neutral"
	WindowsAppXApplicableArchitectures_None    WindowsAppXApplicableArchitectures = "none"
	WindowsAppXApplicableArchitectures_X64     WindowsAppXApplicableArchitectures = "x64"
	WindowsAppXApplicableArchitectures_X86     WindowsAppXApplicableArchitectures = "x86"
)

func PossibleValuesForWindowsAppXApplicableArchitectures() []string {
	return []string{
		string(WindowsAppXApplicableArchitectures_Arm),
		string(WindowsAppXApplicableArchitectures_Arm64),
		string(WindowsAppXApplicableArchitectures_Neutral),
		string(WindowsAppXApplicableArchitectures_None),
		string(WindowsAppXApplicableArchitectures_X64),
		string(WindowsAppXApplicableArchitectures_X86),
	}
}

func (s *WindowsAppXApplicableArchitectures) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsAppXApplicableArchitectures(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsAppXApplicableArchitectures(input string) (*WindowsAppXApplicableArchitectures, error) {
	vals := map[string]WindowsAppXApplicableArchitectures{
		"arm":     WindowsAppXApplicableArchitectures_Arm,
		"arm64":   WindowsAppXApplicableArchitectures_Arm64,
		"neutral": WindowsAppXApplicableArchitectures_Neutral,
		"none":    WindowsAppXApplicableArchitectures_None,
		"x64":     WindowsAppXApplicableArchitectures_X64,
		"x86":     WindowsAppXApplicableArchitectures_X86,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsAppXApplicableArchitectures(input)
	return &out, nil
}
