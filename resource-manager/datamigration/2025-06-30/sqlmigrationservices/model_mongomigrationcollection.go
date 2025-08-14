package sqlmigrationservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MongoMigrationCollection struct {
	MigrationProgressDetails *MongoMigrationProgressDetails `json:"migrationProgressDetails,omitempty"`
	SourceCollection         *string                        `json:"sourceCollection,omitempty"`
	SourceDatabase           *string                        `json:"sourceDatabase,omitempty"`
	TargetCollection         *string                        `json:"targetCollection,omitempty"`
	TargetDatabase           *string                        `json:"targetDatabase,omitempty"`
}
