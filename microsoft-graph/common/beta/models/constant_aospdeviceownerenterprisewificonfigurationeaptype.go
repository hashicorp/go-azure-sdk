package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerEnterpriseWiFiConfigurationEapType string

const (
	AospDeviceOwnerEnterpriseWiFiConfigurationEapTypeeapTls  AospDeviceOwnerEnterpriseWiFiConfigurationEapType = "EapTls"
	AospDeviceOwnerEnterpriseWiFiConfigurationEapTypeeapTtls AospDeviceOwnerEnterpriseWiFiConfigurationEapType = "EapTtls"
	AospDeviceOwnerEnterpriseWiFiConfigurationEapTypepeap    AospDeviceOwnerEnterpriseWiFiConfigurationEapType = "Peap"
)

func PossibleValuesForAospDeviceOwnerEnterpriseWiFiConfigurationEapType() []string {
	return []string{
		string(AospDeviceOwnerEnterpriseWiFiConfigurationEapTypeeapTls),
		string(AospDeviceOwnerEnterpriseWiFiConfigurationEapTypeeapTtls),
		string(AospDeviceOwnerEnterpriseWiFiConfigurationEapTypepeap),
	}
}

func parseAospDeviceOwnerEnterpriseWiFiConfigurationEapType(input string) (*AospDeviceOwnerEnterpriseWiFiConfigurationEapType, error) {
	vals := map[string]AospDeviceOwnerEnterpriseWiFiConfigurationEapType{
		"eaptls":  AospDeviceOwnerEnterpriseWiFiConfigurationEapTypeeapTls,
		"eapttls": AospDeviceOwnerEnterpriseWiFiConfigurationEapTypeeapTtls,
		"peap":    AospDeviceOwnerEnterpriseWiFiConfigurationEapTypepeap,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerEnterpriseWiFiConfigurationEapType(input)
	return &out, nil
}
