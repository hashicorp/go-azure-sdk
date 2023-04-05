package tables

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ColumnDataTypeHintEnum string

const (
	ColumnDataTypeHintEnumArmPath ColumnDataTypeHintEnum = "armPath"
	ColumnDataTypeHintEnumGuid    ColumnDataTypeHintEnum = "guid"
	ColumnDataTypeHintEnumIP      ColumnDataTypeHintEnum = "ip"
	ColumnDataTypeHintEnumUri     ColumnDataTypeHintEnum = "uri"
)

func PossibleValuesForColumnDataTypeHintEnum() []string {
	return []string{
		string(ColumnDataTypeHintEnumArmPath),
		string(ColumnDataTypeHintEnumGuid),
		string(ColumnDataTypeHintEnumIP),
		string(ColumnDataTypeHintEnumUri),
	}
}

type ColumnTypeEnum string

const (
	ColumnTypeEnumBoolean  ColumnTypeEnum = "boolean"
	ColumnTypeEnumDateTime ColumnTypeEnum = "dateTime"
	ColumnTypeEnumDynamic  ColumnTypeEnum = "dynamic"
	ColumnTypeEnumGuid     ColumnTypeEnum = "guid"
	ColumnTypeEnumInt      ColumnTypeEnum = "int"
	ColumnTypeEnumLong     ColumnTypeEnum = "long"
	ColumnTypeEnumReal     ColumnTypeEnum = "real"
	ColumnTypeEnumString   ColumnTypeEnum = "string"
)

func PossibleValuesForColumnTypeEnum() []string {
	return []string{
		string(ColumnTypeEnumBoolean),
		string(ColumnTypeEnumDateTime),
		string(ColumnTypeEnumDynamic),
		string(ColumnTypeEnumGuid),
		string(ColumnTypeEnumInt),
		string(ColumnTypeEnumLong),
		string(ColumnTypeEnumReal),
		string(ColumnTypeEnumString),
	}
}

type ProvisioningStateEnum string

const (
	ProvisioningStateEnumDeleting   ProvisioningStateEnum = "Deleting"
	ProvisioningStateEnumInProgress ProvisioningStateEnum = "InProgress"
	ProvisioningStateEnumSucceeded  ProvisioningStateEnum = "Succeeded"
	ProvisioningStateEnumUpdating   ProvisioningStateEnum = "Updating"
)

func PossibleValuesForProvisioningStateEnum() []string {
	return []string{
		string(ProvisioningStateEnumDeleting),
		string(ProvisioningStateEnumInProgress),
		string(ProvisioningStateEnumSucceeded),
		string(ProvisioningStateEnumUpdating),
	}
}

type SourceEnum string

const (
	SourceEnumCustomer  SourceEnum = "customer"
	SourceEnumMicrosoft SourceEnum = "microsoft"
)

func PossibleValuesForSourceEnum() []string {
	return []string{
		string(SourceEnumCustomer),
		string(SourceEnumMicrosoft),
	}
}

type TablePlanEnum string

const (
	TablePlanEnumAnalytics TablePlanEnum = "Analytics"
	TablePlanEnumBasic     TablePlanEnum = "Basic"
)

func PossibleValuesForTablePlanEnum() []string {
	return []string{
		string(TablePlanEnumAnalytics),
		string(TablePlanEnumBasic),
	}
}

type TableSubTypeEnum string

const (
	TableSubTypeEnumAny                     TableSubTypeEnum = "Any"
	TableSubTypeEnumClassic                 TableSubTypeEnum = "Classic"
	TableSubTypeEnumDataCollectionRuleBased TableSubTypeEnum = "DataCollectionRuleBased"
)

func PossibleValuesForTableSubTypeEnum() []string {
	return []string{
		string(TableSubTypeEnumAny),
		string(TableSubTypeEnumClassic),
		string(TableSubTypeEnumDataCollectionRuleBased),
	}
}

type TableTypeEnum string

const (
	TableTypeEnumCustomLog     TableTypeEnum = "CustomLog"
	TableTypeEnumMicrosoft     TableTypeEnum = "Microsoft"
	TableTypeEnumRestoredLogs  TableTypeEnum = "RestoredLogs"
	TableTypeEnumSearchResults TableTypeEnum = "SearchResults"
)

func PossibleValuesForTableTypeEnum() []string {
	return []string{
		string(TableTypeEnumCustomLog),
		string(TableTypeEnumMicrosoft),
		string(TableTypeEnumRestoredLogs),
		string(TableTypeEnumSearchResults),
	}
}
