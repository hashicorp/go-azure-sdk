package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsWifiEnterpriseEAPConfigurationAuthenticationType string

const (
	WindowsWifiEnterpriseEAPConfigurationAuthenticationType_Guest         WindowsWifiEnterpriseEAPConfigurationAuthenticationType = "guest"
	WindowsWifiEnterpriseEAPConfigurationAuthenticationType_Machine       WindowsWifiEnterpriseEAPConfigurationAuthenticationType = "machine"
	WindowsWifiEnterpriseEAPConfigurationAuthenticationType_MachineOrUser WindowsWifiEnterpriseEAPConfigurationAuthenticationType = "machineOrUser"
	WindowsWifiEnterpriseEAPConfigurationAuthenticationType_None          WindowsWifiEnterpriseEAPConfigurationAuthenticationType = "none"
	WindowsWifiEnterpriseEAPConfigurationAuthenticationType_User          WindowsWifiEnterpriseEAPConfigurationAuthenticationType = "user"
)

func PossibleValuesForWindowsWifiEnterpriseEAPConfigurationAuthenticationType() []string {
	return []string{
		string(WindowsWifiEnterpriseEAPConfigurationAuthenticationType_Guest),
		string(WindowsWifiEnterpriseEAPConfigurationAuthenticationType_Machine),
		string(WindowsWifiEnterpriseEAPConfigurationAuthenticationType_MachineOrUser),
		string(WindowsWifiEnterpriseEAPConfigurationAuthenticationType_None),
		string(WindowsWifiEnterpriseEAPConfigurationAuthenticationType_User),
	}
}

func (s *WindowsWifiEnterpriseEAPConfigurationAuthenticationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsWifiEnterpriseEAPConfigurationAuthenticationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsWifiEnterpriseEAPConfigurationAuthenticationType(input string) (*WindowsWifiEnterpriseEAPConfigurationAuthenticationType, error) {
	vals := map[string]WindowsWifiEnterpriseEAPConfigurationAuthenticationType{
		"guest":         WindowsWifiEnterpriseEAPConfigurationAuthenticationType_Guest,
		"machine":       WindowsWifiEnterpriseEAPConfigurationAuthenticationType_Machine,
		"machineoruser": WindowsWifiEnterpriseEAPConfigurationAuthenticationType_MachineOrUser,
		"none":          WindowsWifiEnterpriseEAPConfigurationAuthenticationType_None,
		"user":          WindowsWifiEnterpriseEAPConfigurationAuthenticationType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsWifiEnterpriseEAPConfigurationAuthenticationType(input)
	return &out, nil
}
