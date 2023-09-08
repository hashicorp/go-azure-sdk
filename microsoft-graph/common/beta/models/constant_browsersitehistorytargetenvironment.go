package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSiteHistoryTargetEnvironment string

const (
	BrowserSiteHistoryTargetEnvironmentconfigurable         BrowserSiteHistoryTargetEnvironment = "Configurable"
	BrowserSiteHistoryTargetEnvironmentinternetExplorer11   BrowserSiteHistoryTargetEnvironment = "InternetExplorer11"
	BrowserSiteHistoryTargetEnvironmentinternetExplorerMode BrowserSiteHistoryTargetEnvironment = "InternetExplorerMode"
	BrowserSiteHistoryTargetEnvironmentmicrosoftEdge        BrowserSiteHistoryTargetEnvironment = "MicrosoftEdge"
	BrowserSiteHistoryTargetEnvironmentnone                 BrowserSiteHistoryTargetEnvironment = "None"
)

func PossibleValuesForBrowserSiteHistoryTargetEnvironment() []string {
	return []string{
		string(BrowserSiteHistoryTargetEnvironmentconfigurable),
		string(BrowserSiteHistoryTargetEnvironmentinternetExplorer11),
		string(BrowserSiteHistoryTargetEnvironmentinternetExplorerMode),
		string(BrowserSiteHistoryTargetEnvironmentmicrosoftEdge),
		string(BrowserSiteHistoryTargetEnvironmentnone),
	}
}

func parseBrowserSiteHistoryTargetEnvironment(input string) (*BrowserSiteHistoryTargetEnvironment, error) {
	vals := map[string]BrowserSiteHistoryTargetEnvironment{
		"configurable":         BrowserSiteHistoryTargetEnvironmentconfigurable,
		"internetexplorer11":   BrowserSiteHistoryTargetEnvironmentinternetExplorer11,
		"internetexplorermode": BrowserSiteHistoryTargetEnvironmentinternetExplorerMode,
		"microsoftedge":        BrowserSiteHistoryTargetEnvironmentmicrosoftEdge,
		"none":                 BrowserSiteHistoryTargetEnvironmentnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BrowserSiteHistoryTargetEnvironment(input)
	return &out, nil
}
