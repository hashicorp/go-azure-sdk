package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkConnectionProtocol string

const (
	NetworkConnectionProtocol_Ggp                               NetworkConnectionProtocol = "ggp"
	NetworkConnectionProtocol_Icmp                              NetworkConnectionProtocol = "icmp"
	NetworkConnectionProtocol_IcmpV6                            NetworkConnectionProtocol = "icmpV6"
	NetworkConnectionProtocol_Idp                               NetworkConnectionProtocol = "idp"
	NetworkConnectionProtocol_Igmp                              NetworkConnectionProtocol = "igmp"
	NetworkConnectionProtocol_Ip                                NetworkConnectionProtocol = "ip"
	NetworkConnectionProtocol_IpSecAuthenticationHeader         NetworkConnectionProtocol = "ipSecAuthenticationHeader"
	NetworkConnectionProtocol_IpSecEncapsulatingSecurityPayload NetworkConnectionProtocol = "ipSecEncapsulatingSecurityPayload"
	NetworkConnectionProtocol_Ipv4                              NetworkConnectionProtocol = "ipv4"
	NetworkConnectionProtocol_Ipv6                              NetworkConnectionProtocol = "ipv6"
	NetworkConnectionProtocol_Ipv6DestinationOptions            NetworkConnectionProtocol = "ipv6DestinationOptions"
	NetworkConnectionProtocol_Ipv6FragmentHeader                NetworkConnectionProtocol = "ipv6FragmentHeader"
	NetworkConnectionProtocol_Ipv6NoNextHeader                  NetworkConnectionProtocol = "ipv6NoNextHeader"
	NetworkConnectionProtocol_Ipv6RoutingHeader                 NetworkConnectionProtocol = "ipv6RoutingHeader"
	NetworkConnectionProtocol_Ipx                               NetworkConnectionProtocol = "ipx"
	NetworkConnectionProtocol_Nd                                NetworkConnectionProtocol = "nd"
	NetworkConnectionProtocol_Pup                               NetworkConnectionProtocol = "pup"
	NetworkConnectionProtocol_Raw                               NetworkConnectionProtocol = "raw"
	NetworkConnectionProtocol_Spx                               NetworkConnectionProtocol = "spx"
	NetworkConnectionProtocol_SpxII                             NetworkConnectionProtocol = "spxII"
	NetworkConnectionProtocol_Tcp                               NetworkConnectionProtocol = "tcp"
	NetworkConnectionProtocol_Udp                               NetworkConnectionProtocol = "udp"
	NetworkConnectionProtocol_Unknown                           NetworkConnectionProtocol = "unknown"
)

func PossibleValuesForNetworkConnectionProtocol() []string {
	return []string{
		string(NetworkConnectionProtocol_Ggp),
		string(NetworkConnectionProtocol_Icmp),
		string(NetworkConnectionProtocol_IcmpV6),
		string(NetworkConnectionProtocol_Idp),
		string(NetworkConnectionProtocol_Igmp),
		string(NetworkConnectionProtocol_Ip),
		string(NetworkConnectionProtocol_IpSecAuthenticationHeader),
		string(NetworkConnectionProtocol_IpSecEncapsulatingSecurityPayload),
		string(NetworkConnectionProtocol_Ipv4),
		string(NetworkConnectionProtocol_Ipv6),
		string(NetworkConnectionProtocol_Ipv6DestinationOptions),
		string(NetworkConnectionProtocol_Ipv6FragmentHeader),
		string(NetworkConnectionProtocol_Ipv6NoNextHeader),
		string(NetworkConnectionProtocol_Ipv6RoutingHeader),
		string(NetworkConnectionProtocol_Ipx),
		string(NetworkConnectionProtocol_Nd),
		string(NetworkConnectionProtocol_Pup),
		string(NetworkConnectionProtocol_Raw),
		string(NetworkConnectionProtocol_Spx),
		string(NetworkConnectionProtocol_SpxII),
		string(NetworkConnectionProtocol_Tcp),
		string(NetworkConnectionProtocol_Udp),
		string(NetworkConnectionProtocol_Unknown),
	}
}

func (s *NetworkConnectionProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkConnectionProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkConnectionProtocol(input string) (*NetworkConnectionProtocol, error) {
	vals := map[string]NetworkConnectionProtocol{
		"ggp":                               NetworkConnectionProtocol_Ggp,
		"icmp":                              NetworkConnectionProtocol_Icmp,
		"icmpv6":                            NetworkConnectionProtocol_IcmpV6,
		"idp":                               NetworkConnectionProtocol_Idp,
		"igmp":                              NetworkConnectionProtocol_Igmp,
		"ip":                                NetworkConnectionProtocol_Ip,
		"ipsecauthenticationheader":         NetworkConnectionProtocol_IpSecAuthenticationHeader,
		"ipsecencapsulatingsecuritypayload": NetworkConnectionProtocol_IpSecEncapsulatingSecurityPayload,
		"ipv4":                              NetworkConnectionProtocol_Ipv4,
		"ipv6":                              NetworkConnectionProtocol_Ipv6,
		"ipv6destinationoptions":            NetworkConnectionProtocol_Ipv6DestinationOptions,
		"ipv6fragmentheader":                NetworkConnectionProtocol_Ipv6FragmentHeader,
		"ipv6nonextheader":                  NetworkConnectionProtocol_Ipv6NoNextHeader,
		"ipv6routingheader":                 NetworkConnectionProtocol_Ipv6RoutingHeader,
		"ipx":                               NetworkConnectionProtocol_Ipx,
		"nd":                                NetworkConnectionProtocol_Nd,
		"pup":                               NetworkConnectionProtocol_Pup,
		"raw":                               NetworkConnectionProtocol_Raw,
		"spx":                               NetworkConnectionProtocol_Spx,
		"spxii":                             NetworkConnectionProtocol_SpxII,
		"tcp":                               NetworkConnectionProtocol_Tcp,
		"udp":                               NetworkConnectionProtocol_Udp,
		"unknown":                           NetworkConnectionProtocol_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkConnectionProtocol(input)
	return &out, nil
}
