package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarAllowedOnlineMeetingProviders string

const (
	CalendarAllowedOnlineMeetingProviders_SkypeForBusiness CalendarAllowedOnlineMeetingProviders = "skypeForBusiness"
	CalendarAllowedOnlineMeetingProviders_SkypeForConsumer CalendarAllowedOnlineMeetingProviders = "skypeForConsumer"
	CalendarAllowedOnlineMeetingProviders_TeamsForBusiness CalendarAllowedOnlineMeetingProviders = "teamsForBusiness"
	CalendarAllowedOnlineMeetingProviders_Unknown          CalendarAllowedOnlineMeetingProviders = "unknown"
)

func PossibleValuesForCalendarAllowedOnlineMeetingProviders() []string {
	return []string{
		string(CalendarAllowedOnlineMeetingProviders_SkypeForBusiness),
		string(CalendarAllowedOnlineMeetingProviders_SkypeForConsumer),
		string(CalendarAllowedOnlineMeetingProviders_TeamsForBusiness),
		string(CalendarAllowedOnlineMeetingProviders_Unknown),
	}
}

func (s *CalendarAllowedOnlineMeetingProviders) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCalendarAllowedOnlineMeetingProviders(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCalendarAllowedOnlineMeetingProviders(input string) (*CalendarAllowedOnlineMeetingProviders, error) {
	vals := map[string]CalendarAllowedOnlineMeetingProviders{
		"skypeforbusiness": CalendarAllowedOnlineMeetingProviders_SkypeForBusiness,
		"skypeforconsumer": CalendarAllowedOnlineMeetingProviders_SkypeForConsumer,
		"teamsforbusiness": CalendarAllowedOnlineMeetingProviders_TeamsForBusiness,
		"unknown":          CalendarAllowedOnlineMeetingProviders_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CalendarAllowedOnlineMeetingProviders(input)
	return &out, nil
}
