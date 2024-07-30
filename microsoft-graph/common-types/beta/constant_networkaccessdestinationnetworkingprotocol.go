package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessDestinationNetworkingProtocol string

const (
	NetworkaccessDestinationNetworkingProtocol_Ggp                               NetworkaccessDestinationNetworkingProtocol = "ggp"
	NetworkaccessDestinationNetworkingProtocol_Icmp                              NetworkaccessDestinationNetworkingProtocol = "icmp"
	NetworkaccessDestinationNetworkingProtocol_IcmpV6                            NetworkaccessDestinationNetworkingProtocol = "icmpV6"
	NetworkaccessDestinationNetworkingProtocol_Idp                               NetworkaccessDestinationNetworkingProtocol = "idp"
	NetworkaccessDestinationNetworkingProtocol_Igmp                              NetworkaccessDestinationNetworkingProtocol = "igmp"
	NetworkaccessDestinationNetworkingProtocol_Ip                                NetworkaccessDestinationNetworkingProtocol = "ip"
	NetworkaccessDestinationNetworkingProtocol_IpSecAuthenticationHeader         NetworkaccessDestinationNetworkingProtocol = "ipSecAuthenticationHeader"
	NetworkaccessDestinationNetworkingProtocol_IpSecEncapsulatingSecurityPayload NetworkaccessDestinationNetworkingProtocol = "ipSecEncapsulatingSecurityPayload"
	NetworkaccessDestinationNetworkingProtocol_Ipv4                              NetworkaccessDestinationNetworkingProtocol = "ipv4"
	NetworkaccessDestinationNetworkingProtocol_Ipv6                              NetworkaccessDestinationNetworkingProtocol = "ipv6"
	NetworkaccessDestinationNetworkingProtocol_Ipv6DestinationOptions            NetworkaccessDestinationNetworkingProtocol = "ipv6DestinationOptions"
	NetworkaccessDestinationNetworkingProtocol_Ipv6FragmentHeader                NetworkaccessDestinationNetworkingProtocol = "ipv6FragmentHeader"
	NetworkaccessDestinationNetworkingProtocol_Ipv6NoNextHeader                  NetworkaccessDestinationNetworkingProtocol = "ipv6NoNextHeader"
	NetworkaccessDestinationNetworkingProtocol_Ipv6RoutingHeader                 NetworkaccessDestinationNetworkingProtocol = "ipv6RoutingHeader"
	NetworkaccessDestinationNetworkingProtocol_Ipx                               NetworkaccessDestinationNetworkingProtocol = "ipx"
	NetworkaccessDestinationNetworkingProtocol_Nd                                NetworkaccessDestinationNetworkingProtocol = "nd"
	NetworkaccessDestinationNetworkingProtocol_Pup                               NetworkaccessDestinationNetworkingProtocol = "pup"
	NetworkaccessDestinationNetworkingProtocol_Raw                               NetworkaccessDestinationNetworkingProtocol = "raw"
	NetworkaccessDestinationNetworkingProtocol_Spx                               NetworkaccessDestinationNetworkingProtocol = "spx"
	NetworkaccessDestinationNetworkingProtocol_SpxII                             NetworkaccessDestinationNetworkingProtocol = "spxII"
	NetworkaccessDestinationNetworkingProtocol_Tcp                               NetworkaccessDestinationNetworkingProtocol = "tcp"
	NetworkaccessDestinationNetworkingProtocol_Udp                               NetworkaccessDestinationNetworkingProtocol = "udp"
)

func PossibleValuesForNetworkaccessDestinationNetworkingProtocol() []string {
	return []string{
		string(NetworkaccessDestinationNetworkingProtocol_Ggp),
		string(NetworkaccessDestinationNetworkingProtocol_Icmp),
		string(NetworkaccessDestinationNetworkingProtocol_IcmpV6),
		string(NetworkaccessDestinationNetworkingProtocol_Idp),
		string(NetworkaccessDestinationNetworkingProtocol_Igmp),
		string(NetworkaccessDestinationNetworkingProtocol_Ip),
		string(NetworkaccessDestinationNetworkingProtocol_IpSecAuthenticationHeader),
		string(NetworkaccessDestinationNetworkingProtocol_IpSecEncapsulatingSecurityPayload),
		string(NetworkaccessDestinationNetworkingProtocol_Ipv4),
		string(NetworkaccessDestinationNetworkingProtocol_Ipv6),
		string(NetworkaccessDestinationNetworkingProtocol_Ipv6DestinationOptions),
		string(NetworkaccessDestinationNetworkingProtocol_Ipv6FragmentHeader),
		string(NetworkaccessDestinationNetworkingProtocol_Ipv6NoNextHeader),
		string(NetworkaccessDestinationNetworkingProtocol_Ipv6RoutingHeader),
		string(NetworkaccessDestinationNetworkingProtocol_Ipx),
		string(NetworkaccessDestinationNetworkingProtocol_Nd),
		string(NetworkaccessDestinationNetworkingProtocol_Pup),
		string(NetworkaccessDestinationNetworkingProtocol_Raw),
		string(NetworkaccessDestinationNetworkingProtocol_Spx),
		string(NetworkaccessDestinationNetworkingProtocol_SpxII),
		string(NetworkaccessDestinationNetworkingProtocol_Tcp),
		string(NetworkaccessDestinationNetworkingProtocol_Udp),
	}
}

func (s *NetworkaccessDestinationNetworkingProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessDestinationNetworkingProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessDestinationNetworkingProtocol(input string) (*NetworkaccessDestinationNetworkingProtocol, error) {
	vals := map[string]NetworkaccessDestinationNetworkingProtocol{
		"ggp":                               NetworkaccessDestinationNetworkingProtocol_Ggp,
		"icmp":                              NetworkaccessDestinationNetworkingProtocol_Icmp,
		"icmpv6":                            NetworkaccessDestinationNetworkingProtocol_IcmpV6,
		"idp":                               NetworkaccessDestinationNetworkingProtocol_Idp,
		"igmp":                              NetworkaccessDestinationNetworkingProtocol_Igmp,
		"ip":                                NetworkaccessDestinationNetworkingProtocol_Ip,
		"ipsecauthenticationheader":         NetworkaccessDestinationNetworkingProtocol_IpSecAuthenticationHeader,
		"ipsecencapsulatingsecuritypayload": NetworkaccessDestinationNetworkingProtocol_IpSecEncapsulatingSecurityPayload,
		"ipv4":                              NetworkaccessDestinationNetworkingProtocol_Ipv4,
		"ipv6":                              NetworkaccessDestinationNetworkingProtocol_Ipv6,
		"ipv6destinationoptions":            NetworkaccessDestinationNetworkingProtocol_Ipv6DestinationOptions,
		"ipv6fragmentheader":                NetworkaccessDestinationNetworkingProtocol_Ipv6FragmentHeader,
		"ipv6nonextheader":                  NetworkaccessDestinationNetworkingProtocol_Ipv6NoNextHeader,
		"ipv6routingheader":                 NetworkaccessDestinationNetworkingProtocol_Ipv6RoutingHeader,
		"ipx":                               NetworkaccessDestinationNetworkingProtocol_Ipx,
		"nd":                                NetworkaccessDestinationNetworkingProtocol_Nd,
		"pup":                               NetworkaccessDestinationNetworkingProtocol_Pup,
		"raw":                               NetworkaccessDestinationNetworkingProtocol_Raw,
		"spx":                               NetworkaccessDestinationNetworkingProtocol_Spx,
		"spxii":                             NetworkaccessDestinationNetworkingProtocol_SpxII,
		"tcp":                               NetworkaccessDestinationNetworkingProtocol_Tcp,
		"udp":                               NetworkaccessDestinationNetworkingProtocol_Udp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessDestinationNetworkingProtocol(input)
	return &out, nil
}
