package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedDeviceProcessorArchitecture string

const (
	WindowsManagedDeviceProcessorArchitecture_ArM64   WindowsManagedDeviceProcessorArchitecture = "arM64"
	WindowsManagedDeviceProcessorArchitecture_Arm     WindowsManagedDeviceProcessorArchitecture = "arm"
	WindowsManagedDeviceProcessorArchitecture_Unknown WindowsManagedDeviceProcessorArchitecture = "unknown"
	WindowsManagedDeviceProcessorArchitecture_X64     WindowsManagedDeviceProcessorArchitecture = "x64"
	WindowsManagedDeviceProcessorArchitecture_X86     WindowsManagedDeviceProcessorArchitecture = "x86"
)

func PossibleValuesForWindowsManagedDeviceProcessorArchitecture() []string {
	return []string{
		string(WindowsManagedDeviceProcessorArchitecture_ArM64),
		string(WindowsManagedDeviceProcessorArchitecture_Arm),
		string(WindowsManagedDeviceProcessorArchitecture_Unknown),
		string(WindowsManagedDeviceProcessorArchitecture_X64),
		string(WindowsManagedDeviceProcessorArchitecture_X86),
	}
}

func (s *WindowsManagedDeviceProcessorArchitecture) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsManagedDeviceProcessorArchitecture(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsManagedDeviceProcessorArchitecture(input string) (*WindowsManagedDeviceProcessorArchitecture, error) {
	vals := map[string]WindowsManagedDeviceProcessorArchitecture{
		"arm64":   WindowsManagedDeviceProcessorArchitecture_ArM64,
		"arm":     WindowsManagedDeviceProcessorArchitecture_Arm,
		"unknown": WindowsManagedDeviceProcessorArchitecture_Unknown,
		"x64":     WindowsManagedDeviceProcessorArchitecture_X64,
		"x86":     WindowsManagedDeviceProcessorArchitecture_X86,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedDeviceProcessorArchitecture(input)
	return &out, nil
}
