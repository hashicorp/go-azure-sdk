package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkEnterpriseWiFiConfigurationEapType string

const (
	AndroidForWorkEnterpriseWiFiConfigurationEapTypeeapTls  AndroidForWorkEnterpriseWiFiConfigurationEapType = "EapTls"
	AndroidForWorkEnterpriseWiFiConfigurationEapTypeeapTtls AndroidForWorkEnterpriseWiFiConfigurationEapType = "EapTtls"
	AndroidForWorkEnterpriseWiFiConfigurationEapTypepeap    AndroidForWorkEnterpriseWiFiConfigurationEapType = "Peap"
)

func PossibleValuesForAndroidForWorkEnterpriseWiFiConfigurationEapType() []string {
	return []string{
		string(AndroidForWorkEnterpriseWiFiConfigurationEapTypeeapTls),
		string(AndroidForWorkEnterpriseWiFiConfigurationEapTypeeapTtls),
		string(AndroidForWorkEnterpriseWiFiConfigurationEapTypepeap),
	}
}

func parseAndroidForWorkEnterpriseWiFiConfigurationEapType(input string) (*AndroidForWorkEnterpriseWiFiConfigurationEapType, error) {
	vals := map[string]AndroidForWorkEnterpriseWiFiConfigurationEapType{
		"eaptls":  AndroidForWorkEnterpriseWiFiConfigurationEapTypeeapTls,
		"eapttls": AndroidForWorkEnterpriseWiFiConfigurationEapTypeeapTtls,
		"peap":    AndroidForWorkEnterpriseWiFiConfigurationEapTypepeap,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkEnterpriseWiFiConfigurationEapType(input)
	return &out, nil
}
