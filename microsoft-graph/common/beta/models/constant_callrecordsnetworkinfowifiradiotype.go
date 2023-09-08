package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsNetworkInfoWifiRadioType string

const (
	CallRecordsNetworkInfoWifiRadioTypeunknown     CallRecordsNetworkInfoWifiRadioType = "Unknown"
	CallRecordsNetworkInfoWifiRadioTypewifi80211a  CallRecordsNetworkInfoWifiRadioType = "Wifi80211a"
	CallRecordsNetworkInfoWifiRadioTypewifi80211ac CallRecordsNetworkInfoWifiRadioType = "Wifi80211ac"
	CallRecordsNetworkInfoWifiRadioTypewifi80211ax CallRecordsNetworkInfoWifiRadioType = "Wifi80211ax"
	CallRecordsNetworkInfoWifiRadioTypewifi80211b  CallRecordsNetworkInfoWifiRadioType = "Wifi80211b"
	CallRecordsNetworkInfoWifiRadioTypewifi80211g  CallRecordsNetworkInfoWifiRadioType = "Wifi80211g"
	CallRecordsNetworkInfoWifiRadioTypewifi80211n  CallRecordsNetworkInfoWifiRadioType = "Wifi80211n"
)

func PossibleValuesForCallRecordsNetworkInfoWifiRadioType() []string {
	return []string{
		string(CallRecordsNetworkInfoWifiRadioTypeunknown),
		string(CallRecordsNetworkInfoWifiRadioTypewifi80211a),
		string(CallRecordsNetworkInfoWifiRadioTypewifi80211ac),
		string(CallRecordsNetworkInfoWifiRadioTypewifi80211ax),
		string(CallRecordsNetworkInfoWifiRadioTypewifi80211b),
		string(CallRecordsNetworkInfoWifiRadioTypewifi80211g),
		string(CallRecordsNetworkInfoWifiRadioTypewifi80211n),
	}
}

func parseCallRecordsNetworkInfoWifiRadioType(input string) (*CallRecordsNetworkInfoWifiRadioType, error) {
	vals := map[string]CallRecordsNetworkInfoWifiRadioType{
		"unknown":     CallRecordsNetworkInfoWifiRadioTypeunknown,
		"wifi80211a":  CallRecordsNetworkInfoWifiRadioTypewifi80211a,
		"wifi80211ac": CallRecordsNetworkInfoWifiRadioTypewifi80211ac,
		"wifi80211ax": CallRecordsNetworkInfoWifiRadioTypewifi80211ax,
		"wifi80211b":  CallRecordsNetworkInfoWifiRadioTypewifi80211b,
		"wifi80211g":  CallRecordsNetworkInfoWifiRadioTypewifi80211g,
		"wifi80211n":  CallRecordsNetworkInfoWifiRadioTypewifi80211n,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsNetworkInfoWifiRadioType(input)
	return &out, nil
}
