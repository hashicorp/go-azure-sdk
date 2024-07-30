package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationPowerLidCloseActionOnBattery string

const (
	Windows10GeneralConfigurationPowerLidCloseActionOnBattery_Hibernate     Windows10GeneralConfigurationPowerLidCloseActionOnBattery = "hibernate"
	Windows10GeneralConfigurationPowerLidCloseActionOnBattery_NoAction      Windows10GeneralConfigurationPowerLidCloseActionOnBattery = "noAction"
	Windows10GeneralConfigurationPowerLidCloseActionOnBattery_NotConfigured Windows10GeneralConfigurationPowerLidCloseActionOnBattery = "notConfigured"
	Windows10GeneralConfigurationPowerLidCloseActionOnBattery_Shutdown      Windows10GeneralConfigurationPowerLidCloseActionOnBattery = "shutdown"
	Windows10GeneralConfigurationPowerLidCloseActionOnBattery_Sleep         Windows10GeneralConfigurationPowerLidCloseActionOnBattery = "sleep"
)

func PossibleValuesForWindows10GeneralConfigurationPowerLidCloseActionOnBattery() []string {
	return []string{
		string(Windows10GeneralConfigurationPowerLidCloseActionOnBattery_Hibernate),
		string(Windows10GeneralConfigurationPowerLidCloseActionOnBattery_NoAction),
		string(Windows10GeneralConfigurationPowerLidCloseActionOnBattery_NotConfigured),
		string(Windows10GeneralConfigurationPowerLidCloseActionOnBattery_Shutdown),
		string(Windows10GeneralConfigurationPowerLidCloseActionOnBattery_Sleep),
	}
}

func (s *Windows10GeneralConfigurationPowerLidCloseActionOnBattery) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationPowerLidCloseActionOnBattery(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationPowerLidCloseActionOnBattery(input string) (*Windows10GeneralConfigurationPowerLidCloseActionOnBattery, error) {
	vals := map[string]Windows10GeneralConfigurationPowerLidCloseActionOnBattery{
		"hibernate":     Windows10GeneralConfigurationPowerLidCloseActionOnBattery_Hibernate,
		"noaction":      Windows10GeneralConfigurationPowerLidCloseActionOnBattery_NoAction,
		"notconfigured": Windows10GeneralConfigurationPowerLidCloseActionOnBattery_NotConfigured,
		"shutdown":      Windows10GeneralConfigurationPowerLidCloseActionOnBattery_Shutdown,
		"sleep":         Windows10GeneralConfigurationPowerLidCloseActionOnBattery_Sleep,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationPowerLidCloseActionOnBattery(input)
	return &out, nil
}
