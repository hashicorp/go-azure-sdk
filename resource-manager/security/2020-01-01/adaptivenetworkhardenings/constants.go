package adaptivenetworkhardenings

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Direction string

const (
	DirectionInbound  Direction = "Inbound"
	DirectionOutbound Direction = "Outbound"
)

func PossibleValuesForDirection() []string {
	return []string{
		string(DirectionInbound),
		string(DirectionOutbound),
	}
}

func parseDirection(input string) (*Direction, error) {
	vals := map[string]Direction{
		"inbound":  DirectionInbound,
		"outbound": DirectionOutbound,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Direction(input)
	return &out, nil
}

type TransportProtocol string

const (
	TransportProtocolTCP TransportProtocol = "TCP"
	TransportProtocolUDP TransportProtocol = "UDP"
)

func PossibleValuesForTransportProtocol() []string {
	return []string{
		string(TransportProtocolTCP),
		string(TransportProtocolUDP),
	}
}

func parseTransportProtocol(input string) (*TransportProtocol, error) {
	vals := map[string]TransportProtocol{
		"tcp": TransportProtocolTCP,
		"udp": TransportProtocolUDP,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TransportProtocol(input)
	return &out, nil
}
