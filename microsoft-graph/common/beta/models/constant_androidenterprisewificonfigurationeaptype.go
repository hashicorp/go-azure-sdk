package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidEnterpriseWiFiConfigurationEapType string

const (
	AndroidEnterpriseWiFiConfigurationEapTypeeapTls  AndroidEnterpriseWiFiConfigurationEapType = "EapTls"
	AndroidEnterpriseWiFiConfigurationEapTypeeapTtls AndroidEnterpriseWiFiConfigurationEapType = "EapTtls"
	AndroidEnterpriseWiFiConfigurationEapTypepeap    AndroidEnterpriseWiFiConfigurationEapType = "Peap"
)

func PossibleValuesForAndroidEnterpriseWiFiConfigurationEapType() []string {
	return []string{
		string(AndroidEnterpriseWiFiConfigurationEapTypeeapTls),
		string(AndroidEnterpriseWiFiConfigurationEapTypeeapTtls),
		string(AndroidEnterpriseWiFiConfigurationEapTypepeap),
	}
}

func parseAndroidEnterpriseWiFiConfigurationEapType(input string) (*AndroidEnterpriseWiFiConfigurationEapType, error) {
	vals := map[string]AndroidEnterpriseWiFiConfigurationEapType{
		"eaptls":  AndroidEnterpriseWiFiConfigurationEapTypeeapTls,
		"eapttls": AndroidEnterpriseWiFiConfigurationEapTypeeapTtls,
		"peap":    AndroidEnterpriseWiFiConfigurationEapTypepeap,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidEnterpriseWiFiConfigurationEapType(input)
	return &out, nil
}
