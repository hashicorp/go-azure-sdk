package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessM365ForwardingRuleProtocol string

const (
	NetworkaccessM365ForwardingRuleProtocol_Ggp                               NetworkaccessM365ForwardingRuleProtocol = "ggp"
	NetworkaccessM365ForwardingRuleProtocol_Icmp                              NetworkaccessM365ForwardingRuleProtocol = "icmp"
	NetworkaccessM365ForwardingRuleProtocol_IcmpV6                            NetworkaccessM365ForwardingRuleProtocol = "icmpV6"
	NetworkaccessM365ForwardingRuleProtocol_Idp                               NetworkaccessM365ForwardingRuleProtocol = "idp"
	NetworkaccessM365ForwardingRuleProtocol_Igmp                              NetworkaccessM365ForwardingRuleProtocol = "igmp"
	NetworkaccessM365ForwardingRuleProtocol_Ip                                NetworkaccessM365ForwardingRuleProtocol = "ip"
	NetworkaccessM365ForwardingRuleProtocol_IpSecAuthenticationHeader         NetworkaccessM365ForwardingRuleProtocol = "ipSecAuthenticationHeader"
	NetworkaccessM365ForwardingRuleProtocol_IpSecEncapsulatingSecurityPayload NetworkaccessM365ForwardingRuleProtocol = "ipSecEncapsulatingSecurityPayload"
	NetworkaccessM365ForwardingRuleProtocol_Ipv4                              NetworkaccessM365ForwardingRuleProtocol = "ipv4"
	NetworkaccessM365ForwardingRuleProtocol_Ipv6                              NetworkaccessM365ForwardingRuleProtocol = "ipv6"
	NetworkaccessM365ForwardingRuleProtocol_Ipv6DestinationOptions            NetworkaccessM365ForwardingRuleProtocol = "ipv6DestinationOptions"
	NetworkaccessM365ForwardingRuleProtocol_Ipv6FragmentHeader                NetworkaccessM365ForwardingRuleProtocol = "ipv6FragmentHeader"
	NetworkaccessM365ForwardingRuleProtocol_Ipv6NoNextHeader                  NetworkaccessM365ForwardingRuleProtocol = "ipv6NoNextHeader"
	NetworkaccessM365ForwardingRuleProtocol_Ipv6RoutingHeader                 NetworkaccessM365ForwardingRuleProtocol = "ipv6RoutingHeader"
	NetworkaccessM365ForwardingRuleProtocol_Ipx                               NetworkaccessM365ForwardingRuleProtocol = "ipx"
	NetworkaccessM365ForwardingRuleProtocol_Nd                                NetworkaccessM365ForwardingRuleProtocol = "nd"
	NetworkaccessM365ForwardingRuleProtocol_Pup                               NetworkaccessM365ForwardingRuleProtocol = "pup"
	NetworkaccessM365ForwardingRuleProtocol_Raw                               NetworkaccessM365ForwardingRuleProtocol = "raw"
	NetworkaccessM365ForwardingRuleProtocol_Spx                               NetworkaccessM365ForwardingRuleProtocol = "spx"
	NetworkaccessM365ForwardingRuleProtocol_SpxII                             NetworkaccessM365ForwardingRuleProtocol = "spxII"
	NetworkaccessM365ForwardingRuleProtocol_Tcp                               NetworkaccessM365ForwardingRuleProtocol = "tcp"
	NetworkaccessM365ForwardingRuleProtocol_Udp                               NetworkaccessM365ForwardingRuleProtocol = "udp"
)

func PossibleValuesForNetworkaccessM365ForwardingRuleProtocol() []string {
	return []string{
		string(NetworkaccessM365ForwardingRuleProtocol_Ggp),
		string(NetworkaccessM365ForwardingRuleProtocol_Icmp),
		string(NetworkaccessM365ForwardingRuleProtocol_IcmpV6),
		string(NetworkaccessM365ForwardingRuleProtocol_Idp),
		string(NetworkaccessM365ForwardingRuleProtocol_Igmp),
		string(NetworkaccessM365ForwardingRuleProtocol_Ip),
		string(NetworkaccessM365ForwardingRuleProtocol_IpSecAuthenticationHeader),
		string(NetworkaccessM365ForwardingRuleProtocol_IpSecEncapsulatingSecurityPayload),
		string(NetworkaccessM365ForwardingRuleProtocol_Ipv4),
		string(NetworkaccessM365ForwardingRuleProtocol_Ipv6),
		string(NetworkaccessM365ForwardingRuleProtocol_Ipv6DestinationOptions),
		string(NetworkaccessM365ForwardingRuleProtocol_Ipv6FragmentHeader),
		string(NetworkaccessM365ForwardingRuleProtocol_Ipv6NoNextHeader),
		string(NetworkaccessM365ForwardingRuleProtocol_Ipv6RoutingHeader),
		string(NetworkaccessM365ForwardingRuleProtocol_Ipx),
		string(NetworkaccessM365ForwardingRuleProtocol_Nd),
		string(NetworkaccessM365ForwardingRuleProtocol_Pup),
		string(NetworkaccessM365ForwardingRuleProtocol_Raw),
		string(NetworkaccessM365ForwardingRuleProtocol_Spx),
		string(NetworkaccessM365ForwardingRuleProtocol_SpxII),
		string(NetworkaccessM365ForwardingRuleProtocol_Tcp),
		string(NetworkaccessM365ForwardingRuleProtocol_Udp),
	}
}

func (s *NetworkaccessM365ForwardingRuleProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessM365ForwardingRuleProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessM365ForwardingRuleProtocol(input string) (*NetworkaccessM365ForwardingRuleProtocol, error) {
	vals := map[string]NetworkaccessM365ForwardingRuleProtocol{
		"ggp":                               NetworkaccessM365ForwardingRuleProtocol_Ggp,
		"icmp":                              NetworkaccessM365ForwardingRuleProtocol_Icmp,
		"icmpv6":                            NetworkaccessM365ForwardingRuleProtocol_IcmpV6,
		"idp":                               NetworkaccessM365ForwardingRuleProtocol_Idp,
		"igmp":                              NetworkaccessM365ForwardingRuleProtocol_Igmp,
		"ip":                                NetworkaccessM365ForwardingRuleProtocol_Ip,
		"ipsecauthenticationheader":         NetworkaccessM365ForwardingRuleProtocol_IpSecAuthenticationHeader,
		"ipsecencapsulatingsecuritypayload": NetworkaccessM365ForwardingRuleProtocol_IpSecEncapsulatingSecurityPayload,
		"ipv4":                              NetworkaccessM365ForwardingRuleProtocol_Ipv4,
		"ipv6":                              NetworkaccessM365ForwardingRuleProtocol_Ipv6,
		"ipv6destinationoptions":            NetworkaccessM365ForwardingRuleProtocol_Ipv6DestinationOptions,
		"ipv6fragmentheader":                NetworkaccessM365ForwardingRuleProtocol_Ipv6FragmentHeader,
		"ipv6nonextheader":                  NetworkaccessM365ForwardingRuleProtocol_Ipv6NoNextHeader,
		"ipv6routingheader":                 NetworkaccessM365ForwardingRuleProtocol_Ipv6RoutingHeader,
		"ipx":                               NetworkaccessM365ForwardingRuleProtocol_Ipx,
		"nd":                                NetworkaccessM365ForwardingRuleProtocol_Nd,
		"pup":                               NetworkaccessM365ForwardingRuleProtocol_Pup,
		"raw":                               NetworkaccessM365ForwardingRuleProtocol_Raw,
		"spx":                               NetworkaccessM365ForwardingRuleProtocol_Spx,
		"spxii":                             NetworkaccessM365ForwardingRuleProtocol_SpxII,
		"tcp":                               NetworkaccessM365ForwardingRuleProtocol_Tcp,
		"udp":                               NetworkaccessM365ForwardingRuleProtocol_Udp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessM365ForwardingRuleProtocol(input)
	return &out, nil
}
