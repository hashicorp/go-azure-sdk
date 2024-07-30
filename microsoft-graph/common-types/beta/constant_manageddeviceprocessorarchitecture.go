package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceProcessorArchitecture string

const (
	ManagedDeviceProcessorArchitecture_ArM64   ManagedDeviceProcessorArchitecture = "arM64"
	ManagedDeviceProcessorArchitecture_Arm     ManagedDeviceProcessorArchitecture = "arm"
	ManagedDeviceProcessorArchitecture_Unknown ManagedDeviceProcessorArchitecture = "unknown"
	ManagedDeviceProcessorArchitecture_X64     ManagedDeviceProcessorArchitecture = "x64"
	ManagedDeviceProcessorArchitecture_X86     ManagedDeviceProcessorArchitecture = "x86"
)

func PossibleValuesForManagedDeviceProcessorArchitecture() []string {
	return []string{
		string(ManagedDeviceProcessorArchitecture_ArM64),
		string(ManagedDeviceProcessorArchitecture_Arm),
		string(ManagedDeviceProcessorArchitecture_Unknown),
		string(ManagedDeviceProcessorArchitecture_X64),
		string(ManagedDeviceProcessorArchitecture_X86),
	}
}

func (s *ManagedDeviceProcessorArchitecture) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceProcessorArchitecture(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceProcessorArchitecture(input string) (*ManagedDeviceProcessorArchitecture, error) {
	vals := map[string]ManagedDeviceProcessorArchitecture{
		"arm64":   ManagedDeviceProcessorArchitecture_ArM64,
		"arm":     ManagedDeviceProcessorArchitecture_Arm,
		"unknown": ManagedDeviceProcessorArchitecture_Unknown,
		"x64":     ManagedDeviceProcessorArchitecture_X64,
		"x86":     ManagedDeviceProcessorArchitecture_X86,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceProcessorArchitecture(input)
	return &out, nil
}
