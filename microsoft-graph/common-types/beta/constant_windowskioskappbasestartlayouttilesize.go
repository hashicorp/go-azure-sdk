package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskAppBaseStartLayoutTileSize string

const (
	WindowsKioskAppBaseStartLayoutTileSize_Hidden WindowsKioskAppBaseStartLayoutTileSize = "hidden"
	WindowsKioskAppBaseStartLayoutTileSize_Large  WindowsKioskAppBaseStartLayoutTileSize = "large"
	WindowsKioskAppBaseStartLayoutTileSize_Medium WindowsKioskAppBaseStartLayoutTileSize = "medium"
	WindowsKioskAppBaseStartLayoutTileSize_Small  WindowsKioskAppBaseStartLayoutTileSize = "small"
	WindowsKioskAppBaseStartLayoutTileSize_Wide   WindowsKioskAppBaseStartLayoutTileSize = "wide"
)

func PossibleValuesForWindowsKioskAppBaseStartLayoutTileSize() []string {
	return []string{
		string(WindowsKioskAppBaseStartLayoutTileSize_Hidden),
		string(WindowsKioskAppBaseStartLayoutTileSize_Large),
		string(WindowsKioskAppBaseStartLayoutTileSize_Medium),
		string(WindowsKioskAppBaseStartLayoutTileSize_Small),
		string(WindowsKioskAppBaseStartLayoutTileSize_Wide),
	}
}

func (s *WindowsKioskAppBaseStartLayoutTileSize) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsKioskAppBaseStartLayoutTileSize(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsKioskAppBaseStartLayoutTileSize(input string) (*WindowsKioskAppBaseStartLayoutTileSize, error) {
	vals := map[string]WindowsKioskAppBaseStartLayoutTileSize{
		"hidden": WindowsKioskAppBaseStartLayoutTileSize_Hidden,
		"large":  WindowsKioskAppBaseStartLayoutTileSize_Large,
		"medium": WindowsKioskAppBaseStartLayoutTileSize_Medium,
		"small":  WindowsKioskAppBaseStartLayoutTileSize_Small,
		"wide":   WindowsKioskAppBaseStartLayoutTileSize_Wide,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsKioskAppBaseStartLayoutTileSize(input)
	return &out, nil
}
