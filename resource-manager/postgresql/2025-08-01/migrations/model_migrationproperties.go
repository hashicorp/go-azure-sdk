package migrations

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrationProperties struct {
	Cancel                                    *Cancel                           `json:"cancel,omitempty"`
	CurrentStatus                             *MigrationStatus                  `json:"currentStatus,omitempty"`
	DbsToCancelMigrationOn                    *[]string                         `json:"dbsToCancelMigrationOn,omitempty"`
	DbsToMigrate                              *[]string                         `json:"dbsToMigrate,omitempty"`
	DbsToTriggerCutoverOn                     *[]string                         `json:"dbsToTriggerCutoverOn,omitempty"`
	MigrateRoles                              *MigrateRolesAndPermissions       `json:"migrateRoles,omitempty"`
	MigrationId                               *string                           `json:"migrationId,omitempty"`
	MigrationInstanceResourceId               *string                           `json:"migrationInstanceResourceId,omitempty"`
	MigrationMode                             *MigrationMode                    `json:"migrationMode,omitempty"`
	MigrationOption                           *MigrationOption                  `json:"migrationOption,omitempty"`
	MigrationWindowEndTimeInUtc               *string                           `json:"migrationWindowEndTimeInUtc,omitempty"`
	MigrationWindowStartTimeInUtc             *string                           `json:"migrationWindowStartTimeInUtc,omitempty"`
	OverwriteDbsInTarget                      *OverwriteDatabasesOnTargetServer `json:"overwriteDbsInTarget,omitempty"`
	SecretParameters                          *MigrationSecretParameters        `json:"secretParameters,omitempty"`
	SetupLogicalReplicationOnSourceDbIfNeeded *LogicalReplicationOnSourceServer `json:"setupLogicalReplicationOnSourceDbIfNeeded,omitempty"`
	SourceDbServerFullyQualifiedDomainName    *string                           `json:"sourceDbServerFullyQualifiedDomainName,omitempty"`
	SourceDbServerMetadata                    *DbServerMetadata                 `json:"sourceDbServerMetadata,omitempty"`
	SourceDbServerResourceId                  *string                           `json:"sourceDbServerResourceId,omitempty"`
	SourceType                                *SourceType                       `json:"sourceType,omitempty"`
	SslMode                                   *SslMode                          `json:"sslMode,omitempty"`
	StartDataMigration                        *StartDataMigration               `json:"startDataMigration,omitempty"`
	TargetDbServerFullyQualifiedDomainName    *string                           `json:"targetDbServerFullyQualifiedDomainName,omitempty"`
	TargetDbServerMetadata                    *DbServerMetadata                 `json:"targetDbServerMetadata,omitempty"`
	TargetDbServerResourceId                  *string                           `json:"targetDbServerResourceId,omitempty"`
	TriggerCutover                            *TriggerCutover                   `json:"triggerCutover,omitempty"`
}

func (o *MigrationProperties) GetMigrationWindowEndTimeInUtcAsTime() (*time.Time, error) {
	if o.MigrationWindowEndTimeInUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.MigrationWindowEndTimeInUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *MigrationProperties) SetMigrationWindowEndTimeInUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.MigrationWindowEndTimeInUtc = &formatted
}

func (o *MigrationProperties) GetMigrationWindowStartTimeInUtcAsTime() (*time.Time, error) {
	if o.MigrationWindowStartTimeInUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.MigrationWindowStartTimeInUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *MigrationProperties) SetMigrationWindowStartTimeInUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.MigrationWindowStartTimeInUtc = &formatted
}
