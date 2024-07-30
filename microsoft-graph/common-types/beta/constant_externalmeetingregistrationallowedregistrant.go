package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalMeetingRegistrationAllowedRegistrant string

const (
	ExternalMeetingRegistrationAllowedRegistrant_Everyone     ExternalMeetingRegistrationAllowedRegistrant = "everyone"
	ExternalMeetingRegistrationAllowedRegistrant_Organization ExternalMeetingRegistrationAllowedRegistrant = "organization"
)

func PossibleValuesForExternalMeetingRegistrationAllowedRegistrant() []string {
	return []string{
		string(ExternalMeetingRegistrationAllowedRegistrant_Everyone),
		string(ExternalMeetingRegistrationAllowedRegistrant_Organization),
	}
}

func (s *ExternalMeetingRegistrationAllowedRegistrant) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternalMeetingRegistrationAllowedRegistrant(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternalMeetingRegistrationAllowedRegistrant(input string) (*ExternalMeetingRegistrationAllowedRegistrant, error) {
	vals := map[string]ExternalMeetingRegistrationAllowedRegistrant{
		"everyone":     ExternalMeetingRegistrationAllowedRegistrant_Everyone,
		"organization": ExternalMeetingRegistrationAllowedRegistrant_Organization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalMeetingRegistrationAllowedRegistrant(input)
	return &out, nil
}
