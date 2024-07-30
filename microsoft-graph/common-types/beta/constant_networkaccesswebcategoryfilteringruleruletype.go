package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessWebCategoryFilteringRuleRuleType string

const (
	NetworkaccessWebCategoryFilteringRuleRuleType_Fqdn        NetworkaccessWebCategoryFilteringRuleRuleType = "fqdn"
	NetworkaccessWebCategoryFilteringRuleRuleType_IpAddress   NetworkaccessWebCategoryFilteringRuleRuleType = "ipAddress"
	NetworkaccessWebCategoryFilteringRuleRuleType_IpRange     NetworkaccessWebCategoryFilteringRuleRuleType = "ipRange"
	NetworkaccessWebCategoryFilteringRuleRuleType_IpSubnet    NetworkaccessWebCategoryFilteringRuleRuleType = "ipSubnet"
	NetworkaccessWebCategoryFilteringRuleRuleType_Url         NetworkaccessWebCategoryFilteringRuleRuleType = "url"
	NetworkaccessWebCategoryFilteringRuleRuleType_WebCategory NetworkaccessWebCategoryFilteringRuleRuleType = "webCategory"
)

func PossibleValuesForNetworkaccessWebCategoryFilteringRuleRuleType() []string {
	return []string{
		string(NetworkaccessWebCategoryFilteringRuleRuleType_Fqdn),
		string(NetworkaccessWebCategoryFilteringRuleRuleType_IpAddress),
		string(NetworkaccessWebCategoryFilteringRuleRuleType_IpRange),
		string(NetworkaccessWebCategoryFilteringRuleRuleType_IpSubnet),
		string(NetworkaccessWebCategoryFilteringRuleRuleType_Url),
		string(NetworkaccessWebCategoryFilteringRuleRuleType_WebCategory),
	}
}

func (s *NetworkaccessWebCategoryFilteringRuleRuleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessWebCategoryFilteringRuleRuleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessWebCategoryFilteringRuleRuleType(input string) (*NetworkaccessWebCategoryFilteringRuleRuleType, error) {
	vals := map[string]NetworkaccessWebCategoryFilteringRuleRuleType{
		"fqdn":        NetworkaccessWebCategoryFilteringRuleRuleType_Fqdn,
		"ipaddress":   NetworkaccessWebCategoryFilteringRuleRuleType_IpAddress,
		"iprange":     NetworkaccessWebCategoryFilteringRuleRuleType_IpRange,
		"ipsubnet":    NetworkaccessWebCategoryFilteringRuleRuleType_IpSubnet,
		"url":         NetworkaccessWebCategoryFilteringRuleRuleType_Url,
		"webcategory": NetworkaccessWebCategoryFilteringRuleRuleType_WebCategory,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessWebCategoryFilteringRuleRuleType(input)
	return &out, nil
}
