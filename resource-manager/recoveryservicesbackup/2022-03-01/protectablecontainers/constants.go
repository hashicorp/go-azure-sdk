package protectablecontainers

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
