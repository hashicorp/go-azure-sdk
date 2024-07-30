package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81AppXBundleApplicableArchitectures string

const (
	WindowsPhone81AppXBundleApplicableArchitectures_Arm     WindowsPhone81AppXBundleApplicableArchitectures = "arm"
	WindowsPhone81AppXBundleApplicableArchitectures_Arm64   WindowsPhone81AppXBundleApplicableArchitectures = "arm64"
	WindowsPhone81AppXBundleApplicableArchitectures_Neutral WindowsPhone81AppXBundleApplicableArchitectures = "neutral"
	WindowsPhone81AppXBundleApplicableArchitectures_None    WindowsPhone81AppXBundleApplicableArchitectures = "none"
	WindowsPhone81AppXBundleApplicableArchitectures_X64     WindowsPhone81AppXBundleApplicableArchitectures = "x64"
	WindowsPhone81AppXBundleApplicableArchitectures_X86     WindowsPhone81AppXBundleApplicableArchitectures = "x86"
)

func PossibleValuesForWindowsPhone81AppXBundleApplicableArchitectures() []string {
	return []string{
		string(WindowsPhone81AppXBundleApplicableArchitectures_Arm),
		string(WindowsPhone81AppXBundleApplicableArchitectures_Arm64),
		string(WindowsPhone81AppXBundleApplicableArchitectures_Neutral),
		string(WindowsPhone81AppXBundleApplicableArchitectures_None),
		string(WindowsPhone81AppXBundleApplicableArchitectures_X64),
		string(WindowsPhone81AppXBundleApplicableArchitectures_X86),
	}
}

func (s *WindowsPhone81AppXBundleApplicableArchitectures) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhone81AppXBundleApplicableArchitectures(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhone81AppXBundleApplicableArchitectures(input string) (*WindowsPhone81AppXBundleApplicableArchitectures, error) {
	vals := map[string]WindowsPhone81AppXBundleApplicableArchitectures{
		"arm":     WindowsPhone81AppXBundleApplicableArchitectures_Arm,
		"arm64":   WindowsPhone81AppXBundleApplicableArchitectures_Arm64,
		"neutral": WindowsPhone81AppXBundleApplicableArchitectures_Neutral,
		"none":    WindowsPhone81AppXBundleApplicableArchitectures_None,
		"x64":     WindowsPhone81AppXBundleApplicableArchitectures_X64,
		"x86":     WindowsPhone81AppXBundleApplicableArchitectures_X86,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81AppXBundleApplicableArchitectures(input)
	return &out, nil
}
