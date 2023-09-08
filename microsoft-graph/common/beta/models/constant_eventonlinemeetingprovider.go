package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventOnlineMeetingProvider string

const (
	EventOnlineMeetingProviderskypeForBusiness EventOnlineMeetingProvider = "SkypeForBusiness"
	EventOnlineMeetingProviderskypeForConsumer EventOnlineMeetingProvider = "SkypeForConsumer"
	EventOnlineMeetingProviderteamsForBusiness EventOnlineMeetingProvider = "TeamsForBusiness"
	EventOnlineMeetingProviderunknown          EventOnlineMeetingProvider = "Unknown"
)

func PossibleValuesForEventOnlineMeetingProvider() []string {
	return []string{
		string(EventOnlineMeetingProviderskypeForBusiness),
		string(EventOnlineMeetingProviderskypeForConsumer),
		string(EventOnlineMeetingProviderteamsForBusiness),
		string(EventOnlineMeetingProviderunknown),
	}
}

func parseEventOnlineMeetingProvider(input string) (*EventOnlineMeetingProvider, error) {
	vals := map[string]EventOnlineMeetingProvider{
		"skypeforbusiness": EventOnlineMeetingProviderskypeForBusiness,
		"skypeforconsumer": EventOnlineMeetingProviderskypeForConsumer,
		"teamsforbusiness": EventOnlineMeetingProviderteamsForBusiness,
		"unknown":          EventOnlineMeetingProviderunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventOnlineMeetingProvider(input)
	return &out, nil
}
