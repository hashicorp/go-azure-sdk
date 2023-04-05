package exports

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExecutionStatus string

const (
	ExecutionStatusCompleted           ExecutionStatus = "Completed"
	ExecutionStatusDataNotAvailable    ExecutionStatus = "DataNotAvailable"
	ExecutionStatusFailed              ExecutionStatus = "Failed"
	ExecutionStatusInProgress          ExecutionStatus = "InProgress"
	ExecutionStatusNewDataNotAvailable ExecutionStatus = "NewDataNotAvailable"
	ExecutionStatusQueued              ExecutionStatus = "Queued"
	ExecutionStatusTimeout             ExecutionStatus = "Timeout"
)

func PossibleValuesForExecutionStatus() []string {
	return []string{
		string(ExecutionStatusCompleted),
		string(ExecutionStatusDataNotAvailable),
		string(ExecutionStatusFailed),
		string(ExecutionStatusInProgress),
		string(ExecutionStatusNewDataNotAvailable),
		string(ExecutionStatusQueued),
		string(ExecutionStatusTimeout),
	}
}

type ExecutionType string

const (
	ExecutionTypeOnDemand  ExecutionType = "OnDemand"
	ExecutionTypeScheduled ExecutionType = "Scheduled"
)

func PossibleValuesForExecutionType() []string {
	return []string{
		string(ExecutionTypeOnDemand),
		string(ExecutionTypeScheduled),
	}
}

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

type FormatType string

const (
	FormatTypeCsv FormatType = "Csv"
)

func PossibleValuesForFormatType() []string {
	return []string{
		string(FormatTypeCsv),
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

type RecurrenceType string

const (
	RecurrenceTypeAnnually RecurrenceType = "Annually"
	RecurrenceTypeDaily    RecurrenceType = "Daily"
	RecurrenceTypeMonthly  RecurrenceType = "Monthly"
	RecurrenceTypeWeekly   RecurrenceType = "Weekly"
)

func PossibleValuesForRecurrenceType() []string {
	return []string{
		string(RecurrenceTypeAnnually),
		string(RecurrenceTypeDaily),
		string(RecurrenceTypeMonthly),
		string(RecurrenceTypeWeekly),
	}
}

type StatusType string

const (
	StatusTypeActive   StatusType = "Active"
	StatusTypeInactive StatusType = "Inactive"
)

func PossibleValuesForStatusType() []string {
	return []string{
		string(StatusTypeActive),
		string(StatusTypeInactive),
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
