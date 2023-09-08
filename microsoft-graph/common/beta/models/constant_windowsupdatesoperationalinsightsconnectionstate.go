package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesOperationalInsightsConnectionState string

const (
	WindowsUpdatesOperationalInsightsConnectionStateconnected     WindowsUpdatesOperationalInsightsConnectionState = "Connected"
	WindowsUpdatesOperationalInsightsConnectionStatenotAuthorized WindowsUpdatesOperationalInsightsConnectionState = "NotAuthorized"
	WindowsUpdatesOperationalInsightsConnectionStatenotFound      WindowsUpdatesOperationalInsightsConnectionState = "NotFound"
)

func PossibleValuesForWindowsUpdatesOperationalInsightsConnectionState() []string {
	return []string{
		string(WindowsUpdatesOperationalInsightsConnectionStateconnected),
		string(WindowsUpdatesOperationalInsightsConnectionStatenotAuthorized),
		string(WindowsUpdatesOperationalInsightsConnectionStatenotFound),
	}
}

func parseWindowsUpdatesOperationalInsightsConnectionState(input string) (*WindowsUpdatesOperationalInsightsConnectionState, error) {
	vals := map[string]WindowsUpdatesOperationalInsightsConnectionState{
		"connected":     WindowsUpdatesOperationalInsightsConnectionStateconnected,
		"notauthorized": WindowsUpdatesOperationalInsightsConnectionStatenotAuthorized,
		"notfound":      WindowsUpdatesOperationalInsightsConnectionStatenotFound,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesOperationalInsightsConnectionState(input)
	return &out, nil
}
