package webapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupItemProperties struct {
	BlobName             *string                  `json:"blobName,omitempty"`
	CorrelationId        *string                  `json:"correlationId,omitempty"`
	Created              *string                  `json:"created,omitempty"`
	Databases            *[]DatabaseBackupSetting `json:"databases,omitempty"`
	FinishedTimeStamp    *string                  `json:"finishedTimeStamp,omitempty"`
	Id                   *int64                   `json:"id,omitempty"`
	LastRestoreTimeStamp *string                  `json:"lastRestoreTimeStamp,omitempty"`
	Log                  *string                  `json:"log,omitempty"`
	Name                 *string                  `json:"name,omitempty"`
	Scheduled            *bool                    `json:"scheduled,omitempty"`
	SizeInBytes          *int64                   `json:"sizeInBytes,omitempty"`
	Status               *BackupItemStatus        `json:"status,omitempty"`
	StorageAccountUrl    *string                  `json:"storageAccountUrl,omitempty"`
	WebsiteSizeInBytes   *int64                   `json:"websiteSizeInBytes,omitempty"`
}
