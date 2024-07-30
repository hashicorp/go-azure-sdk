package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangeNotificationLifecycleEvent string

const (
	ChangeNotificationLifecycleEvent_Missed                  ChangeNotificationLifecycleEvent = "missed"
	ChangeNotificationLifecycleEvent_ReauthorizationRequired ChangeNotificationLifecycleEvent = "reauthorizationRequired"
	ChangeNotificationLifecycleEvent_SubscriptionRemoved     ChangeNotificationLifecycleEvent = "subscriptionRemoved"
)

func PossibleValuesForChangeNotificationLifecycleEvent() []string {
	return []string{
		string(ChangeNotificationLifecycleEvent_Missed),
		string(ChangeNotificationLifecycleEvent_ReauthorizationRequired),
		string(ChangeNotificationLifecycleEvent_SubscriptionRemoved),
	}
}

func (s *ChangeNotificationLifecycleEvent) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChangeNotificationLifecycleEvent(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChangeNotificationLifecycleEvent(input string) (*ChangeNotificationLifecycleEvent, error) {
	vals := map[string]ChangeNotificationLifecycleEvent{
		"missed":                  ChangeNotificationLifecycleEvent_Missed,
		"reauthorizationrequired": ChangeNotificationLifecycleEvent_ReauthorizationRequired,
		"subscriptionremoved":     ChangeNotificationLifecycleEvent_SubscriptionRemoved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChangeNotificationLifecycleEvent(input)
	return &out, nil
}
