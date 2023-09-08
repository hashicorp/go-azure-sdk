package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap string

const (
	AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeapmicrosoftChapVersionTwo AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap = "MicrosoftChapVersionTwo"
	AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeapnone                    AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap = "None"
)

func PossibleValuesForAndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap() []string {
	return []string{
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeapmicrosoftChapVersionTwo),
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeapnone),
	}
}

func parseAndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap(input string) (*AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap, error) {
	vals := map[string]AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap{
		"microsoftchapversiontwo": AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeapmicrosoftChapVersionTwo,
		"none":                    AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeapnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap(input)
	return &out, nil
}
