package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassification string

const (
	WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassificationall         WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassification = "All"
	WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassificationnonSecurity WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassification = "NonSecurity"
	WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassificationsecurity    WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassification = "Security"
)

func PossibleValuesForWindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassification() []string {
	return []string{
		string(WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassificationall),
		string(WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassificationnonSecurity),
		string(WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassificationsecurity),
	}
}

func parseWindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassification(input string) (*WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassification, error) {
	vals := map[string]WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassification{
		"all":         WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassificationall,
		"nonsecurity": WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassificationnonSecurity,
		"security":    WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassificationsecurity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassification(input)
	return &out, nil
}
