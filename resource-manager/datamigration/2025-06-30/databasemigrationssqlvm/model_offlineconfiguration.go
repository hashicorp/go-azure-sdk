package databasemigrationssqlvm

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OfflineConfiguration struct {
	LastBackupName *string `json:"lastBackupName,omitempty"`
	Offline        *bool   `json:"offline,omitempty"`
}
