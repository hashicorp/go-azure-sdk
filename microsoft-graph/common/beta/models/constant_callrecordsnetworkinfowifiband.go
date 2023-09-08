package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsNetworkInfoWifiBand string

const (
	CallRecordsNetworkInfoWifiBandfrequency24GHz CallRecordsNetworkInfoWifiBand = "Frequency24GHz"
	CallRecordsNetworkInfoWifiBandfrequency50GHz CallRecordsNetworkInfoWifiBand = "Frequency50GHz"
	CallRecordsNetworkInfoWifiBandfrequency60GHz CallRecordsNetworkInfoWifiBand = "Frequency60GHz"
	CallRecordsNetworkInfoWifiBandunknown        CallRecordsNetworkInfoWifiBand = "Unknown"
)

func PossibleValuesForCallRecordsNetworkInfoWifiBand() []string {
	return []string{
		string(CallRecordsNetworkInfoWifiBandfrequency24GHz),
		string(CallRecordsNetworkInfoWifiBandfrequency50GHz),
		string(CallRecordsNetworkInfoWifiBandfrequency60GHz),
		string(CallRecordsNetworkInfoWifiBandunknown),
	}
}

func parseCallRecordsNetworkInfoWifiBand(input string) (*CallRecordsNetworkInfoWifiBand, error) {
	vals := map[string]CallRecordsNetworkInfoWifiBand{
		"frequency24ghz": CallRecordsNetworkInfoWifiBandfrequency24GHz,
		"frequency50ghz": CallRecordsNetworkInfoWifiBandfrequency50GHz,
		"frequency60ghz": CallRecordsNetworkInfoWifiBandfrequency60GHz,
		"unknown":        CallRecordsNetworkInfoWifiBandunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsNetworkInfoWifiBand(input)
	return &out, nil
}
