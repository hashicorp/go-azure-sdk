package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerEnterpriseWiFiConfigurationProxySetting string

const (
	AospDeviceOwnerEnterpriseWiFiConfigurationProxySetting_Automatic AospDeviceOwnerEnterpriseWiFiConfigurationProxySetting = "automatic"
	AospDeviceOwnerEnterpriseWiFiConfigurationProxySetting_Manual    AospDeviceOwnerEnterpriseWiFiConfigurationProxySetting = "manual"
	AospDeviceOwnerEnterpriseWiFiConfigurationProxySetting_None      AospDeviceOwnerEnterpriseWiFiConfigurationProxySetting = "none"
)

func PossibleValuesForAospDeviceOwnerEnterpriseWiFiConfigurationProxySetting() []string {
	return []string{
		string(AospDeviceOwnerEnterpriseWiFiConfigurationProxySetting_Automatic),
		string(AospDeviceOwnerEnterpriseWiFiConfigurationProxySetting_Manual),
		string(AospDeviceOwnerEnterpriseWiFiConfigurationProxySetting_None),
	}
}

func (s *AospDeviceOwnerEnterpriseWiFiConfigurationProxySetting) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAospDeviceOwnerEnterpriseWiFiConfigurationProxySetting(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAospDeviceOwnerEnterpriseWiFiConfigurationProxySetting(input string) (*AospDeviceOwnerEnterpriseWiFiConfigurationProxySetting, error) {
	vals := map[string]AospDeviceOwnerEnterpriseWiFiConfigurationProxySetting{
		"automatic": AospDeviceOwnerEnterpriseWiFiConfigurationProxySetting_Automatic,
		"manual":    AospDeviceOwnerEnterpriseWiFiConfigurationProxySetting_Manual,
		"none":      AospDeviceOwnerEnterpriseWiFiConfigurationProxySetting_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerEnterpriseWiFiConfigurationProxySetting(input)
	return &out, nil
}
