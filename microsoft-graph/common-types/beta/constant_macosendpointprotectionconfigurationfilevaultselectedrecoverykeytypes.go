package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypes string

const (
	MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypes_InstitutionalRecoveryKey MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypes = "institutionalRecoveryKey"
	MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypes_NotConfigured            MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypes = "notConfigured"
	MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypes_PersonalRecoveryKey      MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypes = "personalRecoveryKey"
)

func PossibleValuesForMacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypes() []string {
	return []string{
		string(MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypes_InstitutionalRecoveryKey),
		string(MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypes_NotConfigured),
		string(MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypes_PersonalRecoveryKey),
	}
}

func (s *MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypes(input string) (*MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypes, error) {
	vals := map[string]MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypes{
		"institutionalrecoverykey": MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypes_InstitutionalRecoveryKey,
		"notconfigured":            MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypes_NotConfigured,
		"personalrecoverykey":      MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypes_PersonalRecoveryKey,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypes(input)
	return &out, nil
}
