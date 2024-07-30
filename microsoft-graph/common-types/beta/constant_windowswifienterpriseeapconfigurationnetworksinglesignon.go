package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsWifiEnterpriseEAPConfigurationNetworkSingleSignOn string

const (
	WindowsWifiEnterpriseEAPConfigurationNetworkSingleSignOn_Disabled  WindowsWifiEnterpriseEAPConfigurationNetworkSingleSignOn = "disabled"
	WindowsWifiEnterpriseEAPConfigurationNetworkSingleSignOn_Postlogon WindowsWifiEnterpriseEAPConfigurationNetworkSingleSignOn = "postlogon"
	WindowsWifiEnterpriseEAPConfigurationNetworkSingleSignOn_Prelogon  WindowsWifiEnterpriseEAPConfigurationNetworkSingleSignOn = "prelogon"
)

func PossibleValuesForWindowsWifiEnterpriseEAPConfigurationNetworkSingleSignOn() []string {
	return []string{
		string(WindowsWifiEnterpriseEAPConfigurationNetworkSingleSignOn_Disabled),
		string(WindowsWifiEnterpriseEAPConfigurationNetworkSingleSignOn_Postlogon),
		string(WindowsWifiEnterpriseEAPConfigurationNetworkSingleSignOn_Prelogon),
	}
}

func (s *WindowsWifiEnterpriseEAPConfigurationNetworkSingleSignOn) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsWifiEnterpriseEAPConfigurationNetworkSingleSignOn(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsWifiEnterpriseEAPConfigurationNetworkSingleSignOn(input string) (*WindowsWifiEnterpriseEAPConfigurationNetworkSingleSignOn, error) {
	vals := map[string]WindowsWifiEnterpriseEAPConfigurationNetworkSingleSignOn{
		"disabled":  WindowsWifiEnterpriseEAPConfigurationNetworkSingleSignOn_Disabled,
		"postlogon": WindowsWifiEnterpriseEAPConfigurationNetworkSingleSignOn_Postlogon,
		"prelogon":  WindowsWifiEnterpriseEAPConfigurationNetworkSingleSignOn_Prelogon,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsWifiEnterpriseEAPConfigurationNetworkSingleSignOn(input)
	return &out, nil
}
