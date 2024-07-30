package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation string

const (
	Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation_Disabled                   Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation = "disabled"
	Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation_EnabledForAzureAd          Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation = "enabledForAzureAd"
	Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation_EnabledForAzureAdAndHybrid Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation = "enabledForAzureAdAndHybrid"
	Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation_NotConfigured              Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation = "notConfigured"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation_Disabled),
		string(Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation_EnabledForAzureAd),
		string(Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation_EnabledForAzureAdAndHybrid),
		string(Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation_NotConfigured),
	}
}

func (s *Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation(input string) (*Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation{
		"disabled":                   Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation_Disabled,
		"enabledforazuread":          Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation_EnabledForAzureAd,
		"enabledforazureadandhybrid": Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation_EnabledForAzureAdAndHybrid,
		"notconfigured":              Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation(input)
	return &out, nil
}
