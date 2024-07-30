package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessFqdnFilteringRuleRuleType string

const (
	NetworkaccessFqdnFilteringRuleRuleType_Fqdn        NetworkaccessFqdnFilteringRuleRuleType = "fqdn"
	NetworkaccessFqdnFilteringRuleRuleType_IpAddress   NetworkaccessFqdnFilteringRuleRuleType = "ipAddress"
	NetworkaccessFqdnFilteringRuleRuleType_IpRange     NetworkaccessFqdnFilteringRuleRuleType = "ipRange"
	NetworkaccessFqdnFilteringRuleRuleType_IpSubnet    NetworkaccessFqdnFilteringRuleRuleType = "ipSubnet"
	NetworkaccessFqdnFilteringRuleRuleType_Url         NetworkaccessFqdnFilteringRuleRuleType = "url"
	NetworkaccessFqdnFilteringRuleRuleType_WebCategory NetworkaccessFqdnFilteringRuleRuleType = "webCategory"
)

func PossibleValuesForNetworkaccessFqdnFilteringRuleRuleType() []string {
	return []string{
		string(NetworkaccessFqdnFilteringRuleRuleType_Fqdn),
		string(NetworkaccessFqdnFilteringRuleRuleType_IpAddress),
		string(NetworkaccessFqdnFilteringRuleRuleType_IpRange),
		string(NetworkaccessFqdnFilteringRuleRuleType_IpSubnet),
		string(NetworkaccessFqdnFilteringRuleRuleType_Url),
		string(NetworkaccessFqdnFilteringRuleRuleType_WebCategory),
	}
}

func (s *NetworkaccessFqdnFilteringRuleRuleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessFqdnFilteringRuleRuleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessFqdnFilteringRuleRuleType(input string) (*NetworkaccessFqdnFilteringRuleRuleType, error) {
	vals := map[string]NetworkaccessFqdnFilteringRuleRuleType{
		"fqdn":        NetworkaccessFqdnFilteringRuleRuleType_Fqdn,
		"ipaddress":   NetworkaccessFqdnFilteringRuleRuleType_IpAddress,
		"iprange":     NetworkaccessFqdnFilteringRuleRuleType_IpRange,
		"ipsubnet":    NetworkaccessFqdnFilteringRuleRuleType_IpSubnet,
		"url":         NetworkaccessFqdnFilteringRuleRuleType_Url,
		"webcategory": NetworkaccessFqdnFilteringRuleRuleType_WebCategory,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessFqdnFilteringRuleRuleType(input)
	return &out, nil
}
