package sqlpools

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolResourceProperties struct {
	Collation             *string     `json:"collation,omitempty"`
	CreateMode            *CreateMode `json:"createMode,omitempty"`
	CreationDate          *string     `json:"creationDate,omitempty"`
	MaxSizeBytes          *int64      `json:"maxSizeBytes,omitempty"`
	ProvisioningState     *string     `json:"provisioningState,omitempty"`
	RecoverableDatabaseId *string     `json:"recoverableDatabaseId,omitempty"`
	RestorePointInTime    *string     `json:"restorePointInTime,omitempty"`
	SourceDatabaseId      *string     `json:"sourceDatabaseId,omitempty"`
	Status                *string     `json:"status,omitempty"`
}

func (o *SqlPoolResourceProperties) GetCreationDateAsTime() (*time.Time, error) {
	if o.CreationDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreationDate, "2006-01-02T15:04:05Z07:00")
}

func (o *SqlPoolResourceProperties) SetCreationDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreationDate = &formatted
}
