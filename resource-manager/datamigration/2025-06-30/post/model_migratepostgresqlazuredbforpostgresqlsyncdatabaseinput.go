package post

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigratePostgreSqlAzureDbForPostgreSqlSyncDatabaseInput struct {
	Id                 *string                                                        `json:"id,omitempty"`
	MigrationSetting   *map[string]interface{}                                        `json:"migrationSetting,omitempty"`
	Name               *string                                                        `json:"name,omitempty"`
	SelectedTables     *[]MigratePostgreSqlAzureDbForPostgreSqlSyncDatabaseTableInput `json:"selectedTables,omitempty"`
	SourceSetting      *map[string]string                                             `json:"sourceSetting,omitempty"`
	TargetDatabaseName *string                                                        `json:"targetDatabaseName,omitempty"`
	TargetSetting      *map[string]string                                             `json:"targetSetting,omitempty"`
}
