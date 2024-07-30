package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEasEmailProfileConfigurationEasServices string

const (
	IosEasEmailProfileConfigurationEasServices_Calendars IosEasEmailProfileConfigurationEasServices = "calendars"
	IosEasEmailProfileConfigurationEasServices_Contacts  IosEasEmailProfileConfigurationEasServices = "contacts"
	IosEasEmailProfileConfigurationEasServices_Email     IosEasEmailProfileConfigurationEasServices = "email"
	IosEasEmailProfileConfigurationEasServices_None      IosEasEmailProfileConfigurationEasServices = "none"
	IosEasEmailProfileConfigurationEasServices_Notes     IosEasEmailProfileConfigurationEasServices = "notes"
	IosEasEmailProfileConfigurationEasServices_Reminders IosEasEmailProfileConfigurationEasServices = "reminders"
)

func PossibleValuesForIosEasEmailProfileConfigurationEasServices() []string {
	return []string{
		string(IosEasEmailProfileConfigurationEasServices_Calendars),
		string(IosEasEmailProfileConfigurationEasServices_Contacts),
		string(IosEasEmailProfileConfigurationEasServices_Email),
		string(IosEasEmailProfileConfigurationEasServices_None),
		string(IosEasEmailProfileConfigurationEasServices_Notes),
		string(IosEasEmailProfileConfigurationEasServices_Reminders),
	}
}

func (s *IosEasEmailProfileConfigurationEasServices) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosEasEmailProfileConfigurationEasServices(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosEasEmailProfileConfigurationEasServices(input string) (*IosEasEmailProfileConfigurationEasServices, error) {
	vals := map[string]IosEasEmailProfileConfigurationEasServices{
		"calendars": IosEasEmailProfileConfigurationEasServices_Calendars,
		"contacts":  IosEasEmailProfileConfigurationEasServices_Contacts,
		"email":     IosEasEmailProfileConfigurationEasServices_Email,
		"none":      IosEasEmailProfileConfigurationEasServices_None,
		"notes":     IosEasEmailProfileConfigurationEasServices_Notes,
		"reminders": IosEasEmailProfileConfigurationEasServices_Reminders,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosEasEmailProfileConfigurationEasServices(input)
	return &out, nil
}
