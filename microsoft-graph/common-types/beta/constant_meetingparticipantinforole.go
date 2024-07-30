package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingParticipantInfoRole string

const (
	MeetingParticipantInfoRole_Attendee    MeetingParticipantInfoRole = "attendee"
	MeetingParticipantInfoRole_Coorganizer MeetingParticipantInfoRole = "coorganizer"
	MeetingParticipantInfoRole_Presenter   MeetingParticipantInfoRole = "presenter"
	MeetingParticipantInfoRole_Producer    MeetingParticipantInfoRole = "producer"
)

func PossibleValuesForMeetingParticipantInfoRole() []string {
	return []string{
		string(MeetingParticipantInfoRole_Attendee),
		string(MeetingParticipantInfoRole_Coorganizer),
		string(MeetingParticipantInfoRole_Presenter),
		string(MeetingParticipantInfoRole_Producer),
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
		"attendee":    MeetingParticipantInfoRole_Attendee,
		"coorganizer": MeetingParticipantInfoRole_Coorganizer,
		"presenter":   MeetingParticipantInfoRole_Presenter,
		"producer":    MeetingParticipantInfoRole_Producer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MeetingParticipantInfoRole(input)
	return &out, nil
}
