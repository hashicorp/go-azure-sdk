package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap string

const (
	AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap_MicrosoftChapVersionTwo AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap = "microsoftChapVersionTwo"
	AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap_None                    AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap = "none"
)

func PossibleValuesForAndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap() []string {
	return []string{
		string(AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap_MicrosoftChapVersionTwo),
		string(AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap_None),
	}
}

func (s *AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap(input string) (*AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap, error) {
	vals := map[string]AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap{
		"microsoftchapversiontwo": AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap_MicrosoftChapVersionTwo,
		"none":                    AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap(input)
	return &out, nil
}
