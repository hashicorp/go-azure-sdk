package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationEdgeCookiePolicy string

const (
	Windows10GeneralConfigurationEdgeCookiePolicyallow           Windows10GeneralConfigurationEdgeCookiePolicy = "Allow"
	Windows10GeneralConfigurationEdgeCookiePolicyblockAll        Windows10GeneralConfigurationEdgeCookiePolicy = "BlockAll"
	Windows10GeneralConfigurationEdgeCookiePolicyblockThirdParty Windows10GeneralConfigurationEdgeCookiePolicy = "BlockThirdParty"
	Windows10GeneralConfigurationEdgeCookiePolicyuserDefined     Windows10GeneralConfigurationEdgeCookiePolicy = "UserDefined"
)

func PossibleValuesForWindows10GeneralConfigurationEdgeCookiePolicy() []string {
	return []string{
		string(Windows10GeneralConfigurationEdgeCookiePolicyallow),
		string(Windows10GeneralConfigurationEdgeCookiePolicyblockAll),
		string(Windows10GeneralConfigurationEdgeCookiePolicyblockThirdParty),
		string(Windows10GeneralConfigurationEdgeCookiePolicyuserDefined),
	}
}

func parseWindows10GeneralConfigurationEdgeCookiePolicy(input string) (*Windows10GeneralConfigurationEdgeCookiePolicy, error) {
	vals := map[string]Windows10GeneralConfigurationEdgeCookiePolicy{
		"allow":           Windows10GeneralConfigurationEdgeCookiePolicyallow,
		"blockall":        Windows10GeneralConfigurationEdgeCookiePolicyblockAll,
		"blockthirdparty": Windows10GeneralConfigurationEdgeCookiePolicyblockThirdParty,
		"userdefined":     Windows10GeneralConfigurationEdgeCookiePolicyuserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationEdgeCookiePolicy(input)
	return &out, nil
}
