package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationPowerButtonActionPluggedIn string

const (
	Windows10GeneralConfigurationPowerButtonActionPluggedIn_Hibernate     Windows10GeneralConfigurationPowerButtonActionPluggedIn = "hibernate"
	Windows10GeneralConfigurationPowerButtonActionPluggedIn_NoAction      Windows10GeneralConfigurationPowerButtonActionPluggedIn = "noAction"
	Windows10GeneralConfigurationPowerButtonActionPluggedIn_NotConfigured Windows10GeneralConfigurationPowerButtonActionPluggedIn = "notConfigured"
	Windows10GeneralConfigurationPowerButtonActionPluggedIn_Shutdown      Windows10GeneralConfigurationPowerButtonActionPluggedIn = "shutdown"
	Windows10GeneralConfigurationPowerButtonActionPluggedIn_Sleep         Windows10GeneralConfigurationPowerButtonActionPluggedIn = "sleep"
)

func PossibleValuesForWindows10GeneralConfigurationPowerButtonActionPluggedIn() []string {
	return []string{
		string(Windows10GeneralConfigurationPowerButtonActionPluggedIn_Hibernate),
		string(Windows10GeneralConfigurationPowerButtonActionPluggedIn_NoAction),
		string(Windows10GeneralConfigurationPowerButtonActionPluggedIn_NotConfigured),
		string(Windows10GeneralConfigurationPowerButtonActionPluggedIn_Shutdown),
		string(Windows10GeneralConfigurationPowerButtonActionPluggedIn_Sleep),
	}
}

func (s *Windows10GeneralConfigurationPowerButtonActionPluggedIn) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationPowerButtonActionPluggedIn(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationPowerButtonActionPluggedIn(input string) (*Windows10GeneralConfigurationPowerButtonActionPluggedIn, error) {
	vals := map[string]Windows10GeneralConfigurationPowerButtonActionPluggedIn{
		"hibernate":     Windows10GeneralConfigurationPowerButtonActionPluggedIn_Hibernate,
		"noaction":      Windows10GeneralConfigurationPowerButtonActionPluggedIn_NoAction,
		"notconfigured": Windows10GeneralConfigurationPowerButtonActionPluggedIn_NotConfigured,
		"shutdown":      Windows10GeneralConfigurationPowerButtonActionPluggedIn_Shutdown,
		"sleep":         Windows10GeneralConfigurationPowerButtonActionPluggedIn_Sleep,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationPowerButtonActionPluggedIn(input)
	return &out, nil
}
