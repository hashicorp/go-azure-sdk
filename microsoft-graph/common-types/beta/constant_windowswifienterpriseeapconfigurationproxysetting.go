package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsWifiEnterpriseEAPConfigurationProxySetting string

const (
	WindowsWifiEnterpriseEAPConfigurationProxySetting_Automatic WindowsWifiEnterpriseEAPConfigurationProxySetting = "automatic"
	WindowsWifiEnterpriseEAPConfigurationProxySetting_Manual    WindowsWifiEnterpriseEAPConfigurationProxySetting = "manual"
	WindowsWifiEnterpriseEAPConfigurationProxySetting_None      WindowsWifiEnterpriseEAPConfigurationProxySetting = "none"
)

func PossibleValuesForWindowsWifiEnterpriseEAPConfigurationProxySetting() []string {
	return []string{
		string(WindowsWifiEnterpriseEAPConfigurationProxySetting_Automatic),
		string(WindowsWifiEnterpriseEAPConfigurationProxySetting_Manual),
		string(WindowsWifiEnterpriseEAPConfigurationProxySetting_None),
	}
}

func (s *WindowsWifiEnterpriseEAPConfigurationProxySetting) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsWifiEnterpriseEAPConfigurationProxySetting(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsWifiEnterpriseEAPConfigurationProxySetting(input string) (*WindowsWifiEnterpriseEAPConfigurationProxySetting, error) {
	vals := map[string]WindowsWifiEnterpriseEAPConfigurationProxySetting{
		"automatic": WindowsWifiEnterpriseEAPConfigurationProxySetting_Automatic,
		"manual":    WindowsWifiEnterpriseEAPConfigurationProxySetting_Manual,
		"none":      WindowsWifiEnterpriseEAPConfigurationProxySetting_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsWifiEnterpriseEAPConfigurationProxySetting(input)
	return &out, nil
}
