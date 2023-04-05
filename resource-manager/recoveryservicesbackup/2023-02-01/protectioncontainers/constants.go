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
	BackupItemTypeSAPHanaDBInstance BackupItemType = "SAPHanaDBInstance"
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
		string(BackupItemTypeSAPHanaDBInstance),
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

type ProtectableContainerType string

const (
	ProtectableContainerTypeAzureBackupServerContainer                  ProtectableContainerType = "AzureBackupServerContainer"
	ProtectableContainerTypeAzureSqlContainer                           ProtectableContainerType = "AzureSqlContainer"
	ProtectableContainerTypeAzureWorkloadContainer                      ProtectableContainerType = "AzureWorkloadContainer"
	ProtectableContainerTypeCluster                                     ProtectableContainerType = "Cluster"
	ProtectableContainerTypeDPMContainer                                ProtectableContainerType = "DPMContainer"
	ProtectableContainerTypeGenericContainer                            ProtectableContainerType = "GenericContainer"
	ProtectableContainerTypeIaasVMContainer                             ProtectableContainerType = "IaasVMContainer"
	ProtectableContainerTypeIaasVMServiceContainer                      ProtectableContainerType = "IaasVMServiceContainer"
	ProtectableContainerTypeInvalid                                     ProtectableContainerType = "Invalid"
	ProtectableContainerTypeMABContainer                                ProtectableContainerType = "MABContainer"
	ProtectableContainerTypeMicrosoftPointClassicComputeVirtualMachines ProtectableContainerType = "Microsoft.ClassicCompute/virtualMachines"
	ProtectableContainerTypeMicrosoftPointComputeVirtualMachines        ProtectableContainerType = "Microsoft.Compute/virtualMachines"
	ProtectableContainerTypeSQLAGWorkLoadContainer                      ProtectableContainerType = "SQLAGWorkLoadContainer"
	ProtectableContainerTypeStorageContainer                            ProtectableContainerType = "StorageContainer"
	ProtectableContainerTypeUnknown                                     ProtectableContainerType = "Unknown"
	ProtectableContainerTypeVCenter                                     ProtectableContainerType = "VCenter"
	ProtectableContainerTypeVMAppContainer                              ProtectableContainerType = "VMAppContainer"
	ProtectableContainerTypeWindows                                     ProtectableContainerType = "Windows"
)

func PossibleValuesForProtectableContainerType() []string {
	return []string{
		string(ProtectableContainerTypeAzureBackupServerContainer),
		string(ProtectableContainerTypeAzureSqlContainer),
		string(ProtectableContainerTypeAzureWorkloadContainer),
		string(ProtectableContainerTypeCluster),
		string(ProtectableContainerTypeDPMContainer),
		string(ProtectableContainerTypeGenericContainer),
		string(ProtectableContainerTypeIaasVMContainer),
		string(ProtectableContainerTypeIaasVMServiceContainer),
		string(ProtectableContainerTypeInvalid),
		string(ProtectableContainerTypeMABContainer),
		string(ProtectableContainerTypeMicrosoftPointClassicComputeVirtualMachines),
		string(ProtectableContainerTypeMicrosoftPointComputeVirtualMachines),
		string(ProtectableContainerTypeSQLAGWorkLoadContainer),
		string(ProtectableContainerTypeStorageContainer),
		string(ProtectableContainerTypeUnknown),
		string(ProtectableContainerTypeVCenter),
		string(ProtectableContainerTypeVMAppContainer),
		string(ProtectableContainerTypeWindows),
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
