package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity string

const (
	NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegritygcmAes128 NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity = "GcmAes128"
	NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegritygcmAes256 NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity = "GcmAes256"
	NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegritysha256    NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity = "Sha256"
	NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegritysha384    NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity = "Sha384"
)

func PossibleValuesForNetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity() []string {
	return []string{
		string(NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegritygcmAes128),
		string(NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegritygcmAes256),
		string(NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegritysha256),
		string(NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegritysha384),
	}
}

func parseNetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity(input string) (*NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity, error) {
	vals := map[string]NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity{
		"gcmaes128": NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegritygcmAes128,
		"gcmaes256": NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegritygcmAes256,
		"sha256":    NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegritysha256,
		"sha384":    NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegritysha384,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity(input)
	return &out, nil
}
