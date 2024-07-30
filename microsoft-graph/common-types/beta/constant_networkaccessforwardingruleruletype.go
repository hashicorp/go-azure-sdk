package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessForwardingRuleRuleType string

const (
	NetworkaccessForwardingRuleRuleType_Fqdn        NetworkaccessForwardingRuleRuleType = "fqdn"
	NetworkaccessForwardingRuleRuleType_IpAddress   NetworkaccessForwardingRuleRuleType = "ipAddress"
	NetworkaccessForwardingRuleRuleType_IpRange     NetworkaccessForwardingRuleRuleType = "ipRange"
	NetworkaccessForwardingRuleRuleType_IpSubnet    NetworkaccessForwardingRuleRuleType = "ipSubnet"
	NetworkaccessForwardingRuleRuleType_Url         NetworkaccessForwardingRuleRuleType = "url"
	NetworkaccessForwardingRuleRuleType_WebCategory NetworkaccessForwardingRuleRuleType = "webCategory"
)

func PossibleValuesForNetworkaccessForwardingRuleRuleType() []string {
	return []string{
		string(NetworkaccessForwardingRuleRuleType_Fqdn),
		string(NetworkaccessForwardingRuleRuleType_IpAddress),
		string(NetworkaccessForwardingRuleRuleType_IpRange),
		string(NetworkaccessForwardingRuleRuleType_IpSubnet),
		string(NetworkaccessForwardingRuleRuleType_Url),
		string(NetworkaccessForwardingRuleRuleType_WebCategory),
	}
}

func (s *NetworkaccessForwardingRuleRuleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessForwardingRuleRuleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessForwardingRuleRuleType(input string) (*NetworkaccessForwardingRuleRuleType, error) {
	vals := map[string]NetworkaccessForwardingRuleRuleType{
		"fqdn":        NetworkaccessForwardingRuleRuleType_Fqdn,
		"ipaddress":   NetworkaccessForwardingRuleRuleType_IpAddress,
		"iprange":     NetworkaccessForwardingRuleRuleType_IpRange,
		"ipsubnet":    NetworkaccessForwardingRuleRuleType_IpSubnet,
		"url":         NetworkaccessForwardingRuleRuleType_Url,
		"webcategory": NetworkaccessForwardingRuleRuleType_WebCategory,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessForwardingRuleRuleType(input)
	return &out, nil
}
