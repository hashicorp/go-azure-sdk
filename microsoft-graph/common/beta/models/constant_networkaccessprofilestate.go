package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessProfileState string

const (
	NetworkaccessProfileStatedisabled NetworkaccessProfileState = "Disabled"
	NetworkaccessProfileStateenabled  NetworkaccessProfileState = "Enabled"
)

func PossibleValuesForNetworkaccessProfileState() []string {
	return []string{
		string(NetworkaccessProfileStatedisabled),
		string(NetworkaccessProfileStateenabled),
	}
}

func parseNetworkaccessProfileState(input string) (*NetworkaccessProfileState, error) {
	vals := map[string]NetworkaccessProfileState{
		"disabled": NetworkaccessProfileStatedisabled,
		"enabled":  NetworkaccessProfileStateenabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessProfileState(input)
	return &out, nil
}
