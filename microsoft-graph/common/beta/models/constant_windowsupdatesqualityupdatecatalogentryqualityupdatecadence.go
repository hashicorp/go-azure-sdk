package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateCadence string

const (
	WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateCadencemonthly   WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateCadence = "Monthly"
	WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateCadenceoutOfBand WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateCadence = "OutOfBand"
)

func PossibleValuesForWindowsUpdatesQualityUpdateCatalogEntryQualityUpdateCadence() []string {
	return []string{
		string(WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateCadencemonthly),
		string(WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateCadenceoutOfBand),
	}
}

func parseWindowsUpdatesQualityUpdateCatalogEntryQualityUpdateCadence(input string) (*WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateCadence, error) {
	vals := map[string]WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateCadence{
		"monthly":   WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateCadencemonthly,
		"outofband": WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateCadenceoutOfBand,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateCadence(input)
	return &out, nil
}
