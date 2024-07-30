package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap string

const (
	AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap_MicrosoftChapVersionTwo AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap = "microsoftChapVersionTwo"
	AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap_None                    AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap = "none"
)

func PossibleValuesForAospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap() []string {
	return []string{
		string(AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap_MicrosoftChapVersionTwo),
		string(AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap_None),
	}
}

func (s *AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap(input string) (*AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap, error) {
	vals := map[string]AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap{
		"microsoftchapversiontwo": AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap_MicrosoftChapVersionTwo,
		"none":                    AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForPeap(input)
	return &out, nil
}
