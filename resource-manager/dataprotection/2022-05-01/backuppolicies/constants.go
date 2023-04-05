package backuppolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AbsoluteMarker string

const (
	AbsoluteMarkerAllBackup    AbsoluteMarker = "AllBackup"
	AbsoluteMarkerFirstOfDay   AbsoluteMarker = "FirstOfDay"
	AbsoluteMarkerFirstOfMonth AbsoluteMarker = "FirstOfMonth"
	AbsoluteMarkerFirstOfWeek  AbsoluteMarker = "FirstOfWeek"
	AbsoluteMarkerFirstOfYear  AbsoluteMarker = "FirstOfYear"
)

func PossibleValuesForAbsoluteMarker() []string {
	return []string{
		string(AbsoluteMarkerAllBackup),
		string(AbsoluteMarkerFirstOfDay),
		string(AbsoluteMarkerFirstOfMonth),
		string(AbsoluteMarkerFirstOfWeek),
		string(AbsoluteMarkerFirstOfYear),
	}
}

type DataStoreTypes string

const (
	DataStoreTypesArchiveStore     DataStoreTypes = "ArchiveStore"
	DataStoreTypesOperationalStore DataStoreTypes = "OperationalStore"
	DataStoreTypesVaultStore       DataStoreTypes = "VaultStore"
)

func PossibleValuesForDataStoreTypes() []string {
	return []string{
		string(DataStoreTypesArchiveStore),
		string(DataStoreTypesOperationalStore),
		string(DataStoreTypesVaultStore),
	}
}

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

type Month string

const (
	MonthApril     Month = "April"
	MonthAugust    Month = "August"
	MonthDecember  Month = "December"
	MonthFebruary  Month = "February"
	MonthJanuary   Month = "January"
	MonthJuly      Month = "July"
	MonthJune      Month = "June"
	MonthMarch     Month = "March"
	MonthMay       Month = "May"
	MonthNovember  Month = "November"
	MonthOctober   Month = "October"
	MonthSeptember Month = "September"
)

func PossibleValuesForMonth() []string {
	return []string{
		string(MonthApril),
		string(MonthAugust),
		string(MonthDecember),
		string(MonthFebruary),
		string(MonthJanuary),
		string(MonthJuly),
		string(MonthJune),
		string(MonthMarch),
		string(MonthMay),
		string(MonthNovember),
		string(MonthOctober),
		string(MonthSeptember),
	}
}

type WeekNumber string

const (
	WeekNumberFirst  WeekNumber = "First"
	WeekNumberFourth WeekNumber = "Fourth"
	WeekNumberLast   WeekNumber = "Last"
	WeekNumberSecond WeekNumber = "Second"
	WeekNumberThird  WeekNumber = "Third"
)

func PossibleValuesForWeekNumber() []string {
	return []string{
		string(WeekNumberFirst),
		string(WeekNumberFourth),
		string(WeekNumberLast),
		string(WeekNumberSecond),
		string(WeekNumberThird),
	}
}
