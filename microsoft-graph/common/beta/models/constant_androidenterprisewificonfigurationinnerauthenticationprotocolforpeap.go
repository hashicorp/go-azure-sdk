package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap string

const (
	AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeapmicrosoftChapVersionTwo AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap = "MicrosoftChapVersionTwo"
	AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeapnone                    AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap = "None"
)

func PossibleValuesForAndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap() []string {
	return []string{
		string(AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeapmicrosoftChapVersionTwo),
		string(AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeapnone),
	}
}

func parseAndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap(input string) (*AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap, error) {
	vals := map[string]AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap{
		"microsoftchapversiontwo": AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeapmicrosoftChapVersionTwo,
		"none":                    AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeapnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap(input)
	return &out, nil
}
