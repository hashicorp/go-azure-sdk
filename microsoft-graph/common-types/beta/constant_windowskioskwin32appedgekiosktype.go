package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskWin32AppEdgeKioskType string

const (
	WindowsKioskWin32AppEdgeKioskType_FullScreen     WindowsKioskWin32AppEdgeKioskType = "fullScreen"
	WindowsKioskWin32AppEdgeKioskType_PublicBrowsing WindowsKioskWin32AppEdgeKioskType = "publicBrowsing"
)

func PossibleValuesForWindowsKioskWin32AppEdgeKioskType() []string {
	return []string{
		string(WindowsKioskWin32AppEdgeKioskType_FullScreen),
		string(WindowsKioskWin32AppEdgeKioskType_PublicBrowsing),
	}
}

func (s *WindowsKioskWin32AppEdgeKioskType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsKioskWin32AppEdgeKioskType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsKioskWin32AppEdgeKioskType(input string) (*WindowsKioskWin32AppEdgeKioskType, error) {
	vals := map[string]WindowsKioskWin32AppEdgeKioskType{
		"fullscreen":     WindowsKioskWin32AppEdgeKioskType_FullScreen,
		"publicbrowsing": WindowsKioskWin32AppEdgeKioskType_PublicBrowsing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsKioskWin32AppEdgeKioskType(input)
	return &out, nil
}
