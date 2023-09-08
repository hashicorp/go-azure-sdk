package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessForwardingRuleRuleType string

const (
	NetworkaccessForwardingRuleRuleTypefqdn        NetworkaccessForwardingRuleRuleType = "Fqdn"
	NetworkaccessForwardingRuleRuleTypeipAddress   NetworkaccessForwardingRuleRuleType = "IpAddress"
	NetworkaccessForwardingRuleRuleTypeipRange     NetworkaccessForwardingRuleRuleType = "IpRange"
	NetworkaccessForwardingRuleRuleTypeipSubnet    NetworkaccessForwardingRuleRuleType = "IpSubnet"
	NetworkaccessForwardingRuleRuleTypeurl         NetworkaccessForwardingRuleRuleType = "Url"
	NetworkaccessForwardingRuleRuleTypewebCategory NetworkaccessForwardingRuleRuleType = "WebCategory"
)

func PossibleValuesForNetworkaccessForwardingRuleRuleType() []string {
	return []string{
		string(NetworkaccessForwardingRuleRuleTypefqdn),
		string(NetworkaccessForwardingRuleRuleTypeipAddress),
		string(NetworkaccessForwardingRuleRuleTypeipRange),
		string(NetworkaccessForwardingRuleRuleTypeipSubnet),
		string(NetworkaccessForwardingRuleRuleTypeurl),
		string(NetworkaccessForwardingRuleRuleTypewebCategory),
	}
}

func parseNetworkaccessForwardingRuleRuleType(input string) (*NetworkaccessForwardingRuleRuleType, error) {
	vals := map[string]NetworkaccessForwardingRuleRuleType{
		"fqdn":        NetworkaccessForwardingRuleRuleTypefqdn,
		"ipaddress":   NetworkaccessForwardingRuleRuleTypeipAddress,
		"iprange":     NetworkaccessForwardingRuleRuleTypeipRange,
		"ipsubnet":    NetworkaccessForwardingRuleRuleTypeipSubnet,
		"url":         NetworkaccessForwardingRuleRuleTypeurl,
		"webcategory": NetworkaccessForwardingRuleRuleTypewebCategory,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessForwardingRuleRuleType(input)
	return &out, nil
}
