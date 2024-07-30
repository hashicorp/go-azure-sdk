package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsNetworkInfoNetworkTransportProtocol string

const (
	CallRecordsNetworkInfoNetworkTransportProtocol_Tcp     CallRecordsNetworkInfoNetworkTransportProtocol = "tcp"
	CallRecordsNetworkInfoNetworkTransportProtocol_Udp     CallRecordsNetworkInfoNetworkTransportProtocol = "udp"
	CallRecordsNetworkInfoNetworkTransportProtocol_Unknown CallRecordsNetworkInfoNetworkTransportProtocol = "unknown"
)

func PossibleValuesForCallRecordsNetworkInfoNetworkTransportProtocol() []string {
	return []string{
		string(CallRecordsNetworkInfoNetworkTransportProtocol_Tcp),
		string(CallRecordsNetworkInfoNetworkTransportProtocol_Udp),
		string(CallRecordsNetworkInfoNetworkTransportProtocol_Unknown),
	}
}

func (s *CallRecordsNetworkInfoNetworkTransportProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordsNetworkInfoNetworkTransportProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordsNetworkInfoNetworkTransportProtocol(input string) (*CallRecordsNetworkInfoNetworkTransportProtocol, error) {
	vals := map[string]CallRecordsNetworkInfoNetworkTransportProtocol{
		"tcp":     CallRecordsNetworkInfoNetworkTransportProtocol_Tcp,
		"udp":     CallRecordsNetworkInfoNetworkTransportProtocol_Udp,
		"unknown": CallRecordsNetworkInfoNetworkTransportProtocol_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsNetworkInfoNetworkTransportProtocol(input)
	return &out, nil
}
