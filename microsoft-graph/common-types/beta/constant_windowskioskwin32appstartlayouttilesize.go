package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskWin32AppStartLayoutTileSize string

const (
	WindowsKioskWin32AppStartLayoutTileSize_Hidden WindowsKioskWin32AppStartLayoutTileSize = "hidden"
	WindowsKioskWin32AppStartLayoutTileSize_Large  WindowsKioskWin32AppStartLayoutTileSize = "large"
	WindowsKioskWin32AppStartLayoutTileSize_Medium WindowsKioskWin32AppStartLayoutTileSize = "medium"
	WindowsKioskWin32AppStartLayoutTileSize_Small  WindowsKioskWin32AppStartLayoutTileSize = "small"
	WindowsKioskWin32AppStartLayoutTileSize_Wide   WindowsKioskWin32AppStartLayoutTileSize = "wide"
)

func PossibleValuesForWindowsKioskWin32AppStartLayoutTileSize() []string {
	return []string{
		string(WindowsKioskWin32AppStartLayoutTileSize_Hidden),
		string(WindowsKioskWin32AppStartLayoutTileSize_Large),
		string(WindowsKioskWin32AppStartLayoutTileSize_Medium),
		string(WindowsKioskWin32AppStartLayoutTileSize_Small),
		string(WindowsKioskWin32AppStartLayoutTileSize_Wide),
	}
}

func (s *WindowsKioskWin32AppStartLayoutTileSize) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsKioskWin32AppStartLayoutTileSize(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsKioskWin32AppStartLayoutTileSize(input string) (*WindowsKioskWin32AppStartLayoutTileSize, error) {
	vals := map[string]WindowsKioskWin32AppStartLayoutTileSize{
		"hidden": WindowsKioskWin32AppStartLayoutTileSize_Hidden,
		"large":  WindowsKioskWin32AppStartLayoutTileSize_Large,
		"medium": WindowsKioskWin32AppStartLayoutTileSize_Medium,
		"small":  WindowsKioskWin32AppStartLayoutTileSize_Small,
		"wide":   WindowsKioskWin32AppStartLayoutTileSize_Wide,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsKioskWin32AppStartLayoutTileSize(input)
	return &out, nil
}
