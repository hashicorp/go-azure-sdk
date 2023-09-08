package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationEdgeOpensWith string

const (
	Windows10GeneralConfigurationEdgeOpensWithnewTabPage    Windows10GeneralConfigurationEdgeOpensWith = "NewTabPage"
	Windows10GeneralConfigurationEdgeOpensWithnotConfigured Windows10GeneralConfigurationEdgeOpensWith = "NotConfigured"
	Windows10GeneralConfigurationEdgeOpensWithpreviousPages Windows10GeneralConfigurationEdgeOpensWith = "PreviousPages"
	Windows10GeneralConfigurationEdgeOpensWithspecificPages Windows10GeneralConfigurationEdgeOpensWith = "SpecificPages"
	Windows10GeneralConfigurationEdgeOpensWithstartPage     Windows10GeneralConfigurationEdgeOpensWith = "StartPage"
)

func PossibleValuesForWindows10GeneralConfigurationEdgeOpensWith() []string {
	return []string{
		string(Windows10GeneralConfigurationEdgeOpensWithnewTabPage),
		string(Windows10GeneralConfigurationEdgeOpensWithnotConfigured),
		string(Windows10GeneralConfigurationEdgeOpensWithpreviousPages),
		string(Windows10GeneralConfigurationEdgeOpensWithspecificPages),
		string(Windows10GeneralConfigurationEdgeOpensWithstartPage),
	}
}

func parseWindows10GeneralConfigurationEdgeOpensWith(input string) (*Windows10GeneralConfigurationEdgeOpensWith, error) {
	vals := map[string]Windows10GeneralConfigurationEdgeOpensWith{
		"newtabpage":    Windows10GeneralConfigurationEdgeOpensWithnewTabPage,
		"notconfigured": Windows10GeneralConfigurationEdgeOpensWithnotConfigured,
		"previouspages": Windows10GeneralConfigurationEdgeOpensWithpreviousPages,
		"specificpages": Windows10GeneralConfigurationEdgeOpensWithspecificPages,
		"startpage":     Windows10GeneralConfigurationEdgeOpensWithstartPage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationEdgeOpensWith(input)
	return &out, nil
}
