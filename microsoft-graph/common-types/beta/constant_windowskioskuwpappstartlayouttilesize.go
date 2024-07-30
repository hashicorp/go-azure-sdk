package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskUWPAppStartLayoutTileSize string

const (
	WindowsKioskUWPAppStartLayoutTileSize_Hidden WindowsKioskUWPAppStartLayoutTileSize = "hidden"
	WindowsKioskUWPAppStartLayoutTileSize_Large  WindowsKioskUWPAppStartLayoutTileSize = "large"
	WindowsKioskUWPAppStartLayoutTileSize_Medium WindowsKioskUWPAppStartLayoutTileSize = "medium"
	WindowsKioskUWPAppStartLayoutTileSize_Small  WindowsKioskUWPAppStartLayoutTileSize = "small"
	WindowsKioskUWPAppStartLayoutTileSize_Wide   WindowsKioskUWPAppStartLayoutTileSize = "wide"
)

func PossibleValuesForWindowsKioskUWPAppStartLayoutTileSize() []string {
	return []string{
		string(WindowsKioskUWPAppStartLayoutTileSize_Hidden),
		string(WindowsKioskUWPAppStartLayoutTileSize_Large),
		string(WindowsKioskUWPAppStartLayoutTileSize_Medium),
		string(WindowsKioskUWPAppStartLayoutTileSize_Small),
		string(WindowsKioskUWPAppStartLayoutTileSize_Wide),
	}
}

func (s *WindowsKioskUWPAppStartLayoutTileSize) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsKioskUWPAppStartLayoutTileSize(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsKioskUWPAppStartLayoutTileSize(input string) (*WindowsKioskUWPAppStartLayoutTileSize, error) {
	vals := map[string]WindowsKioskUWPAppStartLayoutTileSize{
		"hidden": WindowsKioskUWPAppStartLayoutTileSize_Hidden,
		"large":  WindowsKioskUWPAppStartLayoutTileSize_Large,
		"medium": WindowsKioskUWPAppStartLayoutTileSize_Medium,
		"small":  WindowsKioskUWPAppStartLayoutTileSize_Small,
		"wide":   WindowsKioskUWPAppStartLayoutTileSize_Wide,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsKioskUWPAppStartLayoutTileSize(input)
	return &out, nil
}
