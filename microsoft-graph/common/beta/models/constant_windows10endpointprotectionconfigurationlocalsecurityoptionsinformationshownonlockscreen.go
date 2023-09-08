package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen string

const (
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreendoNotDisplayUser          Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen = "DoNotDisplayUser"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreennotConfigured             Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen = "NotConfigured"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreenuserDisplayNameDomainUser Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen = "UserDisplayNameDomainUser"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreenuserDisplayNameOnly       Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen = "UserDisplayNameOnly"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreendoNotDisplayUser),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreennotConfigured),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreenuserDisplayNameDomainUser),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreenuserDisplayNameOnly),
	}
}

func parseWindows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen(input string) (*Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen{
		"donotdisplayuser":          Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreendoNotDisplayUser,
		"notconfigured":             Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreennotConfigured,
		"userdisplaynamedomainuser": Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreenuserDisplayNameDomainUser,
		"userdisplaynameonly":       Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreenuserDisplayNameOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen(input)
	return &out, nil
}
