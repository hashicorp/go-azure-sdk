package migrations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrationResourcePropertiesForPatch struct {
	Cancel                                    *CancelEnum                       `json:"cancel,omitempty"`
	DbsToCancelMigrationOn                    *[]string                         `json:"dbsToCancelMigrationOn,omitempty"`
	DbsToMigrate                              *[]string                         `json:"dbsToMigrate,omitempty"`
	DbsToTriggerCutoverOn                     *[]string                         `json:"dbsToTriggerCutoverOn,omitempty"`
	MigrationMode                             *MigrationMode                    `json:"migrationMode,omitempty"`
	MigrationWindowStartTimeInUtc             *string                           `json:"migrationWindowStartTimeInUtc,omitempty"`
	OverwriteDbsInTarget                      *OverwriteDbsInTargetEnum         `json:"overwriteDbsInTarget,omitempty"`
	SecretParameters                          *MigrationSecretParameters        `json:"secretParameters,omitempty"`
	SetupLogicalReplicationOnSourceDbIfNeeded *LogicalReplicationOnSourceDbEnum `json:"setupLogicalReplicationOnSourceDbIfNeeded,omitempty"`
	SourceDbServerFullyQualifiedDomainName    *string                           `json:"sourceDbServerFullyQualifiedDomainName,omitempty"`
	SourceDbServerResourceId                  *string                           `json:"sourceDbServerResourceId,omitempty"`
	StartDataMigration                        *StartDataMigrationEnum           `json:"startDataMigration,omitempty"`
	TargetDbServerFullyQualifiedDomainName    *string                           `json:"targetDbServerFullyQualifiedDomainName,omitempty"`
	TriggerCutover                            *TriggerCutoverEnum               `json:"triggerCutover,omitempty"`
}
