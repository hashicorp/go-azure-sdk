package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessForwardingPolicyLinkState string

const (
	NetworkaccessForwardingPolicyLinkStatedisabled NetworkaccessForwardingPolicyLinkState = "Disabled"
	NetworkaccessForwardingPolicyLinkStateenabled  NetworkaccessForwardingPolicyLinkState = "Enabled"
)

func PossibleValuesForNetworkaccessForwardingPolicyLinkState() []string {
	return []string{
		string(NetworkaccessForwardingPolicyLinkStatedisabled),
		string(NetworkaccessForwardingPolicyLinkStateenabled),
	}
}

func parseNetworkaccessForwardingPolicyLinkState(input string) (*NetworkaccessForwardingPolicyLinkState, error) {
	vals := map[string]NetworkaccessForwardingPolicyLinkState{
		"disabled": NetworkaccessForwardingPolicyLinkStatedisabled,
		"enabled":  NetworkaccessForwardingPolicyLinkStateenabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessForwardingPolicyLinkState(input)
	return &out, nil
}
