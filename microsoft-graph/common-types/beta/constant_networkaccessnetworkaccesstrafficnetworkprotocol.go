package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessNetworkAccessTrafficNetworkProtocol string

const (
	NetworkaccessNetworkAccessTrafficNetworkProtocol_Ggp                               NetworkaccessNetworkAccessTrafficNetworkProtocol = "ggp"
	NetworkaccessNetworkAccessTrafficNetworkProtocol_Icmp                              NetworkaccessNetworkAccessTrafficNetworkProtocol = "icmp"
	NetworkaccessNetworkAccessTrafficNetworkProtocol_IcmpV6                            NetworkaccessNetworkAccessTrafficNetworkProtocol = "icmpV6"
	NetworkaccessNetworkAccessTrafficNetworkProtocol_Idp                               NetworkaccessNetworkAccessTrafficNetworkProtocol = "idp"
	NetworkaccessNetworkAccessTrafficNetworkProtocol_Igmp                              NetworkaccessNetworkAccessTrafficNetworkProtocol = "igmp"
	NetworkaccessNetworkAccessTrafficNetworkProtocol_Ip                                NetworkaccessNetworkAccessTrafficNetworkProtocol = "ip"
	NetworkaccessNetworkAccessTrafficNetworkProtocol_IpSecAuthenticationHeader         NetworkaccessNetworkAccessTrafficNetworkProtocol = "ipSecAuthenticationHeader"
	NetworkaccessNetworkAccessTrafficNetworkProtocol_IpSecEncapsulatingSecurityPayload NetworkaccessNetworkAccessTrafficNetworkProtocol = "ipSecEncapsulatingSecurityPayload"
	NetworkaccessNetworkAccessTrafficNetworkProtocol_Ipv4                              NetworkaccessNetworkAccessTrafficNetworkProtocol = "ipv4"
	NetworkaccessNetworkAccessTrafficNetworkProtocol_Ipv6                              NetworkaccessNetworkAccessTrafficNetworkProtocol = "ipv6"
	NetworkaccessNetworkAccessTrafficNetworkProtocol_Ipv6DestinationOptions            NetworkaccessNetworkAccessTrafficNetworkProtocol = "ipv6DestinationOptions"
	NetworkaccessNetworkAccessTrafficNetworkProtocol_Ipv6FragmentHeader                NetworkaccessNetworkAccessTrafficNetworkProtocol = "ipv6FragmentHeader"
	NetworkaccessNetworkAccessTrafficNetworkProtocol_Ipv6NoNextHeader                  NetworkaccessNetworkAccessTrafficNetworkProtocol = "ipv6NoNextHeader"
	NetworkaccessNetworkAccessTrafficNetworkProtocol_Ipv6RoutingHeader                 NetworkaccessNetworkAccessTrafficNetworkProtocol = "ipv6RoutingHeader"
	NetworkaccessNetworkAccessTrafficNetworkProtocol_Ipx                               NetworkaccessNetworkAccessTrafficNetworkProtocol = "ipx"
	NetworkaccessNetworkAccessTrafficNetworkProtocol_Nd                                NetworkaccessNetworkAccessTrafficNetworkProtocol = "nd"
	NetworkaccessNetworkAccessTrafficNetworkProtocol_Pup                               NetworkaccessNetworkAccessTrafficNetworkProtocol = "pup"
	NetworkaccessNetworkAccessTrafficNetworkProtocol_Raw                               NetworkaccessNetworkAccessTrafficNetworkProtocol = "raw"
	NetworkaccessNetworkAccessTrafficNetworkProtocol_Spx                               NetworkaccessNetworkAccessTrafficNetworkProtocol = "spx"
	NetworkaccessNetworkAccessTrafficNetworkProtocol_SpxII                             NetworkaccessNetworkAccessTrafficNetworkProtocol = "spxII"
	NetworkaccessNetworkAccessTrafficNetworkProtocol_Tcp                               NetworkaccessNetworkAccessTrafficNetworkProtocol = "tcp"
	NetworkaccessNetworkAccessTrafficNetworkProtocol_Udp                               NetworkaccessNetworkAccessTrafficNetworkProtocol = "udp"
)

func PossibleValuesForNetworkaccessNetworkAccessTrafficNetworkProtocol() []string {
	return []string{
		string(NetworkaccessNetworkAccessTrafficNetworkProtocol_Ggp),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocol_Icmp),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocol_IcmpV6),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocol_Idp),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocol_Igmp),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocol_Ip),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocol_IpSecAuthenticationHeader),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocol_IpSecEncapsulatingSecurityPayload),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocol_Ipv4),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocol_Ipv6),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocol_Ipv6DestinationOptions),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocol_Ipv6FragmentHeader),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocol_Ipv6NoNextHeader),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocol_Ipv6RoutingHeader),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocol_Ipx),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocol_Nd),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocol_Pup),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocol_Raw),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocol_Spx),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocol_SpxII),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocol_Tcp),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocol_Udp),
	}
}

func (s *NetworkaccessNetworkAccessTrafficNetworkProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessNetworkAccessTrafficNetworkProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessNetworkAccessTrafficNetworkProtocol(input string) (*NetworkaccessNetworkAccessTrafficNetworkProtocol, error) {
	vals := map[string]NetworkaccessNetworkAccessTrafficNetworkProtocol{
		"ggp":                               NetworkaccessNetworkAccessTrafficNetworkProtocol_Ggp,
		"icmp":                              NetworkaccessNetworkAccessTrafficNetworkProtocol_Icmp,
		"icmpv6":                            NetworkaccessNetworkAccessTrafficNetworkProtocol_IcmpV6,
		"idp":                               NetworkaccessNetworkAccessTrafficNetworkProtocol_Idp,
		"igmp":                              NetworkaccessNetworkAccessTrafficNetworkProtocol_Igmp,
		"ip":                                NetworkaccessNetworkAccessTrafficNetworkProtocol_Ip,
		"ipsecauthenticationheader":         NetworkaccessNetworkAccessTrafficNetworkProtocol_IpSecAuthenticationHeader,
		"ipsecencapsulatingsecuritypayload": NetworkaccessNetworkAccessTrafficNetworkProtocol_IpSecEncapsulatingSecurityPayload,
		"ipv4":                              NetworkaccessNetworkAccessTrafficNetworkProtocol_Ipv4,
		"ipv6":                              NetworkaccessNetworkAccessTrafficNetworkProtocol_Ipv6,
		"ipv6destinationoptions":            NetworkaccessNetworkAccessTrafficNetworkProtocol_Ipv6DestinationOptions,
		"ipv6fragmentheader":                NetworkaccessNetworkAccessTrafficNetworkProtocol_Ipv6FragmentHeader,
		"ipv6nonextheader":                  NetworkaccessNetworkAccessTrafficNetworkProtocol_Ipv6NoNextHeader,
		"ipv6routingheader":                 NetworkaccessNetworkAccessTrafficNetworkProtocol_Ipv6RoutingHeader,
		"ipx":                               NetworkaccessNetworkAccessTrafficNetworkProtocol_Ipx,
		"nd":                                NetworkaccessNetworkAccessTrafficNetworkProtocol_Nd,
		"pup":                               NetworkaccessNetworkAccessTrafficNetworkProtocol_Pup,
		"raw":                               NetworkaccessNetworkAccessTrafficNetworkProtocol_Raw,
		"spx":                               NetworkaccessNetworkAccessTrafficNetworkProtocol_Spx,
		"spxii":                             NetworkaccessNetworkAccessTrafficNetworkProtocol_SpxII,
		"tcp":                               NetworkaccessNetworkAccessTrafficNetworkProtocol_Tcp,
		"udp":                               NetworkaccessNetworkAccessTrafficNetworkProtocol_Udp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessNetworkAccessTrafficNetworkProtocol(input)
	return &out, nil
}
