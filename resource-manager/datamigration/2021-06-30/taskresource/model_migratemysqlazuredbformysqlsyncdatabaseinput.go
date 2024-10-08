package taskresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrateMySqlAzureDbForMySqlSyncDatabaseInput struct {
	MigrationSetting   *map[string]string `json:"migrationSetting,omitempty"`
	Name               *string            `json:"name,omitempty"`
	SourceSetting      *map[string]string `json:"sourceSetting,omitempty"`
	TableMap           *map[string]string `json:"tableMap,omitempty"`
	TargetDatabaseName *string            `json:"targetDatabaseName,omitempty"`
	TargetSetting      *map[string]string `json:"targetSetting,omitempty"`
}
