package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessForwardingProfileState string

const (
	NetworkaccessForwardingProfileStatedisabled NetworkaccessForwardingProfileState = "Disabled"
	NetworkaccessForwardingProfileStateenabled  NetworkaccessForwardingProfileState = "Enabled"
)

func PossibleValuesForNetworkaccessForwardingProfileState() []string {
	return []string{
		string(NetworkaccessForwardingProfileStatedisabled),
		string(NetworkaccessForwardingProfileStateenabled),
	}
}

func parseNetworkaccessForwardingProfileState(input string) (*NetworkaccessForwardingProfileState, error) {
	vals := map[string]NetworkaccessForwardingProfileState{
		"disabled": NetworkaccessForwardingProfileStatedisabled,
		"enabled":  NetworkaccessForwardingProfileStateenabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessForwardingProfileState(input)
	return &out, nil
}
