package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsNetworkInfoWifiRadioType string

const (
	CallRecordsNetworkInfoWifiRadioType_Unknown     CallRecordsNetworkInfoWifiRadioType = "unknown"
	CallRecordsNetworkInfoWifiRadioType_Wifi80211a  CallRecordsNetworkInfoWifiRadioType = "wifi80211a"
	CallRecordsNetworkInfoWifiRadioType_Wifi80211ac CallRecordsNetworkInfoWifiRadioType = "wifi80211ac"
	CallRecordsNetworkInfoWifiRadioType_Wifi80211ax CallRecordsNetworkInfoWifiRadioType = "wifi80211ax"
	CallRecordsNetworkInfoWifiRadioType_Wifi80211b  CallRecordsNetworkInfoWifiRadioType = "wifi80211b"
	CallRecordsNetworkInfoWifiRadioType_Wifi80211g  CallRecordsNetworkInfoWifiRadioType = "wifi80211g"
	CallRecordsNetworkInfoWifiRadioType_Wifi80211n  CallRecordsNetworkInfoWifiRadioType = "wifi80211n"
)

func PossibleValuesForCallRecordsNetworkInfoWifiRadioType() []string {
	return []string{
		string(CallRecordsNetworkInfoWifiRadioType_Unknown),
		string(CallRecordsNetworkInfoWifiRadioType_Wifi80211a),
		string(CallRecordsNetworkInfoWifiRadioType_Wifi80211ac),
		string(CallRecordsNetworkInfoWifiRadioType_Wifi80211ax),
		string(CallRecordsNetworkInfoWifiRadioType_Wifi80211b),
		string(CallRecordsNetworkInfoWifiRadioType_Wifi80211g),
		string(CallRecordsNetworkInfoWifiRadioType_Wifi80211n),
	}
}

func (s *CallRecordsNetworkInfoWifiRadioType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordsNetworkInfoWifiRadioType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordsNetworkInfoWifiRadioType(input string) (*CallRecordsNetworkInfoWifiRadioType, error) {
	vals := map[string]CallRecordsNetworkInfoWifiRadioType{
		"unknown":     CallRecordsNetworkInfoWifiRadioType_Unknown,
		"wifi80211a":  CallRecordsNetworkInfoWifiRadioType_Wifi80211a,
		"wifi80211ac": CallRecordsNetworkInfoWifiRadioType_Wifi80211ac,
		"wifi80211ax": CallRecordsNetworkInfoWifiRadioType_Wifi80211ax,
		"wifi80211b":  CallRecordsNetworkInfoWifiRadioType_Wifi80211b,
		"wifi80211g":  CallRecordsNetworkInfoWifiRadioType_Wifi80211g,
		"wifi80211n":  CallRecordsNetworkInfoWifiRadioType_Wifi80211n,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsNetworkInfoWifiRadioType(input)
	return &out, nil
}
