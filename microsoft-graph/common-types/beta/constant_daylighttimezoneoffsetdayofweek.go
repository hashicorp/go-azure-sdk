package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DaylightTimeZoneOffsetDayOfWeek string

const (
	DaylightTimeZoneOffsetDayOfWeek_Friday    DaylightTimeZoneOffsetDayOfWeek = "friday"
	DaylightTimeZoneOffsetDayOfWeek_Monday    DaylightTimeZoneOffsetDayOfWeek = "monday"
	DaylightTimeZoneOffsetDayOfWeek_Saturday  DaylightTimeZoneOffsetDayOfWeek = "saturday"
	DaylightTimeZoneOffsetDayOfWeek_Sunday    DaylightTimeZoneOffsetDayOfWeek = "sunday"
	DaylightTimeZoneOffsetDayOfWeek_Thursday  DaylightTimeZoneOffsetDayOfWeek = "thursday"
	DaylightTimeZoneOffsetDayOfWeek_Tuesday   DaylightTimeZoneOffsetDayOfWeek = "tuesday"
	DaylightTimeZoneOffsetDayOfWeek_Wednesday DaylightTimeZoneOffsetDayOfWeek = "wednesday"
)

func PossibleValuesForDaylightTimeZoneOffsetDayOfWeek() []string {
	return []string{
		string(DaylightTimeZoneOffsetDayOfWeek_Friday),
		string(DaylightTimeZoneOffsetDayOfWeek_Monday),
		string(DaylightTimeZoneOffsetDayOfWeek_Saturday),
		string(DaylightTimeZoneOffsetDayOfWeek_Sunday),
		string(DaylightTimeZoneOffsetDayOfWeek_Thursday),
		string(DaylightTimeZoneOffsetDayOfWeek_Tuesday),
		string(DaylightTimeZoneOffsetDayOfWeek_Wednesday),
	}
}

func (s *DaylightTimeZoneOffsetDayOfWeek) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDaylightTimeZoneOffsetDayOfWeek(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDaylightTimeZoneOffsetDayOfWeek(input string) (*DaylightTimeZoneOffsetDayOfWeek, error) {
	vals := map[string]DaylightTimeZoneOffsetDayOfWeek{
		"friday":    DaylightTimeZoneOffsetDayOfWeek_Friday,
		"monday":    DaylightTimeZoneOffsetDayOfWeek_Monday,
		"saturday":  DaylightTimeZoneOffsetDayOfWeek_Saturday,
		"sunday":    DaylightTimeZoneOffsetDayOfWeek_Sunday,
		"thursday":  DaylightTimeZoneOffsetDayOfWeek_Thursday,
		"tuesday":   DaylightTimeZoneOffsetDayOfWeek_Tuesday,
		"wednesday": DaylightTimeZoneOffsetDayOfWeek_Wednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DaylightTimeZoneOffsetDayOfWeek(input)
	return &out, nil
}
