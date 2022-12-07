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

func parseProtectableContainerType(input string) (*ProtectableContainerType, error) {
	vals := map[string]ProtectableContainerType{
		"azurebackupservercontainer": ProtectableContainerTypeAzureBackupServerContainer,
		"azuresqlcontainer":          ProtectableContainerTypeAzureSqlContainer,
		"azureworkloadcontainer":     ProtectableContainerTypeAzureWorkloadContainer,
		"cluster":                    ProtectableContainerTypeCluster,
		"dpmcontainer":               ProtectableContainerTypeDPMContainer,
		"genericcontainer":           ProtectableContainerTypeGenericContainer,
		"iaasvmcontainer":            ProtectableContainerTypeIaasVMContainer,
		"iaasvmservicecontainer":     ProtectableContainerTypeIaasVMServiceContainer,
		"invalid":                    ProtectableContainerTypeInvalid,
		"mabcontainer":               ProtectableContainerTypeMABContainer,
		"microsoft.classiccompute/virtualmachines": ProtectableContainerTypeMicrosoftPointClassicComputeVirtualMachines,
		"microsoft.compute/virtualmachines":        ProtectableContainerTypeMicrosoftPointComputeVirtualMachines,
		"sqlagworkloadcontainer":                   ProtectableContainerTypeSQLAGWorkLoadContainer,
		"storagecontainer":                         ProtectableContainerTypeStorageContainer,
		"unknown":                                  ProtectableContainerTypeUnknown,
		"vcenter":                                  ProtectableContainerTypeVCenter,
		"vmappcontainer":                           ProtectableContainerTypeVMAppContainer,
		"windows":                                  ProtectableContainerTypeWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtectableContainerType(input)
	return &out, nil
}
