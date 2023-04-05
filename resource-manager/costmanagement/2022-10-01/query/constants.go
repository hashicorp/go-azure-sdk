package query

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExportType string

const (
	ExportTypeActualCost    ExportType = "ActualCost"
	ExportTypeAmortizedCost ExportType = "AmortizedCost"
	ExportTypeUsage         ExportType = "Usage"
)

func PossibleValuesForExportType() []string {
	return []string{
		string(ExportTypeActualCost),
		string(ExportTypeAmortizedCost),
		string(ExportTypeUsage),
	}
}

type ExternalCloudProviderType string

const (
	ExternalCloudProviderTypeExternalBillingAccounts ExternalCloudProviderType = "externalBillingAccounts"
	ExternalCloudProviderTypeExternalSubscriptions   ExternalCloudProviderType = "externalSubscriptions"
)

func PossibleValuesForExternalCloudProviderType() []string {
	return []string{
		string(ExternalCloudProviderTypeExternalBillingAccounts),
		string(ExternalCloudProviderTypeExternalSubscriptions),
	}
}

type FunctionType string

const (
	FunctionTypeSum FunctionType = "Sum"
)

func PossibleValuesForFunctionType() []string {
	return []string{
		string(FunctionTypeSum),
	}
}

type GranularityType string

const (
	GranularityTypeDaily GranularityType = "Daily"
)

func PossibleValuesForGranularityType() []string {
	return []string{
		string(GranularityTypeDaily),
	}
}

type QueryColumnType string

const (
	QueryColumnTypeDimension QueryColumnType = "Dimension"
	QueryColumnTypeTagKey    QueryColumnType = "TagKey"
)

func PossibleValuesForQueryColumnType() []string {
	return []string{
		string(QueryColumnTypeDimension),
		string(QueryColumnTypeTagKey),
	}
}

type QueryOperatorType string

const (
	QueryOperatorTypeIn QueryOperatorType = "In"
)

func PossibleValuesForQueryOperatorType() []string {
	return []string{
		string(QueryOperatorTypeIn),
	}
}

type TimeframeType string

const (
	TimeframeTypeBillingMonthToDate  TimeframeType = "BillingMonthToDate"
	TimeframeTypeCustom              TimeframeType = "Custom"
	TimeframeTypeMonthToDate         TimeframeType = "MonthToDate"
	TimeframeTypeTheLastBillingMonth TimeframeType = "TheLastBillingMonth"
	TimeframeTypeTheLastMonth        TimeframeType = "TheLastMonth"
	TimeframeTypeWeekToDate          TimeframeType = "WeekToDate"
)

func PossibleValuesForTimeframeType() []string {
	return []string{
		string(TimeframeTypeBillingMonthToDate),
		string(TimeframeTypeCustom),
		string(TimeframeTypeMonthToDate),
		string(TimeframeTypeTheLastBillingMonth),
		string(TimeframeTypeTheLastMonth),
		string(TimeframeTypeWeekToDate),
	}
}
