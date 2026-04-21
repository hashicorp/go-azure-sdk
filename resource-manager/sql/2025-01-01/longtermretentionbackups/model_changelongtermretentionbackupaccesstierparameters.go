package longtermretentionbackups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangeLongTermRetentionBackupAccessTierParameters struct {
	BackupStorageAccessTier string `json:"backupStorageAccessTier"`
	OperationMode           string `json:"operationMode"`
}
