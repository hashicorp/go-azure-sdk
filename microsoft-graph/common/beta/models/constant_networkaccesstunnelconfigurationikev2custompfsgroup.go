package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup string

const (
	NetworkaccessTunnelConfigurationIKEv2CustomPfsGroupecp256  NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup = "Ecp256"
	NetworkaccessTunnelConfigurationIKEv2CustomPfsGroupecp384  NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup = "Ecp384"
	NetworkaccessTunnelConfigurationIKEv2CustomPfsGroupnone    NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup = "None"
	NetworkaccessTunnelConfigurationIKEv2CustomPfsGrouppfs1    NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup = "Pfs1"
	NetworkaccessTunnelConfigurationIKEv2CustomPfsGrouppfs14   NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup = "Pfs14"
	NetworkaccessTunnelConfigurationIKEv2CustomPfsGrouppfs2    NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup = "Pfs2"
	NetworkaccessTunnelConfigurationIKEv2CustomPfsGrouppfs2048 NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup = "Pfs2048"
	NetworkaccessTunnelConfigurationIKEv2CustomPfsGrouppfs24   NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup = "Pfs24"
	NetworkaccessTunnelConfigurationIKEv2CustomPfsGrouppfsmm   NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup = "Pfsmm"
)

func PossibleValuesForNetworkaccessTunnelConfigurationIKEv2CustomPfsGroup() []string {
	return []string{
		string(NetworkaccessTunnelConfigurationIKEv2CustomPfsGroupecp256),
		string(NetworkaccessTunnelConfigurationIKEv2CustomPfsGroupecp384),
		string(NetworkaccessTunnelConfigurationIKEv2CustomPfsGroupnone),
		string(NetworkaccessTunnelConfigurationIKEv2CustomPfsGrouppfs1),
		string(NetworkaccessTunnelConfigurationIKEv2CustomPfsGrouppfs14),
		string(NetworkaccessTunnelConfigurationIKEv2CustomPfsGrouppfs2),
		string(NetworkaccessTunnelConfigurationIKEv2CustomPfsGrouppfs2048),
		string(NetworkaccessTunnelConfigurationIKEv2CustomPfsGrouppfs24),
		string(NetworkaccessTunnelConfigurationIKEv2CustomPfsGrouppfsmm),
	}
}

func parseNetworkaccessTunnelConfigurationIKEv2CustomPfsGroup(input string) (*NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup, error) {
	vals := map[string]NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup{
		"ecp256":  NetworkaccessTunnelConfigurationIKEv2CustomPfsGroupecp256,
		"ecp384":  NetworkaccessTunnelConfigurationIKEv2CustomPfsGroupecp384,
		"none":    NetworkaccessTunnelConfigurationIKEv2CustomPfsGroupnone,
		"pfs1":    NetworkaccessTunnelConfigurationIKEv2CustomPfsGrouppfs1,
		"pfs14":   NetworkaccessTunnelConfigurationIKEv2CustomPfsGrouppfs14,
		"pfs2":    NetworkaccessTunnelConfigurationIKEv2CustomPfsGrouppfs2,
		"pfs2048": NetworkaccessTunnelConfigurationIKEv2CustomPfsGrouppfs2048,
		"pfs24":   NetworkaccessTunnelConfigurationIKEv2CustomPfsGrouppfs24,
		"pfsmm":   NetworkaccessTunnelConfigurationIKEv2CustomPfsGrouppfsmm,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup(input)
	return &out, nil
}
