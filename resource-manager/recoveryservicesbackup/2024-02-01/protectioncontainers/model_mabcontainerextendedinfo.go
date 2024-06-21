package protectioncontainers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MabContainerExtendedInfo struct {
	BackupItemType   *BackupItemType `json:"backupItemType,omitempty"`
	BackupItems      *[]string       `json:"backupItems,omitempty"`
	LastBackupStatus *string         `json:"lastBackupStatus,omitempty"`
	LastRefreshedAt  *string         `json:"lastRefreshedAt,omitempty"`
	PolicyName       *string         `json:"policyName,omitempty"`
}
