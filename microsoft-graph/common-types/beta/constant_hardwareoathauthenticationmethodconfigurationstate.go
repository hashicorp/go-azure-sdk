package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HardwareOathAuthenticationMethodConfigurationState string

const (
	HardwareOathAuthenticationMethodConfigurationState_Disabled HardwareOathAuthenticationMethodConfigurationState = "disabled"
	HardwareOathAuthenticationMethodConfigurationState_Enabled  HardwareOathAuthenticationMethodConfigurationState = "enabled"
)

func PossibleValuesForHardwareOathAuthenticationMethodConfigurationState() []string {
	return []string{
		string(HardwareOathAuthenticationMethodConfigurationState_Disabled),
		string(HardwareOathAuthenticationMethodConfigurationState_Enabled),
	}
}

func (s *HardwareOathAuthenticationMethodConfigurationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHardwareOathAuthenticationMethodConfigurationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHardwareOathAuthenticationMethodConfigurationState(input string) (*HardwareOathAuthenticationMethodConfigurationState, error) {
	vals := map[string]HardwareOathAuthenticationMethodConfigurationState{
		"disabled": HardwareOathAuthenticationMethodConfigurationState_Disabled,
		"enabled":  HardwareOathAuthenticationMethodConfigurationState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HardwareOathAuthenticationMethodConfigurationState(input)
	return &out, nil
}
