package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationStartMenuPinnedFolderHomeGroup string

const (
	Windows10GeneralConfigurationStartMenuPinnedFolderHomeGrouphide          Windows10GeneralConfigurationStartMenuPinnedFolderHomeGroup = "Hide"
	Windows10GeneralConfigurationStartMenuPinnedFolderHomeGroupnotConfigured Windows10GeneralConfigurationStartMenuPinnedFolderHomeGroup = "NotConfigured"
	Windows10GeneralConfigurationStartMenuPinnedFolderHomeGroupshow          Windows10GeneralConfigurationStartMenuPinnedFolderHomeGroup = "Show"
)

func PossibleValuesForWindows10GeneralConfigurationStartMenuPinnedFolderHomeGroup() []string {
	return []string{
		string(Windows10GeneralConfigurationStartMenuPinnedFolderHomeGrouphide),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderHomeGroupnotConfigured),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderHomeGroupshow),
	}
}

func parseWindows10GeneralConfigurationStartMenuPinnedFolderHomeGroup(input string) (*Windows10GeneralConfigurationStartMenuPinnedFolderHomeGroup, error) {
	vals := map[string]Windows10GeneralConfigurationStartMenuPinnedFolderHomeGroup{
		"hide":          Windows10GeneralConfigurationStartMenuPinnedFolderHomeGrouphide,
		"notconfigured": Windows10GeneralConfigurationStartMenuPinnedFolderHomeGroupnotConfigured,
		"show":          Windows10GeneralConfigurationStartMenuPinnedFolderHomeGroupshow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationStartMenuPinnedFolderHomeGroup(input)
	return &out, nil
}
