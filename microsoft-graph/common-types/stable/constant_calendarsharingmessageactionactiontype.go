package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarSharingMessageActionActionType string

const (
	CalendarSharingMessageActionActionType_Accept CalendarSharingMessageActionActionType = "accept"
)

func PossibleValuesForCalendarSharingMessageActionActionType() []string {
	return []string{
		string(CalendarSharingMessageActionActionType_Accept),
	}
}

func (s *CalendarSharingMessageActionActionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCalendarSharingMessageActionActionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCalendarSharingMessageActionActionType(input string) (*CalendarSharingMessageActionActionType, error) {
	vals := map[string]CalendarSharingMessageActionActionType{
		"accept": CalendarSharingMessageActionActionType_Accept,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CalendarSharingMessageActionActionType(input)
	return &out, nil
}
