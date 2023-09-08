package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosikEv2VpnConfigurationProviderType string

const (
	IosikEv2VpnConfigurationProviderTypeappProxy      IosikEv2VpnConfigurationProviderType = "AppProxy"
	IosikEv2VpnConfigurationProviderTypenotConfigured IosikEv2VpnConfigurationProviderType = "NotConfigured"
	IosikEv2VpnConfigurationProviderTypepacketTunnel  IosikEv2VpnConfigurationProviderType = "PacketTunnel"
)

func PossibleValuesForIosikEv2VpnConfigurationProviderType() []string {
	return []string{
		string(IosikEv2VpnConfigurationProviderTypeappProxy),
		string(IosikEv2VpnConfigurationProviderTypenotConfigured),
		string(IosikEv2VpnConfigurationProviderTypepacketTunnel),
	}
}

func parseIosikEv2VpnConfigurationProviderType(input string) (*IosikEv2VpnConfigurationProviderType, error) {
	vals := map[string]IosikEv2VpnConfigurationProviderType{
		"appproxy":      IosikEv2VpnConfigurationProviderTypeappProxy,
		"notconfigured": IosikEv2VpnConfigurationProviderTypenotConfigured,
		"packettunnel":  IosikEv2VpnConfigurationProviderTypepacketTunnel,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosikEv2VpnConfigurationProviderType(input)
	return &out, nil
}
