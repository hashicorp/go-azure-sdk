package databasemigrationssqlvm

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MongoMigrationProgressDetails struct {
	DurationInSeconds      *int64                `json:"durationInSeconds,omitempty"`
	MigrationError         *string               `json:"migrationError,omitempty"`
	MigrationStatus        *MongoMigrationStatus `json:"migrationStatus,omitempty"`
	ProcessedDocumentCount *int64                `json:"processedDocumentCount,omitempty"`
	SourceDocumentCount    *int64                `json:"sourceDocumentCount,omitempty"`
}
