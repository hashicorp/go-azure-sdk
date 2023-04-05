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
