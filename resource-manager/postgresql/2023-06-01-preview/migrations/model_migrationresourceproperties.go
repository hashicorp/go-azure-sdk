package migrations

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

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

func (o *MigrationResourceProperties) GetMigrationWindowEndTimeInUtcAsTime() (*time.Time, error) {
	if o.MigrationWindowEndTimeInUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.MigrationWindowEndTimeInUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *MigrationResourceProperties) SetMigrationWindowEndTimeInUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.MigrationWindowEndTimeInUtc = &formatted
}

func (o *MigrationResourceProperties) GetMigrationWindowStartTimeInUtcAsTime() (*time.Time, error) {
	if o.MigrationWindowStartTimeInUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.MigrationWindowStartTimeInUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *MigrationResourceProperties) SetMigrationWindowStartTimeInUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.MigrationWindowStartTimeInUtc = &formatted
}
