package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDeviceGuardLocalSystemAuthorityCredentialGuardSettings string

const (
	Windows10EndpointProtectionConfigurationDeviceGuardLocalSystemAuthorityCredentialGuardSettings_Disable               Windows10EndpointProtectionConfigurationDeviceGuardLocalSystemAuthorityCredentialGuardSettings = "disable"
	Windows10EndpointProtectionConfigurationDeviceGuardLocalSystemAuthorityCredentialGuardSettings_EnableWithUEFILock    Windows10EndpointProtectionConfigurationDeviceGuardLocalSystemAuthorityCredentialGuardSettings = "enableWithUEFILock"
	Windows10EndpointProtectionConfigurationDeviceGuardLocalSystemAuthorityCredentialGuardSettings_EnableWithoutUEFILock Windows10EndpointProtectionConfigurationDeviceGuardLocalSystemAuthorityCredentialGuardSettings = "enableWithoutUEFILock"
	Windows10EndpointProtectionConfigurationDeviceGuardLocalSystemAuthorityCredentialGuardSettings_NotConfigured         Windows10EndpointProtectionConfigurationDeviceGuardLocalSystemAuthorityCredentialGuardSettings = "notConfigured"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDeviceGuardLocalSystemAuthorityCredentialGuardSettings() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDeviceGuardLocalSystemAuthorityCredentialGuardSettings_Disable),
		string(Windows10EndpointProtectionConfigurationDeviceGuardLocalSystemAuthorityCredentialGuardSettings_EnableWithUEFILock),
		string(Windows10EndpointProtectionConfigurationDeviceGuardLocalSystemAuthorityCredentialGuardSettings_EnableWithoutUEFILock),
		string(Windows10EndpointProtectionConfigurationDeviceGuardLocalSystemAuthorityCredentialGuardSettings_NotConfigured),
	}
}

func (s *Windows10EndpointProtectionConfigurationDeviceGuardLocalSystemAuthorityCredentialGuardSettings) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDeviceGuardLocalSystemAuthorityCredentialGuardSettings(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDeviceGuardLocalSystemAuthorityCredentialGuardSettings(input string) (*Windows10EndpointProtectionConfigurationDeviceGuardLocalSystemAuthorityCredentialGuardSettings, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDeviceGuardLocalSystemAuthorityCredentialGuardSettings{
		"disable":               Windows10EndpointProtectionConfigurationDeviceGuardLocalSystemAuthorityCredentialGuardSettings_Disable,
		"enablewithuefilock":    Windows10EndpointProtectionConfigurationDeviceGuardLocalSystemAuthorityCredentialGuardSettings_EnableWithUEFILock,
		"enablewithoutuefilock": Windows10EndpointProtectionConfigurationDeviceGuardLocalSystemAuthorityCredentialGuardSettings_EnableWithoutUEFILock,
		"notconfigured":         Windows10EndpointProtectionConfigurationDeviceGuardLocalSystemAuthorityCredentialGuardSettings_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDeviceGuardLocalSystemAuthorityCredentialGuardSettings(input)
	return &out, nil
}
