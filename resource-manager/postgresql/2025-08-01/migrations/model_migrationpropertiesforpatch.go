package migrations

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrationPropertiesForPatch struct {
	Cancel                                    *Cancel                            `json:"cancel,omitempty"`
	DbsToCancelMigrationOn                    *[]string                          `json:"dbsToCancelMigrationOn,omitempty"`
	DbsToMigrate                              *[]string                          `json:"dbsToMigrate,omitempty"`
	DbsToTriggerCutoverOn                     *[]string                          `json:"dbsToTriggerCutoverOn,omitempty"`
	MigrateRoles                              *MigrateRolesAndPermissions        `json:"migrateRoles,omitempty"`
	MigrationMode                             *MigrationMode                     `json:"migrationMode,omitempty"`
	MigrationWindowStartTimeInUtc             *string                            `json:"migrationWindowStartTimeInUtc,omitempty"`
	OverwriteDbsInTarget                      *OverwriteDatabasesOnTargetServer  `json:"overwriteDbsInTarget,omitempty"`
	SecretParameters                          *MigrationSecretParametersForPatch `json:"secretParameters,omitempty"`
	SetupLogicalReplicationOnSourceDbIfNeeded *LogicalReplicationOnSourceServer  `json:"setupLogicalReplicationOnSourceDbIfNeeded,omitempty"`
	SourceDbServerFullyQualifiedDomainName    *string                            `json:"sourceDbServerFullyQualifiedDomainName,omitempty"`
	SourceDbServerResourceId                  *string                            `json:"sourceDbServerResourceId,omitempty"`
	StartDataMigration                        *StartDataMigration                `json:"startDataMigration,omitempty"`
	TargetDbServerFullyQualifiedDomainName    *string                            `json:"targetDbServerFullyQualifiedDomainName,omitempty"`
	TriggerCutover                            *TriggerCutover                    `json:"triggerCutover,omitempty"`
}

func (o *MigrationPropertiesForPatch) GetMigrationWindowStartTimeInUtcAsTime() (*time.Time, error) {
	if o.MigrationWindowStartTimeInUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.MigrationWindowStartTimeInUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *MigrationPropertiesForPatch) SetMigrationWindowStartTimeInUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.MigrationWindowStartTimeInUtc = &formatted
}
