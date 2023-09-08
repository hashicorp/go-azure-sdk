package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel string

const (
	Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevelhigh        Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel = "High"
	Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevellow         Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel = "Low"
	Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevelmedium      Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel = "Medium"
	Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevelmediumHigh  Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel = "MediumHigh"
	Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevelmediumLow   Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel = "MediumLow"
	Windows81GeneralConfigurationBrowserTrustedSitesSecurityLeveluserDefined Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel = "UserDefined"
)

func PossibleValuesForWindows81GeneralConfigurationBrowserTrustedSitesSecurityLevel() []string {
	return []string{
		string(Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevelhigh),
		string(Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevellow),
		string(Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevelmedium),
		string(Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevelmediumHigh),
		string(Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevelmediumLow),
		string(Windows81GeneralConfigurationBrowserTrustedSitesSecurityLeveluserDefined),
	}
}

func parseWindows81GeneralConfigurationBrowserTrustedSitesSecurityLevel(input string) (*Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel, error) {
	vals := map[string]Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel{
		"high":        Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevelhigh,
		"low":         Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevellow,
		"medium":      Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevelmedium,
		"mediumhigh":  Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevelmediumHigh,
		"mediumlow":   Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevelmediumLow,
		"userdefined": Windows81GeneralConfigurationBrowserTrustedSitesSecurityLeveluserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel(input)
	return &out, nil
}
