package databases

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatabaseUpdateProperties struct {
	AutoPauseDelay                    *int64                       `json:"autoPauseDelay,omitempty"`
	CatalogCollation                  *CatalogCollationType        `json:"catalogCollation,omitempty"`
	Collation                         *string                      `json:"collation,omitempty"`
	CreateMode                        *CreateMode                  `json:"createMode,omitempty"`
	CreationDate                      *string                      `json:"creationDate,omitempty"`
	CurrentBackupStorageRedundancy    *BackupStorageRedundancy     `json:"currentBackupStorageRedundancy,omitempty"`
	CurrentServiceObjectiveName       *string                      `json:"currentServiceObjectiveName,omitempty"`
	CurrentSku                        *Sku                         `json:"currentSku,omitempty"`
	DatabaseId                        *string                      `json:"databaseId,omitempty"`
	DefaultSecondaryLocation          *string                      `json:"defaultSecondaryLocation,omitempty"`
	EarliestRestoreDate               *string                      `json:"earliestRestoreDate,omitempty"`
	ElasticPoolId                     *string                      `json:"elasticPoolId,omitempty"`
	EncryptionProtector               *string                      `json:"encryptionProtector,omitempty"`
	EncryptionProtectorAutoRotation   *bool                        `json:"encryptionProtectorAutoRotation,omitempty"`
	FailoverGroupId                   *string                      `json:"failoverGroupId,omitempty"`
	FederatedClientId                 *string                      `json:"federatedClientId,omitempty"`
	FreeLimitExhaustionBehavior       *FreeLimitExhaustionBehavior `json:"freeLimitExhaustionBehavior,omitempty"`
	HighAvailabilityReplicaCount      *int64                       `json:"highAvailabilityReplicaCount,omitempty"`
	IsInfraEncryptionEnabled          *bool                        `json:"isInfraEncryptionEnabled,omitempty"`
	IsLedgerOn                        *bool                        `json:"isLedgerOn,omitempty"`
	Keys                              *map[string]DatabaseKey      `json:"keys,omitempty"`
	LicenseType                       *DatabaseLicenseType         `json:"licenseType,omitempty"`
	LongTermRetentionBackupResourceId *string                      `json:"longTermRetentionBackupResourceId,omitempty"`
	MaintenanceConfigurationId        *string                      `json:"maintenanceConfigurationId,omitempty"`
	ManualCutover                     *bool                        `json:"manualCutover,omitempty"`
	MaxLogSizeBytes                   *int64                       `json:"maxLogSizeBytes,omitempty"`
	MaxSizeBytes                      *int64                       `json:"maxSizeBytes,omitempty"`
	MinCapacity                       *float64                     `json:"minCapacity,omitempty"`
	PausedDate                        *string                      `json:"pausedDate,omitempty"`
	PerformCutover                    *bool                        `json:"performCutover,omitempty"`
	PreferredEnclaveType              *AlwaysEncryptedEnclaveType  `json:"preferredEnclaveType,omitempty"`
	ReadScale                         *DatabaseReadScale           `json:"readScale,omitempty"`
	RecoverableDatabaseId             *string                      `json:"recoverableDatabaseId,omitempty"`
	RecoveryServicesRecoveryPointId   *string                      `json:"recoveryServicesRecoveryPointId,omitempty"`
	RequestedBackupStorageRedundancy  *BackupStorageRedundancy     `json:"requestedBackupStorageRedundancy,omitempty"`
	RequestedServiceObjectiveName     *string                      `json:"requestedServiceObjectiveName,omitempty"`
	RestorableDroppedDatabaseId       *string                      `json:"restorableDroppedDatabaseId,omitempty"`
	RestorePointInTime                *string                      `json:"restorePointInTime,omitempty"`
	ResumedDate                       *string                      `json:"resumedDate,omitempty"`
	SampleName                        *SampleName                  `json:"sampleName,omitempty"`
	SecondaryType                     *SecondaryType               `json:"secondaryType,omitempty"`
	SourceDatabaseDeletionDate        *string                      `json:"sourceDatabaseDeletionDate,omitempty"`
	SourceDatabaseId                  *string                      `json:"sourceDatabaseId,omitempty"`
	Status                            *DatabaseStatus              `json:"status,omitempty"`
	UseFreeLimit                      *bool                        `json:"useFreeLimit,omitempty"`
	ZoneRedundant                     *bool                        `json:"zoneRedundant,omitempty"`
}
