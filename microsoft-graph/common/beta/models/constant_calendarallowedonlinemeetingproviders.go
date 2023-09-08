package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarAllowedOnlineMeetingProviders string

const (
	CalendarAllowedOnlineMeetingProvidersskypeForBusiness CalendarAllowedOnlineMeetingProviders = "SkypeForBusiness"
	CalendarAllowedOnlineMeetingProvidersskypeForConsumer CalendarAllowedOnlineMeetingProviders = "SkypeForConsumer"
	CalendarAllowedOnlineMeetingProvidersteamsForBusiness CalendarAllowedOnlineMeetingProviders = "TeamsForBusiness"
	CalendarAllowedOnlineMeetingProvidersunknown          CalendarAllowedOnlineMeetingProviders = "Unknown"
)

func PossibleValuesForCalendarAllowedOnlineMeetingProviders() []string {
	return []string{
		string(CalendarAllowedOnlineMeetingProvidersskypeForBusiness),
		string(CalendarAllowedOnlineMeetingProvidersskypeForConsumer),
		string(CalendarAllowedOnlineMeetingProvidersteamsForBusiness),
		string(CalendarAllowedOnlineMeetingProvidersunknown),
	}
}

func parseCalendarAllowedOnlineMeetingProviders(input string) (*CalendarAllowedOnlineMeetingProviders, error) {
	vals := map[string]CalendarAllowedOnlineMeetingProviders{
		"skypeforbusiness": CalendarAllowedOnlineMeetingProvidersskypeForBusiness,
		"skypeforconsumer": CalendarAllowedOnlineMeetingProvidersskypeForConsumer,
		"teamsforbusiness": CalendarAllowedOnlineMeetingProvidersteamsForBusiness,
		"unknown":          CalendarAllowedOnlineMeetingProvidersunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CalendarAllowedOnlineMeetingProviders(input)
	return &out, nil
}
