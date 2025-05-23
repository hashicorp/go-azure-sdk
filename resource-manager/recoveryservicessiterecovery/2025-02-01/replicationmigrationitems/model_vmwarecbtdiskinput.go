package replicationmigrationitems

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VMwareCbtDiskInput struct {
	DiskEncryptionSetId            *string          `json:"diskEncryptionSetId,omitempty"`
	DiskId                         string           `json:"diskId"`
	DiskSizeInGB                   *int64           `json:"diskSizeInGB,omitempty"`
	DiskType                       *DiskAccountType `json:"diskType,omitempty"`
	Iops                           *int64           `json:"iops,omitempty"`
	IsOSDisk                       string           `json:"isOSDisk"`
	LogStorageAccountId            string           `json:"logStorageAccountId"`
	LogStorageAccountSasSecretName string           `json:"logStorageAccountSasSecretName"`
	SectorSizeInBytes              *int64           `json:"sectorSizeInBytes,omitempty"`
	ThroughputInMbps               *int64           `json:"throughputInMbps,omitempty"`
}
