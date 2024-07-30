package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsWifiConfigurationProxySetting string

const (
	WindowsWifiConfigurationProxySetting_Automatic WindowsWifiConfigurationProxySetting = "automatic"
	WindowsWifiConfigurationProxySetting_Manual    WindowsWifiConfigurationProxySetting = "manual"
	WindowsWifiConfigurationProxySetting_None      WindowsWifiConfigurationProxySetting = "none"
)

func PossibleValuesForWindowsWifiConfigurationProxySetting() []string {
	return []string{
		string(WindowsWifiConfigurationProxySetting_Automatic),
		string(WindowsWifiConfigurationProxySetting_Manual),
		string(WindowsWifiConfigurationProxySetting_None),
	}
}

func (s *WindowsWifiConfigurationProxySetting) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsWifiConfigurationProxySetting(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsWifiConfigurationProxySetting(input string) (*WindowsWifiConfigurationProxySetting, error) {
	vals := map[string]WindowsWifiConfigurationProxySetting{
		"automatic": WindowsWifiConfigurationProxySetting_Automatic,
		"manual":    WindowsWifiConfigurationProxySetting_Manual,
		"none":      WindowsWifiConfigurationProxySetting_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsWifiConfigurationProxySetting(input)
	return &out, nil
}
