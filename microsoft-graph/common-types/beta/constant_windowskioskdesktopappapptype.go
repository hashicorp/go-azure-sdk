package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskDesktopAppAppType string

const (
	WindowsKioskDesktopAppAppType_AumId   WindowsKioskDesktopAppAppType = "aumId"
	WindowsKioskDesktopAppAppType_Desktop WindowsKioskDesktopAppAppType = "desktop"
	WindowsKioskDesktopAppAppType_Store   WindowsKioskDesktopAppAppType = "store"
	WindowsKioskDesktopAppAppType_Unknown WindowsKioskDesktopAppAppType = "unknown"
)

func PossibleValuesForWindowsKioskDesktopAppAppType() []string {
	return []string{
		string(WindowsKioskDesktopAppAppType_AumId),
		string(WindowsKioskDesktopAppAppType_Desktop),
		string(WindowsKioskDesktopAppAppType_Store),
		string(WindowsKioskDesktopAppAppType_Unknown),
	}
}

func (s *WindowsKioskDesktopAppAppType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsKioskDesktopAppAppType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsKioskDesktopAppAppType(input string) (*WindowsKioskDesktopAppAppType, error) {
	vals := map[string]WindowsKioskDesktopAppAppType{
		"aumid":   WindowsKioskDesktopAppAppType_AumId,
		"desktop": WindowsKioskDesktopAppAppType_Desktop,
		"store":   WindowsKioskDesktopAppAppType_Store,
		"unknown": WindowsKioskDesktopAppAppType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsKioskDesktopAppAppType(input)
	return &out, nil
}
