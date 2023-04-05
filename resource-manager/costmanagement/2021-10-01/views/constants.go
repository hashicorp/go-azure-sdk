package views

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccumulatedType string

const (
	AccumulatedTypeFalse AccumulatedType = "false"
	AccumulatedTypeTrue  AccumulatedType = "true"
)

func PossibleValuesForAccumulatedType() []string {
	return []string{
		string(AccumulatedTypeFalse),
		string(AccumulatedTypeTrue),
	}
}

type ChartType string

const (
	ChartTypeArea          ChartType = "Area"
	ChartTypeGroupedColumn ChartType = "GroupedColumn"
	ChartTypeLine          ChartType = "Line"
	ChartTypeStackedColumn ChartType = "StackedColumn"
	ChartTypeTable         ChartType = "Table"
)

func PossibleValuesForChartType() []string {
	return []string{
		string(ChartTypeArea),
		string(ChartTypeGroupedColumn),
		string(ChartTypeLine),
		string(ChartTypeStackedColumn),
		string(ChartTypeTable),
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

type KpiTypeType string

const (
	KpiTypeTypeBudget   KpiTypeType = "Budget"
	KpiTypeTypeForecast KpiTypeType = "Forecast"
)

func PossibleValuesForKpiTypeType() []string {
	return []string{
		string(KpiTypeTypeBudget),
		string(KpiTypeTypeForecast),
	}
}

type MetricType string

const (
	MetricTypeAHUB          MetricType = "AHUB"
	MetricTypeActualCost    MetricType = "ActualCost"
	MetricTypeAmortizedCost MetricType = "AmortizedCost"
)

func PossibleValuesForMetricType() []string {
	return []string{
		string(MetricTypeAHUB),
		string(MetricTypeActualCost),
		string(MetricTypeAmortizedCost),
	}
}

type OperatorType string

const (
	OperatorTypeContains OperatorType = "Contains"
	OperatorTypeIn       OperatorType = "In"
)

func PossibleValuesForOperatorType() []string {
	return []string{
		string(OperatorTypeContains),
		string(OperatorTypeIn),
	}
}

type PivotTypeType string

const (
	PivotTypeTypeDimension PivotTypeType = "Dimension"
	PivotTypeTypeTagKey    PivotTypeType = "TagKey"
)

func PossibleValuesForPivotTypeType() []string {
	return []string{
		string(PivotTypeTypeDimension),
		string(PivotTypeTypeTagKey),
	}
}

type ReportConfigColumnType string

const (
	ReportConfigColumnTypeDimension ReportConfigColumnType = "Dimension"
	ReportConfigColumnTypeTag       ReportConfigColumnType = "Tag"
)

func PossibleValuesForReportConfigColumnType() []string {
	return []string{
		string(ReportConfigColumnTypeDimension),
		string(ReportConfigColumnTypeTag),
	}
}

type ReportConfigSortingType string

const (
	ReportConfigSortingTypeAscending  ReportConfigSortingType = "Ascending"
	ReportConfigSortingTypeDescending ReportConfigSortingType = "Descending"
)

func PossibleValuesForReportConfigSortingType() []string {
	return []string{
		string(ReportConfigSortingTypeAscending),
		string(ReportConfigSortingTypeDescending),
	}
}

type ReportGranularityType string

const (
	ReportGranularityTypeDaily   ReportGranularityType = "Daily"
	ReportGranularityTypeMonthly ReportGranularityType = "Monthly"
)

func PossibleValuesForReportGranularityType() []string {
	return []string{
		string(ReportGranularityTypeDaily),
		string(ReportGranularityTypeMonthly),
	}
}

type ReportTimeframeType string

const (
	ReportTimeframeTypeCustom      ReportTimeframeType = "Custom"
	ReportTimeframeTypeMonthToDate ReportTimeframeType = "MonthToDate"
	ReportTimeframeTypeWeekToDate  ReportTimeframeType = "WeekToDate"
	ReportTimeframeTypeYearToDate  ReportTimeframeType = "YearToDate"
)

func PossibleValuesForReportTimeframeType() []string {
	return []string{
		string(ReportTimeframeTypeCustom),
		string(ReportTimeframeTypeMonthToDate),
		string(ReportTimeframeTypeWeekToDate),
		string(ReportTimeframeTypeYearToDate),
	}
}

type ReportType string

const (
	ReportTypeUsage ReportType = "Usage"
)

func PossibleValuesForReportType() []string {
	return []string{
		string(ReportTypeUsage),
	}
}
