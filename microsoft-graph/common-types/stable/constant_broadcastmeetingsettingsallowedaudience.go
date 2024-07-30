package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BroadcastMeetingSettingsAllowedAudience string

const (
	BroadcastMeetingSettingsAllowedAudience_Everyone       BroadcastMeetingSettingsAllowedAudience = "everyone"
	BroadcastMeetingSettingsAllowedAudience_Organization   BroadcastMeetingSettingsAllowedAudience = "organization"
	BroadcastMeetingSettingsAllowedAudience_RoleIsAttendee BroadcastMeetingSettingsAllowedAudience = "roleIsAttendee"
)

func PossibleValuesForBroadcastMeetingSettingsAllowedAudience() []string {
	return []string{
		string(BroadcastMeetingSettingsAllowedAudience_Everyone),
		string(BroadcastMeetingSettingsAllowedAudience_Organization),
		string(BroadcastMeetingSettingsAllowedAudience_RoleIsAttendee),
	}
}

func (s *BroadcastMeetingSettingsAllowedAudience) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBroadcastMeetingSettingsAllowedAudience(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBroadcastMeetingSettingsAllowedAudience(input string) (*BroadcastMeetingSettingsAllowedAudience, error) {
	vals := map[string]BroadcastMeetingSettingsAllowedAudience{
		"everyone":       BroadcastMeetingSettingsAllowedAudience_Everyone,
		"organization":   BroadcastMeetingSettingsAllowedAudience_Organization,
		"roleisattendee": BroadcastMeetingSettingsAllowedAudience_RoleIsAttendee,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BroadcastMeetingSettingsAllowedAudience(input)
	return &out, nil
}
