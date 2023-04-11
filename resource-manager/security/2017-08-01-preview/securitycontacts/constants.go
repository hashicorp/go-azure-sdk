package securitycontacts

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertNotifications string

const (
	AlertNotificationsOff AlertNotifications = "Off"
	AlertNotificationsOn  AlertNotifications = "On"
)

func PossibleValuesForAlertNotifications() []string {
	return []string{
		string(AlertNotificationsOff),
		string(AlertNotificationsOn),
	}
}

func parseAlertNotifications(input string) (*AlertNotifications, error) {
	vals := map[string]AlertNotifications{
		"off": AlertNotificationsOff,
		"on":  AlertNotificationsOn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertNotifications(input)
	return &out, nil
}

type AlertsToAdmins string

const (
	AlertsToAdminsOff AlertsToAdmins = "Off"
	AlertsToAdminsOn  AlertsToAdmins = "On"
)

func PossibleValuesForAlertsToAdmins() []string {
	return []string{
		string(AlertsToAdminsOff),
		string(AlertsToAdminsOn),
	}
}

func parseAlertsToAdmins(input string) (*AlertsToAdmins, error) {
	vals := map[string]AlertsToAdmins{
		"off": AlertsToAdminsOff,
		"on":  AlertsToAdminsOn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertsToAdmins(input)
	return &out, nil
}
