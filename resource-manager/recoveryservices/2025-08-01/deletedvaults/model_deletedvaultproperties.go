package deletedvaults

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletedVaultProperties struct {
	PurgeAt           *string `json:"purgeAt,omitempty"`
	VaultDeletionTime *string `json:"vaultDeletionTime,omitempty"`
	VaultId           *string `json:"vaultId,omitempty"`
}

func (o *DeletedVaultProperties) GetPurgeAtAsTime() (*time.Time, error) {
	if o.PurgeAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.PurgeAt, "2006-01-02T15:04:05Z07:00")
}

func (o *DeletedVaultProperties) SetPurgeAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.PurgeAt = &formatted
}

func (o *DeletedVaultProperties) GetVaultDeletionTimeAsTime() (*time.Time, error) {
	if o.VaultDeletionTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.VaultDeletionTime, "2006-01-02T15:04:05Z07:00")
}

func (o *DeletedVaultProperties) SetVaultDeletionTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.VaultDeletionTime = &formatted
}
