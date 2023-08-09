package recoveryservices

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VaultSubResourceType string

const (
	VaultSubResourceTypeAzureBackup          VaultSubResourceType = "AzureBackup"
	VaultSubResourceTypeAzureBackupSecondary VaultSubResourceType = "AzureBackup_secondary"
	VaultSubResourceTypeAzureSiteRecovery    VaultSubResourceType = "AzureSiteRecovery"
)

func PossibleValuesForVaultSubResourceType() []string {
	return []string{
		string(VaultSubResourceTypeAzureBackup),
		string(VaultSubResourceTypeAzureBackupSecondary),
		string(VaultSubResourceTypeAzureSiteRecovery),
	}
}

func (s *VaultSubResourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVaultSubResourceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVaultSubResourceType(input string) (*VaultSubResourceType, error) {
	vals := map[string]VaultSubResourceType{
		"azurebackup":           VaultSubResourceTypeAzureBackup,
		"azurebackup_secondary": VaultSubResourceTypeAzureBackupSecondary,
		"azuresiterecovery":     VaultSubResourceTypeAzureSiteRecovery,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VaultSubResourceType(input)
	return &out, nil
}
