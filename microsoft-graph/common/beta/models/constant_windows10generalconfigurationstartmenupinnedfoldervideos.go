package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationStartMenuPinnedFolderVideos string

const (
	Windows10GeneralConfigurationStartMenuPinnedFolderVideoshide          Windows10GeneralConfigurationStartMenuPinnedFolderVideos = "Hide"
	Windows10GeneralConfigurationStartMenuPinnedFolderVideosnotConfigured Windows10GeneralConfigurationStartMenuPinnedFolderVideos = "NotConfigured"
	Windows10GeneralConfigurationStartMenuPinnedFolderVideosshow          Windows10GeneralConfigurationStartMenuPinnedFolderVideos = "Show"
)

func PossibleValuesForWindows10GeneralConfigurationStartMenuPinnedFolderVideos() []string {
	return []string{
		string(Windows10GeneralConfigurationStartMenuPinnedFolderVideoshide),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderVideosnotConfigured),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderVideosshow),
	}
}

func parseWindows10GeneralConfigurationStartMenuPinnedFolderVideos(input string) (*Windows10GeneralConfigurationStartMenuPinnedFolderVideos, error) {
	vals := map[string]Windows10GeneralConfigurationStartMenuPinnedFolderVideos{
		"hide":          Windows10GeneralConfigurationStartMenuPinnedFolderVideoshide,
		"notconfigured": Windows10GeneralConfigurationStartMenuPinnedFolderVideosnotConfigured,
		"show":          Windows10GeneralConfigurationStartMenuPinnedFolderVideosshow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationStartMenuPinnedFolderVideos(input)
	return &out, nil
}
