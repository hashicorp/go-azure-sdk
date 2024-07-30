package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap string

const (
	AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap_MicrosoftChapVersionTwo AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap = "microsoftChapVersionTwo"
	AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap_None                    AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap = "none"
)

func PossibleValuesForAndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap() []string {
	return []string{
		string(AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap_MicrosoftChapVersionTwo),
		string(AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap_None),
	}
}

func (s *AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap(input string) (*AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap, error) {
	vals := map[string]AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap{
		"microsoftchapversiontwo": AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap_MicrosoftChapVersionTwo,
		"none":                    AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap(input)
	return &out, nil
}
