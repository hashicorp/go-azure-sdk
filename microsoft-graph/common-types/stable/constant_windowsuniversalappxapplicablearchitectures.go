package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUniversalAppXApplicableArchitectures string

const (
	WindowsUniversalAppXApplicableArchitectures_Arm     WindowsUniversalAppXApplicableArchitectures = "arm"
	WindowsUniversalAppXApplicableArchitectures_Neutral WindowsUniversalAppXApplicableArchitectures = "neutral"
	WindowsUniversalAppXApplicableArchitectures_None    WindowsUniversalAppXApplicableArchitectures = "none"
	WindowsUniversalAppXApplicableArchitectures_X64     WindowsUniversalAppXApplicableArchitectures = "x64"
	WindowsUniversalAppXApplicableArchitectures_X86     WindowsUniversalAppXApplicableArchitectures = "x86"
)

func PossibleValuesForWindowsUniversalAppXApplicableArchitectures() []string {
	return []string{
		string(WindowsUniversalAppXApplicableArchitectures_Arm),
		string(WindowsUniversalAppXApplicableArchitectures_Neutral),
		string(WindowsUniversalAppXApplicableArchitectures_None),
		string(WindowsUniversalAppXApplicableArchitectures_X64),
		string(WindowsUniversalAppXApplicableArchitectures_X86),
	}
}

func (s *WindowsUniversalAppXApplicableArchitectures) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUniversalAppXApplicableArchitectures(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUniversalAppXApplicableArchitectures(input string) (*WindowsUniversalAppXApplicableArchitectures, error) {
	vals := map[string]WindowsUniversalAppXApplicableArchitectures{
		"arm":     WindowsUniversalAppXApplicableArchitectures_Arm,
		"neutral": WindowsUniversalAppXApplicableArchitectures_Neutral,
		"none":    WindowsUniversalAppXApplicableArchitectures_None,
		"x64":     WindowsUniversalAppXApplicableArchitectures_X64,
		"x86":     WindowsUniversalAppXApplicableArchitectures_X86,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUniversalAppXApplicableArchitectures(input)
	return &out, nil
}
