package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption string

const (
	NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryptionaes128    NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption = "Aes128"
	NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryptionaes192    NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption = "Aes192"
	NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryptionaes256    NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption = "Aes256"
	NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryptiongcmAes128 NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption = "GcmAes128"
	NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryptiongcmAes256 NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption = "GcmAes256"
)

func PossibleValuesForNetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption() []string {
	return []string{
		string(NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryptionaes128),
		string(NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryptionaes192),
		string(NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryptionaes256),
		string(NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryptiongcmAes128),
		string(NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryptiongcmAes256),
	}
}

func parseNetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption(input string) (*NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption, error) {
	vals := map[string]NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption{
		"aes128":    NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryptionaes128,
		"aes192":    NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryptionaes192,
		"aes256":    NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryptionaes256,
		"gcmaes128": NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryptiongcmAes128,
		"gcmaes256": NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryptiongcmAes256,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption(input)
	return &out, nil
}
