package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationStartMenuMode string

const (
	Windows10GeneralConfigurationStartMenuModefullScreen    Windows10GeneralConfigurationStartMenuMode = "FullScreen"
	Windows10GeneralConfigurationStartMenuModenonFullScreen Windows10GeneralConfigurationStartMenuMode = "NonFullScreen"
	Windows10GeneralConfigurationStartMenuModeuserDefined   Windows10GeneralConfigurationStartMenuMode = "UserDefined"
)

func PossibleValuesForWindows10GeneralConfigurationStartMenuMode() []string {
	return []string{
		string(Windows10GeneralConfigurationStartMenuModefullScreen),
		string(Windows10GeneralConfigurationStartMenuModenonFullScreen),
		string(Windows10GeneralConfigurationStartMenuModeuserDefined),
	}
}

func parseWindows10GeneralConfigurationStartMenuMode(input string) (*Windows10GeneralConfigurationStartMenuMode, error) {
	vals := map[string]Windows10GeneralConfigurationStartMenuMode{
		"fullscreen":    Windows10GeneralConfigurationStartMenuModefullScreen,
		"nonfullscreen": Windows10GeneralConfigurationStartMenuModenonFullScreen,
		"userdefined":   Windows10GeneralConfigurationStartMenuModeuserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationStartMenuMode(input)
	return &out, nil
}
