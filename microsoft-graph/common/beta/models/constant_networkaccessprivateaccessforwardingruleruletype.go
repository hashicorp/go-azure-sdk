package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessPrivateAccessForwardingRuleRuleType string

const (
	NetworkaccessPrivateAccessForwardingRuleRuleTypefqdn        NetworkaccessPrivateAccessForwardingRuleRuleType = "Fqdn"
	NetworkaccessPrivateAccessForwardingRuleRuleTypeipAddress   NetworkaccessPrivateAccessForwardingRuleRuleType = "IpAddress"
	NetworkaccessPrivateAccessForwardingRuleRuleTypeipRange     NetworkaccessPrivateAccessForwardingRuleRuleType = "IpRange"
	NetworkaccessPrivateAccessForwardingRuleRuleTypeipSubnet    NetworkaccessPrivateAccessForwardingRuleRuleType = "IpSubnet"
	NetworkaccessPrivateAccessForwardingRuleRuleTypeurl         NetworkaccessPrivateAccessForwardingRuleRuleType = "Url"
	NetworkaccessPrivateAccessForwardingRuleRuleTypewebCategory NetworkaccessPrivateAccessForwardingRuleRuleType = "WebCategory"
)

func PossibleValuesForNetworkaccessPrivateAccessForwardingRuleRuleType() []string {
	return []string{
		string(NetworkaccessPrivateAccessForwardingRuleRuleTypefqdn),
		string(NetworkaccessPrivateAccessForwardingRuleRuleTypeipAddress),
		string(NetworkaccessPrivateAccessForwardingRuleRuleTypeipRange),
		string(NetworkaccessPrivateAccessForwardingRuleRuleTypeipSubnet),
		string(NetworkaccessPrivateAccessForwardingRuleRuleTypeurl),
		string(NetworkaccessPrivateAccessForwardingRuleRuleTypewebCategory),
	}
}

func parseNetworkaccessPrivateAccessForwardingRuleRuleType(input string) (*NetworkaccessPrivateAccessForwardingRuleRuleType, error) {
	vals := map[string]NetworkaccessPrivateAccessForwardingRuleRuleType{
		"fqdn":        NetworkaccessPrivateAccessForwardingRuleRuleTypefqdn,
		"ipaddress":   NetworkaccessPrivateAccessForwardingRuleRuleTypeipAddress,
		"iprange":     NetworkaccessPrivateAccessForwardingRuleRuleTypeipRange,
		"ipsubnet":    NetworkaccessPrivateAccessForwardingRuleRuleTypeipSubnet,
		"url":         NetworkaccessPrivateAccessForwardingRuleRuleTypeurl,
		"webcategory": NetworkaccessPrivateAccessForwardingRuleRuleTypewebCategory,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessPrivateAccessForwardingRuleRuleType(input)
	return &out, nil
}
