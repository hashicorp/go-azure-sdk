package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventOnlineMeetingProvider string

const (
	EventOnlineMeetingProvider_SkypeForBusiness EventOnlineMeetingProvider = "skypeForBusiness"
	EventOnlineMeetingProvider_SkypeForConsumer EventOnlineMeetingProvider = "skypeForConsumer"
	EventOnlineMeetingProvider_TeamsForBusiness EventOnlineMeetingProvider = "teamsForBusiness"
	EventOnlineMeetingProvider_Unknown          EventOnlineMeetingProvider = "unknown"
)

func PossibleValuesForEventOnlineMeetingProvider() []string {
	return []string{
		string(EventOnlineMeetingProvider_SkypeForBusiness),
		string(EventOnlineMeetingProvider_SkypeForConsumer),
		string(EventOnlineMeetingProvider_TeamsForBusiness),
		string(EventOnlineMeetingProvider_Unknown),
	}
}

func (s *EventOnlineMeetingProvider) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventOnlineMeetingProvider(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventOnlineMeetingProvider(input string) (*EventOnlineMeetingProvider, error) {
	vals := map[string]EventOnlineMeetingProvider{
		"skypeforbusiness": EventOnlineMeetingProvider_SkypeForBusiness,
		"skypeforconsumer": EventOnlineMeetingProvider_SkypeForConsumer,
		"teamsforbusiness": EventOnlineMeetingProvider_TeamsForBusiness,
		"unknown":          EventOnlineMeetingProvider_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventOnlineMeetingProvider(input)
	return &out, nil
}
