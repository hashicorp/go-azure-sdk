package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsNetworkInfoNetworkTransportProtocol string

const (
	CallRecordsNetworkInfoNetworkTransportProtocoltcp     CallRecordsNetworkInfoNetworkTransportProtocol = "Tcp"
	CallRecordsNetworkInfoNetworkTransportProtocoludp     CallRecordsNetworkInfoNetworkTransportProtocol = "Udp"
	CallRecordsNetworkInfoNetworkTransportProtocolunknown CallRecordsNetworkInfoNetworkTransportProtocol = "Unknown"
)

func PossibleValuesForCallRecordsNetworkInfoNetworkTransportProtocol() []string {
	return []string{
		string(CallRecordsNetworkInfoNetworkTransportProtocoltcp),
		string(CallRecordsNetworkInfoNetworkTransportProtocoludp),
		string(CallRecordsNetworkInfoNetworkTransportProtocolunknown),
	}
}

func parseCallRecordsNetworkInfoNetworkTransportProtocol(input string) (*CallRecordsNetworkInfoNetworkTransportProtocol, error) {
	vals := map[string]CallRecordsNetworkInfoNetworkTransportProtocol{
		"tcp":     CallRecordsNetworkInfoNetworkTransportProtocoltcp,
		"udp":     CallRecordsNetworkInfoNetworkTransportProtocoludp,
		"unknown": CallRecordsNetworkInfoNetworkTransportProtocolunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsNetworkInfoNetworkTransportProtocol(input)
	return &out, nil
}
