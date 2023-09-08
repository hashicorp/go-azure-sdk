package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkConnectionProtocol string

const (
	NetworkConnectionProtocolggp                               NetworkConnectionProtocol = "Ggp"
	NetworkConnectionProtocolicmp                              NetworkConnectionProtocol = "Icmp"
	NetworkConnectionProtocolicmpV6                            NetworkConnectionProtocol = "IcmpV6"
	NetworkConnectionProtocolidp                               NetworkConnectionProtocol = "Idp"
	NetworkConnectionProtocoligmp                              NetworkConnectionProtocol = "Igmp"
	NetworkConnectionProtocolip                                NetworkConnectionProtocol = "Ip"
	NetworkConnectionProtocolipSecAuthenticationHeader         NetworkConnectionProtocol = "IpSecAuthenticationHeader"
	NetworkConnectionProtocolipSecEncapsulatingSecurityPayload NetworkConnectionProtocol = "IpSecEncapsulatingSecurityPayload"
	NetworkConnectionProtocolipv4                              NetworkConnectionProtocol = "Ipv4"
	NetworkConnectionProtocolipv6                              NetworkConnectionProtocol = "Ipv6"
	NetworkConnectionProtocolipv6DestinationOptions            NetworkConnectionProtocol = "Ipv6DestinationOptions"
	NetworkConnectionProtocolipv6FragmentHeader                NetworkConnectionProtocol = "Ipv6FragmentHeader"
	NetworkConnectionProtocolipv6NoNextHeader                  NetworkConnectionProtocol = "Ipv6NoNextHeader"
	NetworkConnectionProtocolipv6RoutingHeader                 NetworkConnectionProtocol = "Ipv6RoutingHeader"
	NetworkConnectionProtocolipx                               NetworkConnectionProtocol = "Ipx"
	NetworkConnectionProtocolnd                                NetworkConnectionProtocol = "Nd"
	NetworkConnectionProtocolpup                               NetworkConnectionProtocol = "Pup"
	NetworkConnectionProtocolraw                               NetworkConnectionProtocol = "Raw"
	NetworkConnectionProtocolspx                               NetworkConnectionProtocol = "Spx"
	NetworkConnectionProtocolspxII                             NetworkConnectionProtocol = "SpxII"
	NetworkConnectionProtocoltcp                               NetworkConnectionProtocol = "Tcp"
	NetworkConnectionProtocoludp                               NetworkConnectionProtocol = "Udp"
	NetworkConnectionProtocolunknown                           NetworkConnectionProtocol = "Unknown"
)

func PossibleValuesForNetworkConnectionProtocol() []string {
	return []string{
		string(NetworkConnectionProtocolggp),
		string(NetworkConnectionProtocolicmp),
		string(NetworkConnectionProtocolicmpV6),
		string(NetworkConnectionProtocolidp),
		string(NetworkConnectionProtocoligmp),
		string(NetworkConnectionProtocolip),
		string(NetworkConnectionProtocolipSecAuthenticationHeader),
		string(NetworkConnectionProtocolipSecEncapsulatingSecurityPayload),
		string(NetworkConnectionProtocolipv4),
		string(NetworkConnectionProtocolipv6),
		string(NetworkConnectionProtocolipv6DestinationOptions),
		string(NetworkConnectionProtocolipv6FragmentHeader),
		string(NetworkConnectionProtocolipv6NoNextHeader),
		string(NetworkConnectionProtocolipv6RoutingHeader),
		string(NetworkConnectionProtocolipx),
		string(NetworkConnectionProtocolnd),
		string(NetworkConnectionProtocolpup),
		string(NetworkConnectionProtocolraw),
		string(NetworkConnectionProtocolspx),
		string(NetworkConnectionProtocolspxII),
		string(NetworkConnectionProtocoltcp),
		string(NetworkConnectionProtocoludp),
		string(NetworkConnectionProtocolunknown),
	}
}

func parseNetworkConnectionProtocol(input string) (*NetworkConnectionProtocol, error) {
	vals := map[string]NetworkConnectionProtocol{
		"ggp":                               NetworkConnectionProtocolggp,
		"icmp":                              NetworkConnectionProtocolicmp,
		"icmpv6":                            NetworkConnectionProtocolicmpV6,
		"idp":                               NetworkConnectionProtocolidp,
		"igmp":                              NetworkConnectionProtocoligmp,
		"ip":                                NetworkConnectionProtocolip,
		"ipsecauthenticationheader":         NetworkConnectionProtocolipSecAuthenticationHeader,
		"ipsecencapsulatingsecuritypayload": NetworkConnectionProtocolipSecEncapsulatingSecurityPayload,
		"ipv4":                              NetworkConnectionProtocolipv4,
		"ipv6":                              NetworkConnectionProtocolipv6,
		"ipv6destinationoptions":            NetworkConnectionProtocolipv6DestinationOptions,
		"ipv6fragmentheader":                NetworkConnectionProtocolipv6FragmentHeader,
		"ipv6nonextheader":                  NetworkConnectionProtocolipv6NoNextHeader,
		"ipv6routingheader":                 NetworkConnectionProtocolipv6RoutingHeader,
		"ipx":                               NetworkConnectionProtocolipx,
		"nd":                                NetworkConnectionProtocolnd,
		"pup":                               NetworkConnectionProtocolpup,
		"raw":                               NetworkConnectionProtocolraw,
		"spx":                               NetworkConnectionProtocolspx,
		"spxii":                             NetworkConnectionProtocolspxII,
		"tcp":                               NetworkConnectionProtocoltcp,
		"udp":                               NetworkConnectionProtocoludp,
		"unknown":                           NetworkConnectionProtocolunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkConnectionProtocol(input)
	return &out, nil
}
