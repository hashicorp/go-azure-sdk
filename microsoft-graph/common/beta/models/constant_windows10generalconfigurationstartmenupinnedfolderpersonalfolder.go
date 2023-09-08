package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFolder string

const (
	Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFolderhide          Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFolder = "Hide"
	Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFoldernotConfigured Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFolder = "NotConfigured"
	Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFoldershow          Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFolder = "Show"
)

func PossibleValuesForWindows10GeneralConfigurationStartMenuPinnedFolderPersonalFolder() []string {
	return []string{
		string(Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFolderhide),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFoldernotConfigured),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFoldershow),
	}
}

func parseWindows10GeneralConfigurationStartMenuPinnedFolderPersonalFolder(input string) (*Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFolder, error) {
	vals := map[string]Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFolder{
		"hide":          Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFolderhide,
		"notconfigured": Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFoldernotConfigured,
		"show":          Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFoldershow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFolder(input)
	return &out, nil
}
