package backupstatus

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
