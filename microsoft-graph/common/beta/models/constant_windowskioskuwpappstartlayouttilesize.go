package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskUWPAppStartLayoutTileSize string

const (
	WindowsKioskUWPAppStartLayoutTileSizehidden WindowsKioskUWPAppStartLayoutTileSize = "Hidden"
	WindowsKioskUWPAppStartLayoutTileSizelarge  WindowsKioskUWPAppStartLayoutTileSize = "Large"
	WindowsKioskUWPAppStartLayoutTileSizemedium WindowsKioskUWPAppStartLayoutTileSize = "Medium"
	WindowsKioskUWPAppStartLayoutTileSizesmall  WindowsKioskUWPAppStartLayoutTileSize = "Small"
	WindowsKioskUWPAppStartLayoutTileSizewide   WindowsKioskUWPAppStartLayoutTileSize = "Wide"
)

func PossibleValuesForWindowsKioskUWPAppStartLayoutTileSize() []string {
	return []string{
		string(WindowsKioskUWPAppStartLayoutTileSizehidden),
		string(WindowsKioskUWPAppStartLayoutTileSizelarge),
		string(WindowsKioskUWPAppStartLayoutTileSizemedium),
		string(WindowsKioskUWPAppStartLayoutTileSizesmall),
		string(WindowsKioskUWPAppStartLayoutTileSizewide),
	}
}

func parseWindowsKioskUWPAppStartLayoutTileSize(input string) (*WindowsKioskUWPAppStartLayoutTileSize, error) {
	vals := map[string]WindowsKioskUWPAppStartLayoutTileSize{
		"hidden": WindowsKioskUWPAppStartLayoutTileSizehidden,
		"large":  WindowsKioskUWPAppStartLayoutTileSizelarge,
		"medium": WindowsKioskUWPAppStartLayoutTileSizemedium,
		"small":  WindowsKioskUWPAppStartLayoutTileSizesmall,
		"wide":   WindowsKioskUWPAppStartLayoutTileSizewide,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsKioskUWPAppStartLayoutTileSize(input)
	return &out, nil
}
