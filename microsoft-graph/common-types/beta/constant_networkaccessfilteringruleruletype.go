package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessFilteringRuleRuleType string

const (
	NetworkaccessFilteringRuleRuleType_Fqdn        NetworkaccessFilteringRuleRuleType = "fqdn"
	NetworkaccessFilteringRuleRuleType_IpAddress   NetworkaccessFilteringRuleRuleType = "ipAddress"
	NetworkaccessFilteringRuleRuleType_IpRange     NetworkaccessFilteringRuleRuleType = "ipRange"
	NetworkaccessFilteringRuleRuleType_IpSubnet    NetworkaccessFilteringRuleRuleType = "ipSubnet"
	NetworkaccessFilteringRuleRuleType_Url         NetworkaccessFilteringRuleRuleType = "url"
	NetworkaccessFilteringRuleRuleType_WebCategory NetworkaccessFilteringRuleRuleType = "webCategory"
)

func PossibleValuesForNetworkaccessFilteringRuleRuleType() []string {
	return []string{
		string(NetworkaccessFilteringRuleRuleType_Fqdn),
		string(NetworkaccessFilteringRuleRuleType_IpAddress),
		string(NetworkaccessFilteringRuleRuleType_IpRange),
		string(NetworkaccessFilteringRuleRuleType_IpSubnet),
		string(NetworkaccessFilteringRuleRuleType_Url),
		string(NetworkaccessFilteringRuleRuleType_WebCategory),
	}
}

func (s *NetworkaccessFilteringRuleRuleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessFilteringRuleRuleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessFilteringRuleRuleType(input string) (*NetworkaccessFilteringRuleRuleType, error) {
	vals := map[string]NetworkaccessFilteringRuleRuleType{
		"fqdn":        NetworkaccessFilteringRuleRuleType_Fqdn,
		"ipaddress":   NetworkaccessFilteringRuleRuleType_IpAddress,
		"iprange":     NetworkaccessFilteringRuleRuleType_IpRange,
		"ipsubnet":    NetworkaccessFilteringRuleRuleType_IpSubnet,
		"url":         NetworkaccessFilteringRuleRuleType_Url,
		"webcategory": NetworkaccessFilteringRuleRuleType_WebCategory,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessFilteringRuleRuleType(input)
	return &out, nil
}
