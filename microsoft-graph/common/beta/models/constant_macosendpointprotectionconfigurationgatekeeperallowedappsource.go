package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource string

const (
	MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSourceanywhere                           MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource = "Anywhere"
	MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSourcemacAppStore                        MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource = "MacAppStore"
	MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSourcemacAppStoreAndIdentifiedDevelopers MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource = "MacAppStoreAndIdentifiedDevelopers"
	MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSourcenotConfigured                      MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource = "NotConfigured"
)

func PossibleValuesForMacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource() []string {
	return []string{
		string(MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSourceanywhere),
		string(MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSourcemacAppStore),
		string(MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSourcemacAppStoreAndIdentifiedDevelopers),
		string(MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSourcenotConfigured),
	}
}

func parseMacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource(input string) (*MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource, error) {
	vals := map[string]MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource{
		"anywhere":                           MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSourceanywhere,
		"macappstore":                        MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSourcemacAppStore,
		"macappstoreandidentifieddevelopers": MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSourcemacAppStoreAndIdentifiedDevelopers,
		"notconfigured":                      MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSourcenotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource(input)
	return &out, nil
}
