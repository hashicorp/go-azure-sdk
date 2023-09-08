package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskAppBaseStartLayoutTileSize string

const (
	WindowsKioskAppBaseStartLayoutTileSizehidden WindowsKioskAppBaseStartLayoutTileSize = "Hidden"
	WindowsKioskAppBaseStartLayoutTileSizelarge  WindowsKioskAppBaseStartLayoutTileSize = "Large"
	WindowsKioskAppBaseStartLayoutTileSizemedium WindowsKioskAppBaseStartLayoutTileSize = "Medium"
	WindowsKioskAppBaseStartLayoutTileSizesmall  WindowsKioskAppBaseStartLayoutTileSize = "Small"
	WindowsKioskAppBaseStartLayoutTileSizewide   WindowsKioskAppBaseStartLayoutTileSize = "Wide"
)

func PossibleValuesForWindowsKioskAppBaseStartLayoutTileSize() []string {
	return []string{
		string(WindowsKioskAppBaseStartLayoutTileSizehidden),
		string(WindowsKioskAppBaseStartLayoutTileSizelarge),
		string(WindowsKioskAppBaseStartLayoutTileSizemedium),
		string(WindowsKioskAppBaseStartLayoutTileSizesmall),
		string(WindowsKioskAppBaseStartLayoutTileSizewide),
	}
}

func parseWindowsKioskAppBaseStartLayoutTileSize(input string) (*WindowsKioskAppBaseStartLayoutTileSize, error) {
	vals := map[string]WindowsKioskAppBaseStartLayoutTileSize{
		"hidden": WindowsKioskAppBaseStartLayoutTileSizehidden,
		"large":  WindowsKioskAppBaseStartLayoutTileSizelarge,
		"medium": WindowsKioskAppBaseStartLayoutTileSizemedium,
		"small":  WindowsKioskAppBaseStartLayoutTileSizesmall,
		"wide":   WindowsKioskAppBaseStartLayoutTileSizewide,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsKioskAppBaseStartLayoutTileSize(input)
	return &out, nil
}
