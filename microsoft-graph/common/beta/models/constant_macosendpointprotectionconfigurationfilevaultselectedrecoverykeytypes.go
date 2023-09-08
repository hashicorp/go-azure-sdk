package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypes string

const (
	MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypesinstitutionalRecoveryKey MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypes = "InstitutionalRecoveryKey"
	MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypesnotConfigured            MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypes = "NotConfigured"
	MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypespersonalRecoveryKey      MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypes = "PersonalRecoveryKey"
)

func PossibleValuesForMacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypes() []string {
	return []string{
		string(MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypesinstitutionalRecoveryKey),
		string(MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypesnotConfigured),
		string(MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypespersonalRecoveryKey),
	}
}

func parseMacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypes(input string) (*MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypes, error) {
	vals := map[string]MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypes{
		"institutionalrecoverykey": MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypesinstitutionalRecoveryKey,
		"notconfigured":            MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypesnotConfigured,
		"personalrecoverykey":      MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypespersonalRecoveryKey,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSEndpointProtectionConfigurationFileVaultSelectedRecoveryKeyTypes(input)
	return &out, nil
}
