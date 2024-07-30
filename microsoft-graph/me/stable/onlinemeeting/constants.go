package onlinemeeting

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingParticipantInfoRole string

const (
	MeetingParticipantInfoRoleAttendee    MeetingParticipantInfoRole = "attendee"
	MeetingParticipantInfoRoleCoorganizer MeetingParticipantInfoRole = "coorganizer"
	MeetingParticipantInfoRolePresenter   MeetingParticipantInfoRole = "presenter"
	MeetingParticipantInfoRoleProducer    MeetingParticipantInfoRole = "producer"
)

func PossibleValuesForMeetingParticipantInfoRole() []string {
	return []string{
		string(MeetingParticipantInfoRoleAttendee),
		string(MeetingParticipantInfoRoleCoorganizer),
		string(MeetingParticipantInfoRolePresenter),
		string(MeetingParticipantInfoRoleProducer),
	}
}

func (s *MeetingParticipantInfoRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMeetingParticipantInfoRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMeetingParticipantInfoRole(input string) (*MeetingParticipantInfoRole, error) {
	vals := map[string]MeetingParticipantInfoRole{
		"attendee":    MeetingParticipantInfoRoleAttendee,
		"coorganizer": MeetingParticipantInfoRoleCoorganizer,
		"presenter":   MeetingParticipantInfoRolePresenter,
		"producer":    MeetingParticipantInfoRoleProducer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MeetingParticipantInfoRole(input)
	return &out, nil
}
