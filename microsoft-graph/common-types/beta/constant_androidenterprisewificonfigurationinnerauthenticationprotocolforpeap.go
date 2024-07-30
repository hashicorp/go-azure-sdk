package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap string

const (
	AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap_MicrosoftChapVersionTwo AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap = "microsoftChapVersionTwo"
	AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap_None                    AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap = "none"
)

func PossibleValuesForAndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap() []string {
	return []string{
		string(AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap_MicrosoftChapVersionTwo),
		string(AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap_None),
	}
}

func (s *AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap(input string) (*AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap, error) {
	vals := map[string]AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap{
		"microsoftchapversiontwo": AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap_MicrosoftChapVersionTwo,
		"none":                    AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap(input)
	return &out, nil
}
