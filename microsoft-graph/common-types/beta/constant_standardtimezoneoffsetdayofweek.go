package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StandardTimeZoneOffsetDayOfWeek string

const (
	StandardTimeZoneOffsetDayOfWeek_Friday    StandardTimeZoneOffsetDayOfWeek = "friday"
	StandardTimeZoneOffsetDayOfWeek_Monday    StandardTimeZoneOffsetDayOfWeek = "monday"
	StandardTimeZoneOffsetDayOfWeek_Saturday  StandardTimeZoneOffsetDayOfWeek = "saturday"
	StandardTimeZoneOffsetDayOfWeek_Sunday    StandardTimeZoneOffsetDayOfWeek = "sunday"
	StandardTimeZoneOffsetDayOfWeek_Thursday  StandardTimeZoneOffsetDayOfWeek = "thursday"
	StandardTimeZoneOffsetDayOfWeek_Tuesday   StandardTimeZoneOffsetDayOfWeek = "tuesday"
	StandardTimeZoneOffsetDayOfWeek_Wednesday StandardTimeZoneOffsetDayOfWeek = "wednesday"
)

func PossibleValuesForStandardTimeZoneOffsetDayOfWeek() []string {
	return []string{
		string(StandardTimeZoneOffsetDayOfWeek_Friday),
		string(StandardTimeZoneOffsetDayOfWeek_Monday),
		string(StandardTimeZoneOffsetDayOfWeek_Saturday),
		string(StandardTimeZoneOffsetDayOfWeek_Sunday),
		string(StandardTimeZoneOffsetDayOfWeek_Thursday),
		string(StandardTimeZoneOffsetDayOfWeek_Tuesday),
		string(StandardTimeZoneOffsetDayOfWeek_Wednesday),
	}
}

func (s *StandardTimeZoneOffsetDayOfWeek) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStandardTimeZoneOffsetDayOfWeek(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStandardTimeZoneOffsetDayOfWeek(input string) (*StandardTimeZoneOffsetDayOfWeek, error) {
	vals := map[string]StandardTimeZoneOffsetDayOfWeek{
		"friday":    StandardTimeZoneOffsetDayOfWeek_Friday,
		"monday":    StandardTimeZoneOffsetDayOfWeek_Monday,
		"saturday":  StandardTimeZoneOffsetDayOfWeek_Saturday,
		"sunday":    StandardTimeZoneOffsetDayOfWeek_Sunday,
		"thursday":  StandardTimeZoneOffsetDayOfWeek_Thursday,
		"tuesday":   StandardTimeZoneOffsetDayOfWeek_Tuesday,
		"wednesday": StandardTimeZoneOffsetDayOfWeek_Wednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StandardTimeZoneOffsetDayOfWeek(input)
	return &out, nil
}
