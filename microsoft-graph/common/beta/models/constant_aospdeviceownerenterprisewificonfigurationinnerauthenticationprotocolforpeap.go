package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap string

const (
	AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeapmicrosoftChapVersionTwo AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap = "MicrosoftChapVersionTwo"
	AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeapnone                    AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap = "None"
)

func PossibleValuesForAospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap() []string {
	return []string{
		string(AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeapmicrosoftChapVersionTwo),
		string(AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeapnone),
	}
}

func parseAospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap(input string) (*AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap, error) {
	vals := map[string]AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap{
		"microsoftchapversiontwo": AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeapmicrosoftChapVersionTwo,
		"none":                    AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeapnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap(input)
	return &out, nil
}
