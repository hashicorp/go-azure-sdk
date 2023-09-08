package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerEnterpriseWiFiConfigurationEapType string

const (
	AndroidDeviceOwnerEnterpriseWiFiConfigurationEapTypeeapTls  AndroidDeviceOwnerEnterpriseWiFiConfigurationEapType = "EapTls"
	AndroidDeviceOwnerEnterpriseWiFiConfigurationEapTypeeapTtls AndroidDeviceOwnerEnterpriseWiFiConfigurationEapType = "EapTtls"
	AndroidDeviceOwnerEnterpriseWiFiConfigurationEapTypepeap    AndroidDeviceOwnerEnterpriseWiFiConfigurationEapType = "Peap"
)

func PossibleValuesForAndroidDeviceOwnerEnterpriseWiFiConfigurationEapType() []string {
	return []string{
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationEapTypeeapTls),
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationEapTypeeapTtls),
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationEapTypepeap),
	}
}

func parseAndroidDeviceOwnerEnterpriseWiFiConfigurationEapType(input string) (*AndroidDeviceOwnerEnterpriseWiFiConfigurationEapType, error) {
	vals := map[string]AndroidDeviceOwnerEnterpriseWiFiConfigurationEapType{
		"eaptls":  AndroidDeviceOwnerEnterpriseWiFiConfigurationEapTypeeapTls,
		"eapttls": AndroidDeviceOwnerEnterpriseWiFiConfigurationEapTypeeapTtls,
		"peap":    AndroidDeviceOwnerEnterpriseWiFiConfigurationEapTypepeap,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerEnterpriseWiFiConfigurationEapType(input)
	return &out, nil
}
