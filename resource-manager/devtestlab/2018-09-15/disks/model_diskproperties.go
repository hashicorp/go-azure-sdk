package disks

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiskProperties struct {
	CreatedDate       *string      `json:"createdDate,omitempty"`
	DiskBlobName      *string      `json:"diskBlobName,omitempty"`
	DiskSizeGiB       *int64       `json:"diskSizeGiB,omitempty"`
	DiskType          *StorageType `json:"diskType,omitempty"`
	DiskUri           *string      `json:"diskUri,omitempty"`
	HostCaching       *string      `json:"hostCaching,omitempty"`
	LeasedByLabVMId   *string      `json:"leasedByLabVmId,omitempty"`
	ManagedDiskId     *string      `json:"managedDiskId,omitempty"`
	ProvisioningState *string      `json:"provisioningState,omitempty"`
	StorageAccountId  *string      `json:"storageAccountId,omitempty"`
	UniqueIdentifier  *string      `json:"uniqueIdentifier,omitempty"`
}

func (o *DiskProperties) GetCreatedDateAsTime() (*time.Time, error) {
	if o.CreatedDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedDate, "2006-01-02T15:04:05Z07:00")
}

func (o *DiskProperties) SetCreatedDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreatedDate = &formatted
}
