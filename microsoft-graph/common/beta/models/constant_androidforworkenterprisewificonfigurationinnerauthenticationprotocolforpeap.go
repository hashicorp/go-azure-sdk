package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap string

const (
	AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeapmicrosoftChapVersionTwo AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap = "MicrosoftChapVersionTwo"
	AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeapnone                    AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap = "None"
)

func PossibleValuesForAndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap() []string {
	return []string{
		string(AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeapmicrosoftChapVersionTwo),
		string(AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeapnone),
	}
}

func parseAndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap(input string) (*AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap, error) {
	vals := map[string]AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap{
		"microsoftchapversiontwo": AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeapmicrosoftChapVersionTwo,
		"none":                    AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeapnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap(input)
	return &out, nil
}
