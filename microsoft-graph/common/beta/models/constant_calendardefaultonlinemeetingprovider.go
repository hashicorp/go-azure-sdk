package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarDefaultOnlineMeetingProvider string

const (
	CalendarDefaultOnlineMeetingProviderskypeForBusiness CalendarDefaultOnlineMeetingProvider = "SkypeForBusiness"
	CalendarDefaultOnlineMeetingProviderskypeForConsumer CalendarDefaultOnlineMeetingProvider = "SkypeForConsumer"
	CalendarDefaultOnlineMeetingProviderteamsForBusiness CalendarDefaultOnlineMeetingProvider = "TeamsForBusiness"
	CalendarDefaultOnlineMeetingProviderunknown          CalendarDefaultOnlineMeetingProvider = "Unknown"
)

func PossibleValuesForCalendarDefaultOnlineMeetingProvider() []string {
	return []string{
		string(CalendarDefaultOnlineMeetingProviderskypeForBusiness),
		string(CalendarDefaultOnlineMeetingProviderskypeForConsumer),
		string(CalendarDefaultOnlineMeetingProviderteamsForBusiness),
		string(CalendarDefaultOnlineMeetingProviderunknown),
	}
}

func parseCalendarDefaultOnlineMeetingProvider(input string) (*CalendarDefaultOnlineMeetingProvider, error) {
	vals := map[string]CalendarDefaultOnlineMeetingProvider{
		"skypeforbusiness": CalendarDefaultOnlineMeetingProviderskypeForBusiness,
		"skypeforconsumer": CalendarDefaultOnlineMeetingProviderskypeForConsumer,
		"teamsforbusiness": CalendarDefaultOnlineMeetingProviderteamsForBusiness,
		"unknown":          CalendarDefaultOnlineMeetingProviderunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CalendarDefaultOnlineMeetingProvider(input)
	return &out, nil
}
