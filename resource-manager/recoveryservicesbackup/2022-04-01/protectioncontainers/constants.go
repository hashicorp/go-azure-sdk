package protectioncontainers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AcquireStorageAccountLock string

const (
	AcquireStorageAccountLockAcquire    AcquireStorageAccountLock = "Acquire"
	AcquireStorageAccountLockNotAcquire AcquireStorageAccountLock = "NotAcquire"
)

func PossibleValuesForAcquireStorageAccountLock() []string {
	return []string{
		string(AcquireStorageAccountLockAcquire),
		string(AcquireStorageAccountLockNotAcquire),
	}
}

type BackupItemType string

const (
	BackupItemTypeAzureFileShare    BackupItemType = "AzureFileShare"
	BackupItemTypeAzureSqlDb        BackupItemType = "AzureSqlDb"
	BackupItemTypeClient            BackupItemType = "Client"
	BackupItemTypeExchange          BackupItemType = "Exchange"
	BackupItemTypeFileFolder        BackupItemType = "FileFolder"
	BackupItemTypeGenericDataSource BackupItemType = "GenericDataSource"
	BackupItemTypeInvalid           BackupItemType = "Invalid"
	BackupItemTypeSAPAseDatabase    BackupItemType = "SAPAseDatabase"
	BackupItemTypeSAPHanaDatabase   BackupItemType = "SAPHanaDatabase"
	BackupItemTypeSQLDB             BackupItemType = "SQLDB"
	BackupItemTypeSQLDataBase       BackupItemType = "SQLDataBase"
	BackupItemTypeSharepoint        BackupItemType = "Sharepoint"
	BackupItemTypeSystemState       BackupItemType = "SystemState"
	BackupItemTypeVM                BackupItemType = "VM"
	BackupItemTypeVMwareVM          BackupItemType = "VMwareVM"
)

func PossibleValuesForBackupItemType() []string {
	return []string{
		string(BackupItemTypeAzureFileShare),
		string(BackupItemTypeAzureSqlDb),
		string(BackupItemTypeClient),
		string(BackupItemTypeExchange),
		string(BackupItemTypeFileFolder),
		string(BackupItemTypeGenericDataSource),
		string(BackupItemTypeInvalid),
		string(BackupItemTypeSAPAseDatabase),
		string(BackupItemTypeSAPHanaDatabase),
		string(BackupItemTypeSQLDB),
		string(BackupItemTypeSQLDataBase),
		string(BackupItemTypeSharepoint),
		string(BackupItemTypeSystemState),
		string(BackupItemTypeVM),
		string(BackupItemTypeVMwareVM),
	}
}

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

type ContainerType string

const (
	ContainerTypeAzureBackupServerContainer                  ContainerType = "AzureBackupServerContainer"
	ContainerTypeAzureSqlContainer                           ContainerType = "AzureSqlContainer"
	ContainerTypeAzureWorkloadContainer                      ContainerType = "AzureWorkloadContainer"
	ContainerTypeCluster                                     ContainerType = "Cluster"
	ContainerTypeDPMContainer                                ContainerType = "DPMContainer"
	ContainerTypeGenericContainer                            ContainerType = "GenericContainer"
	ContainerTypeIaasVMContainer                             ContainerType = "IaasVMContainer"
	ContainerTypeIaasVMServiceContainer                      ContainerType = "IaasVMServiceContainer"
	ContainerTypeInvalid                                     ContainerType = "Invalid"
	ContainerTypeMABContainer                                ContainerType = "MABContainer"
	ContainerTypeMicrosoftPointClassicComputeVirtualMachines ContainerType = "Microsoft.ClassicCompute/virtualMachines"
	ContainerTypeMicrosoftPointComputeVirtualMachines        ContainerType = "Microsoft.Compute/virtualMachines"
	ContainerTypeSQLAGWorkLoadContainer                      ContainerType = "SQLAGWorkLoadContainer"
	ContainerTypeStorageContainer                            ContainerType = "StorageContainer"
	ContainerTypeUnknown                                     ContainerType = "Unknown"
	ContainerTypeVCenter                                     ContainerType = "VCenter"
	ContainerTypeVMAppContainer                              ContainerType = "VMAppContainer"
	ContainerTypeWindows                                     ContainerType = "Windows"
)

func PossibleValuesForContainerType() []string {
	return []string{
		string(ContainerTypeAzureBackupServerContainer),
		string(ContainerTypeAzureSqlContainer),
		string(ContainerTypeAzureWorkloadContainer),
		string(ContainerTypeCluster),
		string(ContainerTypeDPMContainer),
		string(ContainerTypeGenericContainer),
		string(ContainerTypeIaasVMContainer),
		string(ContainerTypeIaasVMServiceContainer),
		string(ContainerTypeInvalid),
		string(ContainerTypeMABContainer),
		string(ContainerTypeMicrosoftPointClassicComputeVirtualMachines),
		string(ContainerTypeMicrosoftPointComputeVirtualMachines),
		string(ContainerTypeSQLAGWorkLoadContainer),
		string(ContainerTypeStorageContainer),
		string(ContainerTypeUnknown),
		string(ContainerTypeVCenter),
		string(ContainerTypeVMAppContainer),
		string(ContainerTypeWindows),
	}
}

type OperationType string

const (
	OperationTypeInvalid    OperationType = "Invalid"
	OperationTypeRegister   OperationType = "Register"
	OperationTypeReregister OperationType = "Reregister"
)

func PossibleValuesForOperationType() []string {
	return []string{
		string(OperationTypeInvalid),
		string(OperationTypeRegister),
		string(OperationTypeReregister),
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
		string(WorkloadTypeSAPHanaDatabase),
		string(WorkloadTypeSQLDB),
		string(WorkloadTypeSQLDataBase),
		string(WorkloadTypeSharepoint),
		string(WorkloadTypeSystemState),
		string(WorkloadTypeVM),
		string(WorkloadTypeVMwareVM),
	}
}
