package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessNetworkingProtocol string

const (
	NetworkaccessNetworkingProtocol_Ggp                               NetworkaccessNetworkingProtocol = "ggp"
	NetworkaccessNetworkingProtocol_Icmp                              NetworkaccessNetworkingProtocol = "icmp"
	NetworkaccessNetworkingProtocol_IcmpV6                            NetworkaccessNetworkingProtocol = "icmpV6"
	NetworkaccessNetworkingProtocol_Idp                               NetworkaccessNetworkingProtocol = "idp"
	NetworkaccessNetworkingProtocol_Igmp                              NetworkaccessNetworkingProtocol = "igmp"
	NetworkaccessNetworkingProtocol_Ip                                NetworkaccessNetworkingProtocol = "ip"
	NetworkaccessNetworkingProtocol_IpSecAuthenticationHeader         NetworkaccessNetworkingProtocol = "ipSecAuthenticationHeader"
	NetworkaccessNetworkingProtocol_IpSecEncapsulatingSecurityPayload NetworkaccessNetworkingProtocol = "ipSecEncapsulatingSecurityPayload"
	NetworkaccessNetworkingProtocol_Ipv4                              NetworkaccessNetworkingProtocol = "ipv4"
	NetworkaccessNetworkingProtocol_Ipv6                              NetworkaccessNetworkingProtocol = "ipv6"
	NetworkaccessNetworkingProtocol_Ipv6DestinationOptions            NetworkaccessNetworkingProtocol = "ipv6DestinationOptions"
	NetworkaccessNetworkingProtocol_Ipv6FragmentHeader                NetworkaccessNetworkingProtocol = "ipv6FragmentHeader"
	NetworkaccessNetworkingProtocol_Ipv6NoNextHeader                  NetworkaccessNetworkingProtocol = "ipv6NoNextHeader"
	NetworkaccessNetworkingProtocol_Ipv6RoutingHeader                 NetworkaccessNetworkingProtocol = "ipv6RoutingHeader"
	NetworkaccessNetworkingProtocol_Ipx                               NetworkaccessNetworkingProtocol = "ipx"
	NetworkaccessNetworkingProtocol_Nd                                NetworkaccessNetworkingProtocol = "nd"
	NetworkaccessNetworkingProtocol_Pup                               NetworkaccessNetworkingProtocol = "pup"
	NetworkaccessNetworkingProtocol_Raw                               NetworkaccessNetworkingProtocol = "raw"
	NetworkaccessNetworkingProtocol_Spx                               NetworkaccessNetworkingProtocol = "spx"
	NetworkaccessNetworkingProtocol_SpxII                             NetworkaccessNetworkingProtocol = "spxII"
	NetworkaccessNetworkingProtocol_Tcp                               NetworkaccessNetworkingProtocol = "tcp"
	NetworkaccessNetworkingProtocol_Udp                               NetworkaccessNetworkingProtocol = "udp"
)

func PossibleValuesForNetworkaccessNetworkingProtocol() []string {
	return []string{
		string(NetworkaccessNetworkingProtocol_Ggp),
		string(NetworkaccessNetworkingProtocol_Icmp),
		string(NetworkaccessNetworkingProtocol_IcmpV6),
		string(NetworkaccessNetworkingProtocol_Idp),
		string(NetworkaccessNetworkingProtocol_Igmp),
		string(NetworkaccessNetworkingProtocol_Ip),
		string(NetworkaccessNetworkingProtocol_IpSecAuthenticationHeader),
		string(NetworkaccessNetworkingProtocol_IpSecEncapsulatingSecurityPayload),
		string(NetworkaccessNetworkingProtocol_Ipv4),
		string(NetworkaccessNetworkingProtocol_Ipv6),
		string(NetworkaccessNetworkingProtocol_Ipv6DestinationOptions),
		string(NetworkaccessNetworkingProtocol_Ipv6FragmentHeader),
		string(NetworkaccessNetworkingProtocol_Ipv6NoNextHeader),
		string(NetworkaccessNetworkingProtocol_Ipv6RoutingHeader),
		string(NetworkaccessNetworkingProtocol_Ipx),
		string(NetworkaccessNetworkingProtocol_Nd),
		string(NetworkaccessNetworkingProtocol_Pup),
		string(NetworkaccessNetworkingProtocol_Raw),
		string(NetworkaccessNetworkingProtocol_Spx),
		string(NetworkaccessNetworkingProtocol_SpxII),
		string(NetworkaccessNetworkingProtocol_Tcp),
		string(NetworkaccessNetworkingProtocol_Udp),
	}
}

func (s *NetworkaccessNetworkingProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessNetworkingProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessNetworkingProtocol(input string) (*NetworkaccessNetworkingProtocol, error) {
	vals := map[string]NetworkaccessNetworkingProtocol{
		"ggp":                               NetworkaccessNetworkingProtocol_Ggp,
		"icmp":                              NetworkaccessNetworkingProtocol_Icmp,
		"icmpv6":                            NetworkaccessNetworkingProtocol_IcmpV6,
		"idp":                               NetworkaccessNetworkingProtocol_Idp,
		"igmp":                              NetworkaccessNetworkingProtocol_Igmp,
		"ip":                                NetworkaccessNetworkingProtocol_Ip,
		"ipsecauthenticationheader":         NetworkaccessNetworkingProtocol_IpSecAuthenticationHeader,
		"ipsecencapsulatingsecuritypayload": NetworkaccessNetworkingProtocol_IpSecEncapsulatingSecurityPayload,
		"ipv4":                              NetworkaccessNetworkingProtocol_Ipv4,
		"ipv6":                              NetworkaccessNetworkingProtocol_Ipv6,
		"ipv6destinationoptions":            NetworkaccessNetworkingProtocol_Ipv6DestinationOptions,
		"ipv6fragmentheader":                NetworkaccessNetworkingProtocol_Ipv6FragmentHeader,
		"ipv6nonextheader":                  NetworkaccessNetworkingProtocol_Ipv6NoNextHeader,
		"ipv6routingheader":                 NetworkaccessNetworkingProtocol_Ipv6RoutingHeader,
		"ipx":                               NetworkaccessNetworkingProtocol_Ipx,
		"nd":                                NetworkaccessNetworkingProtocol_Nd,
		"pup":                               NetworkaccessNetworkingProtocol_Pup,
		"raw":                               NetworkaccessNetworkingProtocol_Raw,
		"spx":                               NetworkaccessNetworkingProtocol_Spx,
		"spxii":                             NetworkaccessNetworkingProtocol_SpxII,
		"tcp":                               NetworkaccessNetworkingProtocol_Tcp,
		"udp":                               NetworkaccessNetworkingProtocol_Udp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessNetworkingProtocol(input)
	return &out, nil
}
