package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndUserNotificationSource string

const (
	EndUserNotificationSourceglobal  EndUserNotificationSource = "Global"
	EndUserNotificationSourcetenant  EndUserNotificationSource = "Tenant"
	EndUserNotificationSourceunknown EndUserNotificationSource = "Unknown"
)

func PossibleValuesForEndUserNotificationSource() []string {
	return []string{
		string(EndUserNotificationSourceglobal),
		string(EndUserNotificationSourcetenant),
		string(EndUserNotificationSourceunknown),
	}
}

func parseEndUserNotificationSource(input string) (*EndUserNotificationSource, error) {
	vals := map[string]EndUserNotificationSource{
		"global":  EndUserNotificationSourceglobal,
		"tenant":  EndUserNotificationSourcetenant,
		"unknown": EndUserNotificationSourceunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EndUserNotificationSource(input)
	return &out, nil
}
