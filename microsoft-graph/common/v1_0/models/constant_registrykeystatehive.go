package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryKeyStateHive string

const (
	RegistryKeyStateHivecurrentConfig        RegistryKeyStateHive = "CurrentConfig"
	RegistryKeyStateHivecurrentUser          RegistryKeyStateHive = "CurrentUser"
	RegistryKeyStateHivelocalMachineSam      RegistryKeyStateHive = "LocalMachineSam"
	RegistryKeyStateHivelocalMachineSecurity RegistryKeyStateHive = "LocalMachineSecurity"
	RegistryKeyStateHivelocalMachineSoftware RegistryKeyStateHive = "LocalMachineSoftware"
	RegistryKeyStateHivelocalMachineSystem   RegistryKeyStateHive = "LocalMachineSystem"
	RegistryKeyStateHiveunknown              RegistryKeyStateHive = "Unknown"
	RegistryKeyStateHiveusersDefault         RegistryKeyStateHive = "UsersDefault"
)

func PossibleValuesForRegistryKeyStateHive() []string {
	return []string{
		string(RegistryKeyStateHivecurrentConfig),
		string(RegistryKeyStateHivecurrentUser),
		string(RegistryKeyStateHivelocalMachineSam),
		string(RegistryKeyStateHivelocalMachineSecurity),
		string(RegistryKeyStateHivelocalMachineSoftware),
		string(RegistryKeyStateHivelocalMachineSystem),
		string(RegistryKeyStateHiveunknown),
		string(RegistryKeyStateHiveusersDefault),
	}
}

func parseRegistryKeyStateHive(input string) (*RegistryKeyStateHive, error) {
	vals := map[string]RegistryKeyStateHive{
		"currentconfig":        RegistryKeyStateHivecurrentConfig,
		"currentuser":          RegistryKeyStateHivecurrentUser,
		"localmachinesam":      RegistryKeyStateHivelocalMachineSam,
		"localmachinesecurity": RegistryKeyStateHivelocalMachineSecurity,
		"localmachinesoftware": RegistryKeyStateHivelocalMachineSoftware,
		"localmachinesystem":   RegistryKeyStateHivelocalMachineSystem,
		"unknown":              RegistryKeyStateHiveunknown,
		"usersdefault":         RegistryKeyStateHiveusersDefault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RegistryKeyStateHive(input)
	return &out, nil
}
