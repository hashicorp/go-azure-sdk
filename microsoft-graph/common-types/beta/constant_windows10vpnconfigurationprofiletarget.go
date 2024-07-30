package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10VpnConfigurationProfileTarget string

const (
	Windows10VpnConfigurationProfileTarget_AutoPilotDevice Windows10VpnConfigurationProfileTarget = "autoPilotDevice"
	Windows10VpnConfigurationProfileTarget_Device          Windows10VpnConfigurationProfileTarget = "device"
	Windows10VpnConfigurationProfileTarget_User            Windows10VpnConfigurationProfileTarget = "user"
)

func PossibleValuesForWindows10VpnConfigurationProfileTarget() []string {
	return []string{
		string(Windows10VpnConfigurationProfileTarget_AutoPilotDevice),
		string(Windows10VpnConfigurationProfileTarget_Device),
		string(Windows10VpnConfigurationProfileTarget_User),
	}
}

func (s *Windows10VpnConfigurationProfileTarget) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10VpnConfigurationProfileTarget(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10VpnConfigurationProfileTarget(input string) (*Windows10VpnConfigurationProfileTarget, error) {
	vals := map[string]Windows10VpnConfigurationProfileTarget{
		"autopilotdevice": Windows10VpnConfigurationProfileTarget_AutoPilotDevice,
		"device":          Windows10VpnConfigurationProfileTarget_Device,
		"user":            Windows10VpnConfigurationProfileTarget_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10VpnConfigurationProfileTarget(input)
	return &out, nil
}
