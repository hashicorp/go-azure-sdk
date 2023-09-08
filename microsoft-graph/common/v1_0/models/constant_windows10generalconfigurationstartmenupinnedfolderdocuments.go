package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationStartMenuPinnedFolderDocuments string

const (
	Windows10GeneralConfigurationStartMenuPinnedFolderDocumentshide          Windows10GeneralConfigurationStartMenuPinnedFolderDocuments = "Hide"
	Windows10GeneralConfigurationStartMenuPinnedFolderDocumentsnotConfigured Windows10GeneralConfigurationStartMenuPinnedFolderDocuments = "NotConfigured"
	Windows10GeneralConfigurationStartMenuPinnedFolderDocumentsshow          Windows10GeneralConfigurationStartMenuPinnedFolderDocuments = "Show"
)

func PossibleValuesForWindows10GeneralConfigurationStartMenuPinnedFolderDocuments() []string {
	return []string{
		string(Windows10GeneralConfigurationStartMenuPinnedFolderDocumentshide),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderDocumentsnotConfigured),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderDocumentsshow),
	}
}

func parseWindows10GeneralConfigurationStartMenuPinnedFolderDocuments(input string) (*Windows10GeneralConfigurationStartMenuPinnedFolderDocuments, error) {
	vals := map[string]Windows10GeneralConfigurationStartMenuPinnedFolderDocuments{
		"hide":          Windows10GeneralConfigurationStartMenuPinnedFolderDocumentshide,
		"notconfigured": Windows10GeneralConfigurationStartMenuPinnedFolderDocumentsnotConfigured,
		"show":          Windows10GeneralConfigurationStartMenuPinnedFolderDocumentsshow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationStartMenuPinnedFolderDocuments(input)
	return &out, nil
}
