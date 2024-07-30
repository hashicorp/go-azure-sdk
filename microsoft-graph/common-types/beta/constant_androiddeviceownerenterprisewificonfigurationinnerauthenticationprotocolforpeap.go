package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap string

const (
	AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap_MicrosoftChapVersionTwo AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap = "microsoftChapVersionTwo"
	AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap_None                    AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap = "none"
)

func PossibleValuesForAndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap() []string {
	return []string{
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap_MicrosoftChapVersionTwo),
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap_None),
	}
}

func (s *AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap(input string) (*AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap, error) {
	vals := map[string]AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap{
		"microsoftchapversiontwo": AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap_MicrosoftChapVersionTwo,
		"none":                    AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap(input)
	return &out, nil
}
