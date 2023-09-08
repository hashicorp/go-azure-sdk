package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser string

const (
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUseradministrators                    Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser = "Administrators"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUseradministratorsAndInteractiveUsers Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser = "AdministratorsAndInteractiveUsers"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUseradministratorsAndPowerUsers       Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser = "AdministratorsAndPowerUsers"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUsernotConfigured                     Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser = "NotConfigured"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUseradministrators),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUseradministratorsAndInteractiveUsers),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUseradministratorsAndPowerUsers),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUsernotConfigured),
	}
}

func parseWindows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser(input string) (*Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser{
		"administrators":                    Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUseradministrators,
		"administratorsandinteractiveusers": Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUseradministratorsAndInteractiveUsers,
		"administratorsandpowerusers":       Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUseradministratorsAndPowerUsers,
		"notconfigured":                     Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUsernotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser(input)
	return &out, nil
}
