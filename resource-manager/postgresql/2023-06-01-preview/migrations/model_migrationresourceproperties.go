package migrations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrationResourceProperties struct {
	Cancel                                    *CancelEnum                       `json:"cancel,omitempty"`
	CurrentStatus                             *MigrationStatus                  `json:"currentStatus,omitempty"`
	DbsToCancelMigrationOn                    *[]string                         `json:"dbsToCancelMigrationOn,omitempty"`
	DbsToMigrate                              *[]string                         `json:"dbsToMigrate,omitempty"`
	DbsToTriggerCutoverOn                     *[]string                         `json:"dbsToTriggerCutoverOn,omitempty"`
	MigrationId                               *string                           `json:"migrationId,omitempty"`
	MigrationMode                             *MigrationMode                    `json:"migrationMode,omitempty"`
	MigrationOption                           *MigrationOption                  `json:"migrationOption,omitempty"`
	MigrationWindowEndTimeInUtc               *string                           `json:"migrationWindowEndTimeInUtc,omitempty"`
	MigrationWindowStartTimeInUtc             *string                           `json:"migrationWindowStartTimeInUtc,omitempty"`
	OverwriteDbsInTarget                      *OverwriteDbsInTargetEnum         `json:"overwriteDbsInTarget,omitempty"`
	SecretParameters                          *MigrationSecretParameters        `json:"secretParameters,omitempty"`
	SetupLogicalReplicationOnSourceDbIfNeeded *LogicalReplicationOnSourceDbEnum `json:"setupLogicalReplicationOnSourceDbIfNeeded,omitempty"`
	SourceDbServerFullyQualifiedDomainName    *string                           `json:"sourceDbServerFullyQualifiedDomainName,omitempty"`
	SourceDbServerMetadata                    *DbServerMetadata                 `json:"sourceDbServerMetadata,omitempty"`
	SourceDbServerResourceId                  *string                           `json:"sourceDbServerResourceId,omitempty"`
	SourceType                                *SourceType                       `json:"sourceType,omitempty"`
	SslMode                                   *SslMode                          `json:"sslMode,omitempty"`
	StartDataMigration                        *StartDataMigrationEnum           `json:"startDataMigration,omitempty"`
	TargetDbServerFullyQualifiedDomainName    *string                           `json:"targetDbServerFullyQualifiedDomainName,omitempty"`
	TargetDbServerMetadata                    *DbServerMetadata                 `json:"targetDbServerMetadata,omitempty"`
	TargetDbServerResourceId                  *string                           `json:"targetDbServerResourceId,omitempty"`
	TriggerCutover                            *TriggerCutoverEnum               `json:"triggerCutover,omitempty"`
}
