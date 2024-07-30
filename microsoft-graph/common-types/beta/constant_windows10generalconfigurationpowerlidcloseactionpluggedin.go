package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationPowerLidCloseActionPluggedIn string

const (
	Windows10GeneralConfigurationPowerLidCloseActionPluggedIn_Hibernate     Windows10GeneralConfigurationPowerLidCloseActionPluggedIn = "hibernate"
	Windows10GeneralConfigurationPowerLidCloseActionPluggedIn_NoAction      Windows10GeneralConfigurationPowerLidCloseActionPluggedIn = "noAction"
	Windows10GeneralConfigurationPowerLidCloseActionPluggedIn_NotConfigured Windows10GeneralConfigurationPowerLidCloseActionPluggedIn = "notConfigured"
	Windows10GeneralConfigurationPowerLidCloseActionPluggedIn_Shutdown      Windows10GeneralConfigurationPowerLidCloseActionPluggedIn = "shutdown"
	Windows10GeneralConfigurationPowerLidCloseActionPluggedIn_Sleep         Windows10GeneralConfigurationPowerLidCloseActionPluggedIn = "sleep"
)

func PossibleValuesForWindows10GeneralConfigurationPowerLidCloseActionPluggedIn() []string {
	return []string{
		string(Windows10GeneralConfigurationPowerLidCloseActionPluggedIn_Hibernate),
		string(Windows10GeneralConfigurationPowerLidCloseActionPluggedIn_NoAction),
		string(Windows10GeneralConfigurationPowerLidCloseActionPluggedIn_NotConfigured),
		string(Windows10GeneralConfigurationPowerLidCloseActionPluggedIn_Shutdown),
		string(Windows10GeneralConfigurationPowerLidCloseActionPluggedIn_Sleep),
	}
}

func (s *Windows10GeneralConfigurationPowerLidCloseActionPluggedIn) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationPowerLidCloseActionPluggedIn(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationPowerLidCloseActionPluggedIn(input string) (*Windows10GeneralConfigurationPowerLidCloseActionPluggedIn, error) {
	vals := map[string]Windows10GeneralConfigurationPowerLidCloseActionPluggedIn{
		"hibernate":     Windows10GeneralConfigurationPowerLidCloseActionPluggedIn_Hibernate,
		"noaction":      Windows10GeneralConfigurationPowerLidCloseActionPluggedIn_NoAction,
		"notconfigured": Windows10GeneralConfigurationPowerLidCloseActionPluggedIn_NotConfigured,
		"shutdown":      Windows10GeneralConfigurationPowerLidCloseActionPluggedIn_Shutdown,
		"sleep":         Windows10GeneralConfigurationPowerLidCloseActionPluggedIn_Sleep,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationPowerLidCloseActionPluggedIn(input)
	return &out, nil
}
