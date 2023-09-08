package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSiteCompatibilityMode string

const (
	BrowserSiteCompatibilityModedefault                     BrowserSiteCompatibilityMode = "Default"
	BrowserSiteCompatibilityModeinternetExplorer10          BrowserSiteCompatibilityMode = "InternetExplorer10"
	BrowserSiteCompatibilityModeinternetExplorer11          BrowserSiteCompatibilityMode = "InternetExplorer11"
	BrowserSiteCompatibilityModeinternetExplorer5           BrowserSiteCompatibilityMode = "InternetExplorer5"
	BrowserSiteCompatibilityModeinternetExplorer7           BrowserSiteCompatibilityMode = "InternetExplorer7"
	BrowserSiteCompatibilityModeinternetExplorer7Enterprise BrowserSiteCompatibilityMode = "InternetExplorer7Enterprise"
	BrowserSiteCompatibilityModeinternetExplorer8           BrowserSiteCompatibilityMode = "InternetExplorer8"
	BrowserSiteCompatibilityModeinternetExplorer8Enterprise BrowserSiteCompatibilityMode = "InternetExplorer8Enterprise"
	BrowserSiteCompatibilityModeinternetExplorer9           BrowserSiteCompatibilityMode = "InternetExplorer9"
)

func PossibleValuesForBrowserSiteCompatibilityMode() []string {
	return []string{
		string(BrowserSiteCompatibilityModedefault),
		string(BrowserSiteCompatibilityModeinternetExplorer10),
		string(BrowserSiteCompatibilityModeinternetExplorer11),
		string(BrowserSiteCompatibilityModeinternetExplorer5),
		string(BrowserSiteCompatibilityModeinternetExplorer7),
		string(BrowserSiteCompatibilityModeinternetExplorer7Enterprise),
		string(BrowserSiteCompatibilityModeinternetExplorer8),
		string(BrowserSiteCompatibilityModeinternetExplorer8Enterprise),
		string(BrowserSiteCompatibilityModeinternetExplorer9),
	}
}

func parseBrowserSiteCompatibilityMode(input string) (*BrowserSiteCompatibilityMode, error) {
	vals := map[string]BrowserSiteCompatibilityMode{
		"default":                     BrowserSiteCompatibilityModedefault,
		"internetexplorer10":          BrowserSiteCompatibilityModeinternetExplorer10,
		"internetexplorer11":          BrowserSiteCompatibilityModeinternetExplorer11,
		"internetexplorer5":           BrowserSiteCompatibilityModeinternetExplorer5,
		"internetexplorer7":           BrowserSiteCompatibilityModeinternetExplorer7,
		"internetexplorer7enterprise": BrowserSiteCompatibilityModeinternetExplorer7Enterprise,
		"internetexplorer8":           BrowserSiteCompatibilityModeinternetExplorer8,
		"internetexplorer8enterprise": BrowserSiteCompatibilityModeinternetExplorer8Enterprise,
		"internetexplorer9":           BrowserSiteCompatibilityModeinternetExplorer9,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BrowserSiteCompatibilityMode(input)
	return &out, nil
}
