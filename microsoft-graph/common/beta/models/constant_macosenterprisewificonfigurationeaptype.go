package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSEnterpriseWiFiConfigurationEapType string

const (
	MacOSEnterpriseWiFiConfigurationEapTypeeapFast MacOSEnterpriseWiFiConfigurationEapType = "EapFast"
	MacOSEnterpriseWiFiConfigurationEapTypeeapSim  MacOSEnterpriseWiFiConfigurationEapType = "EapSim"
	MacOSEnterpriseWiFiConfigurationEapTypeeapTls  MacOSEnterpriseWiFiConfigurationEapType = "EapTls"
	MacOSEnterpriseWiFiConfigurationEapTypeeapTtls MacOSEnterpriseWiFiConfigurationEapType = "EapTtls"
	MacOSEnterpriseWiFiConfigurationEapTypeleap    MacOSEnterpriseWiFiConfigurationEapType = "Leap"
	MacOSEnterpriseWiFiConfigurationEapTypepeap    MacOSEnterpriseWiFiConfigurationEapType = "Peap"
	MacOSEnterpriseWiFiConfigurationEapTypeteap    MacOSEnterpriseWiFiConfigurationEapType = "Teap"
)

func PossibleValuesForMacOSEnterpriseWiFiConfigurationEapType() []string {
	return []string{
		string(MacOSEnterpriseWiFiConfigurationEapTypeeapFast),
		string(MacOSEnterpriseWiFiConfigurationEapTypeeapSim),
		string(MacOSEnterpriseWiFiConfigurationEapTypeeapTls),
		string(MacOSEnterpriseWiFiConfigurationEapTypeeapTtls),
		string(MacOSEnterpriseWiFiConfigurationEapTypeleap),
		string(MacOSEnterpriseWiFiConfigurationEapTypepeap),
		string(MacOSEnterpriseWiFiConfigurationEapTypeteap),
	}
}

func parseMacOSEnterpriseWiFiConfigurationEapType(input string) (*MacOSEnterpriseWiFiConfigurationEapType, error) {
	vals := map[string]MacOSEnterpriseWiFiConfigurationEapType{
		"eapfast": MacOSEnterpriseWiFiConfigurationEapTypeeapFast,
		"eapsim":  MacOSEnterpriseWiFiConfigurationEapTypeeapSim,
		"eaptls":  MacOSEnterpriseWiFiConfigurationEapTypeeapTls,
		"eapttls": MacOSEnterpriseWiFiConfigurationEapTypeeapTtls,
		"leap":    MacOSEnterpriseWiFiConfigurationEapTypeleap,
		"peap":    MacOSEnterpriseWiFiConfigurationEapTypepeap,
		"teap":    MacOSEnterpriseWiFiConfigurationEapTypeteap,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSEnterpriseWiFiConfigurationEapType(input)
	return &out, nil
}
