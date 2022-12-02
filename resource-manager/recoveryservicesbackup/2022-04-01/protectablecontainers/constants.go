package protectablecontainers

import "strings"

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

func parseBackupManagementType(input string) (*BackupManagementType, error) {
	vals := map[string]BackupManagementType{
		"azurebackupserver": BackupManagementTypeAzureBackupServer,
		"azureiaasvm":       BackupManagementTypeAzureIaasVM,
		"azuresql":          BackupManagementTypeAzureSql,
		"azurestorage":      BackupManagementTypeAzureStorage,
		"azureworkload":     BackupManagementTypeAzureWorkload,
		"dpm":               BackupManagementTypeDPM,
		"defaultbackup":     BackupManagementTypeDefaultBackup,
		"invalid":           BackupManagementTypeInvalid,
		"mab":               BackupManagementTypeMAB,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BackupManagementType(input)
	return &out, nil
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

func parseContainerType(input string) (*ContainerType, error) {
	vals := map[string]ContainerType{
		"azurebackupservercontainer": ContainerTypeAzureBackupServerContainer,
		"azuresqlcontainer":          ContainerTypeAzureSqlContainer,
		"azureworkloadcontainer":     ContainerTypeAzureWorkloadContainer,
		"cluster":                    ContainerTypeCluster,
		"dpmcontainer":               ContainerTypeDPMContainer,
		"genericcontainer":           ContainerTypeGenericContainer,
		"iaasvmcontainer":            ContainerTypeIaasVMContainer,
		"iaasvmservicecontainer":     ContainerTypeIaasVMServiceContainer,
		"invalid":                    ContainerTypeInvalid,
		"mabcontainer":               ContainerTypeMABContainer,
		"microsoft.classiccompute/virtualmachines": ContainerTypeMicrosoftPointClassicComputeVirtualMachines,
		"microsoft.compute/virtualmachines":        ContainerTypeMicrosoftPointComputeVirtualMachines,
		"sqlagworkloadcontainer":                   ContainerTypeSQLAGWorkLoadContainer,
		"storagecontainer":                         ContainerTypeStorageContainer,
		"unknown":                                  ContainerTypeUnknown,
		"vcenter":                                  ContainerTypeVCenter,
		"vmappcontainer":                           ContainerTypeVMAppContainer,
		"windows":                                  ContainerTypeWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContainerType(input)
	return &out, nil
}
