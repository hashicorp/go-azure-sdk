package backupprotecteditems

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupManagementType string

const (
	BackupManagementTypeAzureBackupServer BackupManagementType = "AzureBackupServer"
	BackupManagementTypeAzureIaasVM       BackupManagementType = "AzureIaasVM"
	BackupManagementTypeAzureSql          BackupManagementType = "AzureSql"
	BackupManagementTypeAzureStorage      BackupManagementType = "AzureStorage"
	BackupManagementTypeAzureWorkload     BackupManagementType = "AzureWorkload"
	BackupManagementTypeDPM               BackupManagementType = "DPM"
	BackupManagementTypeDefaultBackup     BackupManagementType = "DefaultBackup"
	BackupManagementTypeInvalid           BackupManagementType = "Invalid"
	BackupManagementTypeMAB               BackupManagementType = "MAB"
)

func PossibleValuesForBackupManagementType() []string {
	return []string{
		string(BackupManagementTypeAzureBackupServer),
		string(BackupManagementTypeAzureIaasVM),
		string(BackupManagementTypeAzureSql),
		string(BackupManagementTypeAzureStorage),
		string(BackupManagementTypeAzureWorkload),
		string(BackupManagementTypeDPM),
		string(BackupManagementTypeDefaultBackup),
		string(BackupManagementTypeInvalid),
		string(BackupManagementTypeMAB),
	}
}

type CreateMode string

const (
	CreateModeDefault CreateMode = "Default"
	CreateModeInvalid CreateMode = "Invalid"
	CreateModeRecover CreateMode = "Recover"
)

func PossibleValuesForCreateMode() []string {
	return []string{
		string(CreateModeDefault),
		string(CreateModeInvalid),
		string(CreateModeRecover),
	}
}

type DataSourceType string

const (
	DataSourceTypeAzureFileShare    DataSourceType = "AzureFileShare"
	DataSourceTypeAzureSqlDb        DataSourceType = "AzureSqlDb"
	DataSourceTypeClient            DataSourceType = "Client"
	DataSourceTypeExchange          DataSourceType = "Exchange"
	DataSourceTypeFileFolder        DataSourceType = "FileFolder"
	DataSourceTypeGenericDataSource DataSourceType = "GenericDataSource"
	DataSourceTypeInvalid           DataSourceType = "Invalid"
	DataSourceTypeSAPAseDatabase    DataSourceType = "SAPAseDatabase"
	DataSourceTypeSAPHanaDatabase   DataSourceType = "SAPHanaDatabase"
	DataSourceTypeSQLDB             DataSourceType = "SQLDB"
	DataSourceTypeSQLDataBase       DataSourceType = "SQLDataBase"
	DataSourceTypeSharepoint        DataSourceType = "Sharepoint"
	DataSourceTypeSystemState       DataSourceType = "SystemState"
	DataSourceTypeVM                DataSourceType = "VM"
	DataSourceTypeVMwareVM          DataSourceType = "VMwareVM"
)

func PossibleValuesForDataSourceType() []string {
	return []string{
		string(DataSourceTypeAzureFileShare),
		string(DataSourceTypeAzureSqlDb),
		string(DataSourceTypeClient),
		string(DataSourceTypeExchange),
		string(DataSourceTypeFileFolder),
		string(DataSourceTypeGenericDataSource),
		string(DataSourceTypeInvalid),
		string(DataSourceTypeSAPAseDatabase),
		string(DataSourceTypeSAPHanaDatabase),
		string(DataSourceTypeSQLDB),
		string(DataSourceTypeSQLDataBase),
		string(DataSourceTypeSharepoint),
		string(DataSourceTypeSystemState),
		string(DataSourceTypeVM),
		string(DataSourceTypeVMwareVM),
	}
}

type HealthStatus string

const (
	HealthStatusActionRequired  HealthStatus = "ActionRequired"
	HealthStatusActionSuggested HealthStatus = "ActionSuggested"
	HealthStatusInvalid         HealthStatus = "Invalid"
	HealthStatusPassed          HealthStatus = "Passed"
)

func PossibleValuesForHealthStatus() []string {
	return []string{
		string(HealthStatusActionRequired),
		string(HealthStatusActionSuggested),
		string(HealthStatusInvalid),
		string(HealthStatusPassed),
	}
}

type LastBackupStatus string

const (
	LastBackupStatusHealthy   LastBackupStatus = "Healthy"
	LastBackupStatusIRPending LastBackupStatus = "IRPending"
	LastBackupStatusInvalid   LastBackupStatus = "Invalid"
	LastBackupStatusUnhealthy LastBackupStatus = "Unhealthy"
)

func PossibleValuesForLastBackupStatus() []string {
	return []string{
		string(LastBackupStatusHealthy),
		string(LastBackupStatusIRPending),
		string(LastBackupStatusInvalid),
		string(LastBackupStatusUnhealthy),
	}
}

type ProtectedItemHealthStatus string

const (
	ProtectedItemHealthStatusHealthy      ProtectedItemHealthStatus = "Healthy"
	ProtectedItemHealthStatusIRPending    ProtectedItemHealthStatus = "IRPending"
	ProtectedItemHealthStatusInvalid      ProtectedItemHealthStatus = "Invalid"
	ProtectedItemHealthStatusNotReachable ProtectedItemHealthStatus = "NotReachable"
	ProtectedItemHealthStatusUnhealthy    ProtectedItemHealthStatus = "Unhealthy"
)

func PossibleValuesForProtectedItemHealthStatus() []string {
	return []string{
		string(ProtectedItemHealthStatusHealthy),
		string(ProtectedItemHealthStatusIRPending),
		string(ProtectedItemHealthStatusInvalid),
		string(ProtectedItemHealthStatusNotReachable),
		string(ProtectedItemHealthStatusUnhealthy),
	}
}

type ProtectedItemState string

const (
	ProtectedItemStateIRPending         ProtectedItemState = "IRPending"
	ProtectedItemStateInvalid           ProtectedItemState = "Invalid"
	ProtectedItemStateProtected         ProtectedItemState = "Protected"
	ProtectedItemStateProtectionError   ProtectedItemState = "ProtectionError"
	ProtectedItemStateProtectionPaused  ProtectedItemState = "ProtectionPaused"
	ProtectedItemStateProtectionStopped ProtectedItemState = "ProtectionStopped"
)

func PossibleValuesForProtectedItemState() []string {
	return []string{
		string(ProtectedItemStateIRPending),
		string(ProtectedItemStateInvalid),
		string(ProtectedItemStateProtected),
		string(ProtectedItemStateProtectionError),
		string(ProtectedItemStateProtectionPaused),
		string(ProtectedItemStateProtectionStopped),
	}
}

type ProtectionState string

const (
	ProtectionStateIRPending         ProtectionState = "IRPending"
	ProtectionStateInvalid           ProtectionState = "Invalid"
	ProtectionStateProtected         ProtectionState = "Protected"
	ProtectionStateProtectionError   ProtectionState = "ProtectionError"
	ProtectionStateProtectionPaused  ProtectionState = "ProtectionPaused"
	ProtectionStateProtectionStopped ProtectionState = "ProtectionStopped"
)

func PossibleValuesForProtectionState() []string {
	return []string{
		string(ProtectionStateIRPending),
		string(ProtectionStateInvalid),
		string(ProtectionStateProtected),
		string(ProtectionStateProtectionError),
		string(ProtectionStateProtectionPaused),
		string(ProtectionStateProtectionStopped),
	}
}

type ResourceHealthStatus string

const (
	ResourceHealthStatusHealthy             ResourceHealthStatus = "Healthy"
	ResourceHealthStatusInvalid             ResourceHealthStatus = "Invalid"
	ResourceHealthStatusPersistentDegraded  ResourceHealthStatus = "PersistentDegraded"
	ResourceHealthStatusPersistentUnhealthy ResourceHealthStatus = "PersistentUnhealthy"
	ResourceHealthStatusTransientDegraded   ResourceHealthStatus = "TransientDegraded"
	ResourceHealthStatusTransientUnhealthy  ResourceHealthStatus = "TransientUnhealthy"
)

func PossibleValuesForResourceHealthStatus() []string {
	return []string{
		string(ResourceHealthStatusHealthy),
		string(ResourceHealthStatusInvalid),
		string(ResourceHealthStatusPersistentDegraded),
		string(ResourceHealthStatusPersistentUnhealthy),
		string(ResourceHealthStatusTransientDegraded),
		string(ResourceHealthStatusTransientUnhealthy),
	}
}
