package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileEnterpriseWiFiConfigurationEapType string

const (
	AndroidWorkProfileEnterpriseWiFiConfigurationEapTypeeapTls  AndroidWorkProfileEnterpriseWiFiConfigurationEapType = "EapTls"
	AndroidWorkProfileEnterpriseWiFiConfigurationEapTypeeapTtls AndroidWorkProfileEnterpriseWiFiConfigurationEapType = "EapTtls"
	AndroidWorkProfileEnterpriseWiFiConfigurationEapTypepeap    AndroidWorkProfileEnterpriseWiFiConfigurationEapType = "Peap"
)

func PossibleValuesForAndroidWorkProfileEnterpriseWiFiConfigurationEapType() []string {
	return []string{
		string(AndroidWorkProfileEnterpriseWiFiConfigurationEapTypeeapTls),
		string(AndroidWorkProfileEnterpriseWiFiConfigurationEapTypeeapTtls),
		string(AndroidWorkProfileEnterpriseWiFiConfigurationEapTypepeap),
	}
}

func parseAndroidWorkProfileEnterpriseWiFiConfigurationEapType(input string) (*AndroidWorkProfileEnterpriseWiFiConfigurationEapType, error) {
	vals := map[string]AndroidWorkProfileEnterpriseWiFiConfigurationEapType{
		"eaptls":  AndroidWorkProfileEnterpriseWiFiConfigurationEapTypeeapTls,
		"eapttls": AndroidWorkProfileEnterpriseWiFiConfigurationEapTypeeapTtls,
		"peap":    AndroidWorkProfileEnterpriseWiFiConfigurationEapTypepeap,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileEnterpriseWiFiConfigurationEapType(input)
	return &out, nil
}
