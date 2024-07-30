package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryKeyStateHive string

const (
	RegistryKeyStateHive_CurrentConfig        RegistryKeyStateHive = "currentConfig"
	RegistryKeyStateHive_CurrentUser          RegistryKeyStateHive = "currentUser"
	RegistryKeyStateHive_LocalMachineSam      RegistryKeyStateHive = "localMachineSam"
	RegistryKeyStateHive_LocalMachineSecurity RegistryKeyStateHive = "localMachineSecurity"
	RegistryKeyStateHive_LocalMachineSoftware RegistryKeyStateHive = "localMachineSoftware"
	RegistryKeyStateHive_LocalMachineSystem   RegistryKeyStateHive = "localMachineSystem"
	RegistryKeyStateHive_Unknown              RegistryKeyStateHive = "unknown"
	RegistryKeyStateHive_UsersDefault         RegistryKeyStateHive = "usersDefault"
)

func PossibleValuesForRegistryKeyStateHive() []string {
	return []string{
		string(RegistryKeyStateHive_CurrentConfig),
		string(RegistryKeyStateHive_CurrentUser),
		string(RegistryKeyStateHive_LocalMachineSam),
		string(RegistryKeyStateHive_LocalMachineSecurity),
		string(RegistryKeyStateHive_LocalMachineSoftware),
		string(RegistryKeyStateHive_LocalMachineSystem),
		string(RegistryKeyStateHive_Unknown),
		string(RegistryKeyStateHive_UsersDefault),
	}
}

func (s *RegistryKeyStateHive) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRegistryKeyStateHive(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRegistryKeyStateHive(input string) (*RegistryKeyStateHive, error) {
	vals := map[string]RegistryKeyStateHive{
		"currentconfig":        RegistryKeyStateHive_CurrentConfig,
		"currentuser":          RegistryKeyStateHive_CurrentUser,
		"localmachinesam":      RegistryKeyStateHive_LocalMachineSam,
		"localmachinesecurity": RegistryKeyStateHive_LocalMachineSecurity,
		"localmachinesoftware": RegistryKeyStateHive_LocalMachineSoftware,
		"localmachinesystem":   RegistryKeyStateHive_LocalMachineSystem,
		"unknown":              RegistryKeyStateHive_Unknown,
		"usersdefault":         RegistryKeyStateHive_UsersDefault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RegistryKeyStateHive(input)
	return &out, nil
}
