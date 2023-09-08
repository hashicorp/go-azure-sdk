package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsWiredNetworkConfigurationEapType string

const (
	WindowsWiredNetworkConfigurationEapTypeeapFast WindowsWiredNetworkConfigurationEapType = "EapFast"
	WindowsWiredNetworkConfigurationEapTypeeapSim  WindowsWiredNetworkConfigurationEapType = "EapSim"
	WindowsWiredNetworkConfigurationEapTypeeapTls  WindowsWiredNetworkConfigurationEapType = "EapTls"
	WindowsWiredNetworkConfigurationEapTypeeapTtls WindowsWiredNetworkConfigurationEapType = "EapTtls"
	WindowsWiredNetworkConfigurationEapTypeleap    WindowsWiredNetworkConfigurationEapType = "Leap"
	WindowsWiredNetworkConfigurationEapTypepeap    WindowsWiredNetworkConfigurationEapType = "Peap"
	WindowsWiredNetworkConfigurationEapTypeteap    WindowsWiredNetworkConfigurationEapType = "Teap"
)

func PossibleValuesForWindowsWiredNetworkConfigurationEapType() []string {
	return []string{
		string(WindowsWiredNetworkConfigurationEapTypeeapFast),
		string(WindowsWiredNetworkConfigurationEapTypeeapSim),
		string(WindowsWiredNetworkConfigurationEapTypeeapTls),
		string(WindowsWiredNetworkConfigurationEapTypeeapTtls),
		string(WindowsWiredNetworkConfigurationEapTypeleap),
		string(WindowsWiredNetworkConfigurationEapTypepeap),
		string(WindowsWiredNetworkConfigurationEapTypeteap),
	}
}

func parseWindowsWiredNetworkConfigurationEapType(input string) (*WindowsWiredNetworkConfigurationEapType, error) {
	vals := map[string]WindowsWiredNetworkConfigurationEapType{
		"eapfast": WindowsWiredNetworkConfigurationEapTypeeapFast,
		"eapsim":  WindowsWiredNetworkConfigurationEapTypeeapSim,
		"eaptls":  WindowsWiredNetworkConfigurationEapTypeeapTls,
		"eapttls": WindowsWiredNetworkConfigurationEapTypeeapTtls,
		"leap":    WindowsWiredNetworkConfigurationEapTypeleap,
		"peap":    WindowsWiredNetworkConfigurationEapTypepeap,
		"teap":    WindowsWiredNetworkConfigurationEapTypeteap,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsWiredNetworkConfigurationEapType(input)
	return &out, nil
}
