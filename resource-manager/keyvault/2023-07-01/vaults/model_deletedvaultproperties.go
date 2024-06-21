package vaults

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletedVaultProperties struct {
	DeletionDate           *string            `json:"deletionDate,omitempty"`
	Location               *string            `json:"location,omitempty"`
	PurgeProtectionEnabled *bool              `json:"purgeProtectionEnabled,omitempty"`
	ScheduledPurgeDate     *string            `json:"scheduledPurgeDate,omitempty"`
	Tags                   *map[string]string `json:"tags,omitempty"`
	VaultId                *string            `json:"vaultId,omitempty"`
}
