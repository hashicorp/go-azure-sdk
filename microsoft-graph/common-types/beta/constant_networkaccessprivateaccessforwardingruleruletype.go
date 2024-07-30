package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessPrivateAccessForwardingRuleRuleType string

const (
	NetworkaccessPrivateAccessForwardingRuleRuleType_Fqdn        NetworkaccessPrivateAccessForwardingRuleRuleType = "fqdn"
	NetworkaccessPrivateAccessForwardingRuleRuleType_IpAddress   NetworkaccessPrivateAccessForwardingRuleRuleType = "ipAddress"
	NetworkaccessPrivateAccessForwardingRuleRuleType_IpRange     NetworkaccessPrivateAccessForwardingRuleRuleType = "ipRange"
	NetworkaccessPrivateAccessForwardingRuleRuleType_IpSubnet    NetworkaccessPrivateAccessForwardingRuleRuleType = "ipSubnet"
	NetworkaccessPrivateAccessForwardingRuleRuleType_Url         NetworkaccessPrivateAccessForwardingRuleRuleType = "url"
	NetworkaccessPrivateAccessForwardingRuleRuleType_WebCategory NetworkaccessPrivateAccessForwardingRuleRuleType = "webCategory"
)

func PossibleValuesForNetworkaccessPrivateAccessForwardingRuleRuleType() []string {
	return []string{
		string(NetworkaccessPrivateAccessForwardingRuleRuleType_Fqdn),
		string(NetworkaccessPrivateAccessForwardingRuleRuleType_IpAddress),
		string(NetworkaccessPrivateAccessForwardingRuleRuleType_IpRange),
		string(NetworkaccessPrivateAccessForwardingRuleRuleType_IpSubnet),
		string(NetworkaccessPrivateAccessForwardingRuleRuleType_Url),
		string(NetworkaccessPrivateAccessForwardingRuleRuleType_WebCategory),
	}
}

func (s *NetworkaccessPrivateAccessForwardingRuleRuleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessPrivateAccessForwardingRuleRuleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessPrivateAccessForwardingRuleRuleType(input string) (*NetworkaccessPrivateAccessForwardingRuleRuleType, error) {
	vals := map[string]NetworkaccessPrivateAccessForwardingRuleRuleType{
		"fqdn":        NetworkaccessPrivateAccessForwardingRuleRuleType_Fqdn,
		"ipaddress":   NetworkaccessPrivateAccessForwardingRuleRuleType_IpAddress,
		"iprange":     NetworkaccessPrivateAccessForwardingRuleRuleType_IpRange,
		"ipsubnet":    NetworkaccessPrivateAccessForwardingRuleRuleType_IpSubnet,
		"url":         NetworkaccessPrivateAccessForwardingRuleRuleType_Url,
		"webcategory": NetworkaccessPrivateAccessForwardingRuleRuleType_WebCategory,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessPrivateAccessForwardingRuleRuleType(input)
	return &out, nil
}
