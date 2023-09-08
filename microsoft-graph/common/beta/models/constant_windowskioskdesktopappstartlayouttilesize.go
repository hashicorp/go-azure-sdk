package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskDesktopAppStartLayoutTileSize string

const (
	WindowsKioskDesktopAppStartLayoutTileSizehidden WindowsKioskDesktopAppStartLayoutTileSize = "Hidden"
	WindowsKioskDesktopAppStartLayoutTileSizelarge  WindowsKioskDesktopAppStartLayoutTileSize = "Large"
	WindowsKioskDesktopAppStartLayoutTileSizemedium WindowsKioskDesktopAppStartLayoutTileSize = "Medium"
	WindowsKioskDesktopAppStartLayoutTileSizesmall  WindowsKioskDesktopAppStartLayoutTileSize = "Small"
	WindowsKioskDesktopAppStartLayoutTileSizewide   WindowsKioskDesktopAppStartLayoutTileSize = "Wide"
)

func PossibleValuesForWindowsKioskDesktopAppStartLayoutTileSize() []string {
	return []string{
		string(WindowsKioskDesktopAppStartLayoutTileSizehidden),
		string(WindowsKioskDesktopAppStartLayoutTileSizelarge),
		string(WindowsKioskDesktopAppStartLayoutTileSizemedium),
		string(WindowsKioskDesktopAppStartLayoutTileSizesmall),
		string(WindowsKioskDesktopAppStartLayoutTileSizewide),
	}
}

func parseWindowsKioskDesktopAppStartLayoutTileSize(input string) (*WindowsKioskDesktopAppStartLayoutTileSize, error) {
	vals := map[string]WindowsKioskDesktopAppStartLayoutTileSize{
		"hidden": WindowsKioskDesktopAppStartLayoutTileSizehidden,
		"large":  WindowsKioskDesktopAppStartLayoutTileSizelarge,
		"medium": WindowsKioskDesktopAppStartLayoutTileSizemedium,
		"small":  WindowsKioskDesktopAppStartLayoutTileSizesmall,
		"wide":   WindowsKioskDesktopAppStartLayoutTileSizewide,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsKioskDesktopAppStartLayoutTileSize(input)
	return &out, nil
}
