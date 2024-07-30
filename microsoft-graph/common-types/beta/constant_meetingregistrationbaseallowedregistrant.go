package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingRegistrationBaseAllowedRegistrant string

const (
	MeetingRegistrationBaseAllowedRegistrant_Everyone     MeetingRegistrationBaseAllowedRegistrant = "everyone"
	MeetingRegistrationBaseAllowedRegistrant_Organization MeetingRegistrationBaseAllowedRegistrant = "organization"
)

func PossibleValuesForMeetingRegistrationBaseAllowedRegistrant() []string {
	return []string{
		string(MeetingRegistrationBaseAllowedRegistrant_Everyone),
		string(MeetingRegistrationBaseAllowedRegistrant_Organization),
	}
}

func (s *MeetingRegistrationBaseAllowedRegistrant) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMeetingRegistrationBaseAllowedRegistrant(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMeetingRegistrationBaseAllowedRegistrant(input string) (*MeetingRegistrationBaseAllowedRegistrant, error) {
	vals := map[string]MeetingRegistrationBaseAllowedRegistrant{
		"everyone":     MeetingRegistrationBaseAllowedRegistrant_Everyone,
		"organization": MeetingRegistrationBaseAllowedRegistrant_Organization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MeetingRegistrationBaseAllowedRegistrant(input)
	return &out, nil
}
