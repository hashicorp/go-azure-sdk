package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotificationPriority string

const (
	NotificationPriorityHigh NotificationPriority = "High"
	NotificationPriorityLow  NotificationPriority = "Low"
	NotificationPriorityNone NotificationPriority = "None"
)

func PossibleValuesForNotificationPriority() []string {
	return []string{
		string(NotificationPriorityHigh),
		string(NotificationPriorityLow),
		string(NotificationPriorityNone),
	}
}

func parseNotificationPriority(input string) (*NotificationPriority, error) {
	vals := map[string]NotificationPriority{
		"high": NotificationPriorityHigh,
		"low":  NotificationPriorityLow,
		"none": NotificationPriorityNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NotificationPriority(input)
	return &out, nil
}
