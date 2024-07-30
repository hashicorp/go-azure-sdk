package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUniversalAppXApplicableDeviceTypes string

const (
	WindowsUniversalAppXApplicableDeviceTypes_Desktop     WindowsUniversalAppXApplicableDeviceTypes = "desktop"
	WindowsUniversalAppXApplicableDeviceTypes_Holographic WindowsUniversalAppXApplicableDeviceTypes = "holographic"
	WindowsUniversalAppXApplicableDeviceTypes_Mobile      WindowsUniversalAppXApplicableDeviceTypes = "mobile"
	WindowsUniversalAppXApplicableDeviceTypes_None        WindowsUniversalAppXApplicableDeviceTypes = "none"
	WindowsUniversalAppXApplicableDeviceTypes_Team        WindowsUniversalAppXApplicableDeviceTypes = "team"
)

func PossibleValuesForWindowsUniversalAppXApplicableDeviceTypes() []string {
	return []string{
		string(WindowsUniversalAppXApplicableDeviceTypes_Desktop),
		string(WindowsUniversalAppXApplicableDeviceTypes_Holographic),
		string(WindowsUniversalAppXApplicableDeviceTypes_Mobile),
		string(WindowsUniversalAppXApplicableDeviceTypes_None),
		string(WindowsUniversalAppXApplicableDeviceTypes_Team),
	}
}

func (s *WindowsUniversalAppXApplicableDeviceTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUniversalAppXApplicableDeviceTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUniversalAppXApplicableDeviceTypes(input string) (*WindowsUniversalAppXApplicableDeviceTypes, error) {
	vals := map[string]WindowsUniversalAppXApplicableDeviceTypes{
		"desktop":     WindowsUniversalAppXApplicableDeviceTypes_Desktop,
		"holographic": WindowsUniversalAppXApplicableDeviceTypes_Holographic,
		"mobile":      WindowsUniversalAppXApplicableDeviceTypes_Mobile,
		"none":        WindowsUniversalAppXApplicableDeviceTypes_None,
		"team":        WindowsUniversalAppXApplicableDeviceTypes_Team,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUniversalAppXApplicableDeviceTypes(input)
	return &out, nil
}
