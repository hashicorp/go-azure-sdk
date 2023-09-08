package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndUserNotificationStatus string

const (
	EndUserNotificationStatusarchive EndUserNotificationStatus = "Archive"
	EndUserNotificationStatusdelete  EndUserNotificationStatus = "Delete"
	EndUserNotificationStatusdraft   EndUserNotificationStatus = "Draft"
	EndUserNotificationStatusready   EndUserNotificationStatus = "Ready"
	EndUserNotificationStatusunknown EndUserNotificationStatus = "Unknown"
)

func PossibleValuesForEndUserNotificationStatus() []string {
	return []string{
		string(EndUserNotificationStatusarchive),
		string(EndUserNotificationStatusdelete),
		string(EndUserNotificationStatusdraft),
		string(EndUserNotificationStatusready),
		string(EndUserNotificationStatusunknown),
	}
}

func parseEndUserNotificationStatus(input string) (*EndUserNotificationStatus, error) {
	vals := map[string]EndUserNotificationStatus{
		"archive": EndUserNotificationStatusarchive,
		"delete":  EndUserNotificationStatusdelete,
		"draft":   EndUserNotificationStatusdraft,
		"ready":   EndUserNotificationStatusready,
		"unknown": EndUserNotificationStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EndUserNotificationStatus(input)
	return &out, nil
}
