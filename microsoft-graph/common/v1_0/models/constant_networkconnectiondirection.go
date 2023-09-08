package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkConnectionDirection string

const (
	NetworkConnectionDirectioninbound  NetworkConnectionDirection = "Inbound"
	NetworkConnectionDirectionoutbound NetworkConnectionDirection = "Outbound"
	NetworkConnectionDirectionunknown  NetworkConnectionDirection = "Unknown"
)

func PossibleValuesForNetworkConnectionDirection() []string {
	return []string{
		string(NetworkConnectionDirectioninbound),
		string(NetworkConnectionDirectionoutbound),
		string(NetworkConnectionDirectionunknown),
	}
}

func parseNetworkConnectionDirection(input string) (*NetworkConnectionDirection, error) {
	vals := map[string]NetworkConnectionDirection{
		"inbound":  NetworkConnectionDirectioninbound,
		"outbound": NetworkConnectionDirectionoutbound,
		"unknown":  NetworkConnectionDirectionunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkConnectionDirection(input)
	return &out, nil
}
