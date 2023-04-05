package recoveryservices

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
