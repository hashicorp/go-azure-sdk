package protectionintent

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

type ProtectionIntentItemType string

const (
	ProtectionIntentItemTypeAzureResourceItem                          ProtectionIntentItemType = "AzureResourceItem"
	ProtectionIntentItemTypeAzureWorkloadAutoProtectionIntent          ProtectionIntentItemType = "AzureWorkloadAutoProtectionIntent"
	ProtectionIntentItemTypeAzureWorkloadContainerAutoProtectionIntent ProtectionIntentItemType = "AzureWorkloadContainerAutoProtectionIntent"
	ProtectionIntentItemTypeAzureWorkloadSQLAutoProtectionIntent       ProtectionIntentItemType = "AzureWorkloadSQLAutoProtectionIntent"
	ProtectionIntentItemTypeInvalid                                    ProtectionIntentItemType = "Invalid"
	ProtectionIntentItemTypeRecoveryServiceVaultItem                   ProtectionIntentItemType = "RecoveryServiceVaultItem"
)

func PossibleValuesForProtectionIntentItemType() []string {
	return []string{
		string(ProtectionIntentItemTypeAzureResourceItem),
		string(ProtectionIntentItemTypeAzureWorkloadAutoProtectionIntent),
		string(ProtectionIntentItemTypeAzureWorkloadContainerAutoProtectionIntent),
		string(ProtectionIntentItemTypeAzureWorkloadSQLAutoProtectionIntent),
		string(ProtectionIntentItemTypeInvalid),
		string(ProtectionIntentItemTypeRecoveryServiceVaultItem),
	}
}

type ProtectionStatus string

const (
	ProtectionStatusInvalid          ProtectionStatus = "Invalid"
	ProtectionStatusNotProtected     ProtectionStatus = "NotProtected"
	ProtectionStatusProtected        ProtectionStatus = "Protected"
	ProtectionStatusProtecting       ProtectionStatus = "Protecting"
	ProtectionStatusProtectionFailed ProtectionStatus = "ProtectionFailed"
)

func PossibleValuesForProtectionStatus() []string {
	return []string{
		string(ProtectionStatusInvalid),
		string(ProtectionStatusNotProtected),
		string(ProtectionStatusProtected),
		string(ProtectionStatusProtecting),
		string(ProtectionStatusProtectionFailed),
	}
}

type ValidationStatus string

const (
	ValidationStatusFailed    ValidationStatus = "Failed"
	ValidationStatusInvalid   ValidationStatus = "Invalid"
	ValidationStatusSucceeded ValidationStatus = "Succeeded"
)

func PossibleValuesForValidationStatus() []string {
	return []string{
		string(ValidationStatusFailed),
		string(ValidationStatusInvalid),
		string(ValidationStatusSucceeded),
	}
}
