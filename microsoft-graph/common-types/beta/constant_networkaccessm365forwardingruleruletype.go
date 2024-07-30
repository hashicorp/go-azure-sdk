package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessM365ForwardingRuleRuleType string

const (
	NetworkaccessM365ForwardingRuleRuleType_Fqdn        NetworkaccessM365ForwardingRuleRuleType = "fqdn"
	NetworkaccessM365ForwardingRuleRuleType_IpAddress   NetworkaccessM365ForwardingRuleRuleType = "ipAddress"
	NetworkaccessM365ForwardingRuleRuleType_IpRange     NetworkaccessM365ForwardingRuleRuleType = "ipRange"
	NetworkaccessM365ForwardingRuleRuleType_IpSubnet    NetworkaccessM365ForwardingRuleRuleType = "ipSubnet"
	NetworkaccessM365ForwardingRuleRuleType_Url         NetworkaccessM365ForwardingRuleRuleType = "url"
	NetworkaccessM365ForwardingRuleRuleType_WebCategory NetworkaccessM365ForwardingRuleRuleType = "webCategory"
)

func PossibleValuesForNetworkaccessM365ForwardingRuleRuleType() []string {
	return []string{
		string(NetworkaccessM365ForwardingRuleRuleType_Fqdn),
		string(NetworkaccessM365ForwardingRuleRuleType_IpAddress),
		string(NetworkaccessM365ForwardingRuleRuleType_IpRange),
		string(NetworkaccessM365ForwardingRuleRuleType_IpSubnet),
		string(NetworkaccessM365ForwardingRuleRuleType_Url),
		string(NetworkaccessM365ForwardingRuleRuleType_WebCategory),
	}
}

func (s *NetworkaccessM365ForwardingRuleRuleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessM365ForwardingRuleRuleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessM365ForwardingRuleRuleType(input string) (*NetworkaccessM365ForwardingRuleRuleType, error) {
	vals := map[string]NetworkaccessM365ForwardingRuleRuleType{
		"fqdn":        NetworkaccessM365ForwardingRuleRuleType_Fqdn,
		"ipaddress":   NetworkaccessM365ForwardingRuleRuleType_IpAddress,
		"iprange":     NetworkaccessM365ForwardingRuleRuleType_IpRange,
		"ipsubnet":    NetworkaccessM365ForwardingRuleRuleType_IpSubnet,
		"url":         NetworkaccessM365ForwardingRuleRuleType_Url,
		"webcategory": NetworkaccessM365ForwardingRuleRuleType_WebCategory,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessM365ForwardingRuleRuleType(input)
	return &out, nil
}
