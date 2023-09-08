package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BroadcastMeetingSettingsAllowedAudience string

const (
	BroadcastMeetingSettingsAllowedAudienceeveryone       BroadcastMeetingSettingsAllowedAudience = "Everyone"
	BroadcastMeetingSettingsAllowedAudienceorganization   BroadcastMeetingSettingsAllowedAudience = "Organization"
	BroadcastMeetingSettingsAllowedAudienceroleIsAttendee BroadcastMeetingSettingsAllowedAudience = "RoleIsAttendee"
)

func PossibleValuesForBroadcastMeetingSettingsAllowedAudience() []string {
	return []string{
		string(BroadcastMeetingSettingsAllowedAudienceeveryone),
		string(BroadcastMeetingSettingsAllowedAudienceorganization),
		string(BroadcastMeetingSettingsAllowedAudienceroleIsAttendee),
	}
}

func parseBroadcastMeetingSettingsAllowedAudience(input string) (*BroadcastMeetingSettingsAllowedAudience, error) {
	vals := map[string]BroadcastMeetingSettingsAllowedAudience{
		"everyone":       BroadcastMeetingSettingsAllowedAudienceeveryone,
		"organization":   BroadcastMeetingSettingsAllowedAudienceorganization,
		"roleisattendee": BroadcastMeetingSettingsAllowedAudienceroleIsAttendee,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BroadcastMeetingSettingsAllowedAudience(input)
	return &out, nil
}
