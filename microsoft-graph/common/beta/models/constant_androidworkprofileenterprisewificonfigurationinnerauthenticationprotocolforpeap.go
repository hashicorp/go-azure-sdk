package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap string

const (
	AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeapmicrosoftChapVersionTwo AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap = "MicrosoftChapVersionTwo"
	AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeapnone                    AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap = "None"
)

func PossibleValuesForAndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap() []string {
	return []string{
		string(AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeapmicrosoftChapVersionTwo),
		string(AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeapnone),
	}
}

func parseAndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap(input string) (*AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap, error) {
	vals := map[string]AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap{
		"microsoftchapversiontwo": AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeapmicrosoftChapVersionTwo,
		"none":                    AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeapnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap(input)
	return &out, nil
}
