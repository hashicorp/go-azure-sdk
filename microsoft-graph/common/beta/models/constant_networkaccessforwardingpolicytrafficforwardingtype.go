package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessForwardingPolicyTrafficForwardingType string

const (
	NetworkaccessForwardingPolicyTrafficForwardingTypeinternet NetworkaccessForwardingPolicyTrafficForwardingType = "Internet"
	NetworkaccessForwardingPolicyTrafficForwardingTypem365     NetworkaccessForwardingPolicyTrafficForwardingType = "M365"
	NetworkaccessForwardingPolicyTrafficForwardingTypeprivate  NetworkaccessForwardingPolicyTrafficForwardingType = "Private"
)

func PossibleValuesForNetworkaccessForwardingPolicyTrafficForwardingType() []string {
	return []string{
		string(NetworkaccessForwardingPolicyTrafficForwardingTypeinternet),
		string(NetworkaccessForwardingPolicyTrafficForwardingTypem365),
		string(NetworkaccessForwardingPolicyTrafficForwardingTypeprivate),
	}
}

func parseNetworkaccessForwardingPolicyTrafficForwardingType(input string) (*NetworkaccessForwardingPolicyTrafficForwardingType, error) {
	vals := map[string]NetworkaccessForwardingPolicyTrafficForwardingType{
		"internet": NetworkaccessForwardingPolicyTrafficForwardingTypeinternet,
		"m365":     NetworkaccessForwardingPolicyTrafficForwardingTypem365,
		"private":  NetworkaccessForwardingPolicyTrafficForwardingTypeprivate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessForwardingPolicyTrafficForwardingType(input)
	return &out, nil
}
