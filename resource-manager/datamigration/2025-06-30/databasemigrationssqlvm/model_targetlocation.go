package databasemigrationssqlvm

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetLocation struct {
	AccountKey               *string `json:"accountKey,omitempty"`
	StorageAccountResourceId *string `json:"storageAccountResourceId,omitempty"`
}
