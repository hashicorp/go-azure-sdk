package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessNetworkAccessTrafficTransportProtocol string

const (
	NetworkaccessNetworkAccessTrafficTransportProtocol_Ggp                               NetworkaccessNetworkAccessTrafficTransportProtocol = "ggp"
	NetworkaccessNetworkAccessTrafficTransportProtocol_Icmp                              NetworkaccessNetworkAccessTrafficTransportProtocol = "icmp"
	NetworkaccessNetworkAccessTrafficTransportProtocol_IcmpV6                            NetworkaccessNetworkAccessTrafficTransportProtocol = "icmpV6"
	NetworkaccessNetworkAccessTrafficTransportProtocol_Idp                               NetworkaccessNetworkAccessTrafficTransportProtocol = "idp"
	NetworkaccessNetworkAccessTrafficTransportProtocol_Igmp                              NetworkaccessNetworkAccessTrafficTransportProtocol = "igmp"
	NetworkaccessNetworkAccessTrafficTransportProtocol_Ip                                NetworkaccessNetworkAccessTrafficTransportProtocol = "ip"
	NetworkaccessNetworkAccessTrafficTransportProtocol_IpSecAuthenticationHeader         NetworkaccessNetworkAccessTrafficTransportProtocol = "ipSecAuthenticationHeader"
	NetworkaccessNetworkAccessTrafficTransportProtocol_IpSecEncapsulatingSecurityPayload NetworkaccessNetworkAccessTrafficTransportProtocol = "ipSecEncapsulatingSecurityPayload"
	NetworkaccessNetworkAccessTrafficTransportProtocol_Ipv4                              NetworkaccessNetworkAccessTrafficTransportProtocol = "ipv4"
	NetworkaccessNetworkAccessTrafficTransportProtocol_Ipv6                              NetworkaccessNetworkAccessTrafficTransportProtocol = "ipv6"
	NetworkaccessNetworkAccessTrafficTransportProtocol_Ipv6DestinationOptions            NetworkaccessNetworkAccessTrafficTransportProtocol = "ipv6DestinationOptions"
	NetworkaccessNetworkAccessTrafficTransportProtocol_Ipv6FragmentHeader                NetworkaccessNetworkAccessTrafficTransportProtocol = "ipv6FragmentHeader"
	NetworkaccessNetworkAccessTrafficTransportProtocol_Ipv6NoNextHeader                  NetworkaccessNetworkAccessTrafficTransportProtocol = "ipv6NoNextHeader"
	NetworkaccessNetworkAccessTrafficTransportProtocol_Ipv6RoutingHeader                 NetworkaccessNetworkAccessTrafficTransportProtocol = "ipv6RoutingHeader"
	NetworkaccessNetworkAccessTrafficTransportProtocol_Ipx                               NetworkaccessNetworkAccessTrafficTransportProtocol = "ipx"
	NetworkaccessNetworkAccessTrafficTransportProtocol_Nd                                NetworkaccessNetworkAccessTrafficTransportProtocol = "nd"
	NetworkaccessNetworkAccessTrafficTransportProtocol_Pup                               NetworkaccessNetworkAccessTrafficTransportProtocol = "pup"
	NetworkaccessNetworkAccessTrafficTransportProtocol_Raw                               NetworkaccessNetworkAccessTrafficTransportProtocol = "raw"
	NetworkaccessNetworkAccessTrafficTransportProtocol_Spx                               NetworkaccessNetworkAccessTrafficTransportProtocol = "spx"
	NetworkaccessNetworkAccessTrafficTransportProtocol_SpxII                             NetworkaccessNetworkAccessTrafficTransportProtocol = "spxII"
	NetworkaccessNetworkAccessTrafficTransportProtocol_Tcp                               NetworkaccessNetworkAccessTrafficTransportProtocol = "tcp"
	NetworkaccessNetworkAccessTrafficTransportProtocol_Udp                               NetworkaccessNetworkAccessTrafficTransportProtocol = "udp"
)

func PossibleValuesForNetworkaccessNetworkAccessTrafficTransportProtocol() []string {
	return []string{
		string(NetworkaccessNetworkAccessTrafficTransportProtocol_Ggp),
		string(NetworkaccessNetworkAccessTrafficTransportProtocol_Icmp),
		string(NetworkaccessNetworkAccessTrafficTransportProtocol_IcmpV6),
		string(NetworkaccessNetworkAccessTrafficTransportProtocol_Idp),
		string(NetworkaccessNetworkAccessTrafficTransportProtocol_Igmp),
		string(NetworkaccessNetworkAccessTrafficTransportProtocol_Ip),
		string(NetworkaccessNetworkAccessTrafficTransportProtocol_IpSecAuthenticationHeader),
		string(NetworkaccessNetworkAccessTrafficTransportProtocol_IpSecEncapsulatingSecurityPayload),
		string(NetworkaccessNetworkAccessTrafficTransportProtocol_Ipv4),
		string(NetworkaccessNetworkAccessTrafficTransportProtocol_Ipv6),
		string(NetworkaccessNetworkAccessTrafficTransportProtocol_Ipv6DestinationOptions),
		string(NetworkaccessNetworkAccessTrafficTransportProtocol_Ipv6FragmentHeader),
		string(NetworkaccessNetworkAccessTrafficTransportProtocol_Ipv6NoNextHeader),
		string(NetworkaccessNetworkAccessTrafficTransportProtocol_Ipv6RoutingHeader),
		string(NetworkaccessNetworkAccessTrafficTransportProtocol_Ipx),
		string(NetworkaccessNetworkAccessTrafficTransportProtocol_Nd),
		string(NetworkaccessNetworkAccessTrafficTransportProtocol_Pup),
		string(NetworkaccessNetworkAccessTrafficTransportProtocol_Raw),
		string(NetworkaccessNetworkAccessTrafficTransportProtocol_Spx),
		string(NetworkaccessNetworkAccessTrafficTransportProtocol_SpxII),
		string(NetworkaccessNetworkAccessTrafficTransportProtocol_Tcp),
		string(NetworkaccessNetworkAccessTrafficTransportProtocol_Udp),
	}
}

func (s *NetworkaccessNetworkAccessTrafficTransportProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessNetworkAccessTrafficTransportProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessNetworkAccessTrafficTransportProtocol(input string) (*NetworkaccessNetworkAccessTrafficTransportProtocol, error) {
	vals := map[string]NetworkaccessNetworkAccessTrafficTransportProtocol{
		"ggp":                               NetworkaccessNetworkAccessTrafficTransportProtocol_Ggp,
		"icmp":                              NetworkaccessNetworkAccessTrafficTransportProtocol_Icmp,
		"icmpv6":                            NetworkaccessNetworkAccessTrafficTransportProtocol_IcmpV6,
		"idp":                               NetworkaccessNetworkAccessTrafficTransportProtocol_Idp,
		"igmp":                              NetworkaccessNetworkAccessTrafficTransportProtocol_Igmp,
		"ip":                                NetworkaccessNetworkAccessTrafficTransportProtocol_Ip,
		"ipsecauthenticationheader":         NetworkaccessNetworkAccessTrafficTransportProtocol_IpSecAuthenticationHeader,
		"ipsecencapsulatingsecuritypayload": NetworkaccessNetworkAccessTrafficTransportProtocol_IpSecEncapsulatingSecurityPayload,
		"ipv4":                              NetworkaccessNetworkAccessTrafficTransportProtocol_Ipv4,
		"ipv6":                              NetworkaccessNetworkAccessTrafficTransportProtocol_Ipv6,
		"ipv6destinationoptions":            NetworkaccessNetworkAccessTrafficTransportProtocol_Ipv6DestinationOptions,
		"ipv6fragmentheader":                NetworkaccessNetworkAccessTrafficTransportProtocol_Ipv6FragmentHeader,
		"ipv6nonextheader":                  NetworkaccessNetworkAccessTrafficTransportProtocol_Ipv6NoNextHeader,
		"ipv6routingheader":                 NetworkaccessNetworkAccessTrafficTransportProtocol_Ipv6RoutingHeader,
		"ipx":                               NetworkaccessNetworkAccessTrafficTransportProtocol_Ipx,
		"nd":                                NetworkaccessNetworkAccessTrafficTransportProtocol_Nd,
		"pup":                               NetworkaccessNetworkAccessTrafficTransportProtocol_Pup,
		"raw":                               NetworkaccessNetworkAccessTrafficTransportProtocol_Raw,
		"spx":                               NetworkaccessNetworkAccessTrafficTransportProtocol_Spx,
		"spxii":                             NetworkaccessNetworkAccessTrafficTransportProtocol_SpxII,
		"tcp":                               NetworkaccessNetworkAccessTrafficTransportProtocol_Tcp,
		"udp":                               NetworkaccessNetworkAccessTrafficTransportProtocol_Udp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessNetworkAccessTrafficTransportProtocol(input)
	return &out, nil
}
