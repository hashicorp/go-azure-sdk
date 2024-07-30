package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessInternetAccessForwardingRuleRuleType string

const (
	NetworkaccessInternetAccessForwardingRuleRuleType_Fqdn        NetworkaccessInternetAccessForwardingRuleRuleType = "fqdn"
	NetworkaccessInternetAccessForwardingRuleRuleType_IpAddress   NetworkaccessInternetAccessForwardingRuleRuleType = "ipAddress"
	NetworkaccessInternetAccessForwardingRuleRuleType_IpRange     NetworkaccessInternetAccessForwardingRuleRuleType = "ipRange"
	NetworkaccessInternetAccessForwardingRuleRuleType_IpSubnet    NetworkaccessInternetAccessForwardingRuleRuleType = "ipSubnet"
	NetworkaccessInternetAccessForwardingRuleRuleType_Url         NetworkaccessInternetAccessForwardingRuleRuleType = "url"
	NetworkaccessInternetAccessForwardingRuleRuleType_WebCategory NetworkaccessInternetAccessForwardingRuleRuleType = "webCategory"
)

func PossibleValuesForNetworkaccessInternetAccessForwardingRuleRuleType() []string {
	return []string{
		string(NetworkaccessInternetAccessForwardingRuleRuleType_Fqdn),
		string(NetworkaccessInternetAccessForwardingRuleRuleType_IpAddress),
		string(NetworkaccessInternetAccessForwardingRuleRuleType_IpRange),
		string(NetworkaccessInternetAccessForwardingRuleRuleType_IpSubnet),
		string(NetworkaccessInternetAccessForwardingRuleRuleType_Url),
		string(NetworkaccessInternetAccessForwardingRuleRuleType_WebCategory),
	}
}

func (s *NetworkaccessInternetAccessForwardingRuleRuleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessInternetAccessForwardingRuleRuleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessInternetAccessForwardingRuleRuleType(input string) (*NetworkaccessInternetAccessForwardingRuleRuleType, error) {
	vals := map[string]NetworkaccessInternetAccessForwardingRuleRuleType{
		"fqdn":        NetworkaccessInternetAccessForwardingRuleRuleType_Fqdn,
		"ipaddress":   NetworkaccessInternetAccessForwardingRuleRuleType_IpAddress,
		"iprange":     NetworkaccessInternetAccessForwardingRuleRuleType_IpRange,
		"ipsubnet":    NetworkaccessInternetAccessForwardingRuleRuleType_IpSubnet,
		"url":         NetworkaccessInternetAccessForwardingRuleRuleType_Url,
		"webcategory": NetworkaccessInternetAccessForwardingRuleRuleType_WebCategory,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessInternetAccessForwardingRuleRuleType(input)
	return &out, nil
}
