package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSiteHistoryCompatibilityMode string

const (
	BrowserSiteHistoryCompatibilityModedefault                     BrowserSiteHistoryCompatibilityMode = "Default"
	BrowserSiteHistoryCompatibilityModeinternetExplorer10          BrowserSiteHistoryCompatibilityMode = "InternetExplorer10"
	BrowserSiteHistoryCompatibilityModeinternetExplorer11          BrowserSiteHistoryCompatibilityMode = "InternetExplorer11"
	BrowserSiteHistoryCompatibilityModeinternetExplorer5           BrowserSiteHistoryCompatibilityMode = "InternetExplorer5"
	BrowserSiteHistoryCompatibilityModeinternetExplorer7           BrowserSiteHistoryCompatibilityMode = "InternetExplorer7"
	BrowserSiteHistoryCompatibilityModeinternetExplorer7Enterprise BrowserSiteHistoryCompatibilityMode = "InternetExplorer7Enterprise"
	BrowserSiteHistoryCompatibilityModeinternetExplorer8           BrowserSiteHistoryCompatibilityMode = "InternetExplorer8"
	BrowserSiteHistoryCompatibilityModeinternetExplorer8Enterprise BrowserSiteHistoryCompatibilityMode = "InternetExplorer8Enterprise"
	BrowserSiteHistoryCompatibilityModeinternetExplorer9           BrowserSiteHistoryCompatibilityMode = "InternetExplorer9"
)

func PossibleValuesForBrowserSiteHistoryCompatibilityMode() []string {
	return []string{
		string(BrowserSiteHistoryCompatibilityModedefault),
		string(BrowserSiteHistoryCompatibilityModeinternetExplorer10),
		string(BrowserSiteHistoryCompatibilityModeinternetExplorer11),
		string(BrowserSiteHistoryCompatibilityModeinternetExplorer5),
		string(BrowserSiteHistoryCompatibilityModeinternetExplorer7),
		string(BrowserSiteHistoryCompatibilityModeinternetExplorer7Enterprise),
		string(BrowserSiteHistoryCompatibilityModeinternetExplorer8),
		string(BrowserSiteHistoryCompatibilityModeinternetExplorer8Enterprise),
		string(BrowserSiteHistoryCompatibilityModeinternetExplorer9),
	}
}

func parseBrowserSiteHistoryCompatibilityMode(input string) (*BrowserSiteHistoryCompatibilityMode, error) {
	vals := map[string]BrowserSiteHistoryCompatibilityMode{
		"default":                     BrowserSiteHistoryCompatibilityModedefault,
		"internetexplorer10":          BrowserSiteHistoryCompatibilityModeinternetExplorer10,
		"internetexplorer11":          BrowserSiteHistoryCompatibilityModeinternetExplorer11,
		"internetexplorer5":           BrowserSiteHistoryCompatibilityModeinternetExplorer5,
		"internetexplorer7":           BrowserSiteHistoryCompatibilityModeinternetExplorer7,
		"internetexplorer7enterprise": BrowserSiteHistoryCompatibilityModeinternetExplorer7Enterprise,
		"internetexplorer8":           BrowserSiteHistoryCompatibilityModeinternetExplorer8,
		"internetexplorer8enterprise": BrowserSiteHistoryCompatibilityModeinternetExplorer8Enterprise,
		"internetexplorer9":           BrowserSiteHistoryCompatibilityModeinternetExplorer9,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BrowserSiteHistoryCompatibilityMode(input)
	return &out, nil
}
