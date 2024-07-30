package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessNetworkDestinationType string

const (
	NetworkaccessNetworkDestinationType_Fqdn        NetworkaccessNetworkDestinationType = "fqdn"
	NetworkaccessNetworkDestinationType_IpAddress   NetworkaccessNetworkDestinationType = "ipAddress"
	NetworkaccessNetworkDestinationType_IpRange     NetworkaccessNetworkDestinationType = "ipRange"
	NetworkaccessNetworkDestinationType_IpSubnet    NetworkaccessNetworkDestinationType = "ipSubnet"
	NetworkaccessNetworkDestinationType_Url         NetworkaccessNetworkDestinationType = "url"
	NetworkaccessNetworkDestinationType_WebCategory NetworkaccessNetworkDestinationType = "webCategory"
)

func PossibleValuesForNetworkaccessNetworkDestinationType() []string {
	return []string{
		string(NetworkaccessNetworkDestinationType_Fqdn),
		string(NetworkaccessNetworkDestinationType_IpAddress),
		string(NetworkaccessNetworkDestinationType_IpRange),
		string(NetworkaccessNetworkDestinationType_IpSubnet),
		string(NetworkaccessNetworkDestinationType_Url),
		string(NetworkaccessNetworkDestinationType_WebCategory),
	}
}

func (s *NetworkaccessNetworkDestinationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessNetworkDestinationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessNetworkDestinationType(input string) (*NetworkaccessNetworkDestinationType, error) {
	vals := map[string]NetworkaccessNetworkDestinationType{
		"fqdn":        NetworkaccessNetworkDestinationType_Fqdn,
		"ipaddress":   NetworkaccessNetworkDestinationType_IpAddress,
		"iprange":     NetworkaccessNetworkDestinationType_IpRange,
		"ipsubnet":    NetworkaccessNetworkDestinationType_IpSubnet,
		"url":         NetworkaccessNetworkDestinationType_Url,
		"webcategory": NetworkaccessNetworkDestinationType_WebCategory,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessNetworkDestinationType(input)
	return &out, nil
}
