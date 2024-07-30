package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsWiredNetworkConfigurationAuthenticationType string

const (
	WindowsWiredNetworkConfigurationAuthenticationType_Guest         WindowsWiredNetworkConfigurationAuthenticationType = "guest"
	WindowsWiredNetworkConfigurationAuthenticationType_Machine       WindowsWiredNetworkConfigurationAuthenticationType = "machine"
	WindowsWiredNetworkConfigurationAuthenticationType_MachineOrUser WindowsWiredNetworkConfigurationAuthenticationType = "machineOrUser"
	WindowsWiredNetworkConfigurationAuthenticationType_None          WindowsWiredNetworkConfigurationAuthenticationType = "none"
	WindowsWiredNetworkConfigurationAuthenticationType_User          WindowsWiredNetworkConfigurationAuthenticationType = "user"
)

func PossibleValuesForWindowsWiredNetworkConfigurationAuthenticationType() []string {
	return []string{
		string(WindowsWiredNetworkConfigurationAuthenticationType_Guest),
		string(WindowsWiredNetworkConfigurationAuthenticationType_Machine),
		string(WindowsWiredNetworkConfigurationAuthenticationType_MachineOrUser),
		string(WindowsWiredNetworkConfigurationAuthenticationType_None),
		string(WindowsWiredNetworkConfigurationAuthenticationType_User),
	}
}

func (s *WindowsWiredNetworkConfigurationAuthenticationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsWiredNetworkConfigurationAuthenticationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsWiredNetworkConfigurationAuthenticationType(input string) (*WindowsWiredNetworkConfigurationAuthenticationType, error) {
	vals := map[string]WindowsWiredNetworkConfigurationAuthenticationType{
		"guest":         WindowsWiredNetworkConfigurationAuthenticationType_Guest,
		"machine":       WindowsWiredNetworkConfigurationAuthenticationType_Machine,
		"machineoruser": WindowsWiredNetworkConfigurationAuthenticationType_MachineOrUser,
		"none":          WindowsWiredNetworkConfigurationAuthenticationType_None,
		"user":          WindowsWiredNetworkConfigurationAuthenticationType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsWiredNetworkConfigurationAuthenticationType(input)
	return &out, nil
}
