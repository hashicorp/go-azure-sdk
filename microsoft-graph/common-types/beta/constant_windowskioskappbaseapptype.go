package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskAppBaseAppType string

const (
	WindowsKioskAppBaseAppType_AumId   WindowsKioskAppBaseAppType = "aumId"
	WindowsKioskAppBaseAppType_Desktop WindowsKioskAppBaseAppType = "desktop"
	WindowsKioskAppBaseAppType_Store   WindowsKioskAppBaseAppType = "store"
	WindowsKioskAppBaseAppType_Unknown WindowsKioskAppBaseAppType = "unknown"
)

func PossibleValuesForWindowsKioskAppBaseAppType() []string {
	return []string{
		string(WindowsKioskAppBaseAppType_AumId),
		string(WindowsKioskAppBaseAppType_Desktop),
		string(WindowsKioskAppBaseAppType_Store),
		string(WindowsKioskAppBaseAppType_Unknown),
	}
}

func (s *WindowsKioskAppBaseAppType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsKioskAppBaseAppType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsKioskAppBaseAppType(input string) (*WindowsKioskAppBaseAppType, error) {
	vals := map[string]WindowsKioskAppBaseAppType{
		"aumid":   WindowsKioskAppBaseAppType_AumId,
		"desktop": WindowsKioskAppBaseAppType_Desktop,
		"store":   WindowsKioskAppBaseAppType_Store,
		"unknown": WindowsKioskAppBaseAppType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsKioskAppBaseAppType(input)
	return &out, nil
}
