package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen string

const (
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreenadministrators                    Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen = "Administrators"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreenadministratorsAndInteractiveUsers Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen = "AdministratorsAndInteractiveUsers"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreenadministratorsAndPowerUsers       Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen = "AdministratorsAndPowerUsers"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreennotConfigured                     Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen = "NotConfigured"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreenadministrators),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreenadministratorsAndInteractiveUsers),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreenadministratorsAndPowerUsers),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreennotConfigured),
	}
}

func parseWindows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen(input string) (*Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen{
		"administrators":                    Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreenadministrators,
		"administratorsandinteractiveusers": Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreenadministratorsAndInteractiveUsers,
		"administratorsandpowerusers":       Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreenadministratorsAndPowerUsers,
		"notconfigured":                     Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreennotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen(input)
	return &out, nil
}
