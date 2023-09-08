package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsWifiEnterpriseEAPConfigurationEapType string

const (
	WindowsWifiEnterpriseEAPConfigurationEapTypeeapFast WindowsWifiEnterpriseEAPConfigurationEapType = "EapFast"
	WindowsWifiEnterpriseEAPConfigurationEapTypeeapSim  WindowsWifiEnterpriseEAPConfigurationEapType = "EapSim"
	WindowsWifiEnterpriseEAPConfigurationEapTypeeapTls  WindowsWifiEnterpriseEAPConfigurationEapType = "EapTls"
	WindowsWifiEnterpriseEAPConfigurationEapTypeeapTtls WindowsWifiEnterpriseEAPConfigurationEapType = "EapTtls"
	WindowsWifiEnterpriseEAPConfigurationEapTypeleap    WindowsWifiEnterpriseEAPConfigurationEapType = "Leap"
	WindowsWifiEnterpriseEAPConfigurationEapTypepeap    WindowsWifiEnterpriseEAPConfigurationEapType = "Peap"
	WindowsWifiEnterpriseEAPConfigurationEapTypeteap    WindowsWifiEnterpriseEAPConfigurationEapType = "Teap"
)

func PossibleValuesForWindowsWifiEnterpriseEAPConfigurationEapType() []string {
	return []string{
		string(WindowsWifiEnterpriseEAPConfigurationEapTypeeapFast),
		string(WindowsWifiEnterpriseEAPConfigurationEapTypeeapSim),
		string(WindowsWifiEnterpriseEAPConfigurationEapTypeeapTls),
		string(WindowsWifiEnterpriseEAPConfigurationEapTypeeapTtls),
		string(WindowsWifiEnterpriseEAPConfigurationEapTypeleap),
		string(WindowsWifiEnterpriseEAPConfigurationEapTypepeap),
		string(WindowsWifiEnterpriseEAPConfigurationEapTypeteap),
	}
}

func parseWindowsWifiEnterpriseEAPConfigurationEapType(input string) (*WindowsWifiEnterpriseEAPConfigurationEapType, error) {
	vals := map[string]WindowsWifiEnterpriseEAPConfigurationEapType{
		"eapfast": WindowsWifiEnterpriseEAPConfigurationEapTypeeapFast,
		"eapsim":  WindowsWifiEnterpriseEAPConfigurationEapTypeeapSim,
		"eaptls":  WindowsWifiEnterpriseEAPConfigurationEapTypeeapTls,
		"eapttls": WindowsWifiEnterpriseEAPConfigurationEapTypeeapTtls,
		"leap":    WindowsWifiEnterpriseEAPConfigurationEapTypeleap,
		"peap":    WindowsWifiEnterpriseEAPConfigurationEapTypepeap,
		"teap":    WindowsWifiEnterpriseEAPConfigurationEapTypeteap,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsWifiEnterpriseEAPConfigurationEapType(input)
	return &out, nil
}
