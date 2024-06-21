package backups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupProperties struct {
	BackupId               *string     `json:"backupId,omitempty"`
	BackupPolicyResourceId *string     `json:"backupPolicyResourceId,omitempty"`
	BackupType             *BackupType `json:"backupType,omitempty"`
	CreationDate           *string     `json:"creationDate,omitempty"`
	FailureReason          *string     `json:"failureReason,omitempty"`
	Label                  *string     `json:"label,omitempty"`
	ProvisioningState      *string     `json:"provisioningState,omitempty"`
	Size                   *int64      `json:"size,omitempty"`
	SnapshotName           *string     `json:"snapshotName,omitempty"`
	UseExistingSnapshot    *bool       `json:"useExistingSnapshot,omitempty"`
	VolumeResourceId       string      `json:"volumeResourceId"`
}
