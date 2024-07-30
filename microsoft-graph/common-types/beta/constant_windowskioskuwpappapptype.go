package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskUWPAppAppType string

const (
	WindowsKioskUWPAppAppType_AumId   WindowsKioskUWPAppAppType = "aumId"
	WindowsKioskUWPAppAppType_Desktop WindowsKioskUWPAppAppType = "desktop"
	WindowsKioskUWPAppAppType_Store   WindowsKioskUWPAppAppType = "store"
	WindowsKioskUWPAppAppType_Unknown WindowsKioskUWPAppAppType = "unknown"
)

func PossibleValuesForWindowsKioskUWPAppAppType() []string {
	return []string{
		string(WindowsKioskUWPAppAppType_AumId),
		string(WindowsKioskUWPAppAppType_Desktop),
		string(WindowsKioskUWPAppAppType_Store),
		string(WindowsKioskUWPAppAppType_Unknown),
	}
}

func (s *WindowsKioskUWPAppAppType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsKioskUWPAppAppType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsKioskUWPAppAppType(input string) (*WindowsKioskUWPAppAppType, error) {
	vals := map[string]WindowsKioskUWPAppAppType{
		"aumid":   WindowsKioskUWPAppAppType_AumId,
		"desktop": WindowsKioskUWPAppAppType_Desktop,
		"store":   WindowsKioskUWPAppAppType_Store,
		"unknown": WindowsKioskUWPAppAppType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsKioskUWPAppAppType(input)
	return &out, nil
}
