package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVpnConfigurationProviderType string

const (
	IosVpnConfigurationProviderTypeappProxy      IosVpnConfigurationProviderType = "AppProxy"
	IosVpnConfigurationProviderTypenotConfigured IosVpnConfigurationProviderType = "NotConfigured"
	IosVpnConfigurationProviderTypepacketTunnel  IosVpnConfigurationProviderType = "PacketTunnel"
)

func PossibleValuesForIosVpnConfigurationProviderType() []string {
	return []string{
		string(IosVpnConfigurationProviderTypeappProxy),
		string(IosVpnConfigurationProviderTypenotConfigured),
		string(IosVpnConfigurationProviderTypepacketTunnel),
	}
}

func parseIosVpnConfigurationProviderType(input string) (*IosVpnConfigurationProviderType, error) {
	vals := map[string]IosVpnConfigurationProviderType{
		"appproxy":      IosVpnConfigurationProviderTypeappProxy,
		"notconfigured": IosVpnConfigurationProviderTypenotConfigured,
		"packettunnel":  IosVpnConfigurationProviderTypepacketTunnel,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosVpnConfigurationProviderType(input)
	return &out, nil
}
