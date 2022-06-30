package backups

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupProperties struct {
	BackupId            *string     `json:"backupId,omitempty"`
	BackupType          *BackupType `json:"backupType,omitempty"`
	CreationDate        *string     `json:"creationDate,omitempty"`
	FailureReason       *string     `json:"failureReason,omitempty"`
	Label               *string     `json:"label,omitempty"`
	ProvisioningState   *string     `json:"provisioningState,omitempty"`
	Size                *int64      `json:"size,omitempty"`
	UseExistingSnapshot *bool       `json:"useExistingSnapshot,omitempty"`
	VolumeName          *string     `json:"volumeName,omitempty"`
}

func (o *BackupProperties) GetCreationDateAsTime() (*time.Time, error) {
	if o.CreationDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreationDate, "2006-01-02T15:04:05Z07:00")
}

func (o *BackupProperties) SetCreationDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreationDate = &formatted
}
