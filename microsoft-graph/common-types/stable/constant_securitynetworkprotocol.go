package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityNetworkProtocol string

const (
	SecurityNetworkProtocol_Ggp                               SecurityNetworkProtocol = "ggp"
	SecurityNetworkProtocol_Icmp                              SecurityNetworkProtocol = "icmp"
	SecurityNetworkProtocol_IcmpV6                            SecurityNetworkProtocol = "icmpV6"
	SecurityNetworkProtocol_Idp                               SecurityNetworkProtocol = "idp"
	SecurityNetworkProtocol_Igmp                              SecurityNetworkProtocol = "igmp"
	SecurityNetworkProtocol_Ip                                SecurityNetworkProtocol = "ip"
	SecurityNetworkProtocol_IpSecAuthenticationHeader         SecurityNetworkProtocol = "ipSecAuthenticationHeader"
	SecurityNetworkProtocol_IpSecEncapsulatingSecurityPayload SecurityNetworkProtocol = "ipSecEncapsulatingSecurityPayload"
	SecurityNetworkProtocol_Ipv4                              SecurityNetworkProtocol = "ipv4"
	SecurityNetworkProtocol_Ipv6                              SecurityNetworkProtocol = "ipv6"
	SecurityNetworkProtocol_Ipv6DestinationOptions            SecurityNetworkProtocol = "ipv6DestinationOptions"
	SecurityNetworkProtocol_Ipv6FragmentHeader                SecurityNetworkProtocol = "ipv6FragmentHeader"
	SecurityNetworkProtocol_Ipv6NoNextHeader                  SecurityNetworkProtocol = "ipv6NoNextHeader"
	SecurityNetworkProtocol_Ipv6RoutingHeader                 SecurityNetworkProtocol = "ipv6RoutingHeader"
	SecurityNetworkProtocol_Ipx                               SecurityNetworkProtocol = "ipx"
	SecurityNetworkProtocol_Nd                                SecurityNetworkProtocol = "nd"
	SecurityNetworkProtocol_Pup                               SecurityNetworkProtocol = "pup"
	SecurityNetworkProtocol_Raw                               SecurityNetworkProtocol = "raw"
	SecurityNetworkProtocol_Spx                               SecurityNetworkProtocol = "spx"
	SecurityNetworkProtocol_SpxII                             SecurityNetworkProtocol = "spxII"
	SecurityNetworkProtocol_Tcp                               SecurityNetworkProtocol = "tcp"
	SecurityNetworkProtocol_Udp                               SecurityNetworkProtocol = "udp"
	SecurityNetworkProtocol_Unknown                           SecurityNetworkProtocol = "unknown"
)

func PossibleValuesForSecurityNetworkProtocol() []string {
	return []string{
		string(SecurityNetworkProtocol_Ggp),
		string(SecurityNetworkProtocol_Icmp),
		string(SecurityNetworkProtocol_IcmpV6),
		string(SecurityNetworkProtocol_Idp),
		string(SecurityNetworkProtocol_Igmp),
		string(SecurityNetworkProtocol_Ip),
		string(SecurityNetworkProtocol_IpSecAuthenticationHeader),
		string(SecurityNetworkProtocol_IpSecEncapsulatingSecurityPayload),
		string(SecurityNetworkProtocol_Ipv4),
		string(SecurityNetworkProtocol_Ipv6),
		string(SecurityNetworkProtocol_Ipv6DestinationOptions),
		string(SecurityNetworkProtocol_Ipv6FragmentHeader),
		string(SecurityNetworkProtocol_Ipv6NoNextHeader),
		string(SecurityNetworkProtocol_Ipv6RoutingHeader),
		string(SecurityNetworkProtocol_Ipx),
		string(SecurityNetworkProtocol_Nd),
		string(SecurityNetworkProtocol_Pup),
		string(SecurityNetworkProtocol_Raw),
		string(SecurityNetworkProtocol_Spx),
		string(SecurityNetworkProtocol_SpxII),
		string(SecurityNetworkProtocol_Tcp),
		string(SecurityNetworkProtocol_Udp),
		string(SecurityNetworkProtocol_Unknown),
	}
}

func (s *SecurityNetworkProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityNetworkProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityNetworkProtocol(input string) (*SecurityNetworkProtocol, error) {
	vals := map[string]SecurityNetworkProtocol{
		"ggp":                               SecurityNetworkProtocol_Ggp,
		"icmp":                              SecurityNetworkProtocol_Icmp,
		"icmpv6":                            SecurityNetworkProtocol_IcmpV6,
		"idp":                               SecurityNetworkProtocol_Idp,
		"igmp":                              SecurityNetworkProtocol_Igmp,
		"ip":                                SecurityNetworkProtocol_Ip,
		"ipsecauthenticationheader":         SecurityNetworkProtocol_IpSecAuthenticationHeader,
		"ipsecencapsulatingsecuritypayload": SecurityNetworkProtocol_IpSecEncapsulatingSecurityPayload,
		"ipv4":                              SecurityNetworkProtocol_Ipv4,
		"ipv6":                              SecurityNetworkProtocol_Ipv6,
		"ipv6destinationoptions":            SecurityNetworkProtocol_Ipv6DestinationOptions,
		"ipv6fragmentheader":                SecurityNetworkProtocol_Ipv6FragmentHeader,
		"ipv6nonextheader":                  SecurityNetworkProtocol_Ipv6NoNextHeader,
		"ipv6routingheader":                 SecurityNetworkProtocol_Ipv6RoutingHeader,
		"ipx":                               SecurityNetworkProtocol_Ipx,
		"nd":                                SecurityNetworkProtocol_Nd,
		"pup":                               SecurityNetworkProtocol_Pup,
		"raw":                               SecurityNetworkProtocol_Raw,
		"spx":                               SecurityNetworkProtocol_Spx,
		"spxii":                             SecurityNetworkProtocol_SpxII,
		"tcp":                               SecurityNetworkProtocol_Tcp,
		"udp":                               SecurityNetworkProtocol_Udp,
		"unknown":                           SecurityNetworkProtocol_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityNetworkProtocol(input)
	return &out, nil
}
