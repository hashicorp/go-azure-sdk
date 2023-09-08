package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedDeviceJoinType string

const (
	WindowsManagedDeviceJoinTypeazureADJoined       WindowsManagedDeviceJoinType = "AzureADJoined"
	WindowsManagedDeviceJoinTypeazureADRegistered   WindowsManagedDeviceJoinType = "AzureADRegistered"
	WindowsManagedDeviceJoinTypehybridAzureADJoined WindowsManagedDeviceJoinType = "HybridAzureADJoined"
	WindowsManagedDeviceJoinTypeunknown             WindowsManagedDeviceJoinType = "Unknown"
)

func PossibleValuesForWindowsManagedDeviceJoinType() []string {
	return []string{
		string(WindowsManagedDeviceJoinTypeazureADJoined),
		string(WindowsManagedDeviceJoinTypeazureADRegistered),
		string(WindowsManagedDeviceJoinTypehybridAzureADJoined),
		string(WindowsManagedDeviceJoinTypeunknown),
	}
}

func parseWindowsManagedDeviceJoinType(input string) (*WindowsManagedDeviceJoinType, error) {
	vals := map[string]WindowsManagedDeviceJoinType{
		"azureadjoined":       WindowsManagedDeviceJoinTypeazureADJoined,
		"azureadregistered":   WindowsManagedDeviceJoinTypeazureADRegistered,
		"hybridazureadjoined": WindowsManagedDeviceJoinTypehybridAzureADJoined,
		"unknown":             WindowsManagedDeviceJoinTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedDeviceJoinType(input)
	return &out, nil
}
