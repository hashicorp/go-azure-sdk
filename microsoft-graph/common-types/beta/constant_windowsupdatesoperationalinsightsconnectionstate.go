package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesOperationalInsightsConnectionState string

const (
	WindowsUpdatesOperationalInsightsConnectionState_Connected     WindowsUpdatesOperationalInsightsConnectionState = "connected"
	WindowsUpdatesOperationalInsightsConnectionState_NotAuthorized WindowsUpdatesOperationalInsightsConnectionState = "notAuthorized"
	WindowsUpdatesOperationalInsightsConnectionState_NotFound      WindowsUpdatesOperationalInsightsConnectionState = "notFound"
)

func PossibleValuesForWindowsUpdatesOperationalInsightsConnectionState() []string {
	return []string{
		string(WindowsUpdatesOperationalInsightsConnectionState_Connected),
		string(WindowsUpdatesOperationalInsightsConnectionState_NotAuthorized),
		string(WindowsUpdatesOperationalInsightsConnectionState_NotFound),
	}
}

func (s *WindowsUpdatesOperationalInsightsConnectionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdatesOperationalInsightsConnectionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdatesOperationalInsightsConnectionState(input string) (*WindowsUpdatesOperationalInsightsConnectionState, error) {
	vals := map[string]WindowsUpdatesOperationalInsightsConnectionState{
		"connected":     WindowsUpdatesOperationalInsightsConnectionState_Connected,
		"notauthorized": WindowsUpdatesOperationalInsightsConnectionState_NotAuthorized,
		"notfound":      WindowsUpdatesOperationalInsightsConnectionState_NotFound,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesOperationalInsightsConnectionState(input)
	return &out, nil
}
