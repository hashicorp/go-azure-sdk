package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessBranchSiteConnectivityState string

const (
	NetworkaccessBranchSiteConnectivityStateconnected NetworkaccessBranchSiteConnectivityState = "Connected"
	NetworkaccessBranchSiteConnectivityStateerror     NetworkaccessBranchSiteConnectivityState = "Error"
	NetworkaccessBranchSiteConnectivityStateinactive  NetworkaccessBranchSiteConnectivityState = "Inactive"
	NetworkaccessBranchSiteConnectivityStatepending   NetworkaccessBranchSiteConnectivityState = "Pending"
)

func PossibleValuesForNetworkaccessBranchSiteConnectivityState() []string {
	return []string{
		string(NetworkaccessBranchSiteConnectivityStateconnected),
		string(NetworkaccessBranchSiteConnectivityStateerror),
		string(NetworkaccessBranchSiteConnectivityStateinactive),
		string(NetworkaccessBranchSiteConnectivityStatepending),
	}
}

func parseNetworkaccessBranchSiteConnectivityState(input string) (*NetworkaccessBranchSiteConnectivityState, error) {
	vals := map[string]NetworkaccessBranchSiteConnectivityState{
		"connected": NetworkaccessBranchSiteConnectivityStateconnected,
		"error":     NetworkaccessBranchSiteConnectivityStateerror,
		"inactive":  NetworkaccessBranchSiteConnectivityStateinactive,
		"pending":   NetworkaccessBranchSiteConnectivityStatepending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessBranchSiteConnectivityState(input)
	return &out, nil
}
