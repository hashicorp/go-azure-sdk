package restorabledroppedmanageddatabases

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorableDroppedManagedDatabaseProperties struct {
	CreationDate        *string `json:"creationDate,omitempty"`
	DatabaseName        *string `json:"databaseName,omitempty"`
	DeletionDate        *string `json:"deletionDate,omitempty"`
	EarliestRestoreDate *string `json:"earliestRestoreDate,omitempty"`
}
