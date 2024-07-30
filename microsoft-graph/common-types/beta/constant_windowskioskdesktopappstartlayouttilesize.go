package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskDesktopAppStartLayoutTileSize string

const (
	WindowsKioskDesktopAppStartLayoutTileSize_Hidden WindowsKioskDesktopAppStartLayoutTileSize = "hidden"
	WindowsKioskDesktopAppStartLayoutTileSize_Large  WindowsKioskDesktopAppStartLayoutTileSize = "large"
	WindowsKioskDesktopAppStartLayoutTileSize_Medium WindowsKioskDesktopAppStartLayoutTileSize = "medium"
	WindowsKioskDesktopAppStartLayoutTileSize_Small  WindowsKioskDesktopAppStartLayoutTileSize = "small"
	WindowsKioskDesktopAppStartLayoutTileSize_Wide   WindowsKioskDesktopAppStartLayoutTileSize = "wide"
)

func PossibleValuesForWindowsKioskDesktopAppStartLayoutTileSize() []string {
	return []string{
		string(WindowsKioskDesktopAppStartLayoutTileSize_Hidden),
		string(WindowsKioskDesktopAppStartLayoutTileSize_Large),
		string(WindowsKioskDesktopAppStartLayoutTileSize_Medium),
		string(WindowsKioskDesktopAppStartLayoutTileSize_Small),
		string(WindowsKioskDesktopAppStartLayoutTileSize_Wide),
	}
}

func (s *WindowsKioskDesktopAppStartLayoutTileSize) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsKioskDesktopAppStartLayoutTileSize(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsKioskDesktopAppStartLayoutTileSize(input string) (*WindowsKioskDesktopAppStartLayoutTileSize, error) {
	vals := map[string]WindowsKioskDesktopAppStartLayoutTileSize{
		"hidden": WindowsKioskDesktopAppStartLayoutTileSize_Hidden,
		"large":  WindowsKioskDesktopAppStartLayoutTileSize_Large,
		"medium": WindowsKioskDesktopAppStartLayoutTileSize_Medium,
		"small":  WindowsKioskDesktopAppStartLayoutTileSize_Small,
		"wide":   WindowsKioskDesktopAppStartLayoutTileSize_Wide,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsKioskDesktopAppStartLayoutTileSize(input)
	return &out, nil
}
