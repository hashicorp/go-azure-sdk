package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessPolicyLinkState string

const (
	NetworkaccessPolicyLinkStatedisabled NetworkaccessPolicyLinkState = "Disabled"
	NetworkaccessPolicyLinkStateenabled  NetworkaccessPolicyLinkState = "Enabled"
)

func PossibleValuesForNetworkaccessPolicyLinkState() []string {
	return []string{
		string(NetworkaccessPolicyLinkStatedisabled),
		string(NetworkaccessPolicyLinkStateenabled),
	}
}

func parseNetworkaccessPolicyLinkState(input string) (*NetworkaccessPolicyLinkState, error) {
	vals := map[string]NetworkaccessPolicyLinkState{
		"disabled": NetworkaccessPolicyLinkStatedisabled,
		"enabled":  NetworkaccessPolicyLinkStateenabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessPolicyLinkState(input)
	return &out, nil
}
