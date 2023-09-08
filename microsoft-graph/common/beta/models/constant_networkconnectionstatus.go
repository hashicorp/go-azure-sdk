package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkConnectionStatus string

const (
	NetworkConnectionStatusattempted NetworkConnectionStatus = "Attempted"
	NetworkConnectionStatusblocked   NetworkConnectionStatus = "Blocked"
	NetworkConnectionStatusfailed    NetworkConnectionStatus = "Failed"
	NetworkConnectionStatussucceeded NetworkConnectionStatus = "Succeeded"
	NetworkConnectionStatusunknown   NetworkConnectionStatus = "Unknown"
)

func PossibleValuesForNetworkConnectionStatus() []string {
	return []string{
		string(NetworkConnectionStatusattempted),
		string(NetworkConnectionStatusblocked),
		string(NetworkConnectionStatusfailed),
		string(NetworkConnectionStatussucceeded),
		string(NetworkConnectionStatusunknown),
	}
}

func parseNetworkConnectionStatus(input string) (*NetworkConnectionStatus, error) {
	vals := map[string]NetworkConnectionStatus{
		"attempted": NetworkConnectionStatusattempted,
		"blocked":   NetworkConnectionStatusblocked,
		"failed":    NetworkConnectionStatusfailed,
		"succeeded": NetworkConnectionStatussucceeded,
		"unknown":   NetworkConnectionStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkConnectionStatus(input)
	return &out, nil
}
