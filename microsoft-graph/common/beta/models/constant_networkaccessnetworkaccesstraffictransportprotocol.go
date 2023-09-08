package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessNetworkAccessTrafficTransportProtocol string

const (
	NetworkaccessNetworkAccessTrafficTransportProtocolggp                               NetworkaccessNetworkAccessTrafficTransportProtocol = "Ggp"
	NetworkaccessNetworkAccessTrafficTransportProtocolicmp                              NetworkaccessNetworkAccessTrafficTransportProtocol = "Icmp"
	NetworkaccessNetworkAccessTrafficTransportProtocolicmpV6                            NetworkaccessNetworkAccessTrafficTransportProtocol = "IcmpV6"
	NetworkaccessNetworkAccessTrafficTransportProtocolidp                               NetworkaccessNetworkAccessTrafficTransportProtocol = "Idp"
	NetworkaccessNetworkAccessTrafficTransportProtocoligmp                              NetworkaccessNetworkAccessTrafficTransportProtocol = "Igmp"
	NetworkaccessNetworkAccessTrafficTransportProtocolip                                NetworkaccessNetworkAccessTrafficTransportProtocol = "Ip"
	NetworkaccessNetworkAccessTrafficTransportProtocolipSecAuthenticationHeader         NetworkaccessNetworkAccessTrafficTransportProtocol = "IpSecAuthenticationHeader"
	NetworkaccessNetworkAccessTrafficTransportProtocolipSecEncapsulatingSecurityPayload NetworkaccessNetworkAccessTrafficTransportProtocol = "IpSecEncapsulatingSecurityPayload"
	NetworkaccessNetworkAccessTrafficTransportProtocolipv4                              NetworkaccessNetworkAccessTrafficTransportProtocol = "Ipv4"
	NetworkaccessNetworkAccessTrafficTransportProtocolipv6                              NetworkaccessNetworkAccessTrafficTransportProtocol = "Ipv6"
	NetworkaccessNetworkAccessTrafficTransportProtocolipv6DestinationOptions            NetworkaccessNetworkAccessTrafficTransportProtocol = "Ipv6DestinationOptions"
	NetworkaccessNetworkAccessTrafficTransportProtocolipv6FragmentHeader                NetworkaccessNetworkAccessTrafficTransportProtocol = "Ipv6FragmentHeader"
	NetworkaccessNetworkAccessTrafficTransportProtocolipv6NoNextHeader                  NetworkaccessNetworkAccessTrafficTransportProtocol = "Ipv6NoNextHeader"
	NetworkaccessNetworkAccessTrafficTransportProtocolipv6RoutingHeader                 NetworkaccessNetworkAccessTrafficTransportProtocol = "Ipv6RoutingHeader"
	NetworkaccessNetworkAccessTrafficTransportProtocolipx                               NetworkaccessNetworkAccessTrafficTransportProtocol = "Ipx"
	NetworkaccessNetworkAccessTrafficTransportProtocolnd                                NetworkaccessNetworkAccessTrafficTransportProtocol = "Nd"
	NetworkaccessNetworkAccessTrafficTransportProtocolpup                               NetworkaccessNetworkAccessTrafficTransportProtocol = "Pup"
	NetworkaccessNetworkAccessTrafficTransportProtocolraw                               NetworkaccessNetworkAccessTrafficTransportProtocol = "Raw"
	NetworkaccessNetworkAccessTrafficTransportProtocolspx                               NetworkaccessNetworkAccessTrafficTransportProtocol = "Spx"
	NetworkaccessNetworkAccessTrafficTransportProtocolspxII                             NetworkaccessNetworkAccessTrafficTransportProtocol = "SpxII"
	NetworkaccessNetworkAccessTrafficTransportProtocoltcp                               NetworkaccessNetworkAccessTrafficTransportProtocol = "Tcp"
	NetworkaccessNetworkAccessTrafficTransportProtocoludp                               NetworkaccessNetworkAccessTrafficTransportProtocol = "Udp"
)

func PossibleValuesForNetworkaccessNetworkAccessTrafficTransportProtocol() []string {
	return []string{
		string(NetworkaccessNetworkAccessTrafficTransportProtocolggp),
		string(NetworkaccessNetworkAccessTrafficTransportProtocolicmp),
		string(NetworkaccessNetworkAccessTrafficTransportProtocolicmpV6),
		string(NetworkaccessNetworkAccessTrafficTransportProtocolidp),
		string(NetworkaccessNetworkAccessTrafficTransportProtocoligmp),
		string(NetworkaccessNetworkAccessTrafficTransportProtocolip),
		string(NetworkaccessNetworkAccessTrafficTransportProtocolipSecAuthenticationHeader),
		string(NetworkaccessNetworkAccessTrafficTransportProtocolipSecEncapsulatingSecurityPayload),
		string(NetworkaccessNetworkAccessTrafficTransportProtocolipv4),
		string(NetworkaccessNetworkAccessTrafficTransportProtocolipv6),
		string(NetworkaccessNetworkAccessTrafficTransportProtocolipv6DestinationOptions),
		string(NetworkaccessNetworkAccessTrafficTransportProtocolipv6FragmentHeader),
		string(NetworkaccessNetworkAccessTrafficTransportProtocolipv6NoNextHeader),
		string(NetworkaccessNetworkAccessTrafficTransportProtocolipv6RoutingHeader),
		string(NetworkaccessNetworkAccessTrafficTransportProtocolipx),
		string(NetworkaccessNetworkAccessTrafficTransportProtocolnd),
		string(NetworkaccessNetworkAccessTrafficTransportProtocolpup),
		string(NetworkaccessNetworkAccessTrafficTransportProtocolraw),
		string(NetworkaccessNetworkAccessTrafficTransportProtocolspx),
		string(NetworkaccessNetworkAccessTrafficTransportProtocolspxII),
		string(NetworkaccessNetworkAccessTrafficTransportProtocoltcp),
		string(NetworkaccessNetworkAccessTrafficTransportProtocoludp),
	}
}

func parseNetworkaccessNetworkAccessTrafficTransportProtocol(input string) (*NetworkaccessNetworkAccessTrafficTransportProtocol, error) {
	vals := map[string]NetworkaccessNetworkAccessTrafficTransportProtocol{
		"ggp":                               NetworkaccessNetworkAccessTrafficTransportProtocolggp,
		"icmp":                              NetworkaccessNetworkAccessTrafficTransportProtocolicmp,
		"icmpv6":                            NetworkaccessNetworkAccessTrafficTransportProtocolicmpV6,
		"idp":                               NetworkaccessNetworkAccessTrafficTransportProtocolidp,
		"igmp":                              NetworkaccessNetworkAccessTrafficTransportProtocoligmp,
		"ip":                                NetworkaccessNetworkAccessTrafficTransportProtocolip,
		"ipsecauthenticationheader":         NetworkaccessNetworkAccessTrafficTransportProtocolipSecAuthenticationHeader,
		"ipsecencapsulatingsecuritypayload": NetworkaccessNetworkAccessTrafficTransportProtocolipSecEncapsulatingSecurityPayload,
		"ipv4":                              NetworkaccessNetworkAccessTrafficTransportProtocolipv4,
		"ipv6":                              NetworkaccessNetworkAccessTrafficTransportProtocolipv6,
		"ipv6destinationoptions":            NetworkaccessNetworkAccessTrafficTransportProtocolipv6DestinationOptions,
		"ipv6fragmentheader":                NetworkaccessNetworkAccessTrafficTransportProtocolipv6FragmentHeader,
		"ipv6nonextheader":                  NetworkaccessNetworkAccessTrafficTransportProtocolipv6NoNextHeader,
		"ipv6routingheader":                 NetworkaccessNetworkAccessTrafficTransportProtocolipv6RoutingHeader,
		"ipx":                               NetworkaccessNetworkAccessTrafficTransportProtocolipx,
		"nd":                                NetworkaccessNetworkAccessTrafficTransportProtocolnd,
		"pup":                               NetworkaccessNetworkAccessTrafficTransportProtocolpup,
		"raw":                               NetworkaccessNetworkAccessTrafficTransportProtocolraw,
		"spx":                               NetworkaccessNetworkAccessTrafficTransportProtocolspx,
		"spxii":                             NetworkaccessNetworkAccessTrafficTransportProtocolspxII,
		"tcp":                               NetworkaccessNetworkAccessTrafficTransportProtocoltcp,
		"udp":                               NetworkaccessNetworkAccessTrafficTransportProtocoludp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessNetworkAccessTrafficTransportProtocol(input)
	return &out, nil
}
