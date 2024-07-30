package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskWin32AppAppType string

const (
	WindowsKioskWin32AppAppType_AumId   WindowsKioskWin32AppAppType = "aumId"
	WindowsKioskWin32AppAppType_Desktop WindowsKioskWin32AppAppType = "desktop"
	WindowsKioskWin32AppAppType_Store   WindowsKioskWin32AppAppType = "store"
	WindowsKioskWin32AppAppType_Unknown WindowsKioskWin32AppAppType = "unknown"
)

func PossibleValuesForWindowsKioskWin32AppAppType() []string {
	return []string{
		string(WindowsKioskWin32AppAppType_AumId),
		string(WindowsKioskWin32AppAppType_Desktop),
		string(WindowsKioskWin32AppAppType_Store),
		string(WindowsKioskWin32AppAppType_Unknown),
	}
}

func (s *WindowsKioskWin32AppAppType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsKioskWin32AppAppType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsKioskWin32AppAppType(input string) (*WindowsKioskWin32AppAppType, error) {
	vals := map[string]WindowsKioskWin32AppAppType{
		"aumid":   WindowsKioskWin32AppAppType_AumId,
		"desktop": WindowsKioskWin32AppAppType_Desktop,
		"store":   WindowsKioskWin32AppAppType_Store,
		"unknown": WindowsKioskWin32AppAppType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsKioskWin32AppAppType(input)
	return &out, nil
}
