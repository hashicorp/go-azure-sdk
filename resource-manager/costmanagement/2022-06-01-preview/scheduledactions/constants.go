package scheduledactions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckNameAvailabilityReason string

const (
	CheckNameAvailabilityReasonAlreadyExists CheckNameAvailabilityReason = "AlreadyExists"
	CheckNameAvailabilityReasonInvalid       CheckNameAvailabilityReason = "Invalid"
)

func PossibleValuesForCheckNameAvailabilityReason() []string {
	return []string{
		string(CheckNameAvailabilityReasonAlreadyExists),
		string(CheckNameAvailabilityReasonInvalid),
	}
}

type DaysOfWeek string

const (
	DaysOfWeekFriday    DaysOfWeek = "Friday"
	DaysOfWeekMonday    DaysOfWeek = "Monday"
	DaysOfWeekSaturday  DaysOfWeek = "Saturday"
	DaysOfWeekSunday    DaysOfWeek = "Sunday"
	DaysOfWeekThursday  DaysOfWeek = "Thursday"
	DaysOfWeekTuesday   DaysOfWeek = "Tuesday"
	DaysOfWeekWednesday DaysOfWeek = "Wednesday"
)

func PossibleValuesForDaysOfWeek() []string {
	return []string{
		string(DaysOfWeekFriday),
		string(DaysOfWeekMonday),
		string(DaysOfWeekSaturday),
		string(DaysOfWeekSunday),
		string(DaysOfWeekThursday),
		string(DaysOfWeekTuesday),
		string(DaysOfWeekWednesday),
	}
}

type FileFormat string

const (
	FileFormatCsv FileFormat = "Csv"
)

func PossibleValuesForFileFormat() []string {
	return []string{
		string(FileFormatCsv),
	}
}

type ScheduleFrequency string

const (
	ScheduleFrequencyDaily   ScheduleFrequency = "Daily"
	ScheduleFrequencyMonthly ScheduleFrequency = "Monthly"
	ScheduleFrequencyWeekly  ScheduleFrequency = "Weekly"
)

func PossibleValuesForScheduleFrequency() []string {
	return []string{
		string(ScheduleFrequencyDaily),
		string(ScheduleFrequencyMonthly),
		string(ScheduleFrequencyWeekly),
	}
}

type ScheduledActionKind string

const (
	ScheduledActionKindEmail        ScheduledActionKind = "Email"
	ScheduledActionKindInsightAlert ScheduledActionKind = "InsightAlert"
)

func PossibleValuesForScheduledActionKind() []string {
	return []string{
		string(ScheduledActionKindEmail),
		string(ScheduledActionKindInsightAlert),
	}
}

type ScheduledActionStatus string

const (
	ScheduledActionStatusDisabled ScheduledActionStatus = "Disabled"
	ScheduledActionStatusEnabled  ScheduledActionStatus = "Enabled"
)

func PossibleValuesForScheduledActionStatus() []string {
	return []string{
		string(ScheduledActionStatusDisabled),
		string(ScheduledActionStatusEnabled),
	}
}

type WeeksOfMonth string

const (
	WeeksOfMonthFirst  WeeksOfMonth = "First"
	WeeksOfMonthFourth WeeksOfMonth = "Fourth"
	WeeksOfMonthLast   WeeksOfMonth = "Last"
	WeeksOfMonthSecond WeeksOfMonth = "Second"
	WeeksOfMonthThird  WeeksOfMonth = "Third"
)

func PossibleValuesForWeeksOfMonth() []string {
	return []string{
		string(WeeksOfMonthFirst),
		string(WeeksOfMonthFourth),
		string(WeeksOfMonthLast),
		string(WeeksOfMonthSecond),
		string(WeeksOfMonthThird),
	}
}
