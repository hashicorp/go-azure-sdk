package protectionpolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DayOfWeek string

const (
	DayOfWeekFriday    DayOfWeek = "Friday"
	DayOfWeekMonday    DayOfWeek = "Monday"
	DayOfWeekSaturday  DayOfWeek = "Saturday"
	DayOfWeekSunday    DayOfWeek = "Sunday"
	DayOfWeekThursday  DayOfWeek = "Thursday"
	DayOfWeekTuesday   DayOfWeek = "Tuesday"
	DayOfWeekWednesday DayOfWeek = "Wednesday"
)

func PossibleValuesForDayOfWeek() []string {
	return []string{
		string(DayOfWeekFriday),
		string(DayOfWeekMonday),
		string(DayOfWeekSaturday),
		string(DayOfWeekSunday),
		string(DayOfWeekThursday),
		string(DayOfWeekTuesday),
		string(DayOfWeekWednesday),
	}
}

type IAASVMPolicyType string

const (
	IAASVMPolicyTypeInvalid IAASVMPolicyType = "Invalid"
	IAASVMPolicyTypeVOne    IAASVMPolicyType = "V1"
	IAASVMPolicyTypeVTwo    IAASVMPolicyType = "V2"
)

func PossibleValuesForIAASVMPolicyType() []string {
	return []string{
		string(IAASVMPolicyTypeInvalid),
		string(IAASVMPolicyTypeVOne),
		string(IAASVMPolicyTypeVTwo),
	}
}

type MonthOfYear string

const (
	MonthOfYearApril     MonthOfYear = "April"
	MonthOfYearAugust    MonthOfYear = "August"
	MonthOfYearDecember  MonthOfYear = "December"
	MonthOfYearFebruary  MonthOfYear = "February"
	MonthOfYearInvalid   MonthOfYear = "Invalid"
	MonthOfYearJanuary   MonthOfYear = "January"
	MonthOfYearJuly      MonthOfYear = "July"
	MonthOfYearJune      MonthOfYear = "June"
	MonthOfYearMarch     MonthOfYear = "March"
	MonthOfYearMay       MonthOfYear = "May"
	MonthOfYearNovember  MonthOfYear = "November"
	MonthOfYearOctober   MonthOfYear = "October"
	MonthOfYearSeptember MonthOfYear = "September"
)

func PossibleValuesForMonthOfYear() []string {
	return []string{
		string(MonthOfYearApril),
		string(MonthOfYearAugust),
		string(MonthOfYearDecember),
		string(MonthOfYearFebruary),
		string(MonthOfYearInvalid),
		string(MonthOfYearJanuary),
		string(MonthOfYearJuly),
		string(MonthOfYearJune),
		string(MonthOfYearMarch),
		string(MonthOfYearMay),
		string(MonthOfYearNovember),
		string(MonthOfYearOctober),
		string(MonthOfYearSeptember),
	}
}

type PolicyType string

const (
	PolicyTypeCopyOnlyFull PolicyType = "CopyOnlyFull"
	PolicyTypeDifferential PolicyType = "Differential"
	PolicyTypeFull         PolicyType = "Full"
	PolicyTypeIncremental  PolicyType = "Incremental"
	PolicyTypeInvalid      PolicyType = "Invalid"
	PolicyTypeLog          PolicyType = "Log"
)

func PossibleValuesForPolicyType() []string {
	return []string{
		string(PolicyTypeCopyOnlyFull),
		string(PolicyTypeDifferential),
		string(PolicyTypeFull),
		string(PolicyTypeIncremental),
		string(PolicyTypeInvalid),
		string(PolicyTypeLog),
	}
}

type RetentionDurationType string

const (
	RetentionDurationTypeDays    RetentionDurationType = "Days"
	RetentionDurationTypeInvalid RetentionDurationType = "Invalid"
	RetentionDurationTypeMonths  RetentionDurationType = "Months"
	RetentionDurationTypeWeeks   RetentionDurationType = "Weeks"
	RetentionDurationTypeYears   RetentionDurationType = "Years"
)

func PossibleValuesForRetentionDurationType() []string {
	return []string{
		string(RetentionDurationTypeDays),
		string(RetentionDurationTypeInvalid),
		string(RetentionDurationTypeMonths),
		string(RetentionDurationTypeWeeks),
		string(RetentionDurationTypeYears),
	}
}

type RetentionScheduleFormat string

const (
	RetentionScheduleFormatDaily   RetentionScheduleFormat = "Daily"
	RetentionScheduleFormatInvalid RetentionScheduleFormat = "Invalid"
	RetentionScheduleFormatWeekly  RetentionScheduleFormat = "Weekly"
)

func PossibleValuesForRetentionScheduleFormat() []string {
	return []string{
		string(RetentionScheduleFormatDaily),
		string(RetentionScheduleFormatInvalid),
		string(RetentionScheduleFormatWeekly),
	}
}

type ScheduleRunType string

const (
	ScheduleRunTypeDaily   ScheduleRunType = "Daily"
	ScheduleRunTypeHourly  ScheduleRunType = "Hourly"
	ScheduleRunTypeInvalid ScheduleRunType = "Invalid"
	ScheduleRunTypeWeekly  ScheduleRunType = "Weekly"
)

func PossibleValuesForScheduleRunType() []string {
	return []string{
		string(ScheduleRunTypeDaily),
		string(ScheduleRunTypeHourly),
		string(ScheduleRunTypeInvalid),
		string(ScheduleRunTypeWeekly),
	}
}

type WeekOfMonth string

const (
	WeekOfMonthFirst   WeekOfMonth = "First"
	WeekOfMonthFourth  WeekOfMonth = "Fourth"
	WeekOfMonthInvalid WeekOfMonth = "Invalid"
	WeekOfMonthLast    WeekOfMonth = "Last"
	WeekOfMonthSecond  WeekOfMonth = "Second"
	WeekOfMonthThird   WeekOfMonth = "Third"
)

func PossibleValuesForWeekOfMonth() []string {
	return []string{
		string(WeekOfMonthFirst),
		string(WeekOfMonthFourth),
		string(WeekOfMonthInvalid),
		string(WeekOfMonthLast),
		string(WeekOfMonthSecond),
		string(WeekOfMonthThird),
	}
}

type WorkloadType string

const (
	WorkloadTypeAzureFileShare    WorkloadType = "AzureFileShare"
	WorkloadTypeAzureSqlDb        WorkloadType = "AzureSqlDb"
	WorkloadTypeClient            WorkloadType = "Client"
	WorkloadTypeExchange          WorkloadType = "Exchange"
	WorkloadTypeFileFolder        WorkloadType = "FileFolder"
	WorkloadTypeGenericDataSource WorkloadType = "GenericDataSource"
	WorkloadTypeInvalid           WorkloadType = "Invalid"
	WorkloadTypeSAPAseDatabase    WorkloadType = "SAPAseDatabase"
	WorkloadTypeSAPHanaDatabase   WorkloadType = "SAPHanaDatabase"
	WorkloadTypeSQLDB             WorkloadType = "SQLDB"
	WorkloadTypeSQLDataBase       WorkloadType = "SQLDataBase"
	WorkloadTypeSharepoint        WorkloadType = "Sharepoint"
	WorkloadTypeSystemState       WorkloadType = "SystemState"
	WorkloadTypeVM                WorkloadType = "VM"
	WorkloadTypeVMwareVM          WorkloadType = "VMwareVM"
)

func PossibleValuesForWorkloadType() []string {
	return []string{
		string(WorkloadTypeAzureFileShare),
		string(WorkloadTypeAzureSqlDb),
		string(WorkloadTypeClient),
		string(WorkloadTypeExchange),
		string(WorkloadTypeFileFolder),
		string(WorkloadTypeGenericDataSource),
		string(WorkloadTypeInvalid),
		string(WorkloadTypeSAPAseDatabase),
		string(WorkloadTypeSAPHanaDatabase),
		string(WorkloadTypeSQLDB),
		string(WorkloadTypeSQLDataBase),
		string(WorkloadTypeSharepoint),
		string(WorkloadTypeSystemState),
		string(WorkloadTypeVM),
		string(WorkloadTypeVMwareVM),
	}
}
