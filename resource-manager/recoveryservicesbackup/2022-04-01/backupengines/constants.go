package backupengines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupEngineType string

const (
	BackupEngineTypeAzureBackupServerEngine BackupEngineType = "AzureBackupServerEngine"
	BackupEngineTypeDpmBackupEngine         BackupEngineType = "DpmBackupEngine"
	BackupEngineTypeInvalid                 BackupEngineType = "Invalid"
)

func PossibleValuesForBackupEngineType() []string {
	return []string{
		string(BackupEngineTypeAzureBackupServerEngine),
		string(BackupEngineTypeDpmBackupEngine),
		string(BackupEngineTypeInvalid),
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
