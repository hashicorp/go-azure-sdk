package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationStartMenuPinnedFolderMusic string

const (
	Windows10GeneralConfigurationStartMenuPinnedFolderMusichide          Windows10GeneralConfigurationStartMenuPinnedFolderMusic = "Hide"
	Windows10GeneralConfigurationStartMenuPinnedFolderMusicnotConfigured Windows10GeneralConfigurationStartMenuPinnedFolderMusic = "NotConfigured"
	Windows10GeneralConfigurationStartMenuPinnedFolderMusicshow          Windows10GeneralConfigurationStartMenuPinnedFolderMusic = "Show"
)

func PossibleValuesForWindows10GeneralConfigurationStartMenuPinnedFolderMusic() []string {
	return []string{
		string(Windows10GeneralConfigurationStartMenuPinnedFolderMusichide),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderMusicnotConfigured),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderMusicshow),
	}
}

func parseWindows10GeneralConfigurationStartMenuPinnedFolderMusic(input string) (*Windows10GeneralConfigurationStartMenuPinnedFolderMusic, error) {
	vals := map[string]Windows10GeneralConfigurationStartMenuPinnedFolderMusic{
		"hide":          Windows10GeneralConfigurationStartMenuPinnedFolderMusichide,
		"notconfigured": Windows10GeneralConfigurationStartMenuPinnedFolderMusicnotConfigured,
		"show":          Windows10GeneralConfigurationStartMenuPinnedFolderMusicshow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationStartMenuPinnedFolderMusic(input)
	return &out, nil
}
