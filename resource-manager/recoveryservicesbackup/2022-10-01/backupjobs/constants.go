package backupjobs

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

type JobSupportedAction string

const (
	JobSupportedActionCancellable JobSupportedAction = "Cancellable"
	JobSupportedActionInvalid     JobSupportedAction = "Invalid"
	JobSupportedActionRetriable   JobSupportedAction = "Retriable"
)

func PossibleValuesForJobSupportedAction() []string {
	return []string{
		string(JobSupportedActionCancellable),
		string(JobSupportedActionInvalid),
		string(JobSupportedActionRetriable),
	}
}

type MabServerType string

const (
	MabServerTypeAzureBackupServerContainer MabServerType = "AzureBackupServerContainer"
	MabServerTypeAzureSqlContainer          MabServerType = "AzureSqlContainer"
	MabServerTypeCluster                    MabServerType = "Cluster"
	MabServerTypeDPMContainer               MabServerType = "DPMContainer"
	MabServerTypeGenericContainer           MabServerType = "GenericContainer"
	MabServerTypeIaasVMContainer            MabServerType = "IaasVMContainer"
	MabServerTypeIaasVMServiceContainer     MabServerType = "IaasVMServiceContainer"
	MabServerTypeInvalid                    MabServerType = "Invalid"
	MabServerTypeMABContainer               MabServerType = "MABContainer"
	MabServerTypeSQLAGWorkLoadContainer     MabServerType = "SQLAGWorkLoadContainer"
	MabServerTypeStorageContainer           MabServerType = "StorageContainer"
	MabServerTypeUnknown                    MabServerType = "Unknown"
	MabServerTypeVCenter                    MabServerType = "VCenter"
	MabServerTypeVMAppContainer             MabServerType = "VMAppContainer"
	MabServerTypeWindows                    MabServerType = "Windows"
)

func PossibleValuesForMabServerType() []string {
	return []string{
		string(MabServerTypeAzureBackupServerContainer),
		string(MabServerTypeAzureSqlContainer),
		string(MabServerTypeCluster),
		string(MabServerTypeDPMContainer),
		string(MabServerTypeGenericContainer),
		string(MabServerTypeIaasVMContainer),
		string(MabServerTypeIaasVMServiceContainer),
		string(MabServerTypeInvalid),
		string(MabServerTypeMABContainer),
		string(MabServerTypeSQLAGWorkLoadContainer),
		string(MabServerTypeStorageContainer),
		string(MabServerTypeUnknown),
		string(MabServerTypeVCenter),
		string(MabServerTypeVMAppContainer),
		string(MabServerTypeWindows),
	}
}

type WorkloadType string

const (
	WorkloadTypeAzureFileShare    WorkloadType = "AzureFileShare"
	WorkloadTypeAzureSqlDb        WorkloadType = "AzureSqlDb"
	WorkloadTypeClient            WorkloadType = "Client"
	WorkloadTypeExchange          WorkloadType = "Exchange"
	WorkloadTypeFileFolder        WorkloadType = "FileFolder"
	WorkloadTypeGenericDataSource WorkloadType = "GenericDataSource"
	WorkloadTypeInvalid           WorkloadType = "Invalid"
	WorkloadTypeSAPAseDatabase    WorkloadType = "SAPAseDatabase"
	WorkloadTypeSAPHanaDBInstance WorkloadType = "SAPHanaDBInstance"
	WorkloadTypeSAPHanaDatabase   WorkloadType = "SAPHanaDatabase"
	WorkloadTypeSQLDB             WorkloadType = "SQLDB"
	WorkloadTypeSQLDataBase       WorkloadType = "SQLDataBase"
	WorkloadTypeSharepoint        WorkloadType = "Sharepoint"
	WorkloadTypeSystemState       WorkloadType = "SystemState"
	WorkloadTypeVM                WorkloadType = "VM"
	WorkloadTypeVMwareVM          WorkloadType = "VMwareVM"
)

func PossibleValuesForWorkloadType() []string {
	return []string{
		string(WorkloadTypeAzureFileShare),
		string(WorkloadTypeAzureSqlDb),
		string(WorkloadTypeClient),
		string(WorkloadTypeExchange),
		string(WorkloadTypeFileFolder),
		string(WorkloadTypeGenericDataSource),
		string(WorkloadTypeInvalid),
		string(WorkloadTypeSAPAseDatabase),
		string(WorkloadTypeSAPHanaDBInstance),
		string(WorkloadTypeSAPHanaDatabase),
		string(WorkloadTypeSQLDB),
		string(WorkloadTypeSQLDataBase),
		string(WorkloadTypeSharepoint),
		string(WorkloadTypeSystemState),
		string(WorkloadTypeVM),
		string(WorkloadTypeVMwareVM),
	}
}
