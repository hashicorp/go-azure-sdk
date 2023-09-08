package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkConnectionConnectionStatus string

const (
	TeamworkConnectionConnectionStatusconnected    TeamworkConnectionConnectionStatus = "Connected"
	TeamworkConnectionConnectionStatusdisconnected TeamworkConnectionConnectionStatus = "Disconnected"
	TeamworkConnectionConnectionStatusunknown      TeamworkConnectionConnectionStatus = "Unknown"
)

func PossibleValuesForTeamworkConnectionConnectionStatus() []string {
	return []string{
		string(TeamworkConnectionConnectionStatusconnected),
		string(TeamworkConnectionConnectionStatusdisconnected),
		string(TeamworkConnectionConnectionStatusunknown),
	}
}

func parseTeamworkConnectionConnectionStatus(input string) (*TeamworkConnectionConnectionStatus, error) {
	vals := map[string]TeamworkConnectionConnectionStatus{
		"connected":    TeamworkConnectionConnectionStatusconnected,
		"disconnected": TeamworkConnectionConnectionStatusdisconnected,
		"unknown":      TeamworkConnectionConnectionStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkConnectionConnectionStatus(input)
	return &out, nil
}
