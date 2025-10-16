package backupstatus

import (
	"strings"
)

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

func parseAcquireStorageAccountLock(input string) (*AcquireStorageAccountLock, error) {
	vals := map[string]AcquireStorageAccountLock{
		"acquire":    AcquireStorageAccountLockAcquire,
		"notacquire": AcquireStorageAccountLockNotAcquire,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AcquireStorageAccountLock(input)
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

type FabricName string

const (
	FabricNameAzure   FabricName = "Azure"
	FabricNameInvalid FabricName = "Invalid"
)

func PossibleValuesForFabricName() []string {
	return []string{
		string(FabricNameAzure),
		string(FabricNameInvalid),
	}
}

func parseFabricName(input string) (*FabricName, error) {
	vals := map[string]FabricName{
		"azure":   FabricNameAzure,
		"invalid": FabricNameInvalid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FabricName(input)
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
