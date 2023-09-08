package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81GeneralConfigurationBrowserInternetSecurityLevel string

const (
	Windows81GeneralConfigurationBrowserInternetSecurityLevelhigh        Windows81GeneralConfigurationBrowserInternetSecurityLevel = "High"
	Windows81GeneralConfigurationBrowserInternetSecurityLevelmedium      Windows81GeneralConfigurationBrowserInternetSecurityLevel = "Medium"
	Windows81GeneralConfigurationBrowserInternetSecurityLevelmediumHigh  Windows81GeneralConfigurationBrowserInternetSecurityLevel = "MediumHigh"
	Windows81GeneralConfigurationBrowserInternetSecurityLeveluserDefined Windows81GeneralConfigurationBrowserInternetSecurityLevel = "UserDefined"
)

func PossibleValuesForWindows81GeneralConfigurationBrowserInternetSecurityLevel() []string {
	return []string{
		string(Windows81GeneralConfigurationBrowserInternetSecurityLevelhigh),
		string(Windows81GeneralConfigurationBrowserInternetSecurityLevelmedium),
		string(Windows81GeneralConfigurationBrowserInternetSecurityLevelmediumHigh),
		string(Windows81GeneralConfigurationBrowserInternetSecurityLeveluserDefined),
	}
}

func parseWindows81GeneralConfigurationBrowserInternetSecurityLevel(input string) (*Windows81GeneralConfigurationBrowserInternetSecurityLevel, error) {
	vals := map[string]Windows81GeneralConfigurationBrowserInternetSecurityLevel{
		"high":        Windows81GeneralConfigurationBrowserInternetSecurityLevelhigh,
		"medium":      Windows81GeneralConfigurationBrowserInternetSecurityLevelmedium,
		"mediumhigh":  Windows81GeneralConfigurationBrowserInternetSecurityLevelmediumHigh,
		"userdefined": Windows81GeneralConfigurationBrowserInternetSecurityLeveluserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81GeneralConfigurationBrowserInternetSecurityLevel(input)
	return &out, nil
}
