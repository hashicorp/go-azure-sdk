package protectionintent

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
	DataSourceTypeSAPHanaDBInstance DataSourceType = "SAPHanaDBInstance"
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
		string(DataSourceTypeSAPHanaDBInstance),
		string(DataSourceTypeSAPHanaDatabase),
		string(DataSourceTypeSQLDB),
		string(DataSourceTypeSQLDataBase),
		string(DataSourceTypeSharepoint),
		string(DataSourceTypeSystemState),
		string(DataSourceTypeVM),
		string(DataSourceTypeVMwareVM),
	}
}

func parseDataSourceType(input string) (*DataSourceType, error) {
	vals := map[string]DataSourceType{
		"azurefileshare":    DataSourceTypeAzureFileShare,
		"azuresqldb":        DataSourceTypeAzureSqlDb,
		"client":            DataSourceTypeClient,
		"exchange":          DataSourceTypeExchange,
		"filefolder":        DataSourceTypeFileFolder,
		"genericdatasource": DataSourceTypeGenericDataSource,
		"invalid":           DataSourceTypeInvalid,
		"sapasedatabase":    DataSourceTypeSAPAseDatabase,
		"saphanadbinstance": DataSourceTypeSAPHanaDBInstance,
		"saphanadatabase":   DataSourceTypeSAPHanaDatabase,
		"sqldb":             DataSourceTypeSQLDB,
		"sqldatabase":       DataSourceTypeSQLDataBase,
		"sharepoint":        DataSourceTypeSharepoint,
		"systemstate":       DataSourceTypeSystemState,
		"vm":                DataSourceTypeVM,
		"vmwarevm":          DataSourceTypeVMwareVM,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataSourceType(input)
	return &out, nil
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

func parseProtectionIntentItemType(input string) (*ProtectionIntentItemType, error) {
	vals := map[string]ProtectionIntentItemType{
		"azureresourceitem":                          ProtectionIntentItemTypeAzureResourceItem,
		"azureworkloadautoprotectionintent":          ProtectionIntentItemTypeAzureWorkloadAutoProtectionIntent,
		"azureworkloadcontainerautoprotectionintent": ProtectionIntentItemTypeAzureWorkloadContainerAutoProtectionIntent,
		"azureworkloadsqlautoprotectionintent":       ProtectionIntentItemTypeAzureWorkloadSQLAutoProtectionIntent,
		"invalid":                                    ProtectionIntentItemTypeInvalid,
		"recoveryservicevaultitem":                   ProtectionIntentItemTypeRecoveryServiceVaultItem,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtectionIntentItemType(input)
	return &out, nil
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

func parseProtectionStatus(input string) (*ProtectionStatus, error) {
	vals := map[string]ProtectionStatus{
		"invalid":          ProtectionStatusInvalid,
		"notprotected":     ProtectionStatusNotProtected,
		"protected":        ProtectionStatusProtected,
		"protecting":       ProtectionStatusProtecting,
		"protectionfailed": ProtectionStatusProtectionFailed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtectionStatus(input)
	return &out, nil
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

func parseValidationStatus(input string) (*ValidationStatus, error) {
	vals := map[string]ValidationStatus{
		"failed":    ValidationStatusFailed,
		"invalid":   ValidationStatusInvalid,
		"succeeded": ValidationStatusSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ValidationStatus(input)
	return &out, nil
}
