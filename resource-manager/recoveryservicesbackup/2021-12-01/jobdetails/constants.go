package jobdetails

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

func parseJobSupportedAction(input string) (*JobSupportedAction, error) {
	vals := map[string]JobSupportedAction{
		"cancellable": JobSupportedActionCancellable,
		"invalid":     JobSupportedActionInvalid,
		"retriable":   JobSupportedActionRetriable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JobSupportedAction(input)
	return &out, nil
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

func parseMabServerType(input string) (*MabServerType, error) {
	vals := map[string]MabServerType{
		"azurebackupservercontainer": MabServerTypeAzureBackupServerContainer,
		"azuresqlcontainer":          MabServerTypeAzureSqlContainer,
		"cluster":                    MabServerTypeCluster,
		"dpmcontainer":               MabServerTypeDPMContainer,
		"genericcontainer":           MabServerTypeGenericContainer,
		"iaasvmcontainer":            MabServerTypeIaasVMContainer,
		"iaasvmservicecontainer":     MabServerTypeIaasVMServiceContainer,
		"invalid":                    MabServerTypeInvalid,
		"mabcontainer":               MabServerTypeMABContainer,
		"sqlagworkloadcontainer":     MabServerTypeSQLAGWorkLoadContainer,
		"storagecontainer":           MabServerTypeStorageContainer,
		"unknown":                    MabServerTypeUnknown,
		"vcenter":                    MabServerTypeVCenter,
		"vmappcontainer":             MabServerTypeVMAppContainer,
		"windows":                    MabServerTypeWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MabServerType(input)
	return &out, nil
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

func parseWorkloadType(input string) (*WorkloadType, error) {
	vals := map[string]WorkloadType{
		"azurefileshare":    WorkloadTypeAzureFileShare,
		"azuresqldb":        WorkloadTypeAzureSqlDb,
		"client":            WorkloadTypeClient,
		"exchange":          WorkloadTypeExchange,
		"filefolder":        WorkloadTypeFileFolder,
		"genericdatasource": WorkloadTypeGenericDataSource,
		"invalid":           WorkloadTypeInvalid,
		"sapasedatabase":    WorkloadTypeSAPAseDatabase,
		"saphanadatabase":   WorkloadTypeSAPHanaDatabase,
		"sqldb":             WorkloadTypeSQLDB,
		"sqldatabase":       WorkloadTypeSQLDataBase,
		"sharepoint":        WorkloadTypeSharepoint,
		"systemstate":       WorkloadTypeSystemState,
		"vm":                WorkloadTypeVM,
		"vmwarevm":          WorkloadTypeVMwareVM,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkloadType(input)
	return &out, nil
}
