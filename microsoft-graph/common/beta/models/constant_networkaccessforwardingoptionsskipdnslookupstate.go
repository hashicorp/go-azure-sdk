package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessForwardingOptionsSkipDnsLookupState string

const (
	NetworkaccessForwardingOptionsSkipDnsLookupStatedisabled NetworkaccessForwardingOptionsSkipDnsLookupState = "Disabled"
	NetworkaccessForwardingOptionsSkipDnsLookupStateenabled  NetworkaccessForwardingOptionsSkipDnsLookupState = "Enabled"
)

func PossibleValuesForNetworkaccessForwardingOptionsSkipDnsLookupState() []string {
	return []string{
		string(NetworkaccessForwardingOptionsSkipDnsLookupStatedisabled),
		string(NetworkaccessForwardingOptionsSkipDnsLookupStateenabled),
	}
}

func parseNetworkaccessForwardingOptionsSkipDnsLookupState(input string) (*NetworkaccessForwardingOptionsSkipDnsLookupState, error) {
	vals := map[string]NetworkaccessForwardingOptionsSkipDnsLookupState{
		"disabled": NetworkaccessForwardingOptionsSkipDnsLookupStatedisabled,
		"enabled":  NetworkaccessForwardingOptionsSkipDnsLookupStateenabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessForwardingOptionsSkipDnsLookupState(input)
	return &out, nil
}
