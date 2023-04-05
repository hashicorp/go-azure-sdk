package schedule

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduleDay string

const (
	ScheduleDayFriday    ScheduleDay = "Friday"
	ScheduleDayMonday    ScheduleDay = "Monday"
	ScheduleDaySaturday  ScheduleDay = "Saturday"
	ScheduleDaySunday    ScheduleDay = "Sunday"
	ScheduleDayThursday  ScheduleDay = "Thursday"
	ScheduleDayTuesday   ScheduleDay = "Tuesday"
	ScheduleDayWednesday ScheduleDay = "Wednesday"
)

func PossibleValuesForScheduleDay() []string {
	return []string{
		string(ScheduleDayFriday),
		string(ScheduleDayMonday),
		string(ScheduleDaySaturday),
		string(ScheduleDaySunday),
		string(ScheduleDayThursday),
		string(ScheduleDayTuesday),
		string(ScheduleDayWednesday),
	}
}

type ScheduleFrequency string

const (
	ScheduleFrequencyDay     ScheduleFrequency = "Day"
	ScheduleFrequencyHour    ScheduleFrequency = "Hour"
	ScheduleFrequencyMinute  ScheduleFrequency = "Minute"
	ScheduleFrequencyMonth   ScheduleFrequency = "Month"
	ScheduleFrequencyOneTime ScheduleFrequency = "OneTime"
	ScheduleFrequencyWeek    ScheduleFrequency = "Week"
)

func PossibleValuesForScheduleFrequency() []string {
	return []string{
		string(ScheduleFrequencyDay),
		string(ScheduleFrequencyHour),
		string(ScheduleFrequencyMinute),
		string(ScheduleFrequencyMonth),
		string(ScheduleFrequencyOneTime),
		string(ScheduleFrequencyWeek),
	}
}
