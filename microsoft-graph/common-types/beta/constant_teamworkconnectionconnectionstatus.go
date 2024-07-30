package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkConnectionConnectionStatus string

const (
	TeamworkConnectionConnectionStatus_Connected    TeamworkConnectionConnectionStatus = "connected"
	TeamworkConnectionConnectionStatus_Disconnected TeamworkConnectionConnectionStatus = "disconnected"
	TeamworkConnectionConnectionStatus_Unknown      TeamworkConnectionConnectionStatus = "unknown"
)

func PossibleValuesForTeamworkConnectionConnectionStatus() []string {
	return []string{
		string(TeamworkConnectionConnectionStatus_Connected),
		string(TeamworkConnectionConnectionStatus_Disconnected),
		string(TeamworkConnectionConnectionStatus_Unknown),
	}
}

func (s *TeamworkConnectionConnectionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamworkConnectionConnectionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamworkConnectionConnectionStatus(input string) (*TeamworkConnectionConnectionStatus, error) {
	vals := map[string]TeamworkConnectionConnectionStatus{
		"connected":    TeamworkConnectionConnectionStatus_Connected,
		"disconnected": TeamworkConnectionConnectionStatus_Disconnected,
		"unknown":      TeamworkConnectionConnectionStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkConnectionConnectionStatus(input)
	return &out, nil
}
