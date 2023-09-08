package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationStartMenuPinnedFolderDownloads string

const (
	Windows10GeneralConfigurationStartMenuPinnedFolderDownloadshide          Windows10GeneralConfigurationStartMenuPinnedFolderDownloads = "Hide"
	Windows10GeneralConfigurationStartMenuPinnedFolderDownloadsnotConfigured Windows10GeneralConfigurationStartMenuPinnedFolderDownloads = "NotConfigured"
	Windows10GeneralConfigurationStartMenuPinnedFolderDownloadsshow          Windows10GeneralConfigurationStartMenuPinnedFolderDownloads = "Show"
)

func PossibleValuesForWindows10GeneralConfigurationStartMenuPinnedFolderDownloads() []string {
	return []string{
		string(Windows10GeneralConfigurationStartMenuPinnedFolderDownloadshide),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderDownloadsnotConfigured),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderDownloadsshow),
	}
}

func parseWindows10GeneralConfigurationStartMenuPinnedFolderDownloads(input string) (*Windows10GeneralConfigurationStartMenuPinnedFolderDownloads, error) {
	vals := map[string]Windows10GeneralConfigurationStartMenuPinnedFolderDownloads{
		"hide":          Windows10GeneralConfigurationStartMenuPinnedFolderDownloadshide,
		"notconfigured": Windows10GeneralConfigurationStartMenuPinnedFolderDownloadsnotConfigured,
		"show":          Windows10GeneralConfigurationStartMenuPinnedFolderDownloadsshow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationStartMenuPinnedFolderDownloads(input)
	return &out, nil
}
