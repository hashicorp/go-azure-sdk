package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption string

const (
	NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryptiongcmAes128 NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption = "GcmAes128"
	NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryptiongcmAes192 NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption = "GcmAes192"
	NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryptiongcmAes256 NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption = "GcmAes256"
	NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryptionnone      NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption = "None"
)

func PossibleValuesForNetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption() []string {
	return []string{
		string(NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryptiongcmAes128),
		string(NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryptiongcmAes192),
		string(NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryptiongcmAes256),
		string(NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryptionnone),
	}
}

func parseNetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption(input string) (*NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption, error) {
	vals := map[string]NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption{
		"gcmaes128": NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryptiongcmAes128,
		"gcmaes192": NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryptiongcmAes192,
		"gcmaes256": NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryptiongcmAes256,
		"none":      NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryptionnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption(input)
	return &out, nil
}
