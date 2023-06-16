package backupengines

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupEngineExtendedInfo struct {
	AvailableDiskSpace      *float64 `json:"availableDiskSpace,omitempty"`
	AzureProtectedInstances *int64   `json:"azureProtectedInstances,omitempty"`
	DatabaseName            *string  `json:"databaseName,omitempty"`
	DiskCount               *int64   `json:"diskCount,omitempty"`
	ProtectedItemsCount     *int64   `json:"protectedItemsCount,omitempty"`
	ProtectedServersCount   *int64   `json:"protectedServersCount,omitempty"`
	RefreshedAt             *string  `json:"refreshedAt,omitempty"`
	UsedDiskSpace           *float64 `json:"usedDiskSpace,omitempty"`
}

func (o *BackupEngineExtendedInfo) GetRefreshedAtAsTime() (*time.Time, error) {
	if o.RefreshedAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.RefreshedAt, "2006-01-02T15:04:05Z07:00")
}

func (o *BackupEngineExtendedInfo) SetRefreshedAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.RefreshedAt = &formatted
}
