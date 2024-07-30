package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsNetworkInfoConnectionType string

const (
	CallRecordsNetworkInfoConnectionType_Mobile  CallRecordsNetworkInfoConnectionType = "mobile"
	CallRecordsNetworkInfoConnectionType_Tunnel  CallRecordsNetworkInfoConnectionType = "tunnel"
	CallRecordsNetworkInfoConnectionType_Unknown CallRecordsNetworkInfoConnectionType = "unknown"
	CallRecordsNetworkInfoConnectionType_Wifi    CallRecordsNetworkInfoConnectionType = "wifi"
	CallRecordsNetworkInfoConnectionType_Wired   CallRecordsNetworkInfoConnectionType = "wired"
)

func PossibleValuesForCallRecordsNetworkInfoConnectionType() []string {
	return []string{
		string(CallRecordsNetworkInfoConnectionType_Mobile),
		string(CallRecordsNetworkInfoConnectionType_Tunnel),
		string(CallRecordsNetworkInfoConnectionType_Unknown),
		string(CallRecordsNetworkInfoConnectionType_Wifi),
		string(CallRecordsNetworkInfoConnectionType_Wired),
	}
}

func (s *CallRecordsNetworkInfoConnectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordsNetworkInfoConnectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordsNetworkInfoConnectionType(input string) (*CallRecordsNetworkInfoConnectionType, error) {
	vals := map[string]CallRecordsNetworkInfoConnectionType{
		"mobile":  CallRecordsNetworkInfoConnectionType_Mobile,
		"tunnel":  CallRecordsNetworkInfoConnectionType_Tunnel,
		"unknown": CallRecordsNetworkInfoConnectionType_Unknown,
		"wifi":    CallRecordsNetworkInfoConnectionType_Wifi,
		"wired":   CallRecordsNetworkInfoConnectionType_Wired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsNetworkInfoConnectionType(input)
	return &out, nil
}
