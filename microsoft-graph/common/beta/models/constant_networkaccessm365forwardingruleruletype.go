package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessM365ForwardingRuleRuleType string

const (
	NetworkaccessM365ForwardingRuleRuleTypefqdn        NetworkaccessM365ForwardingRuleRuleType = "Fqdn"
	NetworkaccessM365ForwardingRuleRuleTypeipAddress   NetworkaccessM365ForwardingRuleRuleType = "IpAddress"
	NetworkaccessM365ForwardingRuleRuleTypeipRange     NetworkaccessM365ForwardingRuleRuleType = "IpRange"
	NetworkaccessM365ForwardingRuleRuleTypeipSubnet    NetworkaccessM365ForwardingRuleRuleType = "IpSubnet"
	NetworkaccessM365ForwardingRuleRuleTypeurl         NetworkaccessM365ForwardingRuleRuleType = "Url"
	NetworkaccessM365ForwardingRuleRuleTypewebCategory NetworkaccessM365ForwardingRuleRuleType = "WebCategory"
)

func PossibleValuesForNetworkaccessM365ForwardingRuleRuleType() []string {
	return []string{
		string(NetworkaccessM365ForwardingRuleRuleTypefqdn),
		string(NetworkaccessM365ForwardingRuleRuleTypeipAddress),
		string(NetworkaccessM365ForwardingRuleRuleTypeipRange),
		string(NetworkaccessM365ForwardingRuleRuleTypeipSubnet),
		string(NetworkaccessM365ForwardingRuleRuleTypeurl),
		string(NetworkaccessM365ForwardingRuleRuleTypewebCategory),
	}
}

func parseNetworkaccessM365ForwardingRuleRuleType(input string) (*NetworkaccessM365ForwardingRuleRuleType, error) {
	vals := map[string]NetworkaccessM365ForwardingRuleRuleType{
		"fqdn":        NetworkaccessM365ForwardingRuleRuleTypefqdn,
		"ipaddress":   NetworkaccessM365ForwardingRuleRuleTypeipAddress,
		"iprange":     NetworkaccessM365ForwardingRuleRuleTypeipRange,
		"ipsubnet":    NetworkaccessM365ForwardingRuleRuleTypeipSubnet,
		"url":         NetworkaccessM365ForwardingRuleRuleTypeurl,
		"webcategory": NetworkaccessM365ForwardingRuleRuleTypewebCategory,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessM365ForwardingRuleRuleType(input)
	return &out, nil
}
