package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationStartMenuPinnedFolderNetwork string

const (
	Windows10GeneralConfigurationStartMenuPinnedFolderNetworkhide          Windows10GeneralConfigurationStartMenuPinnedFolderNetwork = "Hide"
	Windows10GeneralConfigurationStartMenuPinnedFolderNetworknotConfigured Windows10GeneralConfigurationStartMenuPinnedFolderNetwork = "NotConfigured"
	Windows10GeneralConfigurationStartMenuPinnedFolderNetworkshow          Windows10GeneralConfigurationStartMenuPinnedFolderNetwork = "Show"
)

func PossibleValuesForWindows10GeneralConfigurationStartMenuPinnedFolderNetwork() []string {
	return []string{
		string(Windows10GeneralConfigurationStartMenuPinnedFolderNetworkhide),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderNetworknotConfigured),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderNetworkshow),
	}
}

func parseWindows10GeneralConfigurationStartMenuPinnedFolderNetwork(input string) (*Windows10GeneralConfigurationStartMenuPinnedFolderNetwork, error) {
	vals := map[string]Windows10GeneralConfigurationStartMenuPinnedFolderNetwork{
		"hide":          Windows10GeneralConfigurationStartMenuPinnedFolderNetworkhide,
		"notconfigured": Windows10GeneralConfigurationStartMenuPinnedFolderNetworknotConfigured,
		"show":          Windows10GeneralConfigurationStartMenuPinnedFolderNetworkshow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationStartMenuPinnedFolderNetwork(input)
	return &out, nil
}
