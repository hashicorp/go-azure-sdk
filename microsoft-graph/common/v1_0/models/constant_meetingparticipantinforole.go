package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingParticipantInfoRole string

const (
	MeetingParticipantInfoRoleattendee    MeetingParticipantInfoRole = "Attendee"
	MeetingParticipantInfoRolecoorganizer MeetingParticipantInfoRole = "Coorganizer"
	MeetingParticipantInfoRolepresenter   MeetingParticipantInfoRole = "Presenter"
	MeetingParticipantInfoRoleproducer    MeetingParticipantInfoRole = "Producer"
)

func PossibleValuesForMeetingParticipantInfoRole() []string {
	return []string{
		string(MeetingParticipantInfoRoleattendee),
		string(MeetingParticipantInfoRolecoorganizer),
		string(MeetingParticipantInfoRolepresenter),
		string(MeetingParticipantInfoRoleproducer),
	}
}

func parseMeetingParticipantInfoRole(input string) (*MeetingParticipantInfoRole, error) {
	vals := map[string]MeetingParticipantInfoRole{
		"attendee":    MeetingParticipantInfoRoleattendee,
		"coorganizer": MeetingParticipantInfoRolecoorganizer,
		"presenter":   MeetingParticipantInfoRolepresenter,
		"producer":    MeetingParticipantInfoRoleproducer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MeetingParticipantInfoRole(input)
	return &out, nil
}
