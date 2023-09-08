package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSWiredNetworkConfigurationNetworkInterface string

const (
	MacOSWiredNetworkConfigurationNetworkInterfaceanyEthernet          MacOSWiredNetworkConfigurationNetworkInterface = "AnyEthernet"
	MacOSWiredNetworkConfigurationNetworkInterfacefirstActiveEthernet  MacOSWiredNetworkConfigurationNetworkInterface = "FirstActiveEthernet"
	MacOSWiredNetworkConfigurationNetworkInterfacefirstEthernet        MacOSWiredNetworkConfigurationNetworkInterface = "FirstEthernet"
	MacOSWiredNetworkConfigurationNetworkInterfacesecondActiveEthernet MacOSWiredNetworkConfigurationNetworkInterface = "SecondActiveEthernet"
	MacOSWiredNetworkConfigurationNetworkInterfacesecondEthernet       MacOSWiredNetworkConfigurationNetworkInterface = "SecondEthernet"
	MacOSWiredNetworkConfigurationNetworkInterfacethirdActiveEthernet  MacOSWiredNetworkConfigurationNetworkInterface = "ThirdActiveEthernet"
	MacOSWiredNetworkConfigurationNetworkInterfacethirdEthernet        MacOSWiredNetworkConfigurationNetworkInterface = "ThirdEthernet"
)

func PossibleValuesForMacOSWiredNetworkConfigurationNetworkInterface() []string {
	return []string{
		string(MacOSWiredNetworkConfigurationNetworkInterfaceanyEthernet),
		string(MacOSWiredNetworkConfigurationNetworkInterfacefirstActiveEthernet),
		string(MacOSWiredNetworkConfigurationNetworkInterfacefirstEthernet),
		string(MacOSWiredNetworkConfigurationNetworkInterfacesecondActiveEthernet),
		string(MacOSWiredNetworkConfigurationNetworkInterfacesecondEthernet),
		string(MacOSWiredNetworkConfigurationNetworkInterfacethirdActiveEthernet),
		string(MacOSWiredNetworkConfigurationNetworkInterfacethirdEthernet),
	}
}

func parseMacOSWiredNetworkConfigurationNetworkInterface(input string) (*MacOSWiredNetworkConfigurationNetworkInterface, error) {
	vals := map[string]MacOSWiredNetworkConfigurationNetworkInterface{
		"anyethernet":          MacOSWiredNetworkConfigurationNetworkInterfaceanyEthernet,
		"firstactiveethernet":  MacOSWiredNetworkConfigurationNetworkInterfacefirstActiveEthernet,
		"firstethernet":        MacOSWiredNetworkConfigurationNetworkInterfacefirstEthernet,
		"secondactiveethernet": MacOSWiredNetworkConfigurationNetworkInterfacesecondActiveEthernet,
		"secondethernet":       MacOSWiredNetworkConfigurationNetworkInterfacesecondEthernet,
		"thirdactiveethernet":  MacOSWiredNetworkConfigurationNetworkInterfacethirdActiveEthernet,
		"thirdethernet":        MacOSWiredNetworkConfigurationNetworkInterfacethirdEthernet,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSWiredNetworkConfigurationNetworkInterface(input)
	return &out, nil
}
