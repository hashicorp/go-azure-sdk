package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingRegistrationAllowedRegistrant string

const (
	MeetingRegistrationAllowedRegistrant_Everyone     MeetingRegistrationAllowedRegistrant = "everyone"
	MeetingRegistrationAllowedRegistrant_Organization MeetingRegistrationAllowedRegistrant = "organization"
)

func PossibleValuesForMeetingRegistrationAllowedRegistrant() []string {
	return []string{
		string(MeetingRegistrationAllowedRegistrant_Everyone),
		string(MeetingRegistrationAllowedRegistrant_Organization),
	}
}

func (s *MeetingRegistrationAllowedRegistrant) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMeetingRegistrationAllowedRegistrant(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMeetingRegistrationAllowedRegistrant(input string) (*MeetingRegistrationAllowedRegistrant, error) {
	vals := map[string]MeetingRegistrationAllowedRegistrant{
		"everyone":     MeetingRegistrationAllowedRegistrant_Everyone,
		"organization": MeetingRegistrationAllowedRegistrant_Organization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MeetingRegistrationAllowedRegistrant(input)
	return &out, nil
}
