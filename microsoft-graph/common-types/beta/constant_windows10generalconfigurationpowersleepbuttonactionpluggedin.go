package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn string

const (
	Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn_Hibernate     Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn = "hibernate"
	Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn_NoAction      Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn = "noAction"
	Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn_NotConfigured Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn = "notConfigured"
	Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn_Shutdown      Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn = "shutdown"
	Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn_Sleep         Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn = "sleep"
)

func PossibleValuesForWindows10GeneralConfigurationPowerSleepButtonActionPluggedIn() []string {
	return []string{
		string(Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn_Hibernate),
		string(Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn_NoAction),
		string(Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn_NotConfigured),
		string(Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn_Shutdown),
		string(Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn_Sleep),
	}
}

func (s *Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationPowerSleepButtonActionPluggedIn(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationPowerSleepButtonActionPluggedIn(input string) (*Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn, error) {
	vals := map[string]Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn{
		"hibernate":     Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn_Hibernate,
		"noaction":      Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn_NoAction,
		"notconfigured": Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn_NotConfigured,
		"shutdown":      Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn_Shutdown,
		"sleep":         Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn_Sleep,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn(input)
	return &out, nil
}
