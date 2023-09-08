package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEasEmailProfileConfigurationEasServices string

const (
	IosEasEmailProfileConfigurationEasServicescalendars IosEasEmailProfileConfigurationEasServices = "Calendars"
	IosEasEmailProfileConfigurationEasServicescontacts  IosEasEmailProfileConfigurationEasServices = "Contacts"
	IosEasEmailProfileConfigurationEasServicesemail     IosEasEmailProfileConfigurationEasServices = "Email"
	IosEasEmailProfileConfigurationEasServicesnone      IosEasEmailProfileConfigurationEasServices = "None"
	IosEasEmailProfileConfigurationEasServicesnotes     IosEasEmailProfileConfigurationEasServices = "Notes"
	IosEasEmailProfileConfigurationEasServicesreminders IosEasEmailProfileConfigurationEasServices = "Reminders"
)

func PossibleValuesForIosEasEmailProfileConfigurationEasServices() []string {
	return []string{
		string(IosEasEmailProfileConfigurationEasServicescalendars),
		string(IosEasEmailProfileConfigurationEasServicescontacts),
		string(IosEasEmailProfileConfigurationEasServicesemail),
		string(IosEasEmailProfileConfigurationEasServicesnone),
		string(IosEasEmailProfileConfigurationEasServicesnotes),
		string(IosEasEmailProfileConfigurationEasServicesreminders),
	}
}

func parseIosEasEmailProfileConfigurationEasServices(input string) (*IosEasEmailProfileConfigurationEasServices, error) {
	vals := map[string]IosEasEmailProfileConfigurationEasServices{
		"calendars": IosEasEmailProfileConfigurationEasServicescalendars,
		"contacts":  IosEasEmailProfileConfigurationEasServicescontacts,
		"email":     IosEasEmailProfileConfigurationEasServicesemail,
		"none":      IosEasEmailProfileConfigurationEasServicesnone,
		"notes":     IosEasEmailProfileConfigurationEasServicesnotes,
		"reminders": IosEasEmailProfileConfigurationEasServicesreminders,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosEasEmailProfileConfigurationEasServices(input)
	return &out, nil
}
