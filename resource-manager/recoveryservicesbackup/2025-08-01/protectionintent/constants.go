package protectionintent

import (
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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
