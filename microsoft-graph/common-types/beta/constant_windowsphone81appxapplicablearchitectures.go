package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81AppXApplicableArchitectures string

const (
	WindowsPhone81AppXApplicableArchitectures_Arm     WindowsPhone81AppXApplicableArchitectures = "arm"
	WindowsPhone81AppXApplicableArchitectures_Arm64   WindowsPhone81AppXApplicableArchitectures = "arm64"
	WindowsPhone81AppXApplicableArchitectures_Neutral WindowsPhone81AppXApplicableArchitectures = "neutral"
	WindowsPhone81AppXApplicableArchitectures_None    WindowsPhone81AppXApplicableArchitectures = "none"
	WindowsPhone81AppXApplicableArchitectures_X64     WindowsPhone81AppXApplicableArchitectures = "x64"
	WindowsPhone81AppXApplicableArchitectures_X86     WindowsPhone81AppXApplicableArchitectures = "x86"
)

func PossibleValuesForWindowsPhone81AppXApplicableArchitectures() []string {
	return []string{
		string(WindowsPhone81AppXApplicableArchitectures_Arm),
		string(WindowsPhone81AppXApplicableArchitectures_Arm64),
		string(WindowsPhone81AppXApplicableArchitectures_Neutral),
		string(WindowsPhone81AppXApplicableArchitectures_None),
		string(WindowsPhone81AppXApplicableArchitectures_X64),
		string(WindowsPhone81AppXApplicableArchitectures_X86),
	}
}

func (s *WindowsPhone81AppXApplicableArchitectures) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhone81AppXApplicableArchitectures(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhone81AppXApplicableArchitectures(input string) (*WindowsPhone81AppXApplicableArchitectures, error) {
	vals := map[string]WindowsPhone81AppXApplicableArchitectures{
		"arm":     WindowsPhone81AppXApplicableArchitectures_Arm,
		"arm64":   WindowsPhone81AppXApplicableArchitectures_Arm64,
		"neutral": WindowsPhone81AppXApplicableArchitectures_Neutral,
		"none":    WindowsPhone81AppXApplicableArchitectures_None,
		"x64":     WindowsPhone81AppXApplicableArchitectures_X64,
		"x86":     WindowsPhone81AppXApplicableArchitectures_X86,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81AppXApplicableArchitectures(input)
	return &out, nil
}
