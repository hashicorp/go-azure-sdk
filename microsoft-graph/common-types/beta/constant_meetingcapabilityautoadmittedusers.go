package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingCapabilityAutoAdmittedUsers string

const (
	MeetingCapabilityAutoAdmittedUsers_Everyone          MeetingCapabilityAutoAdmittedUsers = "everyone"
	MeetingCapabilityAutoAdmittedUsers_EveryoneInCompany MeetingCapabilityAutoAdmittedUsers = "everyoneInCompany"
)

func PossibleValuesForMeetingCapabilityAutoAdmittedUsers() []string {
	return []string{
		string(MeetingCapabilityAutoAdmittedUsers_Everyone),
		string(MeetingCapabilityAutoAdmittedUsers_EveryoneInCompany),
	}
}

func (s *MeetingCapabilityAutoAdmittedUsers) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMeetingCapabilityAutoAdmittedUsers(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMeetingCapabilityAutoAdmittedUsers(input string) (*MeetingCapabilityAutoAdmittedUsers, error) {
	vals := map[string]MeetingCapabilityAutoAdmittedUsers{
		"everyone":          MeetingCapabilityAutoAdmittedUsers_Everyone,
		"everyoneincompany": MeetingCapabilityAutoAdmittedUsers_EveryoneInCompany,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MeetingCapabilityAutoAdmittedUsers(input)
	return &out, nil
}
