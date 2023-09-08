package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSiteTargetEnvironment string

const (
	BrowserSiteTargetEnvironmentconfigurable         BrowserSiteTargetEnvironment = "Configurable"
	BrowserSiteTargetEnvironmentinternetExplorer11   BrowserSiteTargetEnvironment = "InternetExplorer11"
	BrowserSiteTargetEnvironmentinternetExplorerMode BrowserSiteTargetEnvironment = "InternetExplorerMode"
	BrowserSiteTargetEnvironmentmicrosoftEdge        BrowserSiteTargetEnvironment = "MicrosoftEdge"
	BrowserSiteTargetEnvironmentnone                 BrowserSiteTargetEnvironment = "None"
)

func PossibleValuesForBrowserSiteTargetEnvironment() []string {
	return []string{
		string(BrowserSiteTargetEnvironmentconfigurable),
		string(BrowserSiteTargetEnvironmentinternetExplorer11),
		string(BrowserSiteTargetEnvironmentinternetExplorerMode),
		string(BrowserSiteTargetEnvironmentmicrosoftEdge),
		string(BrowserSiteTargetEnvironmentnone),
	}
}

func parseBrowserSiteTargetEnvironment(input string) (*BrowserSiteTargetEnvironment, error) {
	vals := map[string]BrowserSiteTargetEnvironment{
		"configurable":         BrowserSiteTargetEnvironmentconfigurable,
		"internetexplorer11":   BrowserSiteTargetEnvironmentinternetExplorer11,
		"internetexplorermode": BrowserSiteTargetEnvironmentinternetExplorerMode,
		"microsoftedge":        BrowserSiteTargetEnvironmentmicrosoftEdge,
		"none":                 BrowserSiteTargetEnvironmentnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BrowserSiteTargetEnvironment(input)
	return &out, nil
}
