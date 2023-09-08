package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSVpnConfigurationProviderType string

const (
	MacOSVpnConfigurationProviderTypeappProxy      MacOSVpnConfigurationProviderType = "AppProxy"
	MacOSVpnConfigurationProviderTypenotConfigured MacOSVpnConfigurationProviderType = "NotConfigured"
	MacOSVpnConfigurationProviderTypepacketTunnel  MacOSVpnConfigurationProviderType = "PacketTunnel"
)

func PossibleValuesForMacOSVpnConfigurationProviderType() []string {
	return []string{
		string(MacOSVpnConfigurationProviderTypeappProxy),
		string(MacOSVpnConfigurationProviderTypenotConfigured),
		string(MacOSVpnConfigurationProviderTypepacketTunnel),
	}
}

func parseMacOSVpnConfigurationProviderType(input string) (*MacOSVpnConfigurationProviderType, error) {
	vals := map[string]MacOSVpnConfigurationProviderType{
		"appproxy":      MacOSVpnConfigurationProviderTypeappProxy,
		"notconfigured": MacOSVpnConfigurationProviderTypenotConfigured,
		"packettunnel":  MacOSVpnConfigurationProviderTypepacketTunnel,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSVpnConfigurationProviderType(input)
	return &out, nil
}
