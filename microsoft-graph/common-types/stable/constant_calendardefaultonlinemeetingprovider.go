package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarDefaultOnlineMeetingProvider string

const (
	CalendarDefaultOnlineMeetingProvider_SkypeForBusiness CalendarDefaultOnlineMeetingProvider = "skypeForBusiness"
	CalendarDefaultOnlineMeetingProvider_SkypeForConsumer CalendarDefaultOnlineMeetingProvider = "skypeForConsumer"
	CalendarDefaultOnlineMeetingProvider_TeamsForBusiness CalendarDefaultOnlineMeetingProvider = "teamsForBusiness"
	CalendarDefaultOnlineMeetingProvider_Unknown          CalendarDefaultOnlineMeetingProvider = "unknown"
)

func PossibleValuesForCalendarDefaultOnlineMeetingProvider() []string {
	return []string{
		string(CalendarDefaultOnlineMeetingProvider_SkypeForBusiness),
		string(CalendarDefaultOnlineMeetingProvider_SkypeForConsumer),
		string(CalendarDefaultOnlineMeetingProvider_TeamsForBusiness),
		string(CalendarDefaultOnlineMeetingProvider_Unknown),
	}
}

func (s *CalendarDefaultOnlineMeetingProvider) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCalendarDefaultOnlineMeetingProvider(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCalendarDefaultOnlineMeetingProvider(input string) (*CalendarDefaultOnlineMeetingProvider, error) {
	vals := map[string]CalendarDefaultOnlineMeetingProvider{
		"skypeforbusiness": CalendarDefaultOnlineMeetingProvider_SkypeForBusiness,
		"skypeforconsumer": CalendarDefaultOnlineMeetingProvider_SkypeForConsumer,
		"teamsforbusiness": CalendarDefaultOnlineMeetingProvider_TeamsForBusiness,
		"unknown":          CalendarDefaultOnlineMeetingProvider_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CalendarDefaultOnlineMeetingProvider(input)
	return &out, nil
}
