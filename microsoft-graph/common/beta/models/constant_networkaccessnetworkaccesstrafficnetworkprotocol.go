package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessNetworkAccessTrafficNetworkProtocol string

const (
	NetworkaccessNetworkAccessTrafficNetworkProtocolggp                               NetworkaccessNetworkAccessTrafficNetworkProtocol = "Ggp"
	NetworkaccessNetworkAccessTrafficNetworkProtocolicmp                              NetworkaccessNetworkAccessTrafficNetworkProtocol = "Icmp"
	NetworkaccessNetworkAccessTrafficNetworkProtocolicmpV6                            NetworkaccessNetworkAccessTrafficNetworkProtocol = "IcmpV6"
	NetworkaccessNetworkAccessTrafficNetworkProtocolidp                               NetworkaccessNetworkAccessTrafficNetworkProtocol = "Idp"
	NetworkaccessNetworkAccessTrafficNetworkProtocoligmp                              NetworkaccessNetworkAccessTrafficNetworkProtocol = "Igmp"
	NetworkaccessNetworkAccessTrafficNetworkProtocolip                                NetworkaccessNetworkAccessTrafficNetworkProtocol = "Ip"
	NetworkaccessNetworkAccessTrafficNetworkProtocolipSecAuthenticationHeader         NetworkaccessNetworkAccessTrafficNetworkProtocol = "IpSecAuthenticationHeader"
	NetworkaccessNetworkAccessTrafficNetworkProtocolipSecEncapsulatingSecurityPayload NetworkaccessNetworkAccessTrafficNetworkProtocol = "IpSecEncapsulatingSecurityPayload"
	NetworkaccessNetworkAccessTrafficNetworkProtocolipv4                              NetworkaccessNetworkAccessTrafficNetworkProtocol = "Ipv4"
	NetworkaccessNetworkAccessTrafficNetworkProtocolipv6                              NetworkaccessNetworkAccessTrafficNetworkProtocol = "Ipv6"
	NetworkaccessNetworkAccessTrafficNetworkProtocolipv6DestinationOptions            NetworkaccessNetworkAccessTrafficNetworkProtocol = "Ipv6DestinationOptions"
	NetworkaccessNetworkAccessTrafficNetworkProtocolipv6FragmentHeader                NetworkaccessNetworkAccessTrafficNetworkProtocol = "Ipv6FragmentHeader"
	NetworkaccessNetworkAccessTrafficNetworkProtocolipv6NoNextHeader                  NetworkaccessNetworkAccessTrafficNetworkProtocol = "Ipv6NoNextHeader"
	NetworkaccessNetworkAccessTrafficNetworkProtocolipv6RoutingHeader                 NetworkaccessNetworkAccessTrafficNetworkProtocol = "Ipv6RoutingHeader"
	NetworkaccessNetworkAccessTrafficNetworkProtocolipx                               NetworkaccessNetworkAccessTrafficNetworkProtocol = "Ipx"
	NetworkaccessNetworkAccessTrafficNetworkProtocolnd                                NetworkaccessNetworkAccessTrafficNetworkProtocol = "Nd"
	NetworkaccessNetworkAccessTrafficNetworkProtocolpup                               NetworkaccessNetworkAccessTrafficNetworkProtocol = "Pup"
	NetworkaccessNetworkAccessTrafficNetworkProtocolraw                               NetworkaccessNetworkAccessTrafficNetworkProtocol = "Raw"
	NetworkaccessNetworkAccessTrafficNetworkProtocolspx                               NetworkaccessNetworkAccessTrafficNetworkProtocol = "Spx"
	NetworkaccessNetworkAccessTrafficNetworkProtocolspxII                             NetworkaccessNetworkAccessTrafficNetworkProtocol = "SpxII"
	NetworkaccessNetworkAccessTrafficNetworkProtocoltcp                               NetworkaccessNetworkAccessTrafficNetworkProtocol = "Tcp"
	NetworkaccessNetworkAccessTrafficNetworkProtocoludp                               NetworkaccessNetworkAccessTrafficNetworkProtocol = "Udp"
)

func PossibleValuesForNetworkaccessNetworkAccessTrafficNetworkProtocol() []string {
	return []string{
		string(NetworkaccessNetworkAccessTrafficNetworkProtocolggp),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocolicmp),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocolicmpV6),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocolidp),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocoligmp),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocolip),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocolipSecAuthenticationHeader),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocolipSecEncapsulatingSecurityPayload),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocolipv4),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocolipv6),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocolipv6DestinationOptions),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocolipv6FragmentHeader),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocolipv6NoNextHeader),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocolipv6RoutingHeader),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocolipx),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocolnd),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocolpup),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocolraw),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocolspx),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocolspxII),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocoltcp),
		string(NetworkaccessNetworkAccessTrafficNetworkProtocoludp),
	}
}

func parseNetworkaccessNetworkAccessTrafficNetworkProtocol(input string) (*NetworkaccessNetworkAccessTrafficNetworkProtocol, error) {
	vals := map[string]NetworkaccessNetworkAccessTrafficNetworkProtocol{
		"ggp":                               NetworkaccessNetworkAccessTrafficNetworkProtocolggp,
		"icmp":                              NetworkaccessNetworkAccessTrafficNetworkProtocolicmp,
		"icmpv6":                            NetworkaccessNetworkAccessTrafficNetworkProtocolicmpV6,
		"idp":                               NetworkaccessNetworkAccessTrafficNetworkProtocolidp,
		"igmp":                              NetworkaccessNetworkAccessTrafficNetworkProtocoligmp,
		"ip":                                NetworkaccessNetworkAccessTrafficNetworkProtocolip,
		"ipsecauthenticationheader":         NetworkaccessNetworkAccessTrafficNetworkProtocolipSecAuthenticationHeader,
		"ipsecencapsulatingsecuritypayload": NetworkaccessNetworkAccessTrafficNetworkProtocolipSecEncapsulatingSecurityPayload,
		"ipv4":                              NetworkaccessNetworkAccessTrafficNetworkProtocolipv4,
		"ipv6":                              NetworkaccessNetworkAccessTrafficNetworkProtocolipv6,
		"ipv6destinationoptions":            NetworkaccessNetworkAccessTrafficNetworkProtocolipv6DestinationOptions,
		"ipv6fragmentheader":                NetworkaccessNetworkAccessTrafficNetworkProtocolipv6FragmentHeader,
		"ipv6nonextheader":                  NetworkaccessNetworkAccessTrafficNetworkProtocolipv6NoNextHeader,
		"ipv6routingheader":                 NetworkaccessNetworkAccessTrafficNetworkProtocolipv6RoutingHeader,
		"ipx":                               NetworkaccessNetworkAccessTrafficNetworkProtocolipx,
		"nd":                                NetworkaccessNetworkAccessTrafficNetworkProtocolnd,
		"pup":                               NetworkaccessNetworkAccessTrafficNetworkProtocolpup,
		"raw":                               NetworkaccessNetworkAccessTrafficNetworkProtocolraw,
		"spx":                               NetworkaccessNetworkAccessTrafficNetworkProtocolspx,
		"spxii":                             NetworkaccessNetworkAccessTrafficNetworkProtocolspxII,
		"tcp":                               NetworkaccessNetworkAccessTrafficNetworkProtocoltcp,
		"udp":                               NetworkaccessNetworkAccessTrafficNetworkProtocoludp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessNetworkAccessTrafficNetworkProtocol(input)
	return &out, nil
}
