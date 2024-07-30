package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerWiFiConfigurationProxySetting string

const (
	AospDeviceOwnerWiFiConfigurationProxySetting_Automatic AospDeviceOwnerWiFiConfigurationProxySetting = "automatic"
	AospDeviceOwnerWiFiConfigurationProxySetting_Manual    AospDeviceOwnerWiFiConfigurationProxySetting = "manual"
	AospDeviceOwnerWiFiConfigurationProxySetting_None      AospDeviceOwnerWiFiConfigurationProxySetting = "none"
)

func PossibleValuesForAospDeviceOwnerWiFiConfigurationProxySetting() []string {
	return []string{
		string(AospDeviceOwnerWiFiConfigurationProxySetting_Automatic),
		string(AospDeviceOwnerWiFiConfigurationProxySetting_Manual),
		string(AospDeviceOwnerWiFiConfigurationProxySetting_None),
	}
}

func (s *AospDeviceOwnerWiFiConfigurationProxySetting) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAospDeviceOwnerWiFiConfigurationProxySetting(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAospDeviceOwnerWiFiConfigurationProxySetting(input string) (*AospDeviceOwnerWiFiConfigurationProxySetting, error) {
	vals := map[string]AospDeviceOwnerWiFiConfigurationProxySetting{
		"automatic": AospDeviceOwnerWiFiConfigurationProxySetting_Automatic,
		"manual":    AospDeviceOwnerWiFiConfigurationProxySetting_Manual,
		"none":      AospDeviceOwnerWiFiConfigurationProxySetting_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerWiFiConfigurationProxySetting(input)
	return &out, nil
}
