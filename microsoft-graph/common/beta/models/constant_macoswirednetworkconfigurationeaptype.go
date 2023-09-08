package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSWiredNetworkConfigurationEapType string

const (
	MacOSWiredNetworkConfigurationEapTypeeapFast MacOSWiredNetworkConfigurationEapType = "EapFast"
	MacOSWiredNetworkConfigurationEapTypeeapSim  MacOSWiredNetworkConfigurationEapType = "EapSim"
	MacOSWiredNetworkConfigurationEapTypeeapTls  MacOSWiredNetworkConfigurationEapType = "EapTls"
	MacOSWiredNetworkConfigurationEapTypeeapTtls MacOSWiredNetworkConfigurationEapType = "EapTtls"
	MacOSWiredNetworkConfigurationEapTypeleap    MacOSWiredNetworkConfigurationEapType = "Leap"
	MacOSWiredNetworkConfigurationEapTypepeap    MacOSWiredNetworkConfigurationEapType = "Peap"
	MacOSWiredNetworkConfigurationEapTypeteap    MacOSWiredNetworkConfigurationEapType = "Teap"
)

func PossibleValuesForMacOSWiredNetworkConfigurationEapType() []string {
	return []string{
		string(MacOSWiredNetworkConfigurationEapTypeeapFast),
		string(MacOSWiredNetworkConfigurationEapTypeeapSim),
		string(MacOSWiredNetworkConfigurationEapTypeeapTls),
		string(MacOSWiredNetworkConfigurationEapTypeeapTtls),
		string(MacOSWiredNetworkConfigurationEapTypeleap),
		string(MacOSWiredNetworkConfigurationEapTypepeap),
		string(MacOSWiredNetworkConfigurationEapTypeteap),
	}
}

func parseMacOSWiredNetworkConfigurationEapType(input string) (*MacOSWiredNetworkConfigurationEapType, error) {
	vals := map[string]MacOSWiredNetworkConfigurationEapType{
		"eapfast": MacOSWiredNetworkConfigurationEapTypeeapFast,
		"eapsim":  MacOSWiredNetworkConfigurationEapTypeeapSim,
		"eaptls":  MacOSWiredNetworkConfigurationEapTypeeapTls,
		"eapttls": MacOSWiredNetworkConfigurationEapTypeeapTtls,
		"leap":    MacOSWiredNetworkConfigurationEapTypeleap,
		"peap":    MacOSWiredNetworkConfigurationEapTypepeap,
		"teap":    MacOSWiredNetworkConfigurationEapTypeteap,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSWiredNetworkConfigurationEapType(input)
	return &out, nil
}
