package maintenanceconfigurations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Type string

const (
	TypeFirst  Type = "First"
	TypeFourth Type = "Fourth"
	TypeLast   Type = "Last"
	TypeSecond Type = "Second"
	TypeThird  Type = "Third"
)

func PossibleValuesForType() []string {
	return []string{
		string(TypeFirst),
		string(TypeFourth),
		string(TypeLast),
		string(TypeSecond),
		string(TypeThird),
	}
}

type WeekDay string

const (
	WeekDayFriday    WeekDay = "Friday"
	WeekDayMonday    WeekDay = "Monday"
	WeekDaySaturday  WeekDay = "Saturday"
	WeekDaySunday    WeekDay = "Sunday"
	WeekDayThursday  WeekDay = "Thursday"
	WeekDayTuesday   WeekDay = "Tuesday"
	WeekDayWednesday WeekDay = "Wednesday"
)

func PossibleValuesForWeekDay() []string {
	return []string{
		string(WeekDayFriday),
		string(WeekDayMonday),
		string(WeekDaySaturday),
		string(WeekDaySunday),
		string(WeekDayThursday),
		string(WeekDayTuesday),
		string(WeekDayWednesday),
	}
}
