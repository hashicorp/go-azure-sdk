package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsNetworkInfoConnectionType string

const (
	CallRecordsNetworkInfoConnectionTypemobile  CallRecordsNetworkInfoConnectionType = "Mobile"
	CallRecordsNetworkInfoConnectionTypetunnel  CallRecordsNetworkInfoConnectionType = "Tunnel"
	CallRecordsNetworkInfoConnectionTypeunknown CallRecordsNetworkInfoConnectionType = "Unknown"
	CallRecordsNetworkInfoConnectionTypewifi    CallRecordsNetworkInfoConnectionType = "Wifi"
	CallRecordsNetworkInfoConnectionTypewired   CallRecordsNetworkInfoConnectionType = "Wired"
)

func PossibleValuesForCallRecordsNetworkInfoConnectionType() []string {
	return []string{
		string(CallRecordsNetworkInfoConnectionTypemobile),
		string(CallRecordsNetworkInfoConnectionTypetunnel),
		string(CallRecordsNetworkInfoConnectionTypeunknown),
		string(CallRecordsNetworkInfoConnectionTypewifi),
		string(CallRecordsNetworkInfoConnectionTypewired),
	}
}

func parseCallRecordsNetworkInfoConnectionType(input string) (*CallRecordsNetworkInfoConnectionType, error) {
	vals := map[string]CallRecordsNetworkInfoConnectionType{
		"mobile":  CallRecordsNetworkInfoConnectionTypemobile,
		"tunnel":  CallRecordsNetworkInfoConnectionTypetunnel,
		"unknown": CallRecordsNetworkInfoConnectionTypeunknown,
		"wifi":    CallRecordsNetworkInfoConnectionTypewifi,
		"wired":   CallRecordsNetworkInfoConnectionTypewired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsNetworkInfoConnectionType(input)
	return &out, nil
}
