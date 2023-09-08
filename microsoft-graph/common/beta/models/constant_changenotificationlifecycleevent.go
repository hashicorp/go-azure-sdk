package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangeNotificationLifecycleEvent string

const (
	ChangeNotificationLifecycleEventmissed                  ChangeNotificationLifecycleEvent = "Missed"
	ChangeNotificationLifecycleEventreauthorizationRequired ChangeNotificationLifecycleEvent = "ReauthorizationRequired"
	ChangeNotificationLifecycleEventsubscriptionRemoved     ChangeNotificationLifecycleEvent = "SubscriptionRemoved"
)

func PossibleValuesForChangeNotificationLifecycleEvent() []string {
	return []string{
		string(ChangeNotificationLifecycleEventmissed),
		string(ChangeNotificationLifecycleEventreauthorizationRequired),
		string(ChangeNotificationLifecycleEventsubscriptionRemoved),
	}
}

func parseChangeNotificationLifecycleEvent(input string) (*ChangeNotificationLifecycleEvent, error) {
	vals := map[string]ChangeNotificationLifecycleEvent{
		"missed":                  ChangeNotificationLifecycleEventmissed,
		"reauthorizationrequired": ChangeNotificationLifecycleEventreauthorizationRequired,
		"subscriptionremoved":     ChangeNotificationLifecycleEventsubscriptionRemoved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChangeNotificationLifecycleEvent(input)
	return &out, nil
}
