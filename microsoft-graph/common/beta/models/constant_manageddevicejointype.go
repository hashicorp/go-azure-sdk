package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceJoinType string

const (
	ManagedDeviceJoinTypeazureADJoined       ManagedDeviceJoinType = "AzureADJoined"
	ManagedDeviceJoinTypeazureADRegistered   ManagedDeviceJoinType = "AzureADRegistered"
	ManagedDeviceJoinTypehybridAzureADJoined ManagedDeviceJoinType = "HybridAzureADJoined"
	ManagedDeviceJoinTypeunknown             ManagedDeviceJoinType = "Unknown"
)

func PossibleValuesForManagedDeviceJoinType() []string {
	return []string{
		string(ManagedDeviceJoinTypeazureADJoined),
		string(ManagedDeviceJoinTypeazureADRegistered),
		string(ManagedDeviceJoinTypehybridAzureADJoined),
		string(ManagedDeviceJoinTypeunknown),
	}
}

func parseManagedDeviceJoinType(input string) (*ManagedDeviceJoinType, error) {
	vals := map[string]ManagedDeviceJoinType{
		"azureadjoined":       ManagedDeviceJoinTypeazureADJoined,
		"azureadregistered":   ManagedDeviceJoinTypeazureADRegistered,
		"hybridazureadjoined": ManagedDeviceJoinTypehybridAzureADJoined,
		"unknown":             ManagedDeviceJoinTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceJoinType(input)
	return &out, nil
}
