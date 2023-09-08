package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEnterpriseWiFiConfigurationEapType string

const (
	IosEnterpriseWiFiConfigurationEapTypeeapFast IosEnterpriseWiFiConfigurationEapType = "EapFast"
	IosEnterpriseWiFiConfigurationEapTypeeapSim  IosEnterpriseWiFiConfigurationEapType = "EapSim"
	IosEnterpriseWiFiConfigurationEapTypeeapTls  IosEnterpriseWiFiConfigurationEapType = "EapTls"
	IosEnterpriseWiFiConfigurationEapTypeeapTtls IosEnterpriseWiFiConfigurationEapType = "EapTtls"
	IosEnterpriseWiFiConfigurationEapTypeleap    IosEnterpriseWiFiConfigurationEapType = "Leap"
	IosEnterpriseWiFiConfigurationEapTypepeap    IosEnterpriseWiFiConfigurationEapType = "Peap"
	IosEnterpriseWiFiConfigurationEapTypeteap    IosEnterpriseWiFiConfigurationEapType = "Teap"
)

func PossibleValuesForIosEnterpriseWiFiConfigurationEapType() []string {
	return []string{
		string(IosEnterpriseWiFiConfigurationEapTypeeapFast),
		string(IosEnterpriseWiFiConfigurationEapTypeeapSim),
		string(IosEnterpriseWiFiConfigurationEapTypeeapTls),
		string(IosEnterpriseWiFiConfigurationEapTypeeapTtls),
		string(IosEnterpriseWiFiConfigurationEapTypeleap),
		string(IosEnterpriseWiFiConfigurationEapTypepeap),
		string(IosEnterpriseWiFiConfigurationEapTypeteap),
	}
}

func parseIosEnterpriseWiFiConfigurationEapType(input string) (*IosEnterpriseWiFiConfigurationEapType, error) {
	vals := map[string]IosEnterpriseWiFiConfigurationEapType{
		"eapfast": IosEnterpriseWiFiConfigurationEapTypeeapFast,
		"eapsim":  IosEnterpriseWiFiConfigurationEapTypeeapSim,
		"eaptls":  IosEnterpriseWiFiConfigurationEapTypeeapTls,
		"eapttls": IosEnterpriseWiFiConfigurationEapTypeeapTtls,
		"leap":    IosEnterpriseWiFiConfigurationEapTypeleap,
		"peap":    IosEnterpriseWiFiConfigurationEapTypepeap,
		"teap":    IosEnterpriseWiFiConfigurationEapTypeteap,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosEnterpriseWiFiConfigurationEapType(input)
	return &out, nil
}
