package migrationconfigs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrationConfigPropertiesProperties struct {
	MigrationState                    *string `json:"migrationState,omitempty"`
	PendingReplicationOperationsCount *int64  `json:"pendingReplicationOperationsCount,omitempty"`
	PostMigrationName                 string  `json:"postMigrationName"`
	ProvisioningState                 *string `json:"provisioningState,omitempty"`
	TargetNamespace                   string  `json:"targetNamespace"`
}
